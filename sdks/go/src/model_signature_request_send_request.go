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

// SignatureRequestSendRequest struct for SignatureRequestSendRequest
type SignatureRequestSendRequest struct {
	// Use `files[]` to indicate the uploaded file(s) to send for signature.  This endpoint requires either **files** or **file_urls[]**, but not both.
	Files *[]*os.File `json:"files,omitempty"`
	// Use `file_urls[]` to have Dropbox Sign download the file(s) to send for signature.  This endpoint requires either **files** or **file_urls[]**, but not both.
	FileUrls *[]string `json:"file_urls,omitempty"`
	// Add Signers to your Signature Request.  This endpoint requires either **signers** or **grouped_signers**, but not both.
	Signers *[]SubSignatureRequestSigner `json:"signers,omitempty"`
	// Add Grouped Signers to your Signature Request.  This endpoint requires either **signers** or **grouped_signers**, but not both.
	GroupedSigners *[]SubSignatureRequestGroupedSigners `json:"grouped_signers,omitempty"`
	// Allows signers to decline to sign a document if `true`. Defaults to `false`.
	AllowDecline *bool `json:"allow_decline,omitempty"`
	// Allows signers to reassign their signature requests to other signers if set to `true`. Defaults to `false`.  **Note**: Only available for Premium plan and higher.
	AllowReassign *bool `json:"allow_reassign,omitempty"`
	// A list describing the attachments
	Attachments *[]SubAttachment `json:"attachments,omitempty"`
	// The email addresses that should be CCed.
	CcEmailAddresses *[]string `json:"cc_email_addresses,omitempty"`
	// The client id of the API App you want to associate with this request. Used to apply the branding and callback url defined for the app.
	ClientId *string `json:"client_id,omitempty"`
	// When used together with merge fields, `custom_fields` allows users to add pre-filled data to their signature requests.  Pre-filled data can be used with \"send-once\" signature requests by adding merge fields with `form_fields_per_document` or [Text Tags](https://app.hellosign.com/api/textTagsWalkthrough#TextTagIntro) while passing values back with `custom_fields` together in one API call.  For using pre-filled on repeatable signature requests, merge fields are added to templates in the Dropbox Sign UI or by calling [/template/create_embedded_draft](/api/reference/operation/templateCreateEmbeddedDraft) and then passing `custom_fields` on subsequent signature requests referencing that template.
	CustomFields *[]SubCustomField `json:"custom_fields,omitempty"`
	FieldOptions *SubFieldOptions `json:"field_options,omitempty"`
	// Group information for fields defined in `form_fields_per_document`. String-indexed JSON array with `group_label` and `requirement` keys. `form_fields_per_document` must contain fields referencing a group defined in `form_field_groups`.
	FormFieldGroups *[]SubFormFieldGroup `json:"form_field_groups,omitempty"`
	// Conditional Logic rules for fields defined in `form_fields_per_document`.
	FormFieldRules *[]SubFormFieldRule `json:"form_field_rules,omitempty"`
	// The fields that should appear on the document, expressed as an array of objects. (For more details you can read about it here: [Using Form Fields per Document](/docs/openapi/form-fields-per-document).)  **NOTE**: Fields like **text**, **dropdown**, **checkbox**, **radio**, and **hyperlink** have additional required and optional parameters. Check out the list of [additional parameters](/api/reference/constants/#form-fields-per-document) for these field types.  * Text Field use `SubFormFieldsPerDocumentText` * Dropdown Field use `SubFormFieldsPerDocumentDropdown` * Hyperlink Field use `SubFormFieldsPerDocumentHyperlink` * Checkbox Field use `SubFormFieldsPerDocumentCheckbox` * Radio Field use `SubFormFieldsPerDocumentRadio` * Signature Field use `SubFormFieldsPerDocumentSignature` * Date Signed Field use `SubFormFieldsPerDocumentDateSigned` * Initials Field use `SubFormFieldsPerDocumentInitials` * Text Merge Field use `SubFormFieldsPerDocumentTextMerge` * Checkbox Merge Field use `SubFormFieldsPerDocumentCheckboxMerge`
	FormFieldsPerDocument *[]SubFormFieldsPerDocumentBase `json:"form_fields_per_document,omitempty"`
	// Enables automatic Text Tag removal when set to true.  **NOTE**: Removing text tags this way can cause unwanted clipping. We recommend leaving this setting on `false` and instead hiding your text tags using white text or a similar approach. See the [Text Tags Walkthrough](https://app.hellosign.com/api/textTagsWalkthrough#TextTagIntro) for more information.
	HideTextTags *bool `json:"hide_text_tags,omitempty"`
	// Send with a value of `true` if you wish to enable [Qualified Electronic Signatures](https://www.hellosign.com/features/qualified-electronic-signatures) (QES), which requires a face-to-face call to verify the signer's identity.<br> **Note**: QES is only available on the Premium API plan as an add-on purchase. Cannot be used in `test_mode`. Only works on requests with one signer.
	IsQualifiedSignature *bool `json:"is_qualified_signature,omitempty"`
	// The custom message in the email that will be sent to the signers.
	Message *string `json:"message,omitempty"`
	// Key-value data that should be attached to the signature request. This metadata is included in all API responses and events involving the signature request. For example, use the metadata field to store a signer's order number for look up when receiving events for the signature request.  Each request can include up to 10 metadata keys (or 50 nested metadata keys), with key names up to 40 characters long and values up to 1000 characters long.
	Metadata *map[string]interface{} `json:"metadata,omitempty"`
	SigningOptions *SubSigningOptions `json:"signing_options,omitempty"`
	// The URL you want signers redirected to after they successfully sign.
	SigningRedirectUrl *string `json:"signing_redirect_url,omitempty"`
	// The subject in the email that will be sent to the signers.
	Subject *string `json:"subject,omitempty"`
	// Whether this is a test, the signature request will not be legally binding if set to `true`. Defaults to `false`.
	TestMode *bool `json:"test_mode,omitempty"`
	// The title you want to assign to the SignatureRequest.
	Title *string `json:"title,omitempty"`
	// Send with a value of `true` if you wish to enable [Text Tags](https://app.hellosign.com/api/textTagsWalkthrough#TextTagIntro) parsing in your document. Defaults to disabled, or `false`.
	UseTextTags *bool `json:"use_text_tags,omitempty"`
	// When the signature request will expire. Unsigned signatures will be moved to the expired status, and no longer signable. See [Signature Request Expiration Date](https://developers.hellosign.com/docs/signature-request/expiration/) for details.
	ExpiresAt NullableInt32 `json:"expires_at,omitempty"`
}

