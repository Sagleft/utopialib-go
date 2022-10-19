package utopiago

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"time"

	simplecron "github.com/sagleft/simple-cron"
	"gopkg.in/grignaak/tribool.v1"
)

// get API url
func (c *UtopiaClient) getBaseURL() string {
	return c.Protocol + "://" + c.getBaseURLWithoutProtocol()
}

// get API url
func (c *UtopiaClient) getBaseURLWithoutProtocol() string {
	return c.Host + ":" + strconv.Itoa(c.Port) + "/api/1.0/"
}

// get ws API url
func (c *UtopiaClient) getWsURL() string {
	return "ws://" + c.Host + ":" + strconv.Itoa(c.WsPort) + "/UtopiaWSS?token=" + c.Token
}

func (c *UtopiaClient) apiQuery2JSON(
	methodName string,
	params map[string]interface{},
	filters map[string]interface{},
) ([]byte, error) {
	var query = Query{
		Method: methodName,
		Token:  c.Token,
	}
	if params != nil {
		query.Params = params
	}

	var jsonStr, err = json.Marshal(query)
	if err != nil {
		return nil, errors.New("failed to decode response json: " + err.Error())
	}

	req, err := http.NewRequest("POST", c.getBaseURL(), bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	defer closeRequest(resp)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("failed to read response body: " + err.Error())
	}

	return body, nil
}

func (c *UtopiaClient) apiQuery(methodName string, params map[string]interface{}) (map[string]interface{}, error) {
	return c.apiQueryWithFilters(methodName, params, map[string]interface{}{})
}

func (c *UtopiaClient) apiQueryWithFilters(
	methodName string,
	params,
	filters map[string]interface{},
) (map[string]interface{}, error) {

	var err error
	var jsonBody []byte
	var r map[string]interface{}

	if c.RequestTimeoutSeconds == 0 {
		jsonBody, err = c.apiQuery2JSON(methodName, params, filters)
	} else {
		if simplecron.NewRuntimeLimitHandler(time.Duration(c.RequestTimeoutSeconds), func() {
			jsonBody, err = c.apiQuery2JSON(methodName, params, filters)
		}).Run() {
			err = fmt.Errorf("failed to call %s: timeout", methodName)
		}
	}
	if err != nil {
		return r, err
	}

	if !json.Valid(jsonBody) {
		return r, errors.New("failed to validate response")
	}

	if err := json.Unmarshal(jsonBody, &r); err != nil {
		return r, fmt.Errorf("failed to decode response: %w", err)
	}
	return r, nil
}

func closeRequest(resp *http.Response) {
	if resp != nil {
		resp.Body.Close()
	}
}

func (c *UtopiaClient) queryResultToInterfaceArray(methodName string, params map[string]interface{}) ([]interface{}, error) {
	if !c.CheckClientConnection() {
		return nil, errors.New("client disconected")
	}
	response, err := c.apiQuery(methodName, params)
	if result, ok := response["result"]; ok {
		//check type assertion
		IResult, isConvertable := result.([]interface{})
		if !isConvertable {
			return nil, errors.New("failed to get result array")
		}
		return IResult, err
	}
	return nil, errors.New("accaptable result doesn't exists in client response")
}

func (c *UtopiaClient) queryResultToStringsArray(methodName string, params map[string]interface{}) ([]string, error) {
	if !c.CheckClientConnection() {
		return nil, errors.New("client disconected")
	}
	response, err := c.apiQuery(methodName, params)
	if result, ok := response["result"]; ok {
		//check type assertion
		IResult, isConvertable := result.([]string)
		if !isConvertable {
			resultType := reflect.TypeOf(result).String()
			if resultType == "[]interface {}" {
				IResult, isConvertable := result.([]interface{})
				if !isConvertable {
					return nil, errors.New("failed to get result array. can't convert to strings array")
				}
				resultArray := []string{}
				for _, elementRaw := range IResult {
					element, isConvertable := elementRaw.(string)
					if !isConvertable {
						return nil, errors.New("failed to convert result array element to string")
					}
					resultArray = append(resultArray, element)
				}
				return resultArray, nil
			}
			return nil, errors.New("failed to get result array. []string expected, " + resultType + "given")
		}
		return IResult, err
	}
	return nil, errors.New("accaptable result doesn't exists in client response")
}

func (c *UtopiaClient) queryResultToString(methodName string, params map[string]interface{}) (string, error) {
	if !c.CheckClientConnection() {
		return "", errors.New("client disconected")
	}
	response, err := c.apiQuery(methodName, params)
	if err != nil {
		return "", errors.New("failed to send API request: " + err.Error())
	}
	if result, ok := response["result"]; ok {
		resultstr := fmt.Sprintf("%v", result)
		return resultstr, err
	}

	errorInfoRaw, isErrorFound := response["error"]
	if isErrorFound {
		errorInfo, isConvertable := errorInfoRaw.(string)
		if !isConvertable {
			return "", errors.New("failed to parse error (type `" + reflect.ValueOf(errorInfoRaw).String() + "`) from result")
		}
		return "", errors.New(errorInfo)
	}

	return "", errors.New("result & error fields doesn't exists in client response")
}

func (c *UtopiaClient) queryResultToBool(methodName string, params map[string]interface{}) (bool, error) {
	resultstr, err := c.queryResultToString(methodName, params)
	resultBool := tribool.FromString(resultstr).WithMaybeAsTrue()
	return resultBool, err
}

func (c *UtopiaClient) queryResultToFloat64(methodName string, params map[string]interface{}) (float64, error) {
	resultstr, err := c.queryResultToString(methodName, params)
	if err != nil {
		return 0, err
	}
	resultFloat, err := strconv.ParseFloat(resultstr, 64)
	return resultFloat, err
}

func (c *UtopiaClient) queryResultToInt(methodName string, params map[string]interface{}) (int64, error) {
	resultstr, err := c.queryResultToString(methodName, params)
	if err != nil {
		return 0, err
	}
	result, err := strconv.ParseInt(resultstr, 10, 64)
	return result, err
}
