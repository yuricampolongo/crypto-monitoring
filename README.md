# GO Api to monitor crypto currencies
Go API to monitor crypto currency prices, save user preferences to easy investment monitoring and alerts based on thresholds defined by the user

I'm using this API to create an application to monitor my investments in crypto currencies and also to study Go

Feel free to fork and change whatever you need

*Still in development*


## Installation
For now, you must run the services manually, I'll provide soon the docker configurations to execute the project as a container.

`go run currencies_service/main.go`
`go run user_info_service/main.go`

## Endpoints

| Endpoint | Method | Description | Call Example
|--|--|--|--|
| /crypto/currency/:ids/:convert/:interval | GET | Return a list of prices for each crypto currency in the :ids param | http://localhost:8080/crypto/currency/BTC,ETH/BRL/1h
| /user/add | POST | Create/Update a new user to monitor his/hers investments | http://localhost:8081/user/add {"name":"Geralt","email":"geralt@ofrivia.com"}

## Configuring AWS credentials to use DynamoDB

You need to have 2 files in your environment to be able to use DynamoDB to store users information:
    `~/.aws/credentials`
    `~/.aws/config`

To create those files, please follow the AWS documentation: 
https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-quickstart.html#cli-configure-quickstart-config

Remember that using the AWS Free-Tier you can use up to 25gb of the DynamoDB, which is more than sufficient for this application.

**Note**: The user that you need to create to obtain the Access Key and Secret Key from AWS must have all the permissions on DynamoDB, to create tables and CRUD data.

## Benchmarks
To come