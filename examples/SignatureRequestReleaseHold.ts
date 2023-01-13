import * as HelloSignSDK from "hellosign-sdk";

const signatureRequestApi = new HelloSignSDK.SignatureRequestApi();

// Configure HTTP basic authorization: api_key
signatureRequestApi.username = "YOUR_API_KEY";

// or, configure Bearer (JWT) authorization: oauth2
// signatureRequestApi.accessToken = "YOUR_ACCESS_TOKEN";

const signatureRequestId = "2f9781e1a8e2045224d808c153c2e1d3df6f8f2f";

const result = signatureRequestApi.signatureRequestReleaseHold(signatureRequestId);
result.then(response => {
  console.log(response.body);
}).catch(error => {
  console.log("Exception when calling HelloSign API:");
  console.log(error.body);
});
