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
)

// EmailData struct for EmailData
type EmailData struct {
	Preview *EmailView `json:"Preview,omitempty"`
	// Attachments sent with the email
	Attachments []FileInfo `json:"Attachments,omitempty"`
	Status *EmailStatus `json:"Status,omitempty"`
}

// NewEmailData instantiates a new EmailData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailData() *EmailData {
	this := EmailData{}
	return &this
}

// NewEmailDataWithDefaults instantiates a new EmailData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailDataWithDefaults() *EmailData {
	this := EmailData{}
	return &this
}

// GetPreview returns the Preview field value if set, zero value otherwise.
func (o *EmailData) GetPreview() EmailView {
	if o == nil || isNil(o.Preview) {
		var ret EmailView
		return ret
	}
	return *o.Preview
}

// GetPreviewOk returns a tuple with the Preview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailData) GetPreviewOk() (*EmailView, bool) {
	if o == nil || isNil(o.Preview) {
    return nil, false
	}
	return o.Preview, true
}

// HasPreview returns a boolean if a field has been set.
func (o *EmailData) HasPreview() bool {
	if o != nil && !isNil(o.Preview) {
		return true
	}

	return false
}

// SetPreview gets a reference to the given EmailView and assigns it to the Preview field.
func (o *EmailData) SetPreview(v EmailView) {
	o.Preview = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *EmailData) GetAttachments() []FileInfo {
	if o == nil || isNil(o.Attachments) {
		var ret []FileInfo
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailData) GetAttachmentsOk() ([]FileInfo, bool) {
	if o == nil || isNil(o.Attachments) {
    return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *EmailData) HasAttachments() bool {
	if o != nil && !isNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []FileInfo and assigns it to the Attachments field.
func (o *EmailData) SetAttachments(v []FileInfo) {
	o.Attachments = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *EmailData) GetStatus() EmailStatus {
	if o == nil || isNil(o.Status) {
		var ret EmailStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailData) GetStatusOk() (*EmailStatus, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EmailData) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given EmailStatus and assigns it to the Status field.
func (o *EmailData) SetStatus(v EmailStatus) {
	o.Status = &v
}

func (o EmailData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Preview) {
		toSerialize["Preview"] = o.Preview
	}
	if !isNil(o.Attachments) {
		toSerialize["Attachments"] = o.Attachments
	}
	if !isNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableEmailData struct {
	value *EmailData
	isSet bool
}

func (v NullableEmailData) Get() *EmailData {
	return v.value
}

func (v *NullableEmailData) Set(val *EmailData) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailData) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailData(val *EmailData) *NullableEmailData {
	return &NullableEmailData{value: val, isSet: true}
}

func (v NullableEmailData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


