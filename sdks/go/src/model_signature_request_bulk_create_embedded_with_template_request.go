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
	"os"
)

// SignatureRequestBulkCreateEmbeddedWithTemplateRequest struct for SignatureRequestBulkCreateEmbeddedWithTemplateRequest
type SignatureRequestBulkCreateEmbeddedWithTemplateRequest struct {
	// Use `template_ids` to create a SignatureRequest from one or more templates, in the order in which the template will be used.
	TemplateIds []string `json:"template_ids"`
	// Client id of the app you're using to create this embedded signature request. Used for security purposes.
	ClientId string `json:"client_id"`
	// `signer_file` is a CSV file defining values and options for signer fields. Required unless a `signer_list` is used, you may not use both. The CSV can have the following columns:  - `name`: the name of the signer filling the role of RoleName - `email_address`: email address of the signer filling the role of RoleName - `pin`: the 4- to 12-character access code that will secure this signer's signature page (optional) - `sms_phone_number`: An E.164 formatted phone number that will receive a code via SMS to access this signer's signature page. (optional)      **Note**: Not available in test mode and requires a Standard plan or higher. - `*_field`: any column with a _field\" suffix will be treated as a custom field (optional)      You may only specify field values here, any other options should be set in the custom_fields request parameter.  Example CSV:  ``` name, email_address, pin, company_field George, george@example.com, d79a3td, ABC Corp Mary, mary@example.com, gd9as5b, 123 LLC ```
	SignerFile **os.File `json:"signer_file,omitempty"`
	// `signer_list` is an array defining values and options for signer fields. Required unless a `signer_file` is used, you may not use both.
	SignerList *[]SubBulkSignerList `json:"signer_list,omitempty"`
	// Allows signers to decline to sign a document if `true`. Defaults to `false`.
	AllowDecline *bool `json:"allow_decline,omitempty"`
	// Add CC email recipients. Required when a CC role exists for the Template.
	Ccs *[]SubCC `json:"ccs,omitempty"`
	// When used together with merge fields, `custom_fields` allows users to add pre-filled data to their signature requests.  Pre-filled data can be used with \"send-once\" signature requests by adding merge fields with `form_fields_per_document` or [Text Tags](https://app.hellosign.com/api/textTagsWalkthrough#TextTagIntro) while passing values back with `custom_fields` together in one API call.  For using pre-filled on repeatable signature requests, merge fields are added to templates in the Dropbox Sign UI or by calling [/template/create_embedded_draft](/api/reference/operation/templateCreateEmbeddedDraft) and then passing `custom_fields` on subsequent signature requests referencing that template.
	CustomFields *[]SubCustomField `json:"custom_fields,omitempty"`
	// The custom message in the email that will be sent to the signers.
	Message *string `json:"message,omitempty"`
	// Key-value data that should be attached to the signature request. This metadata is included in all API responses and events involving the signature request. For example, use the metadata field to store a signer's order number for look up when receiving events for the signature request.  Each request can include up to 10 metadata keys (or 50 nested metadata keys), with key names up to 40 characters long and values up to 1000 characters long.
	Metadata *map[string]interface{} `json:"metadata,omitempty"`
	// The URL you want signers redirected to after they successfully sign.
	SigningRedirectUrl *string `json:"signing_redirect_url,omitempty"`
	// The subject in the email that will be sent to the signers.
	Subject *string `json:"subject,omitempty"`
	// Whether this is a test, the signature request will not be legally binding if set to `true`. Defaults to `false`.
	TestMode *bool `json:"test_mode,omitempty"`
	// The title you want to assign to the SignatureRequest.
	Title *string `json:"title,omitempty"`
}

// NewSignatureRequestBulkCreateEmbeddedWithTemplateRequest instantiates a new SignatureRequestBulkCreateEmbeddedWithTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignatureRequestBulkCreateEmbeddedWithTemplateRequest(templateIds []string, clientId string) *SignatureRequestBulkCreateEmbeddedWithTemplateRequest {
	this := SignatureRequestBulkCreateEmbeddedWithTemplateRequest{}
	this.TemplateIds = templateIds
	this.ClientId = clientId
	var allowDecline bool = false
	this.AllowDecline = &allowDecline
	var testMode bool = false
	this.TestMode = &testMode
	return &this
}