// NewSignatureRequestSendRequest instantiates a new SignatureRequestSendRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignatureRequestSendRequest() *SignatureRequestSendRequest {
	this := SignatureRequestSendRequest{}
	var allowDecline bool = false
	this.AllowDecline = &allowDecline
	var allowReassign bool = false
	this.AllowReassign = &allowReassign
	var hideTextTags bool = false
	this.HideTextTags = &hideTextTags
	var isQualifiedSignature bool = false
	this.IsQualifiedSignature = &isQualifiedSignature
	var testMode bool = false
	this.TestMode = &testMode
	var useTextTags bool = false
	this.UseTextTags = &useTextTags
	return &this
}

// NewSignatureRequestSendRequestWithDefaults instantiates a new SignatureRequestSendRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignatureRequestSendRequestWithDefaults() *SignatureRequestSendRequest {
	this := SignatureRequestSendRequest{}
	var allowDecline bool = false
	this.AllowDecline = &allowDecline
	var allowReassign bool = false
	this.AllowReassign = &allowReassign
	var hideTextTags bool = false
	this.HideTextTags = &hideTextTags
	var isQualifiedSignature bool = false
	this.IsQualifiedSignature = &isQualifiedSignature
	var testMode bool = false
	this.TestMode = &testMode
	var useTextTags bool = false
	this.UseTextTags = &useTextTags
	return &this
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *SignatureRequestSendRequest) GetFiles() []*os.File {
	if o == nil || o.Files == nil {
		var ret []*os.File
		return ret
	}
	return *o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureRequestSendRequest) GetFilesOk() (*[]*os.File, bool) {
	if o == nil || o.Files == nil {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *SignatureRequestSendRequest) HasFiles() bool {
	if o != nil && o.Files != nil {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []*os.File and assigns it to the Files field.
func (o *SignatureRequestSendRequest) SetFiles(v []*os.File) {
	o.Files = &v
}

// GetFileUrls returns the FileUrls field value if set, zero value otherwise.
func (o *SignatureRequestSendRequest) GetFileUrls() []string {
	if o == nil || o.FileUrls == nil {
		var ret []string
		return ret
	}
	return *o.FileUrls
}

// GetFileUrlsOk returns a tuple with the FileUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureRequestSendRequest) GetFileUrlsOk() (*[]string, bool) {
	if o == nil || o.FileUrls == nil {
		return nil, false
	}
	return o.FileUrls, true
}

// HasFileUrls returns a boolean if a field has been set.
func (o *SignatureRequestSendRequest) HasFileUrls() bool {
	if o != nil && o.FileUrls != nil {
		return true
	}

	return false
}

// SetFileUrls gets a reference to the given []string and assigns it to the FileUrls field.
func (o *SignatureRequestSendRequest) SetFileUrls(v []string) {
	o.FileUrls = &v
}

// GetSigners returns the Signers field value if set, zero value otherwise.
func (o *SignatureRequestSendRequest) GetSigners() []SubSignatureRequestSigner {
	if o == nil || o.Signers == nil {
		var ret []SubSignatureRequestSigner
		return ret
	}
	return *o.Signers
}

// GetSignersOk returns a tuple with the Signers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureRequestSendRequest) GetSignersOk() (*[]SubSignatureRequestSigner, bool) {
	if o == nil || o.Signers == nil {
		return nil, false
	}
	return o.Signers, true
}

// HasSigners returns a boolean if a field has been set.
func (o *SignatureRequestSendRequest) HasSigners() bool {
	if o != nil && o.Signers != nil {
		return true
	}

	return false
}

// SetSigners gets a reference to the given []SubSignatureRequestSigner and assigns it to the Signers field.
func (o *SignatureRequestSendRequest) SetSigners(v []SubSignatureRequestSigner) {
	o.Signers = &v
}

// GetGroupedSigners returns the GroupedSigners field value if set, zero value otherwise.
func (o *SignatureRequestSendRequest) GetGroupedSigners() []SubSignatureRequestGroupedSigners {
	if o == nil || o.GroupedSigners == nil {
		var ret []SubSignatureRequestGroupedSigners
		return ret
	}
	return *o.GroupedSigners
}

// GetGroupedSignersOk returns a tuple with the GroupedSigners field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureRequestSendRequest) GetGroupedSignersOk() (*[]SubSignatureRequestGroupedSigners, bool) {
	if o == nil || o.GroupedSigners == nil {
		return nil, false
	}
	return o.GroupedSigners, true
}

