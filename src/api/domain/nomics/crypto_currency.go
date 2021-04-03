package nomics

type CurrencyTickerRequest struct {
	Ids      string `json:"ids"`
	Convert  string `json:"convert"`
	Interval string `json:"interval"`
}

type CurrencyTickerResponse struct {
	Id        string `json:"id"`
	Currency  string `json:"currency"`
	Name      string `json:"name"`
	Logo      string `json:"logo_url"`
	Price     string `json:"price"`
	PriceDate string `json:"price_date"`
}
