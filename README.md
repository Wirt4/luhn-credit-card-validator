# luhn-credit-card-validator

##Running

###The App
`go run main.go` for a quick run
`go build` to produce the executable

Note that the command line will run the server until you manually kill it.

###The Tests
`go test ./...` from the root


##Using

The endpoint is `http://localhost:8080/validate_credit_card/`.
The payload schema is 
```json
{
    "CreditCardNumber": "<your number sequence here>"
}
```
The endpoint returns
```json
{"ValidCreditCardNumber":"<boolean value>"}
```