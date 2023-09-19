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

// TeamParentResponse Information about the parent team if a team has one, set to `null` otherwise.
type TeamParentResponse struct {
	// The id of a team
	TeamId *string `json:"team_id,omitempty"`
	// The name of a team
	Name *string `json:"name,omitempty"`
}

// NewTeamParentResponse instantiates a new TeamParentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamParentResponse() *TeamParentResponse {
	this := TeamParentResponse{}
	return &this
}

// NewTeamParentResponseWithDefaults instantiates a new TeamParentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamParentResponseWithDefaults() *TeamParentResponse {
	this := TeamParentResponse{}
	return &this
}

// GetTeamId returns the TeamId field value if set, zero value otherwise.
func (o *TeamParentResponse) GetTeamId() string {
	if o == nil || o.TeamId == nil {
		var ret string
		return ret
	}
	return *o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamParentResponse) GetTeamIdOk() (*string, bool) {
	if o == nil || o.TeamId == nil {
		return nil, false
	}
	return o.TeamId, true
}

// HasTeamId returns a boolean if a field has been set.
func (o *TeamParentResponse) HasTeamId() bool {
	if o != nil && o.TeamId != nil {
		return true
	}

	return false
}

// SetTeamId gets a reference to the given string and assigns it to the TeamId field.
func (o *TeamParentResponse) SetTeamId(v string) {
	o.TeamId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TeamParentResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamParentResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TeamParentResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TeamParentResponse) SetName(v string) {
	o.Name = &v
}

func (o TeamParentResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TeamId != nil {
		toSerialize["team_id"] = o.TeamId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableTeamParentResponse struct {
	value *TeamParentResponse
	isSet bool
}

func (v NullableTeamParentResponse) Get() *TeamParentResponse {
	return v.value
}

func (v *NullableTeamParentResponse) Set(val *TeamParentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamParentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamParentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamParentResponse(val *TeamParentResponse) *NullableTeamParentResponse {
	return &NullableTeamParentResponse{value: val, isSet: true}
}

func (v NullableTeamParentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamParentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

