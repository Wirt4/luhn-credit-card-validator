# luhn-credit-card-validator

##Running

###The App
`go run main.go` for a quick run
`go build main.go` to produce the executable

Note that the command line will run the server until you manually kill it.

###The Tests
`go test ./...` from the root

##Using
The app runs on `http://localhost:8080`
The GET endpoints are  `/validate/credit_card/`, and ``
The payload schema is 
```json
{
    "Number": "<your number sequence here>"
}
```
The endpoint returns
```json
{"ValidCreditCardNumber":"<boolean value>"}
```