// NewSignatureRequestBulkCreateEmbeddedWithTemplateRequestWithDefaults instantiates a new SignatureRequestBulkCreateEmbeddedWithTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignatureRequestBulkCreateEmbeddedWithTemplateRequestWithDefaults() *SignatureRequestBulkCreateEmbeddedWithTemplateRequest {
	this := SignatureRequestBulkCreateEmbeddedWithTemplateRequest{}
	var allowDecline bool = false
	this.AllowDecline = &allowDecline
	var testMode bool = false
	this.TestMode = &testMode
	return &this
}

// GetTemplateIds returns the TemplateIds field value
func (o *SignatureRequestBulkCreateEmbeddedWithTemplateRequest) GetTemplateIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TemplateIds
}

// GetTemplateIdsOk returns a tuple with the TemplateIds field value
// and a boolean to check if the value has been set.
func (o *SignatureRequestBulkCreateEmbeddedWithTemplateRequest) GetTemplateIdsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TemplateIds, true
}

// SetTemplateIds sets field value
func (o *SignatureRequestBulkCreateEmbeddedWithTemplateRequest) SetTemplateIds(v []string) {
	o.TemplateIds = v
}

// GetClientId returns the ClientId field value
func (o *SignatureRequestBulkCreateEmbeddedWithTemplateRequest) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *SignatureRequestBulkCreateEmbeddedWithTemplateRequest) GetClientIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *SignatureRequestBulkCreateEmbeddedWithTemplateRequest) SetClientId(v string) {
	o.ClientId = v
}

// GetSignerFile returns the SignerFile field value if set, zero value otherwise.
func (o *SignatureRequestBulkCreateEmbeddedWithTemplateRequest) GetSignerFile() *os.File {
	if o == nil || o.SignerFile == nil {
		var ret *os.File
		return ret
	}
	return *o.SignerFile
}

// GetSignerFileOk returns a tuple with the SignerFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureRequestBulkCreateEmbeddedWithTemplateRequest) GetSignerFileOk() (**os.File, bool) {
	if o == nil || o.SignerFile == nil {
		return nil, false
	}
	return o.SignerFile, true
}

// HasSignerFile returns a boolean if a field has been set.
func (o *SignatureRequestBulkCreateEmbeddedWithTemplateRequest) HasSignerFile() bool {
	if o != nil && o.SignerFile != nil {
		return true
	}

	return false
}

// SetSignerFile gets a reference to the given *os.File and assigns it to the SignerFile field.
func (o *SignatureRequestBulkCreateEmbeddedWithTemplateRequest) SetSignerFile(v *os.File) {
	o.SignerFile = &v
}

// GetSignerList returns the SignerList field value if set, zero value otherwise.
func (o *SignatureRequestBulkCreateEmbeddedWithTemplateRequest) GetSignerList() []SubBulkSignerList {
	if o == nil || o.SignerList == nil {
		var ret []SubBulkSignerList
		return ret
	}
	return *o.SignerList
}

// GetSignerListOk returns a tuple with the SignerList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureRequestBulkCreateEmbeddedWithTemplateRequest) GetSignerListOk() (*[]SubBulkSignerList, bool) {
	if o == nil || o.SignerList == nil {
		return nil, false
	}
	return o.SignerList, true
}

// HasSignerList returns a boolean if a field has been set.
func (o *SignatureRequestBulkCreateEmbeddedWithTemplateRequest) HasSignerList() bool {
	if o != nil && o.SignerList != nil {
		return true
	}

	return false
}

// SetSignerList gets a reference to the given []SubBulkSignerList and assigns it to the SignerList field.
func (o *SignatureRequestBulkCreateEmbeddedWithTemplateRequest) SetSignerList(v []SubBulkSignerList) {
	o.SignerList = &v
}

// GetAllowDecline returns the AllowDecline field value if set, zero value otherwise.
func (o *SignatureRequestBulkCreateEmbeddedWithTemplateRequest) GetAllowDecline() bool {
	if o == nil || o.AllowDecline == nil {
		var ret bool
		return ret
	}
	return *o.AllowDecline
}

