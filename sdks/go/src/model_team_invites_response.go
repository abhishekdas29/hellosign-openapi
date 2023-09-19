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

// TeamInvitesResponse struct for TeamInvitesResponse
type TeamInvitesResponse struct {
	// Contains a list of team invites and their roles.
	TeamInvites *[]TeamInviteResponse `json:"team_invites,omitempty"`
	Warnings *[]WarningResponse `json:"warnings,omitempty"`
}

// NewTeamInvitesResponse instantiates a new TeamInvitesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamInvitesResponse() *TeamInvitesResponse {
	this := TeamInvitesResponse{}
	return &this
}

// NewTeamInvitesResponseWithDefaults instantiates a new TeamInvitesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamInvitesResponseWithDefaults() *TeamInvitesResponse {
	this := TeamInvitesResponse{}
	return &this
}

// GetTeamInvites returns the TeamInvites field value if set, zero value otherwise.
func (o *TeamInvitesResponse) GetTeamInvites() []TeamInviteResponse {
	if o == nil || o.TeamInvites == nil {
		var ret []TeamInviteResponse
		return ret
	}
	return *o.TeamInvites
}

// GetTeamInvitesOk returns a tuple with the TeamInvites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamInvitesResponse) GetTeamInvitesOk() (*[]TeamInviteResponse, bool) {
	if o == nil || o.TeamInvites == nil {
		return nil, false
	}
	return o.TeamInvites, true
}

// HasTeamInvites returns a boolean if a field has been set.
func (o *TeamInvitesResponse) HasTeamInvites() bool {
	if o != nil && o.TeamInvites != nil {
		return true
	}

	return false
}

// SetTeamInvites gets a reference to the given []TeamInviteResponse and assigns it to the TeamInvites field.
func (o *TeamInvitesResponse) SetTeamInvites(v []TeamInviteResponse) {
	o.TeamInvites = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *TeamInvitesResponse) GetWarnings() []WarningResponse {
	if o == nil || o.Warnings == nil {
		var ret []WarningResponse
		return ret
	}
	return *o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamInvitesResponse) GetWarningsOk() (*[]WarningResponse, bool) {
	if o == nil || o.Warnings == nil {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *TeamInvitesResponse) HasWarnings() bool {
	if o != nil && o.Warnings != nil {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningResponse and assigns it to the Warnings field.
func (o *TeamInvitesResponse) SetWarnings(v []WarningResponse) {
	o.Warnings = &v
}

func (o TeamInvitesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TeamInvites != nil {
		toSerialize["team_invites"] = o.TeamInvites
	}
	if o.Warnings != nil {
		toSerialize["warnings"] = o.Warnings
	}
	return json.Marshal(toSerialize)
}

type NullableTeamInvitesResponse struct {
	value *TeamInvitesResponse
	isSet bool
}

func (v NullableTeamInvitesResponse) Get() *TeamInvitesResponse {
	return v.value
}

func (v *NullableTeamInvitesResponse) Set(val *TeamInvitesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamInvitesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamInvitesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamInvitesResponse(val *TeamInvitesResponse) *NullableTeamInvitesResponse {
	return &NullableTeamInvitesResponse{value: val, isSet: true}
}

func (v NullableTeamInvitesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamInvitesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

