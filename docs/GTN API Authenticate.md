* AUTHENTICATE  
* Authenticate

POST

## /trade/auth/token

App Key and App Secret will shared by GTN.

A signed JWT token is used to get the access token to the open API services.

FinTech public key \- need to generate as mentioned [here](https://developer.globaltradingnetwork.com/trade/docs/1.2.1/getting-started/key-concepts#security-keys)

FinTech private key \- need to generate as mentioned [here](https://developer.globaltradingnetwork.com/trade/docs/1.2.1/getting-started/key-concepts#security-keys)

Assertion \- This is a JWT token that will be generated from the client-server side and sent as a request parameter. This will be generated using, Fintech private key, institution app key and claims \- institution code as instCode and server id as userId.

When generating JWT from the server side, you need to add the institution code and server id/ user id as claims.

instCode : FinTech institution code is given by GTN

userId : FinTech backend service's instance unique id

Recommend to use 1 server token over using multiple server tokens.

Support both Base16 and Base64 encoding [Click here to generate the assertion.](https://developer.globaltradingnetwork.com/generate-assertion)

Code Sample \- Base16

* **Java**  
* **.net**  
* **python**  
* **JavaScript**

const crypto \= *require*("crypto");

const *base16Decoder* \= (hexString) \=\> {

   const byteArray \= Buffer.from(hexString, "hex");

   return byteArray;

};

const *base64UrlEncode* \= (buffer) \=\> {

   return buffer

       .*toString*("base64")

       .*replace*(/\\+/g, "-")

       .*replace*(/\\//g, "\_")

       .*replace*(/=+$/, "");

};

const *createToken* \= (privateKeyHex, appKey, institution, userId) \=\> {

   try {

       const privateKeyBuffer \= *base16Decoder*(privateKeyHex);

       const privateKey \= {

           key: privateKeyBuffer,

           format: "der",

           type: "pkcs8"

       };

       const accessTokenExpiryTimeMs \= 123123123234;

       const currentTimeMs \= Date.*now*();

       const accessTokenExpiry \= currentTimeMs \+ accessTokenExpiryTimeMs;

       const exp \= Math.*floor*(accessTokenExpiry / 1000);

       const payload \= {

           iss: appKey,

           instCode: institution,

           exp: exp,

           userId: userId,

           iat: Math.*floor*(currentTimeMs / 1000),

       };

       const header \= {

           alg: "RS256",

           typ: "JWT"

       };

       const encodedHeader \= *base64UrlEncode*(Buffer.from(JSON.*stringify*(header)));

       const encodedPayload \= *base64UrlEncode*(Buffer.from(JSON.*stringify*(payload)));

       const dataToSign \= \`${encodedHeader}.${encodedPayload}\`;

       const signature \= crypto.*sign*("RSA-SHA256", Buffer.from(dataToSign), privateKey);

       const encodedSignature \= *base64UrlEncode*(signature);

       const accessToken \= \`${dataToSign}.${encodedSignature}\`;

       return accessToken;

   } catch (error) {

       console.*error*("Error creating token:", error);

       return null;

   }

};

function *main*() {

   const appKey \= 'app key';

   const institution \= 'inst code';

   const userId \= 'user id'

   const privateKey \= 'private key';

   const appSecret \= 'app secret';

   const assertion \= *createToken*(privateKey, appKey, institution, appSecret, userId);

   console.*log*("assertion : "\+assertion);

   const authorization \= *btoa*(\`${appKey}:${appSecret}\`);

   console.*log*(\`authorization: Basic ${authorization}\`);

}

*main*();

Code Sample \- Base64

* **Java**  
* **.net**  
* **python**  
* **JavaScript**

const crypto \= *require*("crypto");

const *createToken* \= (privateKeyBase64, appKey, institution, userId) \=\> {

   try {

       const privateKeyBuffer \= Buffer.from(privateKeyBase64, "base64");

       const privateKey \= {

           key: privateKeyBuffer,

           format: "der",

           type: "pkcs8"

       };

       const accessTokenExpiryTimeMs \= 123123123234;

       const currentTimeMs \= Date.*now*();

       const accessTokenExpiry \= currentTimeMs \+ accessTokenExpiryTimeMs;

       const exp \= Math.*floor*(accessTokenExpiry / 1000);

       const payload \= {

           iss: appKey,

           instCode: institution,

           exp: exp,

           userId: userId,

           iat: Math.*floor*(currentTimeMs / 1000),

       };

       const header \= {

           alg: "RS256",

           typ: "JWT"

       };

       const encodedHeader \= Buffer.from(JSON.*stringify*(header)).*toString*("base64url");

       const encodedPayload \= Buffer.from(JSON.*stringify*(payload)).*toString*("base64url");

       const dataToSign \= \`${encodedHeader}.${encodedPayload}\`;

       const signature \= crypto.*sign*("RSA-SHA256", Buffer.from(dataToSign), privateKey);

       const encodedSignature \= signature.*toString*("base64url");

       const accessToken \= \`${dataToSign}.${encodedSignature}\`;

       return accessToken;

   } catch (error) {

       console.*error*("Error creating token:", error);

       return null;

   }

};

function *main*() {

   const appKey \= 'app key';

   const institution \= 'inst code';

   const userId \= 'user id';

   const privateKey \= 'base64\_private\_key';

   const appSecret \= 'app secret';

   const assertion \= *createToken*(privateKey, appKey, institution, userId);

   console.*log*("assertion : "\+assertion);

   const authorization \= Buffer.from(\`${appKey}:${appSecret}\`).*toString*("base64");

   console.*log*(\`authorization: Basic ${authorization}\`);

}

*main*();

Request

* Request Body  
* Headers

**Content-Type:** application/json

* assertionstringrequired

signed JWT token at client side using institution code and user id in the payload

## Responses

200 (OK)  
**Content-Type:** application/json

* Schema  
* **Example**  
* statusstringrequired  
* Enum:SUCCESSFAILED  
* reasonstringrequired  
* rejectCodeintegerrequired  
* Find error codes [here](https://developer.globaltradingnetwork.com/trade/docs/1.2.1/error_handling)  
* accessTokenstring  
* refreshTokenstring  
* accessTokenExpiresAtinteger  
* UTC time  
* refreshTokenExpiresAtinteger  
* UTC time  
* tokenTypestring

401 (Unauthorized)  
**Content-Type:** application/json

* Schema  
* **Example**  
* timestampstring  
* statusstring  
* errorstring  
* messagestring  
* pathstring  
* rejectCodestring

### **Sandbox**

* curl  
* python  
* go  
* ⋮  
* CURL

*curl* \-L \-X POST 'https://sandbox.globaltradingnetwork.com/trade/auth/token' \\

\-H 'Content-Type: application/json' \\

\-H 'Accept: application/json' \\

\-H 'Throttle-Key: 10' \\

\-H 'Authorization: Basic YWxleGNvbGxzb3V0dW11cm9AZ21haWwuY29tOiEiKSFtYXhBQ08zMTMhOw==' \\

\--data-raw '{

 "assertion": "string"

}'

Request

Base URL

https://sandbox.globaltradingnetwork.com

Auth

App Key

App Secret

Parameters

Throttle-Key — headerrequired

Body

* Example (from schema)  
* **Example**

{  
  "assertion": "string"  
}

**Send API Request**

Response

Enter your Query and Header parameters and choose `Send API Request` to see the results

Clear

This content was last updated: history 2 months ago  
Did you find it useful?YesNo  
[Token Refresh](https://developer.globaltradingnetwork.com/trade/docs/1.2.1/apis/authenticate/authenticate/token-refresh)

© 2026 GTN Group Holding Ltd. All rights reserved.

GTN API and FIX documentation is provided for informational and integration purposes only. GTN API and FIX services are offered on an "as-is" and "as-available" basis.

GTN makes no warranties of any kind, either express or implied, regarding the accuracy, reliability, or functionality of its API and FIX services or documentation.

Users must ensure compliance with all relevant laws and regulations and GTN shall not be responsible for any losses or damages arising from the use of GTN API and FIX services and/or documentation.

AUTHENTICATE

* Authenticate

# **Token Refresh**

POST

## /trade/auth/token/refresh

This is used to get a new access token using the refresh token. This needs to be done before the refresh token expires. If it has already expired, you need to do the authentication step from the stretch.

Request

* Request Body  
* Headers

**Content-Type:** application/json

* refreshTokenstringrequired

## Responses

200 (OK)

**Content-Type:** application/json

* Schema  
* **Example**  
* statusstringrequired  
* Enum:SUCCESSFAILED  
* reasonstringrequired  
* rejectCodeintegerrequired  
* Find error codes [here](https://developer.globaltradingnetwork.com/trade/docs/1.2.1/error_handling)  
* accessTokenstring  
* refreshTokenstring  
* accessTokenExpiresAtinteger  
* UTC time  
* refreshTokenExpiresAtinteger  
* UTC time  
* tokenTypestring

### **Sandbox**

* curl  
* python  
* go  
* ⋮  
* CURL

*curl* \-L \-X POST 'https://sandbox.globaltradingnetwork.com/trade/auth/token/refresh' \\

\-H 'Content-Type: application/json' \\

\-H 'Accept: application/json' \\

\-H 'Throttle-Key: 10' \\

\--data-raw '{

 "refreshToken": "string"

}'

Request

Base URL

https://sandbox.globaltradingnetwork.com

Parameters

Throttle-Key — headerrequired

Body

* Example (from schema)  
* **Example**

{  
  "refreshToken": "string"  
}

**Send API Request**

Response

Enter your Query and Header parameters and choose `Send API Request` to see the results

Clear

This content was last updated: history 2 months ago

Did you find it useful?YesNo

[Get Token](https://developer.globaltradingnetwork.com/trade/docs/1.2.1/apis/authenticate/authenticate/get-token)  
[Get Customer Token](https://developer.globaltradingnetwork.com/trade/docs/1.2.1/apis/authenticate/authenticate/get-customer-token)

© 2026 GTN Group Holding Ltd. All rights reserved.

GTN API and FIX documentation is provided for informational and integration purposes only. GTN API and FIX services are offered on an "as-is" and "as-available" basis.

GTN makes no warranties of any kind, either express or implied, regarding the accuracy, reliability, or functionality of its API and FIX services or documentation.

Users must ensure compliance with all relevant laws and regulations and GTN shall not be responsible for any losses or damages arising from the use of GTN API and FIX services and/or documentation.

* AUTHENTICATE  
* Authenticate

# **Get Customer Token**

POST

## /trade/auth/customer/token

This is used to get the customer (end-user application) access token using the server access token and customer number created in the GTN trading platform. This access token can be used to access the GTN open API via end-user applications.

Request

* Request Body  
* Headers

**Content-Type:** application/json

* customerNumberstringrequired  
* Customer Number of a customer.  
* accessTokenstringrequired

Server access token which gained from authentication/severAuthToken

## Responses

200 (OK)  
**Content-Type:** application/json

* Schema  
* **Example**  
* statusstringrequired  
* Enum:SUCCESSFAILED  
* reasonstringrequired  
* rejectCodeintegerrequired  
* Find error codes [here](https://developer.globaltradingnetwork.com/trade/docs/1.2.1/error_handling)  
* accessTokenstring  
* refreshTokenstring  
* accessTokenExpiresAtinteger  
* UTC time  
* refreshTokenExpiresAtinteger  
* UTC time  
* tokenTypestring

401 (Unauthorized)  
**Content-Type:** application/json

* Schema  
* **Example**  
* timestampstring  
* statusstring  
* errorstring  
* messagestring  
* pathstring  
* rejectCodestring

### **Sandbox**

* curl  
* python  
* go  
* ⋮  
* CURL

*curl* \-L \-X POST 'https://sandbox.globaltradingnetwork.com/trade/auth/customer/token' \\

\-H 'Content-Type: application/json' \\

\-H 'Accept: application/json' \\

\-H 'Throttle-Key: 10' \\

\--data-raw '{

 "customerNumber": "string",

 "accessToken": "string"

}'

Request

Base URL

https://sandbox.globaltradingnetwork.com

Parameters

Throttle-Key — headerrequired

Body

* Example (from schema)  
* **Example**

{  
  "customerNumber": "string",  
  "accessToken": "string"  
}

**Send API Request**

Response

Enter your Query and Header parameters and choose `Send API Request` to see the results

Clear

This content was last updated: history 2 months ago  
Did you find it useful?YesNo  
[Token Refresh](https://developer.globaltradingnetwork.com/trade/docs/1.2.1/apis/authenticate/authenticate/token-refresh)  
[Customer Token Refresh](https://developer.globaltradingnetwork.com/trade/docs/1.2.1/apis/authenticate/authenticate/customer-token-refresh)

© 2026 GTN Group Holding Ltd. All rights reserved.

GTN API and FIX documentation is provided for informational and integration purposes only. GTN API and FIX services are offered on an "as-is" and "as-available" basis.

GTN makes no warranties of any kind, either express or implied, regarding the accuracy, reliability, or functionality of its API and FIX services or documentation.

Users must ensure compliance with all relevant laws and regulations and GTN shall not be responsible for any losses or damages arising from the use of GTN API and FIX services and/or documentation.

* AUTHENTICATE  
* Authenticate

# **Customer Token Refresh**

POST

## /trade/auth/customer/token/refresh

This is used to get a new end-user access token using the refresh token. This needs to be done before the refresh token expires. If it has already expired, you need to do the authentication step from the stretch.

Request

* Request Body  
* Headers

**Content-Type:** application/json

* refreshTokenstringrequired

## Responses

200 (OK)

**Content-Type:** application/json

* Schema  
* **Example**  
* statusstringrequired  
* reasonstringrequired  
* rejectCodeintegerrequired  
* Find error codes [here](https://developer.globaltradingnetwork.com/trade/docs/1.2.1/error_handling)  
* accessTokenstring  
* refreshTokenstring  
* accessTokenExpiresAtinteger  
* UTC time  
* refreshTokenExpiresAtinteger  
* UTC time  
* tokenTypestring

### **Sandbox**

* curl  
* python  
* go  
* ⋮  
* CURL

*curl* \-L \-X POST 'https://sandbox.globaltradingnetwork.com/trade/auth/customer/token/refresh' \\

\-H 'Content-Type: application/json' \\

\-H 'Accept: application/json' \\

\-H 'Throttle-Key: 10' \\

\--data-raw '{

 "refreshToken": "string"

}'

Request

Base URL

https://sandbox.globaltradingnetwork.com

Parameters

Throttle-Key — headerrequired

Body

* Example (from schema)  
* **Example**

{  
  "refreshToken": "string"  
}

**Send API Request**

Response

Enter your Query and Header parameters and choose `Send API Request` to see the results

Clear

This content was last updated: history 2 months ago

Did you find it useful?YesNo

[Get Customer Token](https://developer.globaltradingnetwork.com/trade/docs/1.2.1/apis/authenticate/authenticate/get-customer-token)  
[Create Customer](https://developer.globaltradingnetwork.com/trade/docs/1.2.1/apis/onboard/customer-creation/create-customer)

© 2026 GTN Group Holding Ltd. All rights reserved.

GTN API and FIX documentation is provided for informational and integration purposes only. GTN API and FIX services are offered on an "as-is" and "as-available" basis.

GTN makes no warranties of any kind, either express or implied, regarding the accuracy, reliability, or functionality of its API and FIX services or documentation.

Users must ensure compliance with all relevant laws and regulations and GTN shall not be responsible for any losses or damages arising from the use of GTN API and FIX services and/or documentation.

