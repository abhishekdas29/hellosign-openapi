# SignatureRequestSendRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Files** | Pointer to **[]*os.File** | Use &#x60;files[]&#x60; to indicate the uploaded file(s) to send for signature.  This endpoint requires either **files** or **file_urls[]**, but not both. | [optional] 
**FileUrls** | Pointer to **[]string** | Use &#x60;file_urls[]&#x60; to have Dropbox Sign download the file(s) to send for signature.  This endpoint requires either **files** or **file_urls[]**, but not both. | [optional] 
**Signers** | Pointer to [**[]SubSignatureRequestSigner**](SubSignatureRequestSigner.md) | Add Signers to your Signature Request.  This endpoint requires either **signers** or **grouped_signers**, but not both. | [optional] 
**GroupedSigners** | Pointer to [**[]SubSignatureRequestGroupedSigners**](SubSignatureRequestGroupedSigners.md) | Add Grouped Signers to your Signature Request.  This endpoint requires either **signers** or **grouped_signers**, but not both. | [optional] 
**AllowDecline** | Pointer to **bool** | Allows signers to decline to sign a document if &#x60;true&#x60;. Defaults to &#x60;false&#x60;. | [optional] [default to false]
**AllowReassign** | Pointer to **bool** | Allows signers to reassign their signature requests to other signers if set to &#x60;true&#x60;. Defaults to &#x60;false&#x60;.  **Note**: Only available for Premium plan and higher. | [optional] [default to false]
**Attachments** | Pointer to [**[]SubAttachment**](SubAttachment.md) | A list describing the attachments | [optional] 
**CcEmailAddresses** | Pointer to **[]string** | The email addresses that should be CCed. | [optional] 
**ClientId** | Pointer to **string** | The client id of the API App you want to associate with this request. Used to apply the branding and callback url defined for the app. | [optional] 
**CustomFields** | Pointer to [**[]SubCustomField**](SubCustomField.md) | When used together with merge fields, &#x60;custom_fields&#x60; allows users to add pre-filled data to their signature requests.  Pre-filled data can be used with \&quot;send-once\&quot; signature requests by adding merge fields with &#x60;form_fields_per_document&#x60; or [Text Tags](https://app.hellosign.com/api/textTagsWalkthrough#TextTagIntro) while passing values back with &#x60;custom_fields&#x60; together in one API call.  For using pre-filled on repeatable signature requests, merge fields are added to templates in the Dropbox Sign UI or by calling [/template/create_embedded_draft](/api/reference/operation/templateCreateEmbeddedDraft) and then passing &#x60;custom_fields&#x60; on subsequent signature requests referencing that template. | [optional] 
**FieldOptions** | Pointer to [**SubFieldOptions**](SubFieldOptions.md) |  | [optional] 
**FormFieldGroups** | Pointer to [**[]SubFormFieldGroup**](SubFormFieldGroup.md) | Group information for fields defined in &#x60;form_fields_per_document&#x60;. String-indexed JSON array with &#x60;group_label&#x60; and &#x60;requirement&#x60; keys. &#x60;form_fields_per_document&#x60; must contain fields referencing a group defined in &#x60;form_field_groups&#x60;. | [optional] 
**FormFieldRules** | Pointer to [**[]SubFormFieldRule**](SubFormFieldRule.md) | Conditional Logic rules for fields defined in &#x60;form_fields_per_document&#x60;. | [optional] 
**FormFieldsPerDocument** | Pointer to [**[]SubFormFieldsPerDocumentBase**](SubFormFieldsPerDocumentBase.md) | The fields that should appear on the document, expressed as an array of objects. (For more details you can read about it here: [Using Form Fields per Document](/docs/openapi/form-fields-per-document).)  **NOTE**: Fields like **text**, **dropdown**, **checkbox**, **radio**, and **hyperlink** have additional required and optional parameters. Check out the list of [additional parameters](/api/reference/constants/#form-fields-per-document) for these field types.  * Text Field use &#x60;SubFormFieldsPerDocumentText&#x60; * Dropdown Field use &#x60;SubFormFieldsPerDocumentDropdown&#x60; * Hyperlink Field use &#x60;SubFormFieldsPerDocumentHyperlink&#x60; * Checkbox Field use &#x60;SubFormFieldsPerDocumentCheckbox&#x60; * Radio Field use &#x60;SubFormFieldsPerDocumentRadio&#x60; * Signature Field use &#x60;SubFormFieldsPerDocumentSignature&#x60; * Date Signed Field use &#x60;SubFormFieldsPerDocumentDateSigned&#x60; * Initials Field use &#x60;SubFormFieldsPerDocumentInitials&#x60; * Text Merge Field use &#x60;SubFormFieldsPerDocumentTextMerge&#x60; * Checkbox Merge Field use &#x60;SubFormFieldsPerDocumentCheckboxMerge&#x60; | [optional] 
**HideTextTags** | Pointer to **bool** | Enables automatic Text Tag removal when set to true.  **NOTE**: Removing text tags this way can cause unwanted clipping. We recommend leaving this setting on &#x60;false&#x60; and instead hiding your text tags using white text or a similar approach. See the [Text Tags Walkthrough](https://app.hellosign.com/api/textTagsWalkthrough#TextTagIntro) for more information. | [optional] [default to false]
**IsQualifiedSignature** | Pointer to **bool** | Send with a value of &#x60;true&#x60; if you wish to enable [Qualified Electronic Signatures](https://www.hellosign.com/features/qualified-electronic-signatures) (QES), which requires a face-to-face call to verify the signer&#39;s identity.&lt;br&gt; **Note**: QES is only available on the Premium API plan as an add-on purchase. Cannot be used in &#x60;test_mode&#x60;. Only works on requests with one signer. | [optional] [default to false]
**Message** | Pointer to **string** | The custom message in the email that will be sent to the signers. | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Key-value data that should be attached to the signature request. This metadata is included in all API responses and events involving the signature request. For example, use the metadata field to store a signer&#39;s order number for look up when receiving events for the signature request.  Each request can include up to 10 metadata keys (or 50 nested metadata keys), with key names up to 40 characters long and values up to 1000 characters long. | [optional] 
**SigningOptions** | Pointer to [**SubSigningOptions**](SubSigningOptions.md) |  | [optional] 
**SigningRedirectUrl** | Pointer to **string** | The URL you want signers redirected to after they successfully sign. | [optional] 
**Subject** | Pointer to **string** | The subject in the email that will be sent to the signers. | [optional] 
**TestMode** | Pointer to **bool** | Whether this is a test, the signature request will not be legally binding if set to &#x60;true&#x60;. Defaults to &#x60;false&#x60;. | [optional] [default to false]
**Title** | Pointer to **string** | The title you want to assign to the SignatureRequest. | [optional] 
**UseTextTags** | Pointer to **bool** | Send with a value of &#x60;true&#x60; if you wish to enable [Text Tags](https://app.hellosign.com/api/textTagsWalkthrough#TextTagIntro) parsing in your document. Defaults to disabled, or &#x60;false&#x60;. | [optional] [default to false]
**ExpiresAt** | Pointer to **NullableInt32** | When the signature request will expire. Unsigned signatures will be moved to the expired status, and no longer signable. See [Signature Request Expiration Date](https://developers.hellosign.com/docs/signature-request/expiration/) for details. | [optional] 

## Methods

### NewSignatureRequestSendRequest

`func NewSignatureRequestSendRequest() *SignatureRequestSendRequest`

NewSignatureRequestSendRequest instantiates a new SignatureRequestSendRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignatureRequestSendRequestWithDefaults

`func NewSignatureRequestSendRequestWithDefaults() *SignatureRequestSendRequest`

NewSignatureRequestSendRequestWithDefaults instantiates a new SignatureRequestSendRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiles

`func (o *SignatureRequestSendRequest) GetFiles() []*os.File`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *SignatureRequestSendRequest) GetFilesOk() (*[]*os.File, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *SignatureRequestSendRequest) SetFiles(v []*os.File)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *SignatureRequestSendRequest) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetFileUrls

