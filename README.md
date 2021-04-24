## APPLICATION NOTES

This application was built with Go 1.16.3. 

## RUNNING

- Clone repository with `git clone git@github.comdoublehops/go-code-challenge.git`
- Run `go run cmd/server/main.go`

This will start a webserver on port 8080. Make a request with cURL:
```bigquery
curl --location --request POST 'localhost:8080/isgood' \
--header 'Content-Type: application/json' \
--data-raw '{
    "deviceCheckDetails": [
        {
            "checkType": "DEVICE",
            "activityType": "PAYMENT",
            "checkSessionKey": "three",
            "activityData": [
                {
                    "kvpKey": "mykeyXXX",
                    "kvpValue": "kvpValueXXX",
                    "kvpType": "kvpTypeXXX"
                }
            ]
        }
    ]
}'
```

## TESTING

Run tests: `go test ./...`

## THIRD PARTY LIBRARIES USED

- github.com/gin-gonic/gin
- github.com/go-playground/validator/v10

## ASSUMPTIONS

- All fields are required. Swagger has no ability to set child properties as
required.

## THINGS MISSING

- Payload devices being sent as an array under the field `deviceCheckDetails`
- Vendor specific list accommodating in validation for `activityType`.
  Likely requires custom validation.
  
## PERSONAL THINGS YET TO UNDERSTAND

- In my experience, Go APIs try to parse JSON before trying to validate. 
  Therefore, if the JSON object doesn't match the struct, the application is likely
  to fail with a 500 response. There must be a better way to handle this.
  
## OTHER

Singular payload
```
{
    "checkType": "DEVICE",
    "activityType": "PAYMENT",
    "checkSessionKey": "three",
    "activityData": [
        {
            "kvpKey": "mykeyXXX",
            "kvpValue": "kvpValueXXX",
            "kvpType": "kvpTypeXXX"
        }
    ]
}
```