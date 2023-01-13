import * as HelloSignSDK from "hellosign-sdk";

const unclaimedDraftApi = new HelloSignSDK.UnclaimedDraftApi();

// Configure HTTP basic authorization: api_key
unclaimedDraftApi.username = "YOUR_API_KEY";

// or, configure Bearer (JWT) authorization: oauth2
// unclaimedDraftApi.accessToken = "YOUR_ACCESS_TOKEN";

const data: HelloSignSDK.UnclaimedDraftEditAndResendRequest = {
  clientId: "ec64a202072370a737edf4a0eb7f4437",
  testMode: true,
};

const signatureRequestId = "2f9781e1a83jdja934d808c153c2e1d3df6f8f2f";

const result = unclaimedDraftApi.unclaimedDraftEditAndResend(signatureRequestId, data);
result.then(response => {
  console.log(response.body);
}).catch(error => {
  console.log("Exception when calling HelloSign API:");
  console.log(error.body);
});
