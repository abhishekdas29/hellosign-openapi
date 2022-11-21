# # UnclaimedDraftCreateEmbeddedRequest



## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
| `clientId`<sup>*_required_</sup> | ```string``` |  Client id of the app used to create the draft. Used to apply the branding and callback url defined for the app.  |  |
| `requesterEmailAddress`<sup>*_required_</sup> | ```string``` |  The email address of the user that should be designated as the requester of this draft, if the draft type is `request_signature`.  |  |
| `file` | ```Array<RequestFile>``` |  Use `file[]` to indicate the uploaded file(s) to send for signature.<br><br>This endpoint requires either **file** or **file_url[]**, but not both.  |  |
| `fileUrl` | ```Array<string>``` |  Use `file_url[]` to have Dropbox Sign download the file(s) to send for signature.<br><br>This endpoint requires either **file** or **file_url[]**, but not both.  |  |
| `allowCcs` | ```boolean``` |  This allows the requester to specify whether the user is allowed to provide email addresses to CC when claiming the draft.  |  [default to true] |
| `allowDecline` | ```boolean``` |  Allows signers to decline to sign a document if `true`. Defaults to `false`.  |  [default to false] |
| `allowReassign` | ```boolean``` |  Allows signers to reassign their signature requests to other signers if set to `true`. Defaults to `false`.<br><br>**Note**: Only available for Premium plan and higher.  |  [default to false] |
| `attachments` | [```Array<SubAttachment>```](SubAttachment.md) |  A list describing the attachments  |  |
| `ccEmailAddresses` | ```Array<string>``` |  The email addresses that should be CCed.  |  |
| `customFields` | [```Array<SubCustomField>```](SubCustomField.md) |  When used together with merge fields, `custom_fields` allows users to add pre-filled data to their signature requests.<br><br>Pre-filled data can be used with &quot;send-once&quot; signature requests by adding merge fields with `form_fields_per_document` or [Text Tags](https://app.hellosign.com/api/textTagsWalkthrough#TextTagIntro) while passing values back with `custom_fields` together in one API call.<br><br>For using pre-filled on repeatable signature requests, merge fields are added to templates in the Dropbox Sign UI or by calling [/template/create_embedded_draft](/api/reference/operation/templateCreateEmbeddedDraft) and then passing `custom_fields` on subsequent signature requests referencing that template.  |  |
| `editorOptions` | [```SubEditorOptions```](SubEditorOptions.md) |    |  |
| `fieldOptions` | [```SubFieldOptions```](SubFieldOptions.md) |    |  |
| `forceSignerPage` | ```boolean``` |  Provide users the ability to review/edit the signers.  |  [default to false] |
| `forceSubjectMessage` | ```boolean``` |  Provide users the ability to review/edit the subject and message.  |  [default to false] |
| `formFieldGroups` | [```Array<SubFormFieldGroup>```](SubFormFieldGroup.md) |  Group information for fields defined in `form_fields_per_document`. String-indexed JSON array with `group_label` and `requirement` keys. `form_fields_per_document` must contain fields referencing a group defined in `form_field_groups`.  |  |
| `formFieldRules` | [```Array<SubFormFieldRule>```](SubFormFieldRule.md) |  Conditional Logic rules for fields defined in `form_fields_per_document`.  |  |
| `formFieldsPerDocument` | [```Array<SubFormFieldsPerDocumentBase>```](SubFormFieldsPerDocumentBase.md) |  The fields that should appear on the document, expressed as an array of objects. (We&#39;re currently fixing a bug where this property only accepts a two-dimensional array. You can read about it here: &lt;a href&#x3D;&quot;/docs/placing-fields/form-fields-per-document&quot; target&#x3D;&quot;_blank&quot;&gt;Using Form Fields per Document&lt;/a&gt;.)<br><br>**NOTE**: Fields like **text**, **dropdown**, **checkbox**, **radio**, and **hyperlink** have additional required and optional parameters. Check out the list of [additional parameters](/api/reference/constants/#form-fields-per-document) for these field types.<br><br>* Text Field use `SubFormFieldsPerDocumentText`<br>* Dropdown Field use `SubFormFieldsPerDocumentDropdown`<br>* Hyperlink Field use `SubFormFieldsPerDocumentHyperlink`<br>* Checkbox Field use `SubFormFieldsPerDocumentCheckbox`<br>* Radio Field use `SubFormFieldsPerDocumentRadio`<br>* Signature Field use `SubFormFieldsPerDocumentSignature`<br>* Date Signed Field use `SubFormFieldsPerDocumentDateSigned`<br>* Initials Field use `SubFormFieldsPerDocumentInitials`<br>* Text Merge Field use `SubFormFieldsPerDocumentTextMerge`<br>* Checkbox Merge Field use `SubFormFieldsPerDocumentCheckboxMerge`  |  |
| `hideTextTags` | ```boolean``` |  Send with a value of `true` if you wish to enable automatic Text Tag removal. Defaults to `false`. When using Text Tags it is preferred that you set this to `false` and hide your tags with white text or something similar because the automatic removal system can cause unwanted clipping. See the [Text Tags](https://app.hellosign.com/api/textTagsWalkthrough#TextTagIntro) walkthrough for more details.  |  [default to false] |
| `holdRequest` | ```boolean``` |  The request from this draft will not automatically send to signers post-claim if set to `true`. Requester must [release](/api/reference/operation/signatureRequestReleaseHold/) the request from hold when ready to send. Defaults to `false`.  |  [default to false] |
| `isForEmbeddedSigning` | ```boolean``` |  The request created from this draft will also be signable in embedded mode if set to `true`. Defaults to `false`.  |  [default to false] |
| `message` | ```string``` |  The custom message in the email that will be sent to the signers.  |  |
| `metadata` | ```{ [key: string]: any; }``` |  Key-value data that should be attached to the signature request. This metadata is included in all API responses and events involving the signature request. For example, use the metadata field to store a signer&#39;s order number for look up when receiving events for the signature request.<br><br>Each request can include up to 10 metadata keys (or 50 nested metadata keys), with key names up to 40 characters long and values up to 1000 characters long.  |  |
| `requestingRedirectUrl` | ```string``` |  The URL you want signers redirected to after they successfully request a signature.  |  |
| `showPreview` | ```boolean``` |  This allows the requester to enable the editor/preview experience.<br><br>- `show_preview&#x3D;true`: Allows requesters to enable the editor/preview experience. - `show_preview&#x3D;false`: Allows requesters to disable the editor/preview experience.  |  |
| `showProgressStepper` | ```boolean``` |  When only one step remains in the signature request process and this parameter is set to `false` then the progress stepper will be hidden.  |  [default to true] |
| `signers` | [```Array<SubUnclaimedDraftSigner>```](SubUnclaimedDraftSigner.md) |  Add Signers to your Unclaimed Draft Signature Request.  |  |
| `signingOptions` | [```SubSigningOptions```](SubSigningOptions.md) |    |  |
| `signingRedirectUrl` | ```string``` |  The URL you want signers redirected to after they successfully sign.  |  |
| `skipMeNow` | ```boolean``` |  Disables the &quot;Me (Now)&quot; option for the person preparing the document. Does not work with type `send_document`. Defaults to `false`.  |  [default to false] |
| `subject` | ```string``` |  The subject in the email that will be sent to the signers.  |  |
| `testMode` | ```boolean``` |  Whether this is a test, the signature request created from this draft will not be legally binding if set to `true`. Defaults to `false`.  |  [default to false] |
| `type` | ```string``` |  The type of the draft. By default this is `request_signature`, but you can set it to `send_document` if you want to self sign a document and download it.  |  [default to TypeEnum_RequestSignature] |
| `usePreexistingFields` | ```boolean``` |  Set `use_text_tags` to `true` to enable [Text Tags](https://app.hellosign.com/api/textTagsWalkthrough#TextTagIntro) parsing in your document (defaults to disabled, or `false`). Alternatively, if your PDF contains pre-defined fields, enable the detection of these fields by setting the `use_preexisting_fields` to `true` (defaults to disabled, or `false`). Currently we only support use of either `use_text_tags` or `use_preexisting_fields` parameter, not both.  |  [default to false] |
| `useTextTags` | ```boolean``` |  Set `use_text_tags` to `true` to enable [Text Tags](https://app.hellosign.com/api/textTagsWalkthrough#TextTagIntro) parsing in your document (defaults to disabled, or `false`). Alternatively, if your PDF contains pre-defined fields, enable the detection of these fields by setting the `use_preexisting_fields` to `true` (defaults to disabled, or `false`). Currently we only support use of either `use_text_tags` or `use_preexisting_fields` parameter, not both.  |  [default to false] |
| `populateAutoFillFields` | ```boolean``` |  Controls whether [auto fill fields](https://faq.hellosign.com/hc/en-us/articles/360051467511-Auto-Fill-Fields) can automatically populate a signer&#39;s information during signing.<br><br>⚠️ **Note** ⚠️: Keep your signer&#39;s information safe by ensuring that the _signer on your signature request is the intended party_ before using this feature.  |  [default to false] |
| `expiresAt` | ```number``` |  When the signature request will expire. Unsigned signatures will be moved to the expired status, and no longer signable.<br><br>**Note** This does not correspond to the **expires_at** returned in the response.  |  |

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)