// HasGroupedSigners returns a boolean if a field has been set.
func (o *SignatureRequestSendRequest) HasGroupedSigners() bool {
	if o != nil && o.GroupedSigners != nil {
		return true
	}

	return false
}

// SetGroupedSigners gets a reference to the given []SubSignatureRequestGroupedSigners and assigns it to the GroupedSigners field.
func (o *SignatureRequestSendRequest) SetGroupedSigners(v []SubSignatureRequestGroupedSigners) {
	o.GroupedSigners = &v
}

// GetAllowDecline returns the AllowDecline field value if set, zero value otherwise.
func (o *SignatureRequestSendRequest) GetAllowDecline() bool {
	if o == nil || o.AllowDecline == nil {
		var ret bool
		return ret
	}
	return *o.AllowDecline
}

// GetAllowDeclineOk returns a tuple with the AllowDecline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureRequestSendRequest) GetAllowDeclineOk() (*bool, bool) {
	if o == nil || o.AllowDecline == nil {
		return nil, false
	}
	return o.AllowDecline, true
}

// HasAllowDecline returns a boolean if a field has been set.
func (o *SignatureRequestSendRequest) HasAllowDecline() bool {
	if o != nil && o.AllowDecline != nil {
		return true
	}

	return false
}

// SetAllowDecline gets a reference to the given bool and assigns it to the AllowDecline field.
func (o *SignatureRequestSendRequest) SetAllowDecline(v bool) {
	o.AllowDecline = &v
}

// GetAllowReassign returns the AllowReassign field value if set, zero value otherwise.
func (o *SignatureRequestSendRequest) GetAllowReassign() bool {
	if o == nil || o.AllowReassign == nil {
		var ret bool
		return ret
	}
	return *o.AllowReassign
}

// GetAllowReassignOk returns a tuple with the AllowReassign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureRequestSendRequest) GetAllowReassignOk() (*bool, bool) {
	if o == nil || o.AllowReassign == nil {
		return nil, false
	}
	return o.AllowReassign, true
}

