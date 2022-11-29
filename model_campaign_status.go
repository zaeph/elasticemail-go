/*
Elastic Email REST API

This API is based on the REST API architecture, allowing the user to easily manage their data with this resource-based approach.    Every API call is established on which specific request type (GET, POST, PUT, DELETE) will be used.    The API has a limit of 20 concurrent connections and a hard timeout of 600 seconds per request.    To start using this API, you will need your Access Token (available <a target=\"_blank\" href=\"https://app.elasticemail.com/marketing/settings/new/manage-api\">here</a>). Remember to keep it safe. Required access levels are listed in the given request’s description.    Downloadable library clients can be found in our Github repository <a target=\"_blank\" href=\"https://github.com/ElasticEmail?tab=repositories&q=%22rest+api%22+in%3Areadme\">here</a>

API version: 4.0.0
Contact: support@elasticemail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ElasticEmail

import (
	"encoding/json"
	"fmt"
)

// CampaignStatus the model 'CampaignStatus'
type CampaignStatus string

// List of CampaignStatus
const (
	DELETED CampaignStatus = "Deleted"
	ACTIVE CampaignStatus = "Active"
	PROCESSING CampaignStatus = "Processing"
	SENDING CampaignStatus = "Sending"
	COMPLETED CampaignStatus = "Completed"
	PAUSED CampaignStatus = "Paused"
	CANCELLED CampaignStatus = "Cancelled"
	DRAFT CampaignStatus = "Draft"
)

// All allowed values of CampaignStatus enum
var AllowedCampaignStatusEnumValues = []CampaignStatus{
	"Deleted",
	"Active",
	"Processing",
	"Sending",
	"Completed",
	"Paused",
	"Cancelled",
	"Draft",
}

func (v *CampaignStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CampaignStatus(value)
	for _, existing := range AllowedCampaignStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CampaignStatus", value)
}

// NewCampaignStatusFromValue returns a pointer to a valid CampaignStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCampaignStatusFromValue(v string) (*CampaignStatus, error) {
	ev := CampaignStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CampaignStatus: valid values are %v", v, AllowedCampaignStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CampaignStatus) IsValid() bool {
	for _, existing := range AllowedCampaignStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CampaignStatus value
func (v CampaignStatus) Ptr() *CampaignStatus {
	return &v
}

type NullableCampaignStatus struct {
	value *CampaignStatus
	isSet bool
}

func (v NullableCampaignStatus) Get() *CampaignStatus {
	return v.value
}

func (v *NullableCampaignStatus) Set(val *CampaignStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignStatus(val *CampaignStatus) *NullableCampaignStatus {
	return &NullableCampaignStatus{value: val, isSet: true}
}

func (v NullableCampaignStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

