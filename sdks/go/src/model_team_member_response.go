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

// TeamMemberResponse struct for TeamMemberResponse
type TeamMemberResponse struct {
	// Account id of the team member.
	AccountId *string `json:"account_id,omitempty"`
	// Email address of the team member.
	EmailAddress *string `json:"email_address,omitempty"`
	// The specific role a member has on the team.
	Role *string `json:"role,omitempty"`
}

// NewTeamMemberResponse instantiates a new TeamMemberResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamMemberResponse() *TeamMemberResponse {
	this := TeamMemberResponse{}
	return &this
}

// NewTeamMemberResponseWithDefaults instantiates a new TeamMemberResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamMemberResponseWithDefaults() *TeamMemberResponse {
	this := TeamMemberResponse{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *TeamMemberResponse) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamMemberResponse) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *TeamMemberResponse) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *TeamMemberResponse) SetAccountId(v string) {
	o.AccountId = &v
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise.
func (o *TeamMemberResponse) GetEmailAddress() string {
	if o == nil || o.EmailAddress == nil {
		var ret string
		return ret
	}
	return *o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamMemberResponse) GetEmailAddressOk() (*string, bool) {
	if o == nil || o.EmailAddress == nil {
		return nil, false
	}
	return o.EmailAddress, true
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *TeamMemberResponse) HasEmailAddress() bool {
	if o != nil && o.EmailAddress != nil {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given string and assigns it to the EmailAddress field.
func (o *TeamMemberResponse) SetEmailAddress(v string) {
	o.EmailAddress = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *TeamMemberResponse) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamMemberResponse) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *TeamMemberResponse) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *TeamMemberResponse) SetRole(v string) {
	o.Role = &v
}

func (o TeamMemberResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
	}
	if o.EmailAddress != nil {
		toSerialize["email_address"] = o.EmailAddress
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	return json.Marshal(toSerialize)
}

type NullableTeamMemberResponse struct {
	value *TeamMemberResponse
	isSet bool
}

func (v NullableTeamMemberResponse) Get() *TeamMemberResponse {
	return v.value
}

func (v *NullableTeamMemberResponse) Set(val *TeamMemberResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamMemberResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamMemberResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamMemberResponse(val *TeamMemberResponse) *NullableTeamMemberResponse {
	return &NullableTeamMemberResponse{value: val, isSet: true}
}

func (v NullableTeamMemberResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamMemberResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

