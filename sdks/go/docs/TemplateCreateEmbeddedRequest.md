# TemplateCreateEmbeddedRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | Client id of the app you&#39;re using to create this draft. Used to apply the branding and callback url defined for the app. | 
**Files** | Pointer to **[]*os.File** | Use &#x60;files[]&#x60; to indicate the uploaded file(s) to send for signature.  This endpoint requires either **files** or **file_urls[]**, but not both. | [optional] 
**FileUrls** | Pointer to **[]string** | Use &#x60;file_urls[]&#x60; to have Dropbox Sign download the file(s) to send for signature.  This endpoint requires either **files** or **file_urls[]**, but not both. | [optional] 
**AllowCcs** | Pointer to **bool** | This allows the requester to specify whether the user is allowed to provide email addresses to CC when creating a template. | [optional] [default to true]
**AllowReassign** | Pointer to **bool** | Allows signers to reassign their signature requests to other signers if set to &#x60;true&#x60;. Defaults to &#x60;false&#x60;.  **Note**: Only available for Premium plan and higher. | [optional] [default to false]
**Attachments** | Pointer to [**[]SubAttachment**](SubAttachment.md) | A list describing the attachments | [optional] 
**CcRoles** | Pointer to **[]string** | The CC roles that must be assigned when using the template to send a signature request | [optional] 
**EditorOptions** | Pointer to [**SubEditorOptions**](SubEditorOptions.md) |  | [optional] 
**FieldOptions** | Pointer to [**SubFieldOptions**](SubFieldOptions.md) |  | [optional] 
**ForceSignerRoles** | Pointer to **bool** | Provide users the ability to review/edit the template signer roles. | [optional] [default to false]
**ForceSubjectMessage** | Pointer to **bool** | Provide users the ability to review/edit the template subject and message. | [optional] [default to false]
**FormFieldGroups** | Pointer to [**[]SubFormFieldGroup**](SubFormFieldGroup.md) | Group information for fields defined in &#x60;form_fields_per_document&#x60;. String-indexed JSON array with &#x60;group_label&#x60; and &#x60;requirement&#x60; keys. &#x60;form_fields_per_document&#x60; must contain fields referencing a group defined in &#x60;form_field_groups&#x60;. | [optional] 
**FormFieldRules** | Pointer to [**[]SubFormFieldRule**](SubFormFieldRule.md) | Conditional Logic rules for fields defined in &#x60;form_fields_per_document&#x60;. | [optional] 
**FormFieldsPerDocument** | Pointer to [**[]SubFormFieldsPerDocumentBase**](SubFormFieldsPerDocumentBase.md) | The fields that should appear on the document, expressed as an array of objects. (For more details you can read about it here: [Using Form Fields per Document](/docs/openapi/form-fields-per-document).)  **NOTE**: Fields like **text**, **dropdown**, **checkbox**, **radio**, and **hyperlink** have additional required and optional parameters. Check out the list of [additional parameters](/api/reference/constants/#form-fields-per-document) for these field types.  * Text Field use &#x60;SubFormFieldsPerDocumentText&#x60; * Dropdown Field use &#x60;SubFormFieldsPerDocumentDropdown&#x60; * Hyperlink Field use &#x60;SubFormFieldsPerDocumentHyperlink&#x60; * Checkbox Field use &#x60;SubFormFieldsPerDocumentCheckbox&#x60; * Radio Field use &#x60;SubFormFieldsPerDocumentRadio&#x60; * Signature Field use &#x60;SubFormFieldsPerDocumentSignature&#x60; * Date Signed Field use &#x60;SubFormFieldsPerDocumentDateSigned&#x60; * Initials Field use &#x60;SubFormFieldsPerDocumentInitials&#x60; * Text Merge Field use &#x60;SubFormFieldsPerDocumentTextMerge&#x60; * Checkbox Merge Field use &#x60;SubFormFieldsPerDocumentCheckboxMerge&#x60; | [optional] 
**MergeFields** | Pointer to [**[]SubMergeField**](SubMergeField.md) | Add merge fields to the template. Merge fields are placed by the user creating the template and used to pre-fill data by passing values into signature requests with the &#x60;custom_fields&#x60; parameter. If the signature request using that template *does not* pass a value into a merge field, then an empty field remains in the document. | [optional] 
**Message** | Pointer to **string** | The default template email message. | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Key-value data that should be attached to the signature request. This metadata is included in all API responses and events involving the signature request. For example, use the metadata field to store a signer&#39;s order number for look up when receiving events for the signature request.  Each request can include up to 10 metadata keys (or 50 nested metadata keys), with key names up to 40 characters long and values up to 1000 characters long. | [optional] 
**ShowPreview** | Pointer to **bool** | This allows the requester to enable the editor/preview experience.  - &#x60;show_preview&#x3D;true&#x60;: Allows requesters to enable the editor/preview experience. - &#x60;show_preview&#x3D;false&#x60;: Allows requesters to disable the editor/preview experience. | [optional] [default to false]
**ShowProgressStepper** | Pointer to **bool** | When only one step remains in the signature request process and this parameter is set to &#x60;false&#x60; then the progress stepper will be hidden. | [optional] [default to true]
**SignerRoles** | Pointer to [**[]SubTemplateRole**](SubTemplateRole.md) | An array of the designated signer roles that must be specified when sending a SignatureRequest using this Template. | [optional] 
**SkipMeNow** | Pointer to **bool** | Disables the \&quot;Me (Now)\&quot; option for the person preparing the document. Does not work with type &#x60;send_document&#x60;. Defaults to &#x60;false&#x60;. | [optional] [default to false]
**Subject** | Pointer to **string** | The template title (alias). | [optional] 
**TestMode** | Pointer to **bool** | Whether this is a test, the signature request created from this draft will not be legally binding if set to &#x60;true&#x60;. Defaults to &#x60;false&#x60;. | [optional] [default to false]
**Title** | Pointer to **string** | The title you want to assign to the SignatureRequest. | [optional] 
**UsePreexistingFields** | Pointer to **bool** | Enable the detection of predefined PDF fields by setting the &#x60;use_preexisting_fields&#x60; to &#x60;true&#x60; (defaults to disabled, or &#x60;false&#x60;). | [optional] [default to false]

## Methods

### NewTemplateCreateEmbeddedRequest

`func NewTemplateCreateEmbeddedRequest(clientId string, ) *TemplateCreateEmbeddedRequest`

NewTemplateCreateEmbeddedRequest instantiates a new TemplateCreateEmbeddedRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateCreateEmbeddedRequestWithDefaults

`func NewTemplateCreateEmbeddedRequestWithDefaults() *TemplateCreateEmbeddedRequest`

NewTemplateCreateEmbeddedRequestWithDefaults instantiates a new TemplateCreateEmbeddedRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *TemplateCreateEmbeddedRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *TemplateCreateEmbeddedRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *TemplateCreateEmbeddedRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetFiles

`func (o *TemplateCreateEmbeddedRequest) GetFiles() []*os.File`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *TemplateCreateEmbeddedRequest) GetFilesOk() (*[]*os.File, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *TemplateCreateEmbeddedRequest) SetFiles(v []*os.File)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *TemplateCreateEmbeddedRequest) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetFileUrls

