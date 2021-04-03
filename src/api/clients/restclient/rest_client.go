package restclient

import (
	"net/http"
)

func Get(url string) (*http.Response, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)

	client := http.Client{}
	return client.Do(request)
}
