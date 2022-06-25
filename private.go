package gogadaddyclient

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (g *GodaddyAPI) execute(method string, url, data string) Result {

	var res Result

	client := &http.Client{}

	req, _ := http.NewRequest(method, url, bytes.NewBuffer([]byte(data)))
	req.Header.Set("Authorization", fmt.Sprintf("sso-key %s:%s", g.Key, g.Secret))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		res.Message = err.Error()
		return res
	}
	body, _ := ioutil.ReadAll(resp.Body)
	req.Body.Close()
	resp.Body.Close()
	client.CloseIdleConnections()

	res.StatusCode = resp.StatusCode
	res.Message = string(body)

	if res.StatusCode == 200 || res.StatusCode == 204 {
		res.Message = "OK"
	}

	return res
}