`func (o *TemplateCreateEmbeddedRequest) GetFileUrls() []string`

GetFileUrls returns the FileUrls field if non-nil, zero value otherwise.

### GetFileUrlsOk

`func (o *TemplateCreateEmbeddedRequest) GetFileUrlsOk() (*[]string, bool)`

GetFileUrlsOk returns a tuple with the FileUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrls

`func (o *TemplateCreateEmbeddedRequest) SetFileUrls(v []string)`

SetFileUrls sets FileUrls field to given value.

### HasFileUrls

`func (o *TemplateCreateEmbeddedRequest) HasFileUrls() bool`

HasFileUrls returns a boolean if a field has been set.

### GetAllowCcs

`func (o *TemplateCreateEmbeddedRequest) GetAllowCcs() bool`

GetAllowCcs returns the AllowCcs field if non-nil, zero value otherwise.

### GetAllowCcsOk

`func (o *TemplateCreateEmbeddedRequest) GetAllowCcsOk() (*bool, bool)`

GetAllowCcsOk returns a tuple with the AllowCcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCcs

`func (o *TemplateCreateEmbeddedRequest) SetAllowCcs(v bool)`

SetAllowCcs sets AllowCcs field to given value.

### HasAllowCcs

`func (o *TemplateCreateEmbeddedRequest) HasAllowCcs() bool`

HasAllowCcs returns a boolean if a field has been set.

### GetAllowReassign

`func (o *TemplateCreateEmbeddedRequest) GetAllowReassign() bool`

GetAllowReassign returns the AllowReassign field if non-nil, zero value otherwise.

### GetAllowReassignOk

