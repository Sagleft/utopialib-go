package utopiago

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"gopkg.in/grignaak/tribool.v1"
)

func (c *UtopiaClient) apiQuery(methodName string, params map[string]interface{}) (map[string]interface{}, error) {
	var responseMap map[string]interface{}
	url := c.Protocol + "://" + c.Host + ":" + strconv.Itoa(c.Port) + "/api/1.0/"
	var query = Query{
		Method: methodName,
		Token:  c.Token,
	}
	if params != nil {
		query.Params = params
	}

	var jsonStr, err = json.Marshal(query)
	if err != nil {
		return responseMap, errors.New("failed to dedcode response json: " + err.Error())
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return responseMap, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	if !json.Valid(body) {
		return responseMap, errors.New("failed to validate json from client")
	}

	json.Unmarshal(body, &responseMap)
	return responseMap, nil
}

func (c *UtopiaClient) queryResultToInterface(methodName string, params map[string]interface{}) (interface{}, error) {
	if !c.CheckClientConnection() {
		return nil, errors.New("client disconected")
	}
	response, err := c.apiQuery(methodName, params)
	if result, ok := response["result"]; ok {
		return result, err
	}
	return nil, errors.New("result field doesn't exists in client response")
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
			return nil, errors.New("failed to get result array")
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
	if result, ok := response["result"]; ok {
		resultstr := fmt.Sprintf("%v", result)
		return resultstr, err
	}
	return "", errors.New("result field doesn't exists in client response")
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
