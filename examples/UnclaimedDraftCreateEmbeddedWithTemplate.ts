import * as HelloSignSDK from "hellosign-sdk";

const unclaimedDraftApi = new HelloSignSDK.UnclaimedDraftApi();

// Configure HTTP basic authorization: api_key
unclaimedDraftApi.username = "YOUR_API_KEY";

// or, configure Bearer (JWT) authorization: oauth2
// unclaimedDraftApi.accessToken = "YOUR_ACCESS_TOKEN";

const signer1: HelloSignSDK.SubUnclaimedDraftTemplateSigner = {
  role: "Client",
  name: "George",
  emailAddress: "george@example.com",
};

const cc1: HelloSignSDK.SubCC = {
  role: "Accounting",
  emailAddress: "accounting@hellosign.com",
};

const data: HelloSignSDK.UnclaimedDraftCreateEmbeddedWithTemplateRequest = {
  clientId: "ec64a202072370a737edf4a0eb7f4437",
  templateIds: ["61a832ff0d8423f91d503e76bfbcc750f7417c78"],
  requesterEmailAddress: "jack@hellosign.com",
  signers: [ signer1 ],
  ccs: [ cc1 ],
  testMode: true,
};

const result = unclaimedDraftApi.unclaimedDraftCreateEmbeddedWithTemplate(data);
result.then(response => {
  console.log(response.body);
}).catch(error => {
  console.log("Exception when calling HelloSign API:");
  console.log(error.body);
});