`func (o *SignatureRequestSendRequest) GetFileUrls() []string`

GetFileUrls returns the FileUrls field if non-nil, zero value otherwise.

### GetFileUrlsOk

`func (o *SignatureRequestSendRequest) GetFileUrlsOk() (*[]string, bool)`

GetFileUrlsOk returns a tuple with the FileUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrls

`func (o *SignatureRequestSendRequest) SetFileUrls(v []string)`

SetFileUrls sets FileUrls field to given value.

### HasFileUrls

`func (o *SignatureRequestSendRequest) HasFileUrls() bool`

HasFileUrls returns a boolean if a field has been set.

### GetSigners

`func (o *SignatureRequestSendRequest) GetSigners() []SubSignatureRequestSigner`

GetSigners returns the Signers field if non-nil, zero value otherwise.

### GetSignersOk

`func (o *SignatureRequestSendRequest) GetSignersOk() (*[]SubSignatureRequestSigner, bool)`

GetSignersOk returns a tuple with the Signers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigners

`func (o *SignatureRequestSendRequest) SetSigners(v []SubSignatureRequestSigner)`

SetSigners sets Signers field to given value.

### HasSigners

`func (o *SignatureRequestSendRequest) HasSigners() bool`

HasSigners returns a boolean if a field has been set.

