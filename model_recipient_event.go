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
	"time"
)

// RecipientEvent Detailed information about message recipient
type RecipientEvent struct {
	// ID number of transaction
	TransactionID *string `json:"TransactionID,omitempty"`
	// ID number of selected message.
	MsgID *string `json:"MsgID,omitempty"`
	// Default From: email address.
	FromEmail *string `json:"FromEmail,omitempty"`
	// Ending date for search in YYYY-MM-DDThh:mm:ss format.
	To *string `json:"To,omitempty"`
	// Default subject of email.
	Subject *string `json:"Subject,omitempty"`
	// Type of an Event
	EventType *EventType `json:"EventType,omitempty"`
	// Creation date
	EventDate *time.Time `json:"EventDate,omitempty"`
	// Name of selected channel.
	ChannelName *string `json:"ChannelName,omitempty"`
	// Message category
	MessageCategory *MessageCategory `json:"MessageCategory,omitempty"`
	// Date of next try
	NextTryOn NullableTime `json:"NextTryOn,omitempty"`
	// Content of message, HTML encoded
	Message *string `json:"Message,omitempty"`
	// IP which this email was sent through
	IPAddress *string `json:"IPAddress,omitempty"`
	// Name of an IP pool this email was sent through
	PoolName *string `json:"PoolName,omitempty"`
}

// NewRecipientEvent instantiates a new RecipientEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecipientEvent() *RecipientEvent {
	this := RecipientEvent{}
	return &this
}

// NewRecipientEventWithDefaults instantiates a new RecipientEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecipientEventWithDefaults() *RecipientEvent {
	this := RecipientEvent{}
	return &this
}

// GetTransactionID returns the TransactionID field value if set, zero value otherwise.
func (o *RecipientEvent) GetTransactionID() string {
	if o == nil || o.TransactionID == nil {
		var ret string
		return ret
	}
	return *o.TransactionID
}

// GetTransactionIDOk returns a tuple with the TransactionID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientEvent) GetTransactionIDOk() (*string, bool) {
	if o == nil || o.TransactionID == nil {
		return nil, false
	}
	return o.TransactionID, true
}

// HasTransactionID returns a boolean if a field has been set.
func (o *RecipientEvent) HasTransactionID() bool {
	if o != nil && o.TransactionID != nil {
		return true
	}

	return false
}

// SetTransactionID gets a reference to the given string and assigns it to the TransactionID field.
func (o *RecipientEvent) SetTransactionID(v string) {
	o.TransactionID = &v
}

// GetMsgID returns the MsgID field value if set, zero value otherwise.
func (o *RecipientEvent) GetMsgID() string {
	if o == nil || o.MsgID == nil {
		var ret string
		return ret
	}
	return *o.MsgID
}

// GetMsgIDOk returns a tuple with the MsgID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientEvent) GetMsgIDOk() (*string, bool) {
	if o == nil || o.MsgID == nil {
		return nil, false
	}
	return o.MsgID, true
}

// HasMsgID returns a boolean if a field has been set.
func (o *RecipientEvent) HasMsgID() bool {
	if o != nil && o.MsgID != nil {
		return true
	}

	return false
}

// SetMsgID gets a reference to the given string and assigns it to the MsgID field.
func (o *RecipientEvent) SetMsgID(v string) {
	o.MsgID = &v
}

// GetFromEmail returns the FromEmail field value if set, zero value otherwise.
func (o *RecipientEvent) GetFromEmail() string {
	if o == nil || o.FromEmail == nil {
		var ret string
		return ret
	}
	return *o.FromEmail
}

// GetFromEmailOk returns a tuple with the FromEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientEvent) GetFromEmailOk() (*string, bool) {
	if o == nil || o.FromEmail == nil {
		return nil, false
	}
	return o.FromEmail, true
}

// HasFromEmail returns a boolean if a field has been set.
func (o *RecipientEvent) HasFromEmail() bool {
	if o != nil && o.FromEmail != nil {
		return true
	}

	return false
}

