* ONBOARD  
* Customer Creation

# **Create Customer**

POST

## /trade/bo/v1.2/customer/account

The request is used to create a customer and his accounts in the system. This will create customers, their login profile, cash account to manage the cash, portfolio account to manage the holdings, and all the exchange accounts customers allow to trade.

supports:  server token   
---

Customer accounts will be created according to the configured default customer profile or subscription package.

Customer profile ; is similar to a blue print of customer accounts. If this is enabled at the institution level all the customer accounts will be created according to that customer profile.

Subscription package ; enables to define the set of exchanges to which the exchange/ trading accounts should be created for a customer to trade with currencies. Exchange accounts will be created for all the exchanges defined in the subscription package, the cash accounts will be created according to the respective exchange and configured currency and security account will be created with type equity for each cash account.

Note: Please try again, If you receive error code \- 1123\.

Request

* Request Body  
* Headers

**Content-Type:** application/json

* referenceNumberstringrequired  
* Fintech's Reference Number ; this can be any value and value format can be decided by Fintech itself.  
* institutionCodestringrequired  
* firstNamestring  
* lastNamestring  
* passportNumberstring  
* ninstring  
* drivingLicensestring  
* homeTelstring  
* Format: `^\+[1-9]\d{1,14}$`  
* officeTelstring  
* Format: `^\+[1-9]\d{1,14}$`  
* mobilestring  
* Format: `^\+[1-9]\d{1,16}$`  
* emailstring  
* professionstring  
* address1string  
* address2string  
* citystring  
* countryCodestring  
* Find country codes [here](https://developer.globaltradingnetwork.com/trade/docs/1.2.1/apis/manage/master-data/get-country-list)  
* genderstring  
* Enum:M-MaleF-Female-1-rather not say  
* birthDatestring  
* Format: `yyyy/MM/dd`  
* nationalitystring  
* Find country codes [here](https://developer.globaltradingnetwork.com/trade/docs/1.2.1/apis/manage/master-data/get-country-list)  
* profileIdstring  
* masterAccountNumberstring  
* master account number of the customer  
* preferredLanguagestring

preferred language of the customer. only works for DIFC customers  
Enum:ENARRUSPPRITHY

## Responses

200 (OK)  
**Content-Type:** \*/\*

* Schema  
* **Example**  
* statusstringrequired  
* Enum:SUCCESSFAILED  
* reasonstring  
* rejectCodeinteger  
* Find error codes [here](https://developer.globaltradingnetwork.com/trade/docs/1.2.1/error_handling)  
* customerNumberstring  
* GTN customer number of the Customer account which was created at GTN side.  
* cashAccountNumbersarray  
* List of Cash Account Numbers of the created GTN cash accounts under created customer account.  
* accountNumbersarray  
* List of Security Account Numbers of the created GTN security accounts under created customer account.  
* exchangeAccountIdsarray

List of Exchange Account Ids of the created GTN exchange accounts under created customer account.  
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

*curl* \-L \-X POST 'https://sandbox.globaltradingnetwork.com/trade/bo/v1.2/customer/account' \\

\-H 'Content-Type: application/json' \\

\-H 'Accept: \*/\*' \\

\-H 'Throttle-Key: 10' \\

\-H 'Authorization: Bearer \<TOKEN\>' \\

\--data-raw '{

 "referenceNumber": "string",

 "institutionCode": "string",

 "firstName": "string",

 "lastName": "string",

 "passportNumber": "string",

 "nin": "string",

 "drivingLicense": "string",

 "homeTel": "string",

 "officeTel": "string",

 "mobile": "string",

 "email": "string",

 "profession": "string",

 "address1": "string",

 "address2": "string",

 "city": "string",

 "countryCode": "string",

 "gender": "M-Male",

 "birthDate": "string",

 "nationality": "string",

 "profileId": "string",

 "masterAccountNumber": "string",

 "preferredLanguage": "EN"

}'

Request

Base URL

https://sandbox.globaltradingnetwork.com

Auth

Bearer Token

Parameters

Throttle-Key — headerrequired

Body

* Example (from schema)  
* **Example**

{  
  "referenceNumber": "string",  
  "institutionCode": "string",  
  "firstName": "string",  
  "lastName": "string",  
  "passportNumber": "string",  
  "nin": "string",  
  "drivingLicense": "string",  
  "homeTel": "string",  
  "officeTel": "string",  
  "mobile": "string",  
  "email": "string",  
  "profession": "string",  
  "address1": "string",  
  "address2": "string",  
  "city": "string",  
  "countryCode": "string",  
  "gender": "M-Male",  
  "birthDate": "string",  
  "nationality": "string",  
  "profileId": "string",  
  "masterAccountNumber": "string",  
  "preferredLanguage": "EN"  
}

**Send API Request**

Response

Enter your Query and Header parameters and choose `Send API Request` to see the results

Clear

This content was last updated: history 2 months ago  
Did you find it useful?YesNo  
[Customer Token Refresh](https://developer.globaltradingnetwork.com/trade/docs/1.2.1/apis/authenticate/authenticate/customer-token-refresh)  
[Update Customer](https://developer.globaltradingnetwork.com/trade/docs/1.2.1/apis/onboard/customer-creation/update-customer)

© 2026 GTN Group Holding Ltd. All rights reserved.

GTN API and FIX documentation is provided for informational and integration purposes only. GTN API and FIX services are offered on an "as-is" and "as-available" basis.

GTN makes no warranties of any kind, either express or implied, regarding the accuracy, reliability, or functionality of its API and FIX services or documentation.

Users must ensure compliance with all relevant laws and regulations and GTN shall not be responsible for any losses or damages arising from the use of GTN API and FIX services and/or documentation.

ONBOARD

* Customer Creation

# **Update Customer**

PATCH

## /trade/bo/v1.2/customer/account

This is used to update customer details.

supports:  server token customer token  
---

Request

* Request Body  
* Headers

**Content-Type:** application/json

* customerNumberstringrequired  
* Customer number of the customer account which wants to amend.  
* firstNamestring  
* First name  
* lastNamestring  
* Last name  
* passportNumberstring  
* Passport number  
* ninstring  
* National id number  
* drivingLicensestring  
* Driving license  
* homeTelstring  
* Home telephone number  
  Format: `^\+[1-9]\d{1,14}$`  
* officeTelstring  
* Office telephone number  
  Format: `^\+[1-9]\d{1,14}$`  
* mobilestring  
* Mobile number  
  Format: `^\+[1-9]\d{1,16}$`  
* emailstring  
* Email  
* professionstring  
* Profession  
* address1string  
* address2string  
* citystring  
* City  
* countryCodestring  
* Find country codes [here](https://developer.globaltradingnetwork.com/trade/docs/1.2.1/apis/manage/master-data/get-country-list)  
* genderstring  
* M-Male | F-Female | \-1-rather not say  
  Enum:MF-1  
* birthDatestring  
* Birth date  
  Format: `yyyy/MM/dd`  
* nationalitystring  
* Find country codes [here](https://developer.globaltradingnetwork.com/trade/docs/1.2.1/apis/manage/master-data/get-country-list)  
* preferredLanguagestring

preferred language of the customer. only works for DIFC customers  
Enum:ENARRUSPPRITHY

## Responses

200 (OK)  
**Content-Type:** application/json

* Schema  
* **Example**  
* statusstringrequired  
* Enum:SUCCESSFAILED  
* reasonstring  
* rejectCodenumber  
* Find error codes [here](https://developer.globaltradingnetwork.com/trade/docs/1.2.1/error_handling)  
* customerNumberstring

Customer account number  
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

*curl* \-L \-X PATCH 'https://sandbox.globaltradingnetwork.com/trade/bo/v1.2/customer/account' \\

\-H 'Content-Type: application/json' \\

\-H 'Accept: application/json' \\

\-H 'Throttle-Key: 10' \\

\-H 'Authorization: Bearer \<TOKEN\>' \\

\--data-raw '{

 "customerNumber": "string",

 "firstName": "string",

 "lastName": "string",

 "passportNumber": "string",

 "nin": "string",

 "drivingLicense": "string",

 "homeTel": "string",

 "officeTel": "string",

 "mobile": "string",

 "email": "string",

 "profession": "string",

 "address1": "string",

 "address2": "string",

 "city": "string",

 "countryCode": "string",

 "gender": "M",

 "birthDate": "string",

 "nationality": "string",

 "preferredLanguage": "EN"

}'

Request

Base URL

https://sandbox.globaltradingnetwork.com

Auth

Bearer Token

Parameters

Throttle-Key — headerrequired

Body

* Example (from schema)  
* **Example**

{  
  "customerNumber": "string",  
  "firstName": "string",  
  "lastName": "string",  
  "passportNumber": "string",  
  "nin": "string",  
  "drivingLicense": "string",  
  "homeTel": "string",  
  "officeTel": "string",  
  "mobile": "string",  
  "email": "string",  
  "profession": "string",  
  "address1": "string",  
  "address2": "string",  
  "city": "string",  
  "countryCode": "string",  
  "gender": "M",  
  "birthDate": "string",  
  "nationality": "string",  
  "preferredLanguage": "EN"  
}

**Send API Request**

Response

Enter your Query and Header parameters and choose `Send API Request` to see the results

Clear

This content was last updated: history 2 months ago  
Did you find it useful?YesNo  
[Create Customer](https://developer.globaltradingnetwork.com/trade/docs/1.2.1/apis/onboard/customer-creation/create-customer)  
[Get Customer Details](https://developer.globaltradingnetwork.com/trade/docs/1.2.1/apis/onboard/customer-creation/get-customer-details)

© 2026 GTN Group Holding Ltd. All rights reserved.

GTN API and FIX documentation is provided for informational and integration purposes only. GTN API and FIX services are offered on an "as-is" and "as-available" basis.

GTN makes no warranties of any kind, either express or implied, regarding the accuracy, reliability, or functionality of its API and FIX services or documentation.

Users must ensure compliance with all relevant laws and regulations and GTN shall not be responsible for any losses or damages arising from the use of GTN API and FIX services and/or documentation.

* ONBOARD  
* Customer Creation

# **Get Customer Details**

GET

## /trade/bo/v1.2.1/customer/account

This is used to request to get the overall details about a customer with a given customer number. This will give the customer KYC details as well as information about all the accounts under the given customer account.

supports:  server token customer token  
---

T\> Customer number or Reference number or customer token is required.

Request

* Query Params  
* Headers  
* customerNumberstring  
* Customer number of the customer account which wants to get the details.  
  If customer token is used to send the api request, customer number used to create that token will be considered.  
* Possible values:`<= 20 characters`  
* Default: `ASI173455897`  
* referenceNumberstring  
* Reference number of the customer account which wants to get the details.  
* Possible values:`<= 50 characters`

## Responses

200 (OK)

**Content-Type:** application/json

* Schema  
* **Example**  
* statusstringrequired  
* Enum:FAILEDSUCCESS  
* reasonstring  
* rejectCodeinteger  
* Find error codes [here](https://developer.globaltradingnetwork.com/trade/docs/1.2.1/error_handling)  
* customerNumberstring  
* firstNamestring  
* lastNamestring  
* passportNumberstring  
* ninstring  
* drivingLicensestring  
* homeTelstring  
* Format: `^\+[1-9]\d{1,14}$`  
* officeTelstring  
* Format: `^\+[1-9]\d{1,14}$`  
* mobilestring  
* Format: `^\+[1-9]\d{1,16}$`  
* emailstring  
* professionstring  
* address1string  
* address2string  
* citystring  
* countryCodestring  
* genderstring  
* birthDatestring  
* Format: `^[0-9]{4}/(0[1-9]|1[0-2])/(0[1-9]|[1-2][0-9]|3[0-1])$`  
* nationalitystring  
* usernamestring  
* referenceNumberstring  
* masterAccountNumberstring  
* master account number of the customer  
* cashAccountsobject\[\]  
* isStoreEnabledinteger  
* 0 (not enabled) | 1 (enabled)  
  Enum:01  
* isTradingPasswordEnabledinteger  
* 0 (not enabled) | 1 (password each time) | 2 (password once) | \-1 (undefined)  
  Enum:012-1  
* preferredLanguagestring

preferred language of the customer. only works for DIFC customers  
Enum:ENARRUSPPRITHY

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

*curl* \-L \-X GET 'https://sandbox.globaltradingnetwork.com/trade/bo/v1.2.1/customer/account' \\

\-H 'Accept: application/json' \\

\-H 'Throttle-Key: 10' \\

\-H 'Authorization: Bearer \<TOKEN\>'

Request

Base URL

https://sandbox.globaltradingnetwork.com

Auth

Bearer Token

Parameters

Throttle-Key — headerrequired

**Send API Request**

Response

Enter your Query and Header parameters and choose `Send API Request` to see the results

Clear

This content was last updated: history 11 days ago

Did you find it useful?YesNo

[Update Customer](https://developer.globaltradingnetwork.com/trade/docs/1.2.1/apis/onboard/customer-creation/update-customer)  
[Update Account Setup](https://developer.globaltradingnetwork.com/trade/docs/1.2.1/apis/onboard/customer-creation/update-account-setup)

© 2026 GTN Group Holding Ltd. All rights reserved.

GTN API and FIX documentation is provided for informational and integration purposes only. GTN API and FIX services are offered on an "as-is" and "as-available" basis.

GTN makes no warranties of any kind, either express or implied, regarding the accuracy, reliability, or functionality of its API and FIX services or documentation.

Users must ensure compliance with all relevant laws and regulations and GTN shall not be responsible for any losses or damages arising from the use of GTN API and FIX services and/or documentation.

* ONBOARD  
* Customer Creation

# **Update Account Setup**

PATCH

## /trade/bo/v1.2.1/customer/account/profile

This API is used to update customer account(given customer account, if calling using server token or associated customer account when calling using customer token) by using given profile. So the customer account will be updated according to the given profile.

supports:  server token customer token  
---

Request

* Request Body  
* Headers

**Content-Type:** application/json

* customerNumberstring  
* profileIdintegerrequired

## Responses

200 (OK)

**Content-Type:** \*/\*

* Schema  
* **Example**  
* statusstringrequired  
* Enum:SUCCESSFAILED  
* reasonstring  
* rejectCodeinteger

Find error codes [here](https://developer.globaltradingnetwork.com/trade/docs/1.2.1/error_handling)

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

*curl* \-L \-X PATCH 'https://sandbox.globaltradingnetwork.com/trade/bo/v1.2.1/customer/account/profile' \\

\-H 'Content-Type: application/json' \\

\-H 'Accept: \*/\*' \\

\-H 'Throttle-Key: 10' \\

\-H 'Authorization: Bearer \<TOKEN\>' \\

\--data-raw '{

 "customerNumber": "string",

 "profileId": 0

}'

Request

Base URL

https://sandbox.globaltradingnetwork.com

Auth

Bearer Token

Parameters

Throttle-Key — headerrequired

Body

* Example (from schema)  
* **Example**

{  
  "customerNumber": "string",  
  "profileId": 0  
}

**Send API Request**

Response

Enter your Query and Header parameters and choose `Send API Request` to see the results

Clear

This content was last updated: history 2 months ago

Did you find it useful?YesNo

[Get Customer Details](https://developer.globaltradingnetwork.com/trade/docs/1.2.1/apis/onboard/customer-creation/get-customer-details)  
[Estimate Order Charges](https://developer.globaltradingnetwork.com/trade/docs/1.2.1/apis/trade/fees-and-commissions/estimate-order-charges)

© 2026 GTN Group Holding Ltd. All rights reserved.

GTN API and FIX documentation is provided for informational and integration purposes only. GTN API and FIX services are offered on an "as-is" and "as-available" basis.

GTN makes no warranties of any kind, either express or implied, regarding the accuracy, reliability, or functionality of its API and FIX services or documentation.

Users must ensure compliance with all relevant laws and regulations and GTN shall not be responsible for any losses or damages arising from the use of GTN API and FIX services and/or documentation.  