// GetAllowDeclineOk returns a tuple with the AllowDecline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureRequestBulkCreateEmbeddedWithTemplateRequest) GetAllowDeclineOk() (*bool, bool) {
	if o == nil || o.AllowDecline == nil {
		return nil, false
	}
	return o.AllowDecline, true
}

// HasAllowDecline returns a boolean if a field has been set.
func (o *SignatureRequestBulkCreateEmbeddedWithTemplateRequest) HasAllowDecline() bool {
	if o != nil && o.AllowDecline != nil {
		return true
	}

	return false
}

// SetAllowDecline gets a reference to the given bool and assigns it to the AllowDecline field.
func (o *SignatureRequestBulkCreateEmbeddedWithTemplateRequest) SetAllowDecline(v bool) {
	o.AllowDecline = &v
}

// GetCcs returns the Ccs field value if set, zero value otherwise.
func (o *SignatureRequestBulkCreateEmbeddedWithTemplateRequest) GetCcs() []SubCC {
	if o == nil || o.Ccs == nil {
		var ret []SubCC
		return ret
	}
	return *o.Ccs
}

// GetCcsOk returns a tuple with the Ccs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureRequestBulkCreateEmbeddedWithTemplateRequest) GetCcsOk() (*[]SubCC, bool) {
	if o == nil || o.Ccs == nil {
		return nil, false
	}
	return o.Ccs, true
}

// HasCcs returns a boolean if a field has been set.
func (o *SignatureRequestBulkCreateEmbeddedWithTemplateRequest) HasCcs() bool {
	if o != nil && o.Ccs != nil {
		return true
	}

	return false
}

// SetCcs gets a reference to the given []SubCC and assigns it to the Ccs field.
func (o *SignatureRequestBulkCreateEmbeddedWithTemplateRequest) SetCcs(v []SubCC) {
	o.Ccs = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *SignatureRequestBulkCreateEmbeddedWithTemplateRequest) GetCustomFields() []SubCustomField {
	if o == nil || o.CustomFields == nil {
		var ret []SubCustomField
		return ret
	}
	return *o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureRequestBulkCreateEmbeddedWithTemplateRequest) GetCustomFieldsOk() (*[]SubCustomField, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *SignatureRequestBulkCreateEmbeddedWithTemplateRequest) HasCustomFields() bool {
	if o != nil && o.CustomFields != nil {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given []SubCustomField and assigns it to the CustomFields field.
func (o *SignatureRequestBulkCreateEmbeddedWithTemplateRequest) SetCustomFields(v []SubCustomField) {
	o.CustomFields = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *SignatureRequestBulkCreateEmbeddedWithTemplateRequest) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureRequestBulkCreateEmbeddedWithTemplateRequest) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *SignatureRequestBulkCreateEmbeddedWithTemplateRequest) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *SignatureRequestBulkCreateEmbeddedWithTemplateRequest) SetMessage(v string) {
	o.Message = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *SignatureRequestBulkCreateEmbeddedWithTemplateRequest) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureRequestBulkCreateEmbeddedWithTemplateRequest) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *SignatureRequestBulkCreateEmbeddedWithTemplateRequest) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *SignatureRequestBulkCreateEmbeddedWithTemplateRequest) SetMetadata(v map[string]interface{}) {
	o.Metadata = &v
}

// GetSigningRedirectUrl returns the SigningRedirectUrl field value if set, zero value otherwise.
func (o *SignatureRequestBulkCreateEmbeddedWithTemplateRequest) GetSigningRedirectUrl() string {
	if o == nil || o.SigningRedirectUrl == nil {
		var ret string
		return ret
	}
	return *o.SigningRedirectUrl
}

// GetSigningRedirectUrlOk returns a tuple with the SigningRedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureRequestBulkCreateEmbeddedWithTemplateRequest) GetSigningRedirectUrlOk() (*string, bool) {
	if o == nil || o.SigningRedirectUrl == nil {
		return nil, false
	}
	return o.SigningRedirectUrl, true
}

// HasSigningRedirectUrl returns a boolean if a field has been set.
func (o *SignatureRequestBulkCreateEmbeddedWithTemplateRequest) HasSigningRedirectUrl() bool {
	if o != nil && o.SigningRedirectUrl != nil {
		return true
	}

	return false
}