// SetFromEmail gets a reference to the given string and assigns it to the FromEmail field.
func (o *RecipientEvent) SetFromEmail(v string) {
	o.FromEmail = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *RecipientEvent) GetTo() string {
	if o == nil || o.To == nil {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientEvent) GetToOk() (*string, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *RecipientEvent) HasTo() bool {
	if o != nil && o.To != nil {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *RecipientEvent) SetTo(v string) {
	o.To = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *RecipientEvent) GetSubject() string {
	if o == nil || o.Subject == nil {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientEvent) GetSubjectOk() (*string, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *RecipientEvent) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *RecipientEvent) SetSubject(v string) {
	o.Subject = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *RecipientEvent) GetEventType() EventType {
	if o == nil || o.EventType == nil {
		var ret EventType
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientEvent) GetEventTypeOk() (*EventType, bool) {
	if o == nil || o.EventType == nil {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *RecipientEvent) HasEventType() bool {
	if o != nil && o.EventType != nil {
		return true
	}

	return false
}

// SetEventType gets a reference to the given EventType and assigns it to the EventType field.
func (o *RecipientEvent) SetEventType(v EventType) {
	o.EventType = &v
}

// GetEventDate returns the EventDate field value if set, zero value otherwise.
func (o *RecipientEvent) GetEventDate() time.Time {
	if o == nil || o.EventDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EventDate
}

// GetEventDateOk returns a tuple with the EventDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientEvent) GetEventDateOk() (*time.Time, bool) {
	if o == nil || o.EventDate == nil {
		return nil, false
	}
	return o.EventDate, true
}

// HasEventDate returns a boolean if a field has been set.
func (o *RecipientEvent) HasEventDate() bool {
	if o != nil && o.EventDate != nil {
		return true
	}

	return false
}

// SetEventDate gets a reference to the given time.Time and assigns it to the EventDate field.
func (o *RecipientEvent) SetEventDate(v time.Time) {
	o.EventDate = &v
}

// GetChannelName returns the ChannelName field value if set, zero value otherwise.
func (o *RecipientEvent) GetChannelName() string {
	if o == nil || o.ChannelName == nil {
		var ret string
		return ret
	}
	return *o.ChannelName
}

// GetChannelNameOk returns a tuple with the ChannelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientEvent) GetChannelNameOk() (*string, bool) {
	if o == nil || o.ChannelName == nil {
		return nil, false
	}
	return o.ChannelName, true
}

// HasChannelName returns a boolean if a field has been set.
func (o *RecipientEvent) HasChannelName() bool {
	if o != nil && o.ChannelName != nil {
		return true
	}

	return false
}

// SetChannelName gets a reference to the given string and assigns it to the ChannelName field.
func (o *RecipientEvent) SetChannelName(v string) {
	o.ChannelName = &v
}

// GetMessageCategory returns the MessageCategory field value if set, zero value otherwise.
func (o *RecipientEvent) GetMessageCategory() MessageCategory {
	if o == nil || o.MessageCategory == nil {
		var ret MessageCategory
		return ret
	}
	return *o.MessageCategory
}

// GetMessageCategoryOk returns a tuple with the MessageCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientEvent) GetMessageCategoryOk() (*MessageCategory, bool) {
	if o == nil || o.MessageCategory == nil {
		return nil, false
	}
	return o.MessageCategory, true
}

// HasMessageCategory returns a boolean if a field has been set.
func (o *RecipientEvent) HasMessageCategory() bool {
	if o != nil && o.MessageCategory != nil {
		return true
	}

	return false
}

// SetMessageCategory gets a reference to the given MessageCategory and assigns it to the MessageCategory field.
func (o *RecipientEvent) SetMessageCategory(v MessageCategory) {
	o.MessageCategory = &v
}

// GetNextTryOn returns the NextTryOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecipientEvent) GetNextTryOn() time.Time {
	if o == nil || o.NextTryOn.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.NextTryOn.Get()
}

// GetNextTryOnOk returns a tuple with the NextTryOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecipientEvent) GetNextTryOnOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NextTryOn.Get(), o.NextTryOn.IsSet()
}

// HasNextTryOn returns a boolean if a field has been set.
func (o *RecipientEvent) HasNextTryOn() bool {
	if o != nil && o.NextTryOn.IsSet() {
		return true
	}

	return false
}

