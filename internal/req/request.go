package req

import (
	"io"
	"net/http"
)

func getURL(params map[string]string) string {
	var url = "http://www.google.com"
	if len(params) == 0 {
		return url
	}

	url = url + "?"

	var b = true
	for k, v := range params {
		if b {
			b = false
		} else {
			url = url + "&"
		}
		url = url + k + "=" + v
	}
	return url
}

func GetRequest1(params map[string]string) (string, error) {
	url := getURL(params)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	//req.Header.Set()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
