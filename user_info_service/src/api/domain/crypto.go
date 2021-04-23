package domain

type Crypto struct {
	Id      string   `json:"id"`
	UserID  string   `json:"user_id"`
	Cryptos []string `json:"cryptos"`
}

func (c Crypto) GetId() string {
	return c.Id
}