// HasAllowReassign returns a boolean if a field has been set.
func (o *SignatureRequestSendRequest) HasAllowReassign() bool {
	if o != nil && o.AllowReassign != nil {
		return true
	}

	return false
}

// SetAllowReassign gets a reference to the given bool and assigns it to the AllowReassign field.
func (o *SignatureRequestSendRequest) SetAllowReassign(v bool) {
	o.AllowReassign = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *SignatureRequestSendRequest) GetAttachments() []SubAttachment {
	if o == nil || o.Attachments == nil {
		var ret []SubAttachment
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureRequestSendRequest) GetAttachmentsOk() (*[]SubAttachment, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *SignatureRequestSendRequest) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []SubAttachment and assigns it to the Attachments field.
func (o *SignatureRequestSendRequest) SetAttachments(v []SubAttachment) {
	o.Attachments = &v
}

// GetCcEmailAddresses returns the CcEmailAddresses field value if set, zero value otherwise.
func (o *SignatureRequestSendRequest) GetCcEmailAddresses() []string {
	if o == nil || o.CcEmailAddresses == nil {
		var ret []string
		return ret
	}
	return *o.CcEmailAddresses
}

// GetCcEmailAddressesOk returns a tuple with the CcEmailAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureRequestSendRequest) GetCcEmailAddressesOk() (*[]string, bool) {
	if o == nil || o.CcEmailAddresses == nil {
		return nil, false
	}
	return o.CcEmailAddresses, true
}

// HasCcEmailAddresses returns a boolean if a field has been set.
func (o *SignatureRequestSendRequest) HasCcEmailAddresses() bool {
	if o != nil && o.CcEmailAddresses != nil {
		return true
	}

	return false
}

// SetCcEmailAddresses gets a reference to the given []string and assigns it to the CcEmailAddresses field.
func (o *SignatureRequestSendRequest) SetCcEmailAddresses(v []string) {
	o.CcEmailAddresses = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *SignatureRequestSendRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureRequestSendRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *SignatureRequestSendRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *SignatureRequestSendRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *SignatureRequestSendRequest) GetCustomFields() []SubCustomField {
	if o == nil || o.CustomFields == nil {
		var ret []SubCustomField
		return ret
	}
	return *o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureRequestSendRequest) GetCustomFieldsOk() (*[]SubCustomField, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *SignatureRequestSendRequest) HasCustomFields() bool {
	if o != nil && o.CustomFields != nil {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given []SubCustomField and assigns it to the CustomFields field.
func (o *SignatureRequestSendRequest) SetCustomFields(v []SubCustomField) {
	o.CustomFields = &v
}

// GetFieldOptions returns the FieldOptions field value if set, zero value otherwise.
func (o *SignatureRequestSendRequest) GetFieldOptions() SubFieldOptions {
	if o == nil || o.FieldOptions == nil {
		var ret SubFieldOptions
		return ret
	}
	return *o.FieldOptions
}

// GetFieldOptionsOk returns a tuple with the FieldOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureRequestSendRequest) GetFieldOptionsOk() (*SubFieldOptions, bool) {
	if o == nil || o.FieldOptions == nil {
		return nil, false
	}
	return o.FieldOptions, true
}

// HasFieldOptions returns a boolean if a field has been set.
func (o *SignatureRequestSendRequest) HasFieldOptions() bool {
	if o != nil && o.FieldOptions != nil {
		return true
	}

	return false
}

// SetFieldOptions gets a reference to the given SubFieldOptions and assigns it to the FieldOptions field.
func (o *SignatureRequestSendRequest) SetFieldOptions(v SubFieldOptions) {
	o.FieldOptions = &v
}

// GetFormFieldGroups returns the FormFieldGroups field value if set, zero value otherwise.
func (o *SignatureRequestSendRequest) GetFormFieldGroups() []SubFormFieldGroup {
	if o == nil || o.FormFieldGroups == nil {
		var ret []SubFormFieldGroup
		return ret
	}
	return *o.FormFieldGroups
}

// GetFormFieldGroupsOk returns a tuple with the FormFieldGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureRequestSendRequest) GetFormFieldGroupsOk() (*[]SubFormFieldGroup, bool) {
	if o == nil || o.FormFieldGroups == nil {
		return nil, false
	}
	return o.FormFieldGroups, true
}

