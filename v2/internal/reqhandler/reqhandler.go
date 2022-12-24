package reqhandler

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type RequestHandler interface {
	Send(reqType, URL string, data []byte) ([]byte, error)
}

type defaultHandler struct {
	client *http.Client
}

func NewDefaultHandler(timeout time.Duration) RequestHandler {
	return &defaultHandler{
		client: &http.Client{Timeout: timeout},
	}
}

func (h *defaultHandler) Send(reqType, URL string, data []byte) ([]byte, error) {

	req, err := http.NewRequest(reqType, URL, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := h.client.Do(req)
	defer closeRequest(resp)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return body, nil
}
