package reqhandler

import "net/http"

func closeRequest(resp *http.Response) {
	if resp != nil {
		resp.Body.Close()
	}
}