// HasFormFieldGroups returns a boolean if a field has been set.
func (o *SignatureRequestSendRequest) HasFormFieldGroups() bool {
	if o != nil && o.FormFieldGroups != nil {
		return true
	}

	return false
}

// SetFormFieldGroups gets a reference to the given []SubFormFieldGroup and assigns it to the FormFieldGroups field.
func (o *SignatureRequestSendRequest) SetFormFieldGroups(v []SubFormFieldGroup) {
	o.FormFieldGroups = &v
}

// GetFormFieldRules returns the FormFieldRules field value if set, zero value otherwise.
func (o *SignatureRequestSendRequest) GetFormFieldRules() []SubFormFieldRule {
	if o == nil || o.FormFieldRules == nil {
		var ret []SubFormFieldRule
		return ret
	}
	return *o.FormFieldRules
}

// GetFormFieldRulesOk returns a tuple with the FormFieldRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureRequestSendRequest) GetFormFieldRulesOk() (*[]SubFormFieldRule, bool) {
	if o == nil || o.FormFieldRules == nil {
		return nil, false
	}
	return o.FormFieldRules, true
}

// HasFormFieldRules returns a boolean if a field has been set.
func (o *SignatureRequestSendRequest) HasFormFieldRules() bool {
	if o != nil && o.FormFieldRules != nil {
		return true
	}

	return false
}

// SetFormFieldRules gets a reference to the given []SubFormFieldRule and assigns it to the FormFieldRules field.
func (o *SignatureRequestSendRequest) SetFormFieldRules(v []SubFormFieldRule) {
	o.FormFieldRules = &v
}

// GetFormFieldsPerDocument returns the FormFieldsPerDocument field value if set, zero value otherwise.
func (o *SignatureRequestSendRequest) GetFormFieldsPerDocument() []SubFormFieldsPerDocumentBase {
	if o == nil || o.FormFieldsPerDocument == nil {
		var ret []SubFormFieldsPerDocumentBase
		return ret
	}
	return *o.FormFieldsPerDocument
}

// GetFormFieldsPerDocumentOk returns a tuple with the FormFieldsPerDocument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureRequestSendRequest) GetFormFieldsPerDocumentOk() (*[]SubFormFieldsPerDocumentBase, bool) {
	if o == nil || o.FormFieldsPerDocument == nil {
		return nil, false
	}
	return o.FormFieldsPerDocument, true
}

// HasFormFieldsPerDocument returns a boolean if a field has been set.
func (o *SignatureRequestSendRequest) HasFormFieldsPerDocument() bool {
	if o != nil && o.FormFieldsPerDocument != nil {
		return true
	}

	return false
}

// SetFormFieldsPerDocument gets a reference to the given []SubFormFieldsPerDocumentBase and assigns it to the FormFieldsPerDocument field.
func (o *SignatureRequestSendRequest) SetFormFieldsPerDocument(v []SubFormFieldsPerDocumentBase) {
	o.FormFieldsPerDocument = &v
}

// GetHideTextTags returns the HideTextTags field value if set, zero value otherwise.
func (o *SignatureRequestSendRequest) GetHideTextTags() bool {
	if o == nil || o.HideTextTags == nil {
		var ret bool
		return ret
	}
	return *o.HideTextTags
}

// GetHideTextTagsOk returns a tuple with the HideTextTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureRequestSendRequest) GetHideTextTagsOk() (*bool, bool) {
	if o == nil || o.HideTextTags == nil {
		return nil, false
	}
	return o.HideTextTags, true
}

// HasHideTextTags returns a boolean if a field has been set.
func (o *SignatureRequestSendRequest) HasHideTextTags() bool {
	if o != nil && o.HideTextTags != nil {
		return true
	}

	return false
}

// SetHideTextTags gets a reference to the given bool and assigns it to the HideTextTags field.
func (o *SignatureRequestSendRequest) SetHideTextTags(v bool) {
	o.HideTextTags = &v
}

// GetIsQualifiedSignature returns the IsQualifiedSignature field value if set, zero value otherwise.
func (o *SignatureRequestSendRequest) GetIsQualifiedSignature() bool {
	if o == nil || o.IsQualifiedSignature == nil {
		var ret bool
		return ret
	}
	return *o.IsQualifiedSignature
}

