package restclient

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/yuricampolongo/crypto-monitoring/src/api/utils"
)

func Get(endpoint string, path string, queryParams map[string]string) (*http.Response, error) {
	var finalUrl *url.URL
	finalUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, errors.New("Invalid endpoint")
	}

	finalUrl.Path = path
	finalUrl.RawQuery = utils.BuildQueryParams(queryParams)

	request, err := http.NewRequest(http.MethodGet, finalUrl.String(), nil)
	fmt.Println(finalUrl.String())
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	return client.Do(request)
}
