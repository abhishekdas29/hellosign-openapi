# SubFormFieldsPerDocumentTextAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | A text input field. Use the &#x60;SubFormFieldsPerDocumentText&#x60; class. | [default to "text"]
**Placeholder** | Pointer to **string** | Placeholder value for text field. | [optional] 
**AutoFillType** | Pointer to **string** | Auto fill type for populating fields automatically. Check out the list of [auto fill types](/api/reference/constants/#auto-fill-types) to learn more about the possible values. | [optional] 
**LinkId** | Pointer to **string** | Link two or more text fields. Enter data into one linked text field, which automatically fill all other linked text fields. | [optional] 
**Masked** | Pointer to **bool** | Masks entered data. For more information see [Masking sensitive information](https://faq.hellosign.com/hc/en-us/articles/360040742811-Masking-sensitive-information). &#x60;true&#x60; for masking the data in a text field, otherwise &#x60;false&#x60;. | [optional] 
**ValidationType** | Pointer to **string** | Each text field may contain a &#x60;validation_type&#x60; parameter. Check out the list of [validation types](https://faq.hellosign.com/hc/en-us/articles/217115577) to learn more about the possible values.  **NOTE**: When using &#x60;custom_regex&#x60; you are required to pass a second parameter &#x60;validation_custom_regex&#x60; and you can optionally provide &#x60;validation_custom_regex_format_label&#x60; for the error message the user will see in case of an invalid value. | [optional] 
**ValidationCustomRegex** | Pointer to **string** |  | [optional] 
**ValidationCustomRegexFormatLabel** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **string** | Content of a &#x60;me_now&#x60; text field | [optional] 
**FontFamily** | Pointer to **string** | Font family for the field. | [optional] 
**FontSize** | Pointer to **int32** | The initial px font size for the field contents. Can be any integer value between &#x60;7&#x60; and &#x60;49&#x60;.  **NOTE**: Font size may be reduced during processing in order to fit the contents within the dimensions of the field. | [optional] 

## Methods

### NewSubFormFieldsPerDocumentTextAllOf

`func NewSubFormFieldsPerDocumentTextAllOf(type_ string, ) *SubFormFieldsPerDocumentTextAllOf`

NewSubFormFieldsPerDocumentTextAllOf instantiates a new SubFormFieldsPerDocumentTextAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubFormFieldsPerDocumentTextAllOfWithDefaults

`func NewSubFormFieldsPerDocumentTextAllOfWithDefaults() *SubFormFieldsPerDocumentTextAllOf`

NewSubFormFieldsPerDocumentTextAllOfWithDefaults instantiates a new SubFormFieldsPerDocumentTextAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SubFormFieldsPerDocumentTextAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SubFormFieldsPerDocumentTextAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SubFormFieldsPerDocumentTextAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetPlaceholder

`func (o *SubFormFieldsPerDocumentTextAllOf) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *SubFormFieldsPerDocumentTextAllOf) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *SubFormFieldsPerDocumentTextAllOf) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *SubFormFieldsPerDocumentTextAllOf) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.

### GetAutoFillType

`func (o *SubFormFieldsPerDocumentTextAllOf) GetAutoFillType() string`

GetAutoFillType returns the AutoFillType field if non-nil, zero value otherwise.

### GetAutoFillTypeOk

`func (o *SubFormFieldsPerDocumentTextAllOf) GetAutoFillTypeOk() (*string, bool)`

GetAutoFillTypeOk returns a tuple with the AutoFillType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoFillType

`func (o *SubFormFieldsPerDocumentTextAllOf) SetAutoFillType(v string)`

SetAutoFillType sets AutoFillType field to given value.

### HasAutoFillType

`func (o *SubFormFieldsPerDocumentTextAllOf) HasAutoFillType() bool`

HasAutoFillType returns a boolean if a field has been set.

### GetLinkId

`func (o *SubFormFieldsPerDocumentTextAllOf) GetLinkId() string`

GetLinkId returns the LinkId field if non-nil, zero value otherwise.

### GetLinkIdOk

`func (o *SubFormFieldsPerDocumentTextAllOf) GetLinkIdOk() (*string, bool)`

GetLinkIdOk returns a tuple with the LinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkId

`func (o *SubFormFieldsPerDocumentTextAllOf) SetLinkId(v string)`

SetLinkId sets LinkId field to given value.

### HasLinkId

`func (o *SubFormFieldsPerDocumentTextAllOf) HasLinkId() bool`

HasLinkId returns a boolean if a field has been set.

### GetMasked

`func (o *SubFormFieldsPerDocumentTextAllOf) GetMasked() bool`

GetMasked returns the Masked field if non-nil, zero value otherwise.

### GetMaskedOk

`func (o *SubFormFieldsPerDocumentTextAllOf) GetMaskedOk() (*bool, bool)`

GetMaskedOk returns a tuple with the Masked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasked

`func (o *SubFormFieldsPerDocumentTextAllOf) SetMasked(v bool)`

SetMasked sets Masked field to given value.

### HasMasked

`func (o *SubFormFieldsPerDocumentTextAllOf) HasMasked() bool`

HasMasked returns a boolean if a field has been set.

### GetValidationType

`func (o *SubFormFieldsPerDocumentTextAllOf) GetValidationType() string`

GetValidationType returns the ValidationType field if non-nil, zero value otherwise.

### GetValidationTypeOk

`func (o *SubFormFieldsPerDocumentTextAllOf) GetValidationTypeOk() (*string, bool)`

GetValidationTypeOk returns a tuple with the ValidationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationType

`func (o *SubFormFieldsPerDocumentTextAllOf) SetValidationType(v string)`

SetValidationType sets ValidationType field to given value.

### HasValidationType

`func (o *SubFormFieldsPerDocumentTextAllOf) HasValidationType() bool`

HasValidationType returns a boolean if a field has been set.

### GetValidationCustomRegex

`func (o *SubFormFieldsPerDocumentTextAllOf) GetValidationCustomRegex() string`

GetValidationCustomRegex returns the ValidationCustomRegex field if non-nil, zero value otherwise.

### GetValidationCustomRegexOk

`func (o *SubFormFieldsPerDocumentTextAllOf) GetValidationCustomRegexOk() (*string, bool)`

GetValidationCustomRegexOk returns a tuple with the ValidationCustomRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationCustomRegex

`func (o *SubFormFieldsPerDocumentTextAllOf) SetValidationCustomRegex(v string)`

SetValidationCustomRegex sets ValidationCustomRegex field to given value.

### HasValidationCustomRegex

`func (o *SubFormFieldsPerDocumentTextAllOf) HasValidationCustomRegex() bool`

HasValidationCustomRegex returns a boolean if a field has been set.

### GetValidationCustomRegexFormatLabel

`func (o *SubFormFieldsPerDocumentTextAllOf) GetValidationCustomRegexFormatLabel() string`

GetValidationCustomRegexFormatLabel returns the ValidationCustomRegexFormatLabel field if non-nil, zero value otherwise.

### GetValidationCustomRegexFormatLabelOk

`func (o *SubFormFieldsPerDocumentTextAllOf) GetValidationCustomRegexFormatLabelOk() (*string, bool)`

GetValidationCustomRegexFormatLabelOk returns a tuple with the ValidationCustomRegexFormatLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationCustomRegexFormatLabel

`func (o *SubFormFieldsPerDocumentTextAllOf) SetValidationCustomRegexFormatLabel(v string)`

SetValidationCustomRegexFormatLabel sets ValidationCustomRegexFormatLabel field to given value.

### HasValidationCustomRegexFormatLabel

`func (o *SubFormFieldsPerDocumentTextAllOf) HasValidationCustomRegexFormatLabel() bool`

HasValidationCustomRegexFormatLabel returns a boolean if a field has been set.

### GetContent

`func (o *SubFormFieldsPerDocumentTextAllOf) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SubFormFieldsPerDocumentTextAllOf) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SubFormFieldsPerDocumentTextAllOf) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *SubFormFieldsPerDocumentTextAllOf) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetFontFamily

`func (o *SubFormFieldsPerDocumentTextAllOf) GetFontFamily() string`

GetFontFamily returns the FontFamily field if non-nil, zero value otherwise.

### GetFontFamilyOk

`func (o *SubFormFieldsPerDocumentTextAllOf) GetFontFamilyOk() (*string, bool)`

GetFontFamilyOk returns a tuple with the FontFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFontFamily

`func (o *SubFormFieldsPerDocumentTextAllOf) SetFontFamily(v string)`

SetFontFamily sets FontFamily field to given value.

### HasFontFamily

`func (o *SubFormFieldsPerDocumentTextAllOf) HasFontFamily() bool`

HasFontFamily returns a boolean if a field has been set.

### GetFontSize

`func (o *SubFormFieldsPerDocumentTextAllOf) GetFontSize() int32`

GetFontSize returns the FontSize field if non-nil, zero value otherwise.

### GetFontSizeOk

`func (o *SubFormFieldsPerDocumentTextAllOf) GetFontSizeOk() (*int32, bool)`

GetFontSizeOk returns a tuple with the FontSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFontSize

`func (o *SubFormFieldsPerDocumentTextAllOf) SetFontSize(v int32)`

SetFontSize sets FontSize field to given value.

### HasFontSize

`func (o *SubFormFieldsPerDocumentTextAllOf) HasFontSize() bool`

HasFontSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