// GetIsQualifiedSignatureOk returns a tuple with the IsQualifiedSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureRequestSendRequest) GetIsQualifiedSignatureOk() (*bool, bool) {
	if o == nil || o.IsQualifiedSignature == nil {
		return nil, false
	}
	return o.IsQualifiedSignature, true
}

// HasIsQualifiedSignature returns a boolean if a field has been set.
func (o *SignatureRequestSendRequest) HasIsQualifiedSignature() bool {
	if o != nil && o.IsQualifiedSignature != nil {
		return true
	}

	return false
}

// SetIsQualifiedSignature gets a reference to the given bool and assigns it to the IsQualifiedSignature field.
func (o *SignatureRequestSendRequest) SetIsQualifiedSignature(v bool) {
	o.IsQualifiedSignature = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *SignatureRequestSendRequest) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureRequestSendRequest) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *SignatureRequestSendRequest) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *SignatureRequestSendRequest) SetMessage(v string) {
	o.Message = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *SignatureRequestSendRequest) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureRequestSendRequest) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *SignatureRequestSendRequest) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *SignatureRequestSendRequest) SetMetadata(v map[string]interface{}) {
	o.Metadata = &v
}

// GetSigningOptions returns the SigningOptions field value if set, zero value otherwise.
func (o *SignatureRequestSendRequest) GetSigningOptions() SubSigningOptions {
	if o == nil || o.SigningOptions == nil {
		var ret SubSigningOptions
		return ret
	}
	return *o.SigningOptions
}

// GetSigningOptionsOk returns a tuple with the SigningOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureRequestSendRequest) GetSigningOptionsOk() (*SubSigningOptions, bool) {
	if o == nil || o.SigningOptions == nil {
		return nil, false
	}
	return o.SigningOptions, true
}

// HasSigningOptions returns a boolean if a field has been set.
func (o *SignatureRequestSendRequest) HasSigningOptions() bool {
	if o != nil && o.SigningOptions != nil {
		return true
	}

	return false
}

// SetSigningOptions gets a reference to the given SubSigningOptions and assigns it to the SigningOptions field.
func (o *SignatureRequestSendRequest) SetSigningOptions(v SubSigningOptions) {
	o.SigningOptions = &v
}

// GetSigningRedirectUrl returns the SigningRedirectUrl field value if set, zero value otherwise.
func (o *SignatureRequestSendRequest) GetSigningRedirectUrl() string {
	if o == nil || o.SigningRedirectUrl == nil {
		var ret string
		return ret
	}
	return *o.SigningRedirectUrl
}

// GetSigningRedirectUrlOk returns a tuple with the SigningRedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureRequestSendRequest) GetSigningRedirectUrlOk() (*string, bool) {
	if o == nil || o.SigningRedirectUrl == nil {
		return nil, false
	}
	return o.SigningRedirectUrl, true
}

// HasSigningRedirectUrl returns a boolean if a field has been set.
func (o *SignatureRequestSendRequest) HasSigningRedirectUrl() bool {
	if o != nil && o.SigningRedirectUrl != nil {
		return true
	}

	return false
}

// SetSigningRedirectUrl gets a reference to the given string and assigns it to the SigningRedirectUrl field.
func (o *SignatureRequestSendRequest) SetSigningRedirectUrl(v string) {
	o.SigningRedirectUrl = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *SignatureRequestSendRequest) GetSubject() string {
	if o == nil || o.Subject == nil {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureRequestSendRequest) GetSubjectOk() (*string, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *SignatureRequestSendRequest) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *SignatureRequestSendRequest) SetSubject(v string) {
	o.Subject = &v
}

// GetTestMode returns the TestMode field value if set, zero value otherwise.
func (o *SignatureRequestSendRequest) GetTestMode() bool {
	if o == nil || o.TestMode == nil {
		var ret bool
		return ret
	}
	return *o.TestMode
}

// GetTestModeOk returns a tuple with the TestMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureRequestSendRequest) GetTestModeOk() (*bool, bool) {
	if o == nil || o.TestMode == nil {
		return nil, false
	}
	return o.TestMode, true
}

// HasTestMode returns a boolean if a field has been set.
func (o *SignatureRequestSendRequest) HasTestMode() bool {
	if o != nil && o.TestMode != nil {
		return true
	}

	return false
}

// SetTestMode gets a reference to the given bool and assigns it to the TestMode field.
func (o *SignatureRequestSendRequest) SetTestMode(v bool) {
	o.TestMode = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *SignatureRequestSendRequest) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureRequestSendRequest) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *SignatureRequestSendRequest) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *SignatureRequestSendRequest) SetTitle(v string) {
	o.Title = &v
}

