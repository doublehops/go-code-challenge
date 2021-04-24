package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

	body, _ := json.Marshal(&payload)

	pl := bytes.NewBuffer(body)

	req, err := http.NewRequest("POST", "http://localhost:8080/isgood", pl)
	if err != nil {
		t.Fatal(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}

	puppy := api.GoodResponse{}

	rb, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(rb, &puppy)

	assert.Equal(t, "200 OK", resp.Status, "Record response is good")
	assert.Equal(t, "true", puppy.Puppy,  "Record response is good")
}