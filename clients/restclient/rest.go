package restclient

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/go-resty/resty/v2"
)

func Get(endpoint string, path string, queryParams map[string]string) (*http.Response, error) {
	var finalUrl *url.URL

	finalUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, errors.New("Invalid endpoint")
	}

	finalUrl.Path = path
	finalUrl.RawQuery = buildQueryParams(queryParams)

	client := resty.New()

	response, err := client.R().Get(finalUrl.String())
	fmt.Println(finalUrl.String())
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error to perform GET request to endpoint (%v)", finalUrl.String()))
	}

	return response.RawResponse, nil
}

func buildQueryParams(queryParams map[string]string) string {
	params := url.Values{}
	for k, v := range queryParams {
		params.Add(k, v)
	}
	return params.Encode()
}
