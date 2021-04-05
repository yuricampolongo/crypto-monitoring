# GO Api to monitor crypto currencies
Go API to monitor crypto currency prices and save user preferences to easy investment monitoring
I'm using this API to create an application to monitor my investments in crypto currencies and also to study Go
Feel free to fork and change whatever you need

*Still in development*


## Installation
Just run the application with `go run main.go`

## Endpoints

| Endpoint | Method | Description | Call Example
|--|--|--|--|
| /crypto/currency/:ids/:convert/:interval | GET | Return a list of prices for each crypto currency in the :ids param | /crypto/currency/BTC,ETH/BRL/1h

## Benchmarks
To come