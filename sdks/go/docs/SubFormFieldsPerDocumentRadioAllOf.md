# SubFormFieldsPerDocumentRadioAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | An input field for radios. Use the &#x60;SubFormFieldsPerDocumentRadio&#x60; class. | [default to "radio"]
**Group** | **string** | String referencing group defined in &#x60;form_field_groups&#x60; parameter. | 
**IsChecked** | **bool** | &#x60;true&#x60; for checking the radio field by default, otherwise &#x60;false&#x60;. Only one radio field per group can be &#x60;true&#x60;. | 

## Methods

### NewSubFormFieldsPerDocumentRadioAllOf

`func NewSubFormFieldsPerDocumentRadioAllOf(type_ string, group string, isChecked bool, ) *SubFormFieldsPerDocumentRadioAllOf`

NewSubFormFieldsPerDocumentRadioAllOf instantiates a new SubFormFieldsPerDocumentRadioAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubFormFieldsPerDocumentRadioAllOfWithDefaults

`func NewSubFormFieldsPerDocumentRadioAllOfWithDefaults() *SubFormFieldsPerDocumentRadioAllOf`

NewSubFormFieldsPerDocumentRadioAllOfWithDefaults instantiates a new SubFormFieldsPerDocumentRadioAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SubFormFieldsPerDocumentRadioAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SubFormFieldsPerDocumentRadioAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SubFormFieldsPerDocumentRadioAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetGroup

`func (o *SubFormFieldsPerDocumentRadioAllOf) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *SubFormFieldsPerDocumentRadioAllOf) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *SubFormFieldsPerDocumentRadioAllOf) SetGroup(v string)`

SetGroup sets Group field to given value.


### GetIsChecked

`func (o *SubFormFieldsPerDocumentRadioAllOf) GetIsChecked() bool`

GetIsChecked returns the IsChecked field if non-nil, zero value otherwise.

### GetIsCheckedOk

`func (o *SubFormFieldsPerDocumentRadioAllOf) GetIsCheckedOk() (*bool, bool)`

GetIsCheckedOk returns a tuple with the IsChecked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsChecked

`func (o *SubFormFieldsPerDocumentRadioAllOf) SetIsChecked(v bool)`

SetIsChecked sets IsChecked field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


