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