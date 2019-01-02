Shortener API
=============
This is a simple API for Shortener URL

**Version:** 1.0.0

**Contact information:**  
syrchikov_max@mail.ru  

**License:** [Apache 2.0](http://www.apache.org/licenses/LICENSE-2.0.html)

### Security
---
**basicAuth**  

|basic|*Basic*|
|---|---|

### /users
---
##### ***POST***
**Summary:** User registrations

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| body | body | Created user object | Yes | object |

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Successful operation | [User](#user) |
| 400 | Bad input parameter | object |
| default | Operation error |  |

### /users/login
---
##### ***POST***
**Summary:** Logs user into the system

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| body | body | Created user object | Yes | object |

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Successful operation | [User](#user) |
| 400 | Invalid username/password supplied | object |
| default | Operation error |  |

### /users/logout
---
##### ***GET***
**Summary:** Logs out current logged in user session

**Responses**

| Code | Description |
| ---- | ----------- |
| 200 | Successful operation |
| 401 |  |
| default | Operation error |

**Security**

| Security Schema | Scopes |
| --- | --- |
| basicAuth | |

### /users/me
---
##### ***GET***
**Summary:** Info about current user

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Successful operation | [User](#user) |
| 401 |  |  |
| default | Operation error |  |

**Security**

| Security Schema | Scopes |
| --- | --- |
| basicAuth | |

### /users/me/shorten_urls
---
##### ***POST***
**Summary:** Creating a new short link

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| body | body | Created user object | Yes | object |

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Successful operation | [Link](#link) |
| 400 | Bad input parameter | object |
| 401 |  |  |
| default | Operation error |  |

**Security**

| Security Schema | Scopes |
| --- | --- |
| basicAuth | |

##### ***GET***
**Summary:** Get all short links for this user

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Successful operation | object |
| 401 |  |  |
| default | Operation error |  |

**Security**

| Security Schema | Scopes |
| --- | --- |
| basicAuth | |

### /users/me/shorten_urls/{hash}
---
##### ***GET***
**Summary:** Get information about a specific short user link

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| hash | path | Hash of short url | Yes | string |

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Successful operation | object |
| 401 |  |  |
| 404 | Bad input url | object |
| default | Operation error |  |

**Security**

| Security Schema | Scopes |
| --- | --- |
| basicAuth | |

##### ***DELETE***
**Summary:** Remove short link

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| hash | path | Hash of short url | Yes | string |

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Successful operation |  |
| 401 |  |  |
| 404 | Bad input url | object |
| default | Operation error |  |

**Security**

| Security Schema | Scopes |
| --- | --- |
| basicAuth | |

### /users/me/shorten_urls/{hash}/referers
---
##### ***GET***
**Summary:** Get top from 20 referring sites

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| hash | path | Hash of short url | Yes | string |

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Successful operation | object |
| 401 |  |  |
| 404 | Bad input url | object |
| default | Operation error |  |

**Security**

| Security Schema | Scopes |
| --- | --- |
| basicAuth | |

### /shorten_urls/{hash}
---
##### ***GET***
**Summary:** Get link for redirect

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| hash | path | Hash of short url | Yes | string |

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Successful operation | object |
| 401 |  |  |
| 404 | Bad input url | object |
| default | Operation error |  |

**Security**

| Security Schema | Scopes |
| --- | --- |
| basicAuth | |

### Models
---

### User  

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | long |  | Yes |
| username | string |  | Yes |
| password | string |  | Yes |
| email | string |  | Yes |
| dateCreated | dateTime |  | No |
| lastLoginDate | dateTime |  | No |
| timezone | string |  | No |
| language | string |  | No |

### Link  

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | long |  | Yes |
| hash | string | Short url | Yes |
| fullTitle | string | Full link | Yes |
| user | [User](#user) |  | Yes |
| views | long | Views this link | No |
| dateCreated | dateTime |  | No |

### Transition  

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | long |  | Yes |
| link | [Link](#link) |  | Yes |
| referer | string | Where the transition came from | No |
| dateCreated | dateTime |  | No |

___  
PS
[My private SwaggerHub](https://app.swaggerhub.com/apis/Intercross/shortener/1.0.0)

|  |
| ---- |
| **swagger-markdown -i swagger.yml** - сгенерировать md-файл |

|  |
| ---- |
| **swagger generate server -A shortener -f ./swagger.yml** - сгенерировать код |