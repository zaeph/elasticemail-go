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
	"time"
)

// SmtpCredentials SMTP Credentials info
type SmtpCredentials struct {
	AccessLevel *AccessLevel `json:"AccessLevel,omitempty"`
	// Name of the key.
	Name *string `json:"Name,omitempty"`
	// Date this SmtpCredential was created.
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// Date this SmtpCredential was last used.
	LastUse NullableTime `json:"LastUse,omitempty"`
	// Date this SmtpCredential expires.
	Expires NullableTime `json:"Expires,omitempty"`
	// Which IPs can use this SmtpCredential
	RestrictAccessToIPRange []string `json:"RestrictAccessToIPRange,omitempty"`
}

// NewSmtpCredentials instantiates a new SmtpCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmtpCredentials() *SmtpCredentials {
	this := SmtpCredentials{}
	var accessLevel AccessLevel = NONE
	this.AccessLevel = &accessLevel
	return &this
}

// NewSmtpCredentialsWithDefaults instantiates a new SmtpCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmtpCredentialsWithDefaults() *SmtpCredentials {
	this := SmtpCredentials{}
	var accessLevel AccessLevel = NONE
	this.AccessLevel = &accessLevel
	return &this
}

// GetAccessLevel returns the AccessLevel field value if set, zero value otherwise.
func (o *SmtpCredentials) GetAccessLevel() AccessLevel {
	if o == nil || isNil(o.AccessLevel) {
		var ret AccessLevel
		return ret
	}
	return *o.AccessLevel
}

// GetAccessLevelOk returns a tuple with the AccessLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpCredentials) GetAccessLevelOk() (*AccessLevel, bool) {
	if o == nil || isNil(o.AccessLevel) {
    return nil, false
	}
	return o.AccessLevel, true
}

// HasAccessLevel returns a boolean if a field has been set.
func (o *SmtpCredentials) HasAccessLevel() bool {
	if o != nil && !isNil(o.AccessLevel) {
		return true
	}

	return false
}

// SetAccessLevel gets a reference to the given AccessLevel and assigns it to the AccessLevel field.
func (o *SmtpCredentials) SetAccessLevel(v AccessLevel) {
	o.AccessLevel = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SmtpCredentials) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpCredentials) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SmtpCredentials) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SmtpCredentials) SetName(v string) {
	o.Name = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *SmtpCredentials) GetDateCreated() time.Time {
	if o == nil || isNil(o.DateCreated) {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpCredentials) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || isNil(o.DateCreated) {
    return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *SmtpCredentials) HasDateCreated() bool {
	if o != nil && !isNil(o.DateCreated) {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *SmtpCredentials) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetLastUse returns the LastUse field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmtpCredentials) GetLastUse() time.Time {
	if o == nil || isNil(o.LastUse.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastUse.Get()
}

// GetLastUseOk returns a tuple with the LastUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmtpCredentials) GetLastUseOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return o.LastUse.Get(), o.LastUse.IsSet()
}

// HasLastUse returns a boolean if a field has been set.
func (o *SmtpCredentials) HasLastUse() bool {
	if o != nil && o.LastUse.IsSet() {
		return true
	}

	return false
}

// SetLastUse gets a reference to the given NullableTime and assigns it to the LastUse field.
func (o *SmtpCredentials) SetLastUse(v time.Time) {
	o.LastUse.Set(&v)
}
// SetLastUseNil sets the value for LastUse to be an explicit nil
func (o *SmtpCredentials) SetLastUseNil() {
	o.LastUse.Set(nil)
}

// UnsetLastUse ensures that no value is present for LastUse, not even an explicit nil
func (o *SmtpCredentials) UnsetLastUse() {
	o.LastUse.Unset()
}

// GetExpires returns the Expires field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmtpCredentials) GetExpires() time.Time {
	if o == nil || isNil(o.Expires.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Expires.Get()
}

// GetExpiresOk returns a tuple with the Expires field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmtpCredentials) GetExpiresOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return o.Expires.Get(), o.Expires.IsSet()
}

// HasExpires returns a boolean if a field has been set.
func (o *SmtpCredentials) HasExpires() bool {
	if o != nil && o.Expires.IsSet() {
		return true
	}

	return false
}

// SetExpires gets a reference to the given NullableTime and assigns it to the Expires field.
func (o *SmtpCredentials) SetExpires(v time.Time) {
	o.Expires.Set(&v)
}
// SetExpiresNil sets the value for Expires to be an explicit nil
func (o *SmtpCredentials) SetExpiresNil() {
	o.Expires.Set(nil)
}

// UnsetExpires ensures that no value is present for Expires, not even an explicit nil
func (o *SmtpCredentials) UnsetExpires() {
	o.Expires.Unset()
}

// GetRestrictAccessToIPRange returns the RestrictAccessToIPRange field value if set, zero value otherwise.
func (o *SmtpCredentials) GetRestrictAccessToIPRange() []string {
	if o == nil || isNil(o.RestrictAccessToIPRange) {
		var ret []string
		return ret
	}
	return o.RestrictAccessToIPRange
}

// GetRestrictAccessToIPRangeOk returns a tuple with the RestrictAccessToIPRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpCredentials) GetRestrictAccessToIPRangeOk() ([]string, bool) {
	if o == nil || isNil(o.RestrictAccessToIPRange) {
    return nil, false
	}
	return o.RestrictAccessToIPRange, true
}

// HasRestrictAccessToIPRange returns a boolean if a field has been set.
func (o *SmtpCredentials) HasRestrictAccessToIPRange() bool {
	if o != nil && !isNil(o.RestrictAccessToIPRange) {
		return true
	}

	return false
}

// SetRestrictAccessToIPRange gets a reference to the given []string and assigns it to the RestrictAccessToIPRange field.
func (o *SmtpCredentials) SetRestrictAccessToIPRange(v []string) {
	o.RestrictAccessToIPRange = v
}

func (o SmtpCredentials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AccessLevel) {
		toSerialize["AccessLevel"] = o.AccessLevel
	}
	if !isNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !isNil(o.DateCreated) {
		toSerialize["DateCreated"] = o.DateCreated
	}
	if o.LastUse.IsSet() {
		toSerialize["LastUse"] = o.LastUse.Get()
	}
	if o.Expires.IsSet() {
		toSerialize["Expires"] = o.Expires.Get()
	}
	if !isNil(o.RestrictAccessToIPRange) {
		toSerialize["RestrictAccessToIPRange"] = o.RestrictAccessToIPRange
	}
	return json.Marshal(toSerialize)
}

type NullableSmtpCredentials struct {
	value *SmtpCredentials
	isSet bool
}

func (v NullableSmtpCredentials) Get() *SmtpCredentials {
	return v.value
}

func (v *NullableSmtpCredentials) Set(val *SmtpCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableSmtpCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableSmtpCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmtpCredentials(val *SmtpCredentials) *NullableSmtpCredentials {
	return &NullableSmtpCredentials{value: val, isSet: true}
}

func (v NullableSmtpCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmtpCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