// SetNextTryOn gets a reference to the given NullableTime and assigns it to the NextTryOn field.
func (o *RecipientEvent) SetNextTryOn(v time.Time) {
	o.NextTryOn.Set(&v)
}
// SetNextTryOnNil sets the value for NextTryOn to be an explicit nil
func (o *RecipientEvent) SetNextTryOnNil() {
	o.NextTryOn.Set(nil)
}

// UnsetNextTryOn ensures that no value is present for NextTryOn, not even an explicit nil
func (o *RecipientEvent) UnsetNextTryOn() {
	o.NextTryOn.Unset()
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *RecipientEvent) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientEvent) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *RecipientEvent) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *RecipientEvent) SetMessage(v string) {
	o.Message = &v
}

// GetIPAddress returns the IPAddress field value if set, zero value otherwise.
func (o *RecipientEvent) GetIPAddress() string {
	if o == nil || o.IPAddress == nil {
		var ret string
		return ret
	}
	return *o.IPAddress
}

// GetIPAddressOk returns a tuple with the IPAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientEvent) GetIPAddressOk() (*string, bool) {
	if o == nil || o.IPAddress == nil {
		return nil, false
	}
	return o.IPAddress, true
}

// HasIPAddress returns a boolean if a field has been set.
func (o *RecipientEvent) HasIPAddress() bool {
	if o != nil && o.IPAddress != nil {
		return true
	}

	return false
}

// SetIPAddress gets a reference to the given string and assigns it to the IPAddress field.
func (o *RecipientEvent) SetIPAddress(v string) {
	o.IPAddress = &v
}

// GetPoolName returns the PoolName field value if set, zero value otherwise.
func (o *RecipientEvent) GetPoolName() string {
	if o == nil || o.PoolName == nil {
		var ret string
		return ret
	}
	return *o.PoolName
}

// GetPoolNameOk returns a tuple with the PoolName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientEvent) GetPoolNameOk() (*string, bool) {
	if o == nil || o.PoolName == nil {
		return nil, false
	}
	return o.PoolName, true
}

// HasPoolName returns a boolean if a field has been set.
func (o *RecipientEvent) HasPoolName() bool {
	if o != nil && o.PoolName != nil {
		return true
	}

	return false
}

// SetPoolName gets a reference to the given string and assigns it to the PoolName field.
func (o *RecipientEvent) SetPoolName(v string) {
	o.PoolName = &v
}

func (o RecipientEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TransactionID != nil {
		toSerialize["TransactionID"] = o.TransactionID
	}
	if o.MsgID != nil {
		toSerialize["MsgID"] = o.MsgID
	}
	if o.FromEmail != nil {
		toSerialize["FromEmail"] = o.FromEmail
	}
	if o.To != nil {
		toSerialize["To"] = o.To
	}
	if o.Subject != nil {
		toSerialize["Subject"] = o.Subject
	}
	if o.EventType != nil {
		toSerialize["EventType"] = o.EventType
	}
	if o.EventDate != nil {
		toSerialize["EventDate"] = o.EventDate
	}
	if o.ChannelName != nil {
		toSerialize["ChannelName"] = o.ChannelName
	}
	if o.MessageCategory != nil {
		toSerialize["MessageCategory"] = o.MessageCategory
	}
	if o.NextTryOn.IsSet() {
		toSerialize["NextTryOn"] = o.NextTryOn.Get()
	}
	if o.Message != nil {
		toSerialize["Message"] = o.Message
	}
	if o.IPAddress != nil {
		toSerialize["IPAddress"] = o.IPAddress
	}
	if o.PoolName != nil {
		toSerialize["PoolName"] = o.PoolName
	}
	return json.Marshal(toSerialize)
}

type NullableRecipientEvent struct {
	value *RecipientEvent
	isSet bool
}

func (v NullableRecipientEvent) Get() *RecipientEvent {
	return v.value
}

func (v *NullableRecipientEvent) Set(val *RecipientEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableRecipientEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableRecipientEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecipientEvent(val *RecipientEvent) *NullableRecipientEvent {
	return &NullableRecipientEvent{value: val, isSet: true}
}

func (v NullableRecipientEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecipientEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

