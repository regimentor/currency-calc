package pkg

import (
	"fmt"
	"io"
	"net/http"
)

func HttpRequest(method, url string) ([]byte, error) {
	client := http.DefaultClient

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("create request due err: %v", err)
	}

	response, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request due err: %v", err)
	}

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("read response due err: %v", err)
	}

	return bytes, nil
}
