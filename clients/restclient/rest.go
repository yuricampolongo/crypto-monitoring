package restclient

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/go-resty/resty/v2"
)

var (
	Do RestInterface
)

type RestInterface interface {
	Get(endpoint string, path string, queryParams map[string]string) (*Response, error)
}

type rest struct {
}

func init() {
	Do = &rest{}
}

func (c *rest) Get(endpoint string, path string, queryParams map[string]string) (*Response, error) {
	var finalUrl *url.URL

	finalUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, errors.New("invalid endpoint")
	}

	finalUrl.Path = path
	finalUrl.RawQuery = buildQueryParams(queryParams)

	client := resty.New()

	response, err := client.R().Get(finalUrl.String())
	fmt.Println(finalUrl.String())
	if err != nil {
		errorDesc := fmt.Sprintf("error to perform GET request to endpoint (%v)", finalUrl.String())
		return nil, errors.New(errorDesc)
	}

	return &Response{
		StatusCode: response.StatusCode(),
		Body:       response.String(),
	}, nil
}

func buildQueryParams(queryParams map[string]string) string {
	params := url.Values{}
	for k, v := range queryParams {
		params.Add(k, v)
	}
	return params.Encode()
}
