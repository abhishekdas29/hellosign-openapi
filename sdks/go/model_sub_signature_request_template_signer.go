/*
Dropbox Sign API

Dropbox Sign v3 API

API version: 3.0.0
Contact: apisupport@hellosign.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dropboxsign

import (
	"encoding/json"
)

// SubSignatureRequestTemplateSigner struct for SubSignatureRequestTemplateSigner
type SubSignatureRequestTemplateSigner struct {
	// Must match an existing role in chosen Template(s). It's case-sensitive.
	Role string `json:"role"`
	// The name of the signer.
	Name string `json:"name"`
	// The email address of the signer.
	EmailAddress string `json:"email_address"`
	// The 4- to 12-character access code that will secure this signer's signature page.
	Pin *string `json:"pin,omitempty"`
	// An E.164 formatted phone number.  **Note**: Not available in test mode and requires a Standard plan or higher.
	SmsPhoneNumber *string `json:"sms_phone_number,omitempty"`
	// Specifies the feature used with the `sms_phone_number`. Default `authentication`.  If `authentication`, signer is sent a verification code via SMS that is required to access the document.  If `delivery`, a link to complete the signature request is delivered via SMS (_and_ email).
	SmsPhoneNumberType *string `json:"sms_phone_number_type,omitempty"`
}

// NewSubSignatureRequestTemplateSigner instantiates a new SubSignatureRequestTemplateSigner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubSignatureRequestTemplateSigner(role string, name string, emailAddress string) *SubSignatureRequestTemplateSigner {
	this := SubSignatureRequestTemplateSigner{}
	this.Role = role
	this.Name = name
	this.EmailAddress = emailAddress
	return &this
}

// NewSubSignatureRequestTemplateSignerWithDefaults instantiates a new SubSignatureRequestTemplateSigner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubSignatureRequestTemplateSignerWithDefaults() *SubSignatureRequestTemplateSigner {
	this := SubSignatureRequestTemplateSigner{}
	return &this
}

// GetRole returns the Role field value
func (o *SubSignatureRequestTemplateSigner) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *SubSignatureRequestTemplateSigner) GetRoleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *SubSignatureRequestTemplateSigner) SetRole(v string) {
	o.Role = v
}

// GetName returns the Name field value
func (o *SubSignatureRequestTemplateSigner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SubSignatureRequestTemplateSigner) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SubSignatureRequestTemplateSigner) SetName(v string) {
	o.Name = v
}

// GetEmailAddress returns the EmailAddress field value
func (o *SubSignatureRequestTemplateSigner) GetEmailAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value
// and a boolean to check if the value has been set.
func (o *SubSignatureRequestTemplateSigner) GetEmailAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EmailAddress, true
}

// SetEmailAddress sets field value
func (o *SubSignatureRequestTemplateSigner) SetEmailAddress(v string) {
	o.EmailAddress = v
}

// GetPin returns the Pin field value if set, zero value otherwise.
func (o *SubSignatureRequestTemplateSigner) GetPin() string {
	if o == nil || o.Pin == nil {
		var ret string
		return ret
	}
	return *o.Pin
}

// GetPinOk returns a tuple with the Pin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubSignatureRequestTemplateSigner) GetPinOk() (*string, bool) {
	if o == nil || o.Pin == nil {
		return nil, false
	}
	return o.Pin, true
}

// HasPin returns a boolean if a field has been set.
func (o *SubSignatureRequestTemplateSigner) HasPin() bool {
	if o != nil && o.Pin != nil {
		return true
	}

	return false
}

// SetPin gets a reference to the given string and assigns it to the Pin field.
func (o *SubSignatureRequestTemplateSigner) SetPin(v string) {
	o.Pin = &v
}

// GetSmsPhoneNumber returns the SmsPhoneNumber field value if set, zero value otherwise.
func (o *SubSignatureRequestTemplateSigner) GetSmsPhoneNumber() string {
	if o == nil || o.SmsPhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.SmsPhoneNumber
}

// GetSmsPhoneNumberOk returns a tuple with the SmsPhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubSignatureRequestTemplateSigner) GetSmsPhoneNumberOk() (*string, bool) {
	if o == nil || o.SmsPhoneNumber == nil {
		return nil, false
	}
	return o.SmsPhoneNumber, true
}

// HasSmsPhoneNumber returns a boolean if a field has been set.
func (o *SubSignatureRequestTemplateSigner) HasSmsPhoneNumber() bool {
	if o != nil && o.SmsPhoneNumber != nil {
		return true
	}

	return false
}

// SetSmsPhoneNumber gets a reference to the given string and assigns it to the SmsPhoneNumber field.
func (o *SubSignatureRequestTemplateSigner) SetSmsPhoneNumber(v string) {
	o.SmsPhoneNumber = &v
}

// GetSmsPhoneNumberType returns the SmsPhoneNumberType field value if set, zero value otherwise.
func (o *SubSignatureRequestTemplateSigner) GetSmsPhoneNumberType() string {
	if o == nil || o.SmsPhoneNumberType == nil {
		var ret string
		return ret
	}
	return *o.SmsPhoneNumberType
}

// GetSmsPhoneNumberTypeOk returns a tuple with the SmsPhoneNumberType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubSignatureRequestTemplateSigner) GetSmsPhoneNumberTypeOk() (*string, bool) {
	if o == nil || o.SmsPhoneNumberType == nil {
		return nil, false
	}
	return o.SmsPhoneNumberType, true
}

// HasSmsPhoneNumberType returns a boolean if a field has been set.
func (o *SubSignatureRequestTemplateSigner) HasSmsPhoneNumberType() bool {
	if o != nil && o.SmsPhoneNumberType != nil {
		return true
	}

	return false
}

// SetSmsPhoneNumberType gets a reference to the given string and assigns it to the SmsPhoneNumberType field.
func (o *SubSignatureRequestTemplateSigner) SetSmsPhoneNumberType(v string) {
	o.SmsPhoneNumberType = &v
}

func (o SubSignatureRequestTemplateSigner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["role"] = o.Role
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["email_address"] = o.EmailAddress
	}
	if o.Pin != nil {
		toSerialize["pin"] = o.Pin
	}
	if o.SmsPhoneNumber != nil {
		toSerialize["sms_phone_number"] = o.SmsPhoneNumber
	}
	if o.SmsPhoneNumberType != nil {
		toSerialize["sms_phone_number_type"] = o.SmsPhoneNumberType
	}
	return json.Marshal(toSerialize)
}

type NullableSubSignatureRequestTemplateSigner struct {
	value *SubSignatureRequestTemplateSigner
	isSet bool
}

func (v NullableSubSignatureRequestTemplateSigner) Get() *SubSignatureRequestTemplateSigner {
	return v.value
}

func (v *NullableSubSignatureRequestTemplateSigner) Set(val *SubSignatureRequestTemplateSigner) {
	v.value = val
	v.isSet = true
}

func (v NullableSubSignatureRequestTemplateSigner) IsSet() bool {
	return v.isSet
}

func (v *NullableSubSignatureRequestTemplateSigner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubSignatureRequestTemplateSigner(val *SubSignatureRequestTemplateSigner) *NullableSubSignatureRequestTemplateSigner {
	return &NullableSubSignatureRequestTemplateSigner{value: val, isSet: true}
}

func (v NullableSubSignatureRequestTemplateSigner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubSignatureRequestTemplateSigner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


