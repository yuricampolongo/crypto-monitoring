package utils

import (
	"net/url"
)

func BuildQueryParams(queryParams map[string]string) string {
	params := url.Values{}
	for k, v := range queryParams {
		params.Add(k, v)
	}
	return params.Encode()
}