`func (o *TemplateCreateEmbeddedRequest) GetAllowReassignOk() (*bool, bool)`

GetAllowReassignOk returns a tuple with the AllowReassign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowReassign

`func (o *TemplateCreateEmbeddedRequest) SetAllowReassign(v bool)`

SetAllowReassign sets AllowReassign field to given value.

### HasAllowReassign

`func (o *TemplateCreateEmbeddedRequest) HasAllowReassign() bool`

HasAllowReassign returns a boolean if a field has been set.

### GetAttachments

`func (o *TemplateCreateEmbeddedRequest) GetAttachments() []SubAttachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *TemplateCreateEmbeddedRequest) GetAttachmentsOk() (*[]SubAttachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *TemplateCreateEmbeddedRequest) SetAttachments(v []SubAttachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *TemplateCreateEmbeddedRequest) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetCcRoles

`func (o *TemplateCreateEmbeddedRequest) GetCcRoles() []string`

GetCcRoles returns the CcRoles field if non-nil, zero value otherwise.

### GetCcRolesOk

`func (o *TemplateCreateEmbeddedRequest) GetCcRolesOk() (*[]string, bool)`

GetCcRolesOk returns a tuple with the CcRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcRoles

`func (o *TemplateCreateEmbeddedRequest) SetCcRoles(v []string)`

SetCcRoles sets CcRoles field to given value.

### HasCcRoles

`func (o *TemplateCreateEmbeddedRequest) HasCcRoles() bool`

HasCcRoles returns a boolean if a field has been set.

### GetEditorOptions

`func (o *TemplateCreateEmbeddedRequest) GetEditorOptions() SubEditorOptions`

GetEditorOptions returns the EditorOptions field if non-nil, zero value otherwise.

### GetEditorOptionsOk

`func (o *TemplateCreateEmbeddedRequest) GetEditorOptionsOk() (*SubEditorOptions, bool)`

GetEditorOptionsOk returns a tuple with the EditorOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditorOptions

`func (o *TemplateCreateEmbeddedRequest) SetEditorOptions(v SubEditorOptions)`

SetEditorOptions sets EditorOptions field to given value.

### HasEditorOptions

`func (o *TemplateCreateEmbeddedRequest) HasEditorOptions() bool`

HasEditorOptions returns a boolean if a field has been set.

### GetFieldOptions

`func (o *TemplateCreateEmbeddedRequest) GetFieldOptions() SubFieldOptions`

GetFieldOptions returns the FieldOptions field if non-nil, zero value otherwise.

### GetFieldOptionsOk

`func (o *TemplateCreateEmbeddedRequest) GetFieldOptionsOk() (*SubFieldOptions, bool)`

GetFieldOptionsOk returns a tuple with the FieldOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldOptions

`func (o *TemplateCreateEmbeddedRequest) SetFieldOptions(v SubFieldOptions)`

SetFieldOptions sets FieldOptions field to given value.

### HasFieldOptions

`func (o *TemplateCreateEmbeddedRequest) HasFieldOptions() bool`

HasFieldOptions returns a boolean if a field has been set.

### GetForceSignerRoles

`func (o *TemplateCreateEmbeddedRequest) GetForceSignerRoles() bool`

GetForceSignerRoles returns the ForceSignerRoles field if non-nil, zero value otherwise.

### GetForceSignerRolesOk

`func (o *TemplateCreateEmbeddedRequest) GetForceSignerRolesOk() (*bool, bool)`

GetForceSignerRolesOk returns a tuple with the ForceSignerRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSignerRoles

`func (o *TemplateCreateEmbeddedRequest) SetForceSignerRoles(v bool)`

SetForceSignerRoles sets ForceSignerRoles field to given value.

### HasForceSignerRoles

`func (o *TemplateCreateEmbeddedRequest) HasForceSignerRoles() bool`

HasForceSignerRoles returns a boolean if a field has been set.

### GetForceSubjectMessage

`func (o *TemplateCreateEmbeddedRequest) GetForceSubjectMessage() bool`

GetForceSubjectMessage returns the ForceSubjectMessage field if non-nil, zero value otherwise.

### GetForceSubjectMessageOk

`func (o *TemplateCreateEmbeddedRequest) GetForceSubjectMessageOk() (*bool, bool)`

GetForceSubjectMessageOk returns a tuple with the ForceSubjectMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSubjectMessage

`func (o *TemplateCreateEmbeddedRequest) SetForceSubjectMessage(v bool)`

SetForceSubjectMessage sets ForceSubjectMessage field to given value.

### HasForceSubjectMessage

`func (o *TemplateCreateEmbeddedRequest) HasForceSubjectMessage() bool`

HasForceSubjectMessage returns a boolean if a field has been set.

### GetFormFieldGroups

`func (o *TemplateCreateEmbeddedRequest) GetFormFieldGroups() []SubFormFieldGroup`

GetFormFieldGroups returns the FormFieldGroups field if non-nil, zero value otherwise.

### GetFormFieldGroupsOk

`func (o *TemplateCreateEmbeddedRequest) GetFormFieldGroupsOk() (*[]SubFormFieldGroup, bool)`

GetFormFieldGroupsOk returns a tuple with the FormFieldGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormFieldGroups

`func (o *TemplateCreateEmbeddedRequest) SetFormFieldGroups(v []SubFormFieldGroup)`

SetFormFieldGroups sets FormFieldGroups field to given value.

### HasFormFieldGroups

`func (o *TemplateCreateEmbeddedRequest) HasFormFieldGroups() bool`

HasFormFieldGroups returns a boolean if a field has been set.

### GetFormFieldRules

`func (o *TemplateCreateEmbeddedRequest) GetFormFieldRules() []SubFormFieldRule`

GetFormFieldRules returns the FormFieldRules field if non-nil, zero value otherwise.

### GetFormFieldRulesOk

`func (o *TemplateCreateEmbeddedRequest) GetFormFieldRulesOk() (*[]SubFormFieldRule, bool)`

GetFormFieldRulesOk returns a tuple with the FormFieldRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormFieldRules

`func (o *TemplateCreateEmbeddedRequest) SetFormFieldRules(v []SubFormFieldRule)`

SetFormFieldRules sets FormFieldRules field to given value.

### HasFormFieldRules

`func (o *TemplateCreateEmbeddedRequest) HasFormFieldRules() bool`

HasFormFieldRules returns a boolean if a field has been set.

### GetFormFieldsPerDocument

`func (o *TemplateCreateEmbeddedRequest) GetFormFieldsPerDocument() []SubFormFieldsPerDocumentBase`

GetFormFieldsPerDocument returns the FormFieldsPerDocument field if non-nil, zero value otherwise.

### GetFormFieldsPerDocumentOk

`func (o *TemplateCreateEmbeddedRequest) GetFormFieldsPerDocumentOk() (*[]SubFormFieldsPerDocumentBase, bool)`

GetFormFieldsPerDocumentOk returns a tuple with the FormFieldsPerDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormFieldsPerDocument

`func (o *TemplateCreateEmbeddedRequest) SetFormFieldsPerDocument(v []SubFormFieldsPerDocumentBase)`

SetFormFieldsPerDocument sets FormFieldsPerDocument field to given value.

### HasFormFieldsPerDocument

`func (o *TemplateCreateEmbeddedRequest) HasFormFieldsPerDocument() bool`

HasFormFieldsPerDocument returns a boolean if a field has been set.

### GetMergeFields

`func (o *TemplateCreateEmbeddedRequest) GetMergeFields() []SubMergeField`

GetMergeFields returns the MergeFields field if non-nil, zero value otherwise.

### GetMergeFieldsOk

`func (o *TemplateCreateEmbeddedRequest) GetMergeFieldsOk() (*[]SubMergeField, bool)`

GetMergeFieldsOk returns a tuple with the MergeFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeFields

`func (o *TemplateCreateEmbeddedRequest) SetMergeFields(v []SubMergeField)`

SetMergeFields sets MergeFields field to given value.

### HasMergeFields

`func (o *TemplateCreateEmbeddedRequest) HasMergeFields() bool`

HasMergeFields returns a boolean if a field has been set.

### GetMessage

`func (o *TemplateCreateEmbeddedRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TemplateCreateEmbeddedRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TemplateCreateEmbeddedRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *TemplateCreateEmbeddedRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMetadata

`func (o *TemplateCreateEmbeddedRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TemplateCreateEmbeddedRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TemplateCreateEmbeddedRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *TemplateCreateEmbeddedRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetShowPreview

`func (o *TemplateCreateEmbeddedRequest) GetShowPreview() bool`

GetShowPreview returns the ShowPreview field if non-nil, zero value otherwise.

### GetShowPreviewOk

`func (o *TemplateCreateEmbeddedRequest) GetShowPreviewOk() (*bool, bool)`

GetShowPreviewOk returns a tuple with the ShowPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowPreview

`func (o *TemplateCreateEmbeddedRequest) SetShowPreview(v bool)`

SetShowPreview sets ShowPreview field to given value.

### HasShowPreview

`func (o *TemplateCreateEmbeddedRequest) HasShowPreview() bool`

HasShowPreview returns a boolean if a field has been set.

### GetShowProgressStepper

`func (o *TemplateCreateEmbeddedRequest) GetShowProgressStepper() bool`

GetShowProgressStepper returns the ShowProgressStepper field if non-nil, zero value otherwise.

### GetShowProgressStepperOk

`func (o *TemplateCreateEmbeddedRequest) GetShowProgressStepperOk() (*bool, bool)`

GetShowProgressStepperOk returns a tuple with the ShowProgressStepper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowProgressStepper

`func (o *TemplateCreateEmbeddedRequest) SetShowProgressStepper(v bool)`

SetShowProgressStepper sets ShowProgressStepper field to given value.

### HasShowProgressStepper

`func (o *TemplateCreateEmbeddedRequest) HasShowProgressStepper() bool`

HasShowProgressStepper returns a boolean if a field has been set.

### GetSignerRoles

`func (o *TemplateCreateEmbeddedRequest) GetSignerRoles() []SubTemplateRole`

GetSignerRoles returns the SignerRoles field if non-nil, zero value otherwise.

### GetSignerRolesOk

`func (o *TemplateCreateEmbeddedRequest) GetSignerRolesOk() (*[]SubTemplateRole, bool)`

GetSignerRolesOk returns a tuple with the SignerRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerRoles

`func (o *TemplateCreateEmbeddedRequest) SetSignerRoles(v []SubTemplateRole)`

SetSignerRoles sets SignerRoles field to given value.

### HasSignerRoles

`func (o *TemplateCreateEmbeddedRequest) HasSignerRoles() bool`

HasSignerRoles returns a boolean if a field has been set.

### GetSkipMeNow

`func (o *TemplateCreateEmbeddedRequest) GetSkipMeNow() bool`

GetSkipMeNow returns the SkipMeNow field if non-nil, zero value otherwise.

### GetSkipMeNowOk

`func (o *TemplateCreateEmbeddedRequest) GetSkipMeNowOk() (*bool, bool)`

GetSkipMeNowOk returns a tuple with the SkipMeNow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipMeNow

`func (o *TemplateCreateEmbeddedRequest) SetSkipMeNow(v bool)`

SetSkipMeNow sets SkipMeNow field to given value.

### HasSkipMeNow

`func (o *TemplateCreateEmbeddedRequest) HasSkipMeNow() bool`

HasSkipMeNow returns a boolean if a field has been set.

### GetSubject

`func (o *TemplateCreateEmbeddedRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *TemplateCreateEmbeddedRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *TemplateCreateEmbeddedRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *TemplateCreateEmbeddedRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTestMode

`func (o *TemplateCreateEmbeddedRequest) GetTestMode() bool`

GetTestMode returns the TestMode field if non-nil, zero value otherwise.

### GetTestModeOk

`func (o *TemplateCreateEmbeddedRequest) GetTestModeOk() (*bool, bool)`

GetTestModeOk returns a tuple with the TestMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestMode

`func (o *TemplateCreateEmbeddedRequest) SetTestMode(v bool)`

SetTestMode sets TestMode field to given value.

### HasTestMode

`func (o *TemplateCreateEmbeddedRequest) HasTestMode() bool`

HasTestMode returns a boolean if a field has been set.

### GetTitle

`func (o *TemplateCreateEmbeddedRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TemplateCreateEmbeddedRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TemplateCreateEmbeddedRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TemplateCreateEmbeddedRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUsePreexistingFields

`func (o *TemplateCreateEmbeddedRequest) GetUsePreexistingFields() bool`

GetUsePreexistingFields returns the UsePreexistingFields field if non-nil, zero value otherwise.

### GetUsePreexistingFieldsOk

`func (o *TemplateCreateEmbeddedRequest) GetUsePreexistingFieldsOk() (*bool, bool)`

GetUsePreexistingFieldsOk returns a tuple with the UsePreexistingFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePreexistingFields

`func (o *TemplateCreateEmbeddedRequest) SetUsePreexistingFields(v bool)`

SetUsePreexistingFields sets UsePreexistingFields field to given value.

### HasUsePreexistingFields

`func (o *TemplateCreateEmbeddedRequest) HasUsePreexistingFields() bool`

HasUsePreexistingFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