### GetGroupedSigners

`func (o *SignatureRequestSendRequest) GetGroupedSigners() []SubSignatureRequestGroupedSigners`

GetGroupedSigners returns the GroupedSigners field if non-nil, zero value otherwise.

### GetGroupedSignersOk

`func (o *SignatureRequestSendRequest) GetGroupedSignersOk() (*[]SubSignatureRequestGroupedSigners, bool)`

GetGroupedSignersOk returns a tuple with the GroupedSigners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupedSigners

`func (o *SignatureRequestSendRequest) SetGroupedSigners(v []SubSignatureRequestGroupedSigners)`

SetGroupedSigners sets GroupedSigners field to given value.

### HasGroupedSigners

`func (o *SignatureRequestSendRequest) HasGroupedSigners() bool`

HasGroupedSigners returns a boolean if a field has been set.

### GetAllowDecline

`func (o *SignatureRequestSendRequest) GetAllowDecline() bool`

GetAllowDecline returns the AllowDecline field if non-nil, zero value otherwise.

### GetAllowDeclineOk

`func (o *SignatureRequestSendRequest) GetAllowDeclineOk() (*bool, bool)`

GetAllowDeclineOk returns a tuple with the AllowDecline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDecline

`func (o *SignatureRequestSendRequest) SetAllowDecline(v bool)`

SetAllowDecline sets AllowDecline field to given value.

### HasAllowDecline

`func (o *SignatureRequestSendRequest) HasAllowDecline() bool`

HasAllowDecline returns a boolean if a field has been set.

### GetAllowReassign

`func (o *SignatureRequestSendRequest) GetAllowReassign() bool`

GetAllowReassign returns the AllowReassign field if non-nil, zero value otherwise.

### GetAllowReassignOk

`func (o *SignatureRequestSendRequest) GetAllowReassignOk() (*bool, bool)`

GetAllowReassignOk returns a tuple with the AllowReassign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowReassign

`func (o *SignatureRequestSendRequest) SetAllowReassign(v bool)`

SetAllowReassign sets AllowReassign field to given value.

### HasAllowReassign

`func (o *SignatureRequestSendRequest) HasAllowReassign() bool`

HasAllowReassign returns a boolean if a field has been set.

### GetAttachments

