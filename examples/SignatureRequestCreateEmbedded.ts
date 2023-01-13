import * as HelloSignSDK from "hellosign-sdk";

const fs = require('fs');

const signatureRequestApi = new HelloSignSDK.SignatureRequestApi();

// Configure HTTP basic authorization: api_key
signatureRequestApi.username = "YOUR_API_KEY";

// or, configure Bearer (JWT) authorization: oauth2
// signatureRequestApi.accessToken = "YOUR_ACCESS_TOKEN";

const signer1: HelloSignSDK.SubSignatureRequestSigner = {
  emailAddress: "jack@example.com",
  name: "Jack",
  order: 0,
};

const signer2: HelloSignSDK.SubSignatureRequestSigner = {
  emailAddress: "jill@example.com",
  name: "Jill",
  order: 1,
};

const signingOptions: HelloSignSDK.SubSigningOptions = {
  draw: true,
  type: true,
  upload: true,
  phone: true,
  defaultType: HelloSignSDK.SubSigningOptions.DefaultTypeEnum.Draw,
};

const data: HelloSignSDK.SignatureRequestCreateEmbeddedRequest = {
  clientId: "ec64a202072370a737edf4a0eb7f4437",
  title: "NDA with Acme Co.",
  subject: "The NDA we talked about",
  message: "Please sign this NDA and then we can discuss more. Let me know if you have any questions.",
  signers: [ signer1, signer2 ],
  ccEmailAddresses: [
    "lawyer@hellosign.com",
    "lawyer@example.com",
  ],
  files: [fs.createReadStream("example_signature_request.pdf")],
  signingOptions,
  testMode: true,
};

const result = signatureRequestApi.signatureRequestCreateEmbedded(data);
result.then(response => {
  console.log(response.body);
}).catch(error => {
  console.log("Exception when calling HelloSign API:");
  console.log(error.body);
});
