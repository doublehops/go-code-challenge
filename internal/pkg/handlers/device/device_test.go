package device

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"codechallenge.local/internal/pkg/types/api"
)

func TestCreateRecordSuccess(t *testing.T) {

	payload := &api.DeviceCheckDetails{
		Devices: api.Device{
			CheckType:       "DEVICE",
			ActivityType:    "PAYMENT",
			CheckSessionKey: "Test",
			ActivityData: []api.ActivityData{{
				KvpKey:   "testKey",
				KvpValue: "testValue",
				KvpType:  "testType",
			}},
		},
	}

	d := New()
	err := d.CheckDevice(payload)

	assert.Nil(t, err, "Record created without error")
}

func TestCreateRecordValidationFail(t *testing.T) {

	payload := &api.DeviceCheckDetails{
		Devices: api.Device{
			CheckType:       "INCORRECT_VALUE",
			ActivityType:    "INCORRECT_ALSO",
			CheckSessionKey: "Test",
			ActivityData: []api.ActivityData{{
				KvpKey:   "testKey",
				KvpValue: "testValue",
				KvpType:  "testType",
			}},
		},
	}

	d := New()
	err := d.CheckDevice(payload)

	assert.Contains(t, err.Error(), "Field validation for 'CheckType' failed on the 'oneof' tag")
	assert.Contains(t, err.Error(), "Field validation for 'ActivityType' failed on the 'oneof' tag")
}