`func (o *SignatureRequestSendRequest) GetAttachments() []SubAttachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *SignatureRequestSendRequest) GetAttachmentsOk() (*[]SubAttachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *SignatureRequestSendRequest) SetAttachments(v []SubAttachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *SignatureRequestSendRequest) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetCcEmailAddresses

`func (o *SignatureRequestSendRequest) GetCcEmailAddresses() []string`

GetCcEmailAddresses returns the CcEmailAddresses field if non-nil, zero value otherwise.

### GetCcEmailAddressesOk

`func (o *SignatureRequestSendRequest) GetCcEmailAddressesOk() (*[]string, bool)`

GetCcEmailAddressesOk returns a tuple with the CcEmailAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcEmailAddresses

`func (o *SignatureRequestSendRequest) SetCcEmailAddresses(v []string)`

SetCcEmailAddresses sets CcEmailAddresses field to given value.

### HasCcEmailAddresses

`func (o *SignatureRequestSendRequest) HasCcEmailAddresses() bool`

HasCcEmailAddresses returns a boolean if a field has been set.

### GetClientId

`func (o *SignatureRequestSendRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *SignatureRequestSendRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *SignatureRequestSendRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *SignatureRequestSendRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetCustomFields

`func (o *SignatureRequestSendRequest) GetCustomFields() []SubCustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *SignatureRequestSendRequest) GetCustomFieldsOk() (*[]SubCustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *SignatureRequestSendRequest) SetCustomFields(v []SubCustomField)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *SignatureRequestSendRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetFieldOptions

`func (o *SignatureRequestSendRequest) GetFieldOptions() SubFieldOptions`

GetFieldOptions returns the FieldOptions field if non-nil, zero value otherwise.

### GetFieldOptionsOk

`func (o *SignatureRequestSendRequest) GetFieldOptionsOk() (*SubFieldOptions, bool)`

GetFieldOptionsOk returns a tuple with the FieldOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldOptions

`func (o *SignatureRequestSendRequest) SetFieldOptions(v SubFieldOptions)`

SetFieldOptions sets FieldOptions field to given value.

### HasFieldOptions

`func (o *SignatureRequestSendRequest) HasFieldOptions() bool`

HasFieldOptions returns a boolean if a field has been set.

### GetFormFieldGroups

`func (o *SignatureRequestSendRequest) GetFormFieldGroups() []SubFormFieldGroup`

GetFormFieldGroups returns the FormFieldGroups field if non-nil, zero value otherwise.

### GetFormFieldGroupsOk

`func (o *SignatureRequestSendRequest) GetFormFieldGroupsOk() (*[]SubFormFieldGroup, bool)`

GetFormFieldGroupsOk returns a tuple with the FormFieldGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormFieldGroups

`func (o *SignatureRequestSendRequest) SetFormFieldGroups(v []SubFormFieldGroup)`

SetFormFieldGroups sets FormFieldGroups field to given value.

### HasFormFieldGroups

`func (o *SignatureRequestSendRequest) HasFormFieldGroups() bool`

HasFormFieldGroups returns a boolean if a field has been set.

### GetFormFieldRules

`func (o *SignatureRequestSendRequest) GetFormFieldRules() []SubFormFieldRule`

GetFormFieldRules returns the FormFieldRules field if non-nil, zero value otherwise.

### GetFormFieldRulesOk

`func (o *SignatureRequestSendRequest) GetFormFieldRulesOk() (*[]SubFormFieldRule, bool)`

GetFormFieldRulesOk returns a tuple with the FormFieldRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormFieldRules

`func (o *SignatureRequestSendRequest) SetFormFieldRules(v []SubFormFieldRule)`

SetFormFieldRules sets FormFieldRules field to given value.

### HasFormFieldRules

`func (o *SignatureRequestSendRequest) HasFormFieldRules() bool`

HasFormFieldRules returns a boolean if a field has been set.

### GetFormFieldsPerDocument

`func (o *SignatureRequestSendRequest) GetFormFieldsPerDocument() []SubFormFieldsPerDocumentBase`

GetFormFieldsPerDocument returns the FormFieldsPerDocument field if non-nil, zero value otherwise.

### GetFormFieldsPerDocumentOk

`func (o *SignatureRequestSendRequest) GetFormFieldsPerDocumentOk() (*[]SubFormFieldsPerDocumentBase, bool)`

GetFormFieldsPerDocumentOk returns a tuple with the FormFieldsPerDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormFieldsPerDocument

`func (o *SignatureRequestSendRequest) SetFormFieldsPerDocument(v []SubFormFieldsPerDocumentBase)`

SetFormFieldsPerDocument sets FormFieldsPerDocument field to given value.

### HasFormFieldsPerDocument

`func (o *SignatureRequestSendRequest) HasFormFieldsPerDocument() bool`

HasFormFieldsPerDocument returns a boolean if a field has been set.

### GetHideTextTags

`func (o *SignatureRequestSendRequest) GetHideTextTags() bool`

GetHideTextTags returns the HideTextTags field if non-nil, zero value otherwise.

### GetHideTextTagsOk

`func (o *SignatureRequestSendRequest) GetHideTextTagsOk() (*bool, bool)`

GetHideTextTagsOk returns a tuple with the HideTextTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideTextTags

`func (o *SignatureRequestSendRequest) SetHideTextTags(v bool)`

SetHideTextTags sets HideTextTags field to given value.

### HasHideTextTags

`func (o *SignatureRequestSendRequest) HasHideTextTags() bool`

HasHideTextTags returns a boolean if a field has been set.

### GetIsQualifiedSignature

`func (o *SignatureRequestSendRequest) GetIsQualifiedSignature() bool`

GetIsQualifiedSignature returns the IsQualifiedSignature field if non-nil, zero value otherwise.

### GetIsQualifiedSignatureOk

`func (o *SignatureRequestSendRequest) GetIsQualifiedSignatureOk() (*bool, bool)`

GetIsQualifiedSignatureOk returns a tuple with the IsQualifiedSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQualifiedSignature

`func (o *SignatureRequestSendRequest) SetIsQualifiedSignature(v bool)`

SetIsQualifiedSignature sets IsQualifiedSignature field to given value.

### HasIsQualifiedSignature

`func (o *SignatureRequestSendRequest) HasIsQualifiedSignature() bool`

HasIsQualifiedSignature returns a boolean if a field has been set.

### GetMessage

`func (o *SignatureRequestSendRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SignatureRequestSendRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SignatureRequestSendRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SignatureRequestSendRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMetadata

`func (o *SignatureRequestSendRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SignatureRequestSendRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SignatureRequestSendRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SignatureRequestSendRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSigningOptions

`func (o *SignatureRequestSendRequest) GetSigningOptions() SubSigningOptions`

GetSigningOptions returns the SigningOptions field if non-nil, zero value otherwise.

### GetSigningOptionsOk

`func (o *SignatureRequestSendRequest) GetSigningOptionsOk() (*SubSigningOptions, bool)`

GetSigningOptionsOk returns a tuple with the SigningOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningOptions

`func (o *SignatureRequestSendRequest) SetSigningOptions(v SubSigningOptions)`

SetSigningOptions sets SigningOptions field to given value.

### HasSigningOptions

`func (o *SignatureRequestSendRequest) HasSigningOptions() bool`

HasSigningOptions returns a boolean if a field has been set.

### GetSigningRedirectUrl

`func (o *SignatureRequestSendRequest) GetSigningRedirectUrl() string`

GetSigningRedirectUrl returns the SigningRedirectUrl field if non-nil, zero value otherwise.

### GetSigningRedirectUrlOk

`func (o *SignatureRequestSendRequest) GetSigningRedirectUrlOk() (*string, bool)`

GetSigningRedirectUrlOk returns a tuple with the SigningRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningRedirectUrl

`func (o *SignatureRequestSendRequest) SetSigningRedirectUrl(v string)`

SetSigningRedirectUrl sets SigningRedirectUrl field to given value.

### HasSigningRedirectUrl

`func (o *SignatureRequestSendRequest) HasSigningRedirectUrl() bool`

HasSigningRedirectUrl returns a boolean if a field has been set.

### GetSubject

`func (o *SignatureRequestSendRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *SignatureRequestSendRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *SignatureRequestSendRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *SignatureRequestSendRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTestMode

`func (o *SignatureRequestSendRequest) GetTestMode() bool`

GetTestMode returns the TestMode field if non-nil, zero value otherwise.

### GetTestModeOk

`func (o *SignatureRequestSendRequest) GetTestModeOk() (*bool, bool)`

GetTestModeOk returns a tuple with the TestMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestMode

`func (o *SignatureRequestSendRequest) SetTestMode(v bool)`

SetTestMode sets TestMode field to given value.

### HasTestMode

`func (o *SignatureRequestSendRequest) HasTestMode() bool`

HasTestMode returns a boolean if a field has been set.

### GetTitle

`func (o *SignatureRequestSendRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SignatureRequestSendRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SignatureRequestSendRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SignatureRequestSendRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUseTextTags

`func (o *SignatureRequestSendRequest) GetUseTextTags() bool`

GetUseTextTags returns the UseTextTags field if non-nil, zero value otherwise.

### GetUseTextTagsOk

`func (o *SignatureRequestSendRequest) GetUseTextTagsOk() (*bool, bool)`

GetUseTextTagsOk returns a tuple with the UseTextTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTextTags

`func (o *SignatureRequestSendRequest) SetUseTextTags(v bool)`

SetUseTextTags sets UseTextTags field to given value.

### HasUseTextTags

`func (o *SignatureRequestSendRequest) HasUseTextTags() bool`

HasUseTextTags returns a boolean if a field has been set.

### GetExpiresAt

`func (o *SignatureRequestSendRequest) GetExpiresAt() int32`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *SignatureRequestSendRequest) GetExpiresAtOk() (*int32, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *SignatureRequestSendRequest) SetExpiresAt(v int32)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *SignatureRequestSendRequest) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *SignatureRequestSendRequest) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *SignatureRequestSendRequest) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

