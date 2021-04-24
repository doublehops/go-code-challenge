package api

type DeviceCheckDetails struct {
	Devices []Device `json:"deviceCheckDetails"`
}

type ActivityData struct {
	KvpKey   string `json:"kvpKey"`
	KvpValue string `json:"kvpValue"`
	KvpType  string `json:"kvpType"`
}

type Device struct {
	CheckType       string         `json:"checkType" validate:"required,oneof=DEVICE BIOMETRIC COMBO"`
	ActivityType    string         `json:"activityType" validate:"required,oneof=SIGNUP LOGIN PAYMENT CONFIRMATION"`
	CheckSessionKey string         `json:"checkSessionKey" validate:"required"`
	ActivityData    []ActivityData `json:"activityData"`
}