// GetUseTextTags returns the UseTextTags field value if set, zero value otherwise.
func (o *SignatureRequestSendRequest) GetUseTextTags() bool {
	if o == nil || o.UseTextTags == nil {
		var ret bool
		return ret
	}
	return *o.UseTextTags
}

// GetUseTextTagsOk returns a tuple with the UseTextTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureRequestSendRequest) GetUseTextTagsOk() (*bool, bool) {
	if o == nil || o.UseTextTags == nil {
		return nil, false
	}
	return o.UseTextTags, true
}

// HasUseTextTags returns a boolean if a field has been set.
func (o *SignatureRequestSendRequest) HasUseTextTags() bool {
	if o != nil && o.UseTextTags != nil {
		return true
	}

	return false
}

// SetUseTextTags gets a reference to the given bool and assigns it to the UseTextTags field.
func (o *SignatureRequestSendRequest) SetUseTextTags(v bool) {
	o.UseTextTags = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SignatureRequestSendRequest) GetExpiresAt() int32 {
	if o == nil || o.ExpiresAt.Get() == nil {
		var ret int32
		return ret
	}
	return *o.ExpiresAt.Get()
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SignatureRequestSendRequest) GetExpiresAtOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExpiresAt.Get(), o.ExpiresAt.IsSet()
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *SignatureRequestSendRequest) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt.IsSet() {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given NullableInt32 and assigns it to the ExpiresAt field.
func (o *SignatureRequestSendRequest) SetExpiresAt(v int32) {
	o.ExpiresAt.Set(&v)
}
// SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil
func (o *SignatureRequestSendRequest) SetExpiresAtNil() {
	o.ExpiresAt.Set(nil)
}

// UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
func (o *SignatureRequestSendRequest) UnsetExpiresAt() {
	o.ExpiresAt.Unset()
}

func (o SignatureRequestSendRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Files != nil {
		toSerialize["files"] = o.Files
	}
	if o.FileUrls != nil {
		toSerialize["file_urls"] = o.FileUrls
	}
	if o.Signers != nil {
		toSerialize["signers"] = o.Signers
	}
	if o.GroupedSigners != nil {
		toSerialize["grouped_signers"] = o.GroupedSigners
	}
	if o.AllowDecline != nil {
		toSerialize["allow_decline"] = o.AllowDecline
	}
	if o.AllowReassign != nil {
		toSerialize["allow_reassign"] = o.AllowReassign
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	if o.CcEmailAddresses != nil {
		toSerialize["cc_email_addresses"] = o.CcEmailAddresses
	}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.CustomFields != nil {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if o.FieldOptions != nil {
		toSerialize["field_options"] = o.FieldOptions
	}
	if o.FormFieldGroups != nil {
		toSerialize["form_field_groups"] = o.FormFieldGroups
	}
	if o.FormFieldRules != nil {
		toSerialize["form_field_rules"] = o.FormFieldRules
	}
	if o.FormFieldsPerDocument != nil {
		toSerialize["form_fields_per_document"] = o.FormFieldsPerDocument
	}
	if o.HideTextTags != nil {
		toSerialize["hide_text_tags"] = o.HideTextTags
	}
	if o.IsQualifiedSignature != nil {
		toSerialize["is_qualified_signature"] = o.IsQualifiedSignature
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.SigningOptions != nil {
		toSerialize["signing_options"] = o.SigningOptions
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
	if o.UseTextTags != nil {
		toSerialize["use_text_tags"] = o.UseTextTags
	}
	if o.ExpiresAt.IsSet() {
		toSerialize["expires_at"] = o.ExpiresAt.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSignatureRequestSendRequest struct {
	value *SignatureRequestSendRequest
	isSet bool
}

func (v NullableSignatureRequestSendRequest) Get() *SignatureRequestSendRequest {
	return v.value
}

func (v *NullableSignatureRequestSendRequest) Set(val *SignatureRequestSendRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSignatureRequestSendRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSignatureRequestSendRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignatureRequestSendRequest(val *SignatureRequestSendRequest) *NullableSignatureRequestSendRequest {
	return &NullableSignatureRequestSendRequest{value: val, isSet: true}
}

func (v NullableSignatureRequestSendRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignatureRequestSendRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

