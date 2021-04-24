package api

type ActivityData struct {
	KvpKey   string `json:"kvpKey" validate:"required"`
	KvpValue string `json:"kvpValue" validate:"required"`
	KvpType  string `json:"kvpType" validate:"required,oneof=general.string general.integer general.float general.bool"`
}

type Device struct {
	CheckType       string         `json:"checkType" validate:"required,oneof=DEVICE BIOMETRIC COMBO"`
	ActivityType    string         `json:"activityType" validate:"required,oneof=SIGNUP LOGIN PAYMENT CONFIRMATION"`
	CheckSessionKey string         `json:"checkSessionKey" validate:"required"`
	ActivityData    []ActivityData `json:"activityData" validate:"required"`
}

type DeviceCheckDetails struct {
	Devices []Device `json:"deviceCheckDetails" validate:"required"`
}