// SetSigningRedirectUrl gets a reference to the given string and assigns it to the SigningRedirectUrl field.
func (o *SignatureRequestBulkCreateEmbeddedWithTemplateRequest) SetSigningRedirectUrl(v string) {
	o.SigningRedirectUrl = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *SignatureRequestBulkCreateEmbeddedWithTemplateRequest) GetSubject() string {
	if o == nil || o.Subject == nil {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureRequestBulkCreateEmbeddedWithTemplateRequest) GetSubjectOk() (*string, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *SignatureRequestBulkCreateEmbeddedWithTemplateRequest) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *SignatureRequestBulkCreateEmbeddedWithTemplateRequest) SetSubject(v string) {
	o.Subject = &v
}

// GetTestMode returns the TestMode field value if set, zero value otherwise.
func (o *SignatureRequestBulkCreateEmbeddedWithTemplateRequest) GetTestMode() bool {
	if o == nil || o.TestMode == nil {
		var ret bool
		return ret
	}
	return *o.TestMode
}

// GetTestModeOk returns a tuple with the TestMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureRequestBulkCreateEmbeddedWithTemplateRequest) GetTestModeOk() (*bool, bool) {
	if o == nil || o.TestMode == nil {
		return nil, false
	}
	return o.TestMode, true
}

// HasTestMode returns a boolean if a field has been set.
func (o *SignatureRequestBulkCreateEmbeddedWithTemplateRequest) HasTestMode() bool {
	if o != nil && o.TestMode != nil {
		return true
	}

	return false
}

// SetTestMode gets a reference to the given bool and assigns it to the TestMode field.
func (o *SignatureRequestBulkCreateEmbeddedWithTemplateRequest) SetTestMode(v bool) {
	o.TestMode = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *SignatureRequestBulkCreateEmbeddedWithTemplateRequest) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureRequestBulkCreateEmbeddedWithTemplateRequest) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *SignatureRequestBulkCreateEmbeddedWithTemplateRequest) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *SignatureRequestBulkCreateEmbeddedWithTemplateRequest) SetTitle(v string) {
	o.Title = &v
}

func (o SignatureRequestBulkCreateEmbeddedWithTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["template_ids"] = o.TemplateIds
	}
	if true {
		toSerialize["client_id"] = o.ClientId
	}
	if o.SignerFile != nil {
		toSerialize["signer_file"] = o.SignerFile
	}
	if o.SignerList != nil {
		toSerialize["signer_list"] = o.SignerList
	}
	if o.AllowDecline != nil {
		toSerialize["allow_decline"] = o.AllowDecline
	}
	if o.Ccs != nil {
		toSerialize["ccs"] = o.Ccs
	}
	if o.CustomFields != nil {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.SigningRedirectUrl != nil {
		toSerialize["signing_redirect_url"] = o.SigningRedirectUrl
	}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}
	if o.TestMode != nil {
		toSerialize["test_mode"] = o.TestMode
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	return json.Marshal(toSerialize)
}

type NullableSignatureRequestBulkCreateEmbeddedWithTemplateRequest struct {
	value *SignatureRequestBulkCreateEmbeddedWithTemplateRequest
	isSet bool
}

func (v NullableSignatureRequestBulkCreateEmbeddedWithTemplateRequest) Get() *SignatureRequestBulkCreateEmbeddedWithTemplateRequest {
	return v.value
}

func (v *NullableSignatureRequestBulkCreateEmbeddedWithTemplateRequest) Set(val *SignatureRequestBulkCreateEmbeddedWithTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSignatureRequestBulkCreateEmbeddedWithTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSignatureRequestBulkCreateEmbeddedWithTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignatureRequestBulkCreateEmbeddedWithTemplateRequest(val *SignatureRequestBulkCreateEmbeddedWithTemplateRequest) *NullableSignatureRequestBulkCreateEmbeddedWithTemplateRequest {
	return &NullableSignatureRequestBulkCreateEmbeddedWithTemplateRequest{value: val, isSet: true}
}

func (v NullableSignatureRequestBulkCreateEmbeddedWithTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignatureRequestBulkCreateEmbeddedWithTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

