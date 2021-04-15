package domain

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCurrencyTickerRequestSucess(t *testing.T) {
	request := CurrencyRequest{
		Ids:      "BTC,ETH",
		Convert:  "CAD",
		Interval: "1h",
	}

	bytes, err := json.Marshal(request)
	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	var target CurrencyRequest

	err = json.Unmarshal(bytes, &target)
	assert.Nil(t, err)
	assert.EqualValues(t, target.Ids, request.Ids)
}
