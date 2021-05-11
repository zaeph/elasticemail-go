/*
 * Elastic Email REST API
 *
 * This API is based on the REST API architecture, allowing the user to easily manage their data with this resource-based approach.    Every API call is established on which specific request type (GET, POST, PUT, DELETE) will be used.    To start using this API, you will need your Access Token (available <a href=\"https://elasticemail.com/account#/settings/new/manage-api\">here</a>). Remember to keep it safe. Required access levels are listed in the given request’s description.    This is the documentation for REST API. If you’d like to read our legacy documentation regarding Web API v2 click <a href=\"https://api.elasticemail.com/public/help\">here</a>.
 *
 * API version: 4.0.0
 * Contact: support@elasticemail.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ElasticEmail

import (
	"encoding/json"
	"fmt"
)

// LogJobStatus the model 'LogJobStatus'
type LogJobStatus string

// List of LogJobStatus
const (
	ALL LogJobStatus = "All"
	READY_TO_SEND LogJobStatus = "ReadyToSend"
	WAITING_TO_RETRY LogJobStatus = "WaitingToRetry"
	SENDING LogJobStatus = "Sending"
	ERROR LogJobStatus = "Error"
	SENT LogJobStatus = "Sent"
	OPENED LogJobStatus = "Opened"
	CLICKED LogJobStatus = "Clicked"
	UNSUBSCRIBED LogJobStatus = "Unsubscribed"
	ABUSE_REPORT LogJobStatus = "AbuseReport"
)

func (v *LogJobStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LogJobStatus(value)
	for _, existing := range []LogJobStatus{ "All", "ReadyToSend", "WaitingToRetry", "Sending", "Error", "Sent", "Opened", "Clicked", "Unsubscribed", "AbuseReport",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LogJobStatus", value)
}

// Ptr returns reference to LogJobStatus value
func (v LogJobStatus) Ptr() *LogJobStatus {
	return &v
}

type NullableLogJobStatus struct {
	value *LogJobStatus
	isSet bool
}

func (v NullableLogJobStatus) Get() *LogJobStatus {
	return v.value
}

func (v *NullableLogJobStatus) Set(val *LogJobStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableLogJobStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableLogJobStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogJobStatus(val *LogJobStatus) *NullableLogJobStatus {
	return &NullableLogJobStatus{value: val, isSet: true}
}

func (v NullableLogJobStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogJobStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
