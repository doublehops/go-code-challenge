## APPLICATION NOTES

This application was built with Go 1.16.3. 

I experienced some issues getting the payload to validate when I was sending devices
as an array so my examples here are all singular device requests.

I have been unable to have the test http request test the response body properly. I could 
continue to investigate but I've gone beyond the expected time to work on this so 
I'll leave as is. I think this work includes a good example of how I go about development.

PLEASE REFER to the `singlar-device-payload` branch for the completed work 
including tests.


## RUNNING

- Clone repository with `git clone git@github.comdoublehops/go-code-challenge.git`
- Possibly need to run `go mod vendor && go get -u ./...` to download dependencies.
- Run `go run cmd/server/main.go`

This will start a webserver on port 8080. Make a request with cURL:
```bigquery
{
    "deviceCheckDetails": {
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
}
```

## TESTING

Run tests: `go test ./...`

## DEPLOYMENT

There is currently no deployment.

## THIRD PARTY LIBRARIES USED

- github.com/gin-gonic/gin
- github.com/go-playground/validator/v10
- github.com/stretchr/testify/assert

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