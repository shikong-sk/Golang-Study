package httpMethod

import (
	"io/ioutil"
	"net/http"
	"time"
)

func Get(url string,timeout int32) (string, error) {
	client := &http.Client{Timeout: time.Duration(timeout) * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	baseBody, _ := ioutil.ReadAll(resp.Body)
	body := string(baseBody)
	return body, nil
}
