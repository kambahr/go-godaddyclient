package godaddyclient

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// IGodaddyAPI is the GoaddyClient interface.
type IGodaddyAPI interface {
	HTTPExec(method HTTPMethod, url, data string) Result
}

func (g *GodaddyAPI) HTTPExec(method HTTPMethod, url, data string) Result {

	var res Result

	client := &http.Client{}

	urlx := fmt.Sprintf("%s/%s/%s", g.BaseURL, g.Version, url)

	// clean up the url
	urlx = strings.ReplaceAll(urlx, "//", "/")
	urlx = strings.ReplaceAll(urlx, "https:/", "https://")
	urlx = strings.ReplaceAll(urlx, "http:/", "https://")

	httpMethod := fmt.Sprintf("%v", method)

	req, _ := http.NewRequest(httpMethod, urlx, bytes.NewBuffer([]byte(data)))
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

	return res
}
