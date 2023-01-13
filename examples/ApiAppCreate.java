import com.hellosign.openapi.ApiClient;
import com.hellosign.openapi.ApiException;
import com.hellosign.openapi.Configuration;
import com.hellosign.openapi.api.*;
import com.hellosign.openapi.auth.HttpBasicAuth;
import com.hellosign.openapi.auth.HttpBearerAuth;
import com.hellosign.openapi.model.*;
import com.hellosign.openapi.model.SubOAuth.ScopesEnum;

import java.io.File;
import java.util.Arrays;
import java.util.Collections;

public class Example {
    public static void main(String[] args) {
        ApiClient apiClient = Configuration.getDefaultApiClient();

        // Configure HTTP basic authorization: api_key
        HttpBasicAuth apiKey = (HttpBasicAuth) apiClient
            .getAuthentication("api_key");
        apiKey.setUsername("YOUR_API_KEY");

        // or, configure Bearer (JWT) authorization: oauth2
        /*
        HttpBearerAuth oauth2 = (HttpBearerAuth) apiClient
            .getAuthentication("oauth2");

        oauth2.setBearerToken("YOUR_ACCESS_TOKEN");
        */

        ApiAppApi apiAppApi = new ApiAppApi(apiClient);

        SubOAuth oauth = new SubOAuth()
            .callbackUrl("https://example.com/oauth")
            .scopes(Arrays.asList(ScopesEnum.BASIC_ACCOUNT_INFO, ScopesEnum.REQUEST_SIGNATURE));

        SubWhiteLabelingOptions whiteLabelingOptions = new SubWhiteLabelingOptions()
            .primaryButtonColor("#00b3e6")
            .primaryButtonTextColor("#ffffff");

        File customLogoFile = new File("CustomLogoFile.png");

        ApiAppCreateRequest data = new ApiAppCreateRequest()
            .name("My Production App")
            .domains(Collections.singletonList("example.com"))
            .oauth(oauth)
            .whiteLabelingOptions(whiteLabelingOptions)
            .customLogoFile(customLogoFile);

        try {
            ApiAppGetResponse result = apiAppApi.apiAppCreate(data);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountApi#accountCreate");
            System.err.println("Status code: " + e.getCode());
            System.err.println("Reason: " + e.getResponseBody());
            System.err.println("Response headers: " + e.getResponseHeaders());
            e.printStackTrace();
        }
    }
}
