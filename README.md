Shortener API
=============
This is a simple API for Shortener URL

**Version:** 1.0.0

**Contact information:**  
syrchikov_max@mail.ru  

**License:** [Apache 2.0](http://www.apache.org/licenses/LICENSE-2.0.html)

### Security
---
**BasicAuth**  

|basic|*Basic*|
|---|---|

### /users
---
##### ***POST***
**Summary:** User registrations

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| body | body | User registration | Yes | object |

**Responses**

| Code | Description |
| ---- | ----------- |
| 200 |  |
| 400 |  |

### /users/login
---
##### ***POST***
**Summary:** Logs user into the system

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| body | body | User login | Yes | object |

**Responses**

| Code | Description |
| ---- | ----------- |
| 200 |  |
| 400 |  |

### /users/logout
---
##### ***GET***
**Summary:** Logs out current logged in user session

**Responses**

| Code | Description |
| ---- | ----------- |
| 204 |  |
| 401 |  |

**Security**

| Security Schema | Scopes |
| --- | --- |
| BasicAuth | |

### /users/me
---
##### ***GET***
**Summary:** Info about current user

**Responses**

| Code | Description |
| ---- | ----------- |
| 200 |  |
| 401 |  |

**Security**

| Security Schema | Scopes |
| --- | --- |
| BasicAuth | |

### /users/me/shorten_urls
---
##### ***POST***
**Summary:** Creating a new short link

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| body | body | Full link | Yes | object |

**Responses**

| Code | Description |
| ---- | ----------- |
| 200 |  |
| 400 |  |
| 401 |  |

**Security**

| Security Schema | Scopes |
| --- | --- |
| BasicAuth | |

##### ***GET***
**Summary:** Get all short links for this user

**Responses**

| Code | Description |
| ---- | ----------- |
| 200 |  |
| 401 |  |

**Security**

| Security Schema | Scopes |
| --- | --- |
| BasicAuth | |

### /users/me/shorten_urls/{hash}
---
##### ***GET***
**Summary:** Get information about a specific short user link

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| hash | path | Hash of short url | Yes | string |

**Responses**

| Code | Description |
| ---- | ----------- |
| 200 |  |
| 400 |  |
| 404 |  |

**Security**

| Security Schema | Scopes |
| --- | --- |
| BasicAuth | |

##### ***DELETE***
**Summary:** Remove short link

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| hash | path | Hash of short url | Yes | string |

**Responses**

| Code | Description |
| ---- | ----------- |
| 204 |  |
| 401 |  |
| 404 |  |

**Security**

| Security Schema | Scopes |
| --- | --- |
| BasicAuth | |

### /users/me/shorten_urls/{hash}/referers
---
##### ***GET***
**Summary:** Get top from 20 referring sites

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| hash | path | Hash of short url | Yes | string |

**Responses**

| Code | Description |
| ---- | ----------- |
| 200 |  |
| 401 |  |
| 404 |  |

**Security**

| Security Schema | Scopes |
| --- | --- |
| BasicAuth | |

### /shorten_urls/{hash}
---
##### ***GET***
**Summary:** Get link for redirect

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| hash | path | Hash of short url | Yes | string |

**Responses**

| Code | Description |
| ---- | ----------- |
| 301 |  |
| 404 |  |

### Models
---

### User  

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | long |  | Yes |
| username | string |  | Yes |
| hash | string |  | Yes |
| email | string |  | Yes |
| dateCreated | string |  | Yes |
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

### Error  

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| code | integer |  | Yes |
| message | string |  | No |

___  
PS
[My private SwaggerHub](https://app.swaggerhub.com/apis/Intercross/shortener/1.0.0)

|  |
| ---- |
| **swagger-markdown -i ./swagger/swagger.yml** - сгенерировать md-файл |

|  |
| ---- |
| **swagger generate server -A shortener -f ./swagger/swagger.yml** - сгенерировать код |