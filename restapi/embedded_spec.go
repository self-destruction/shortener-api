// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is a simple API for Shortener URL",
    "title": "Shortener API",
    "contact": {
      "email": "syrchikov_max@mail.ru"
    },
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "host": "localhost:8080",
  "basePath": "/api/v1",
  "paths": {
    "/shorten_urls/{hash}": {
      "get": {
        "tags": [
          "Link"
        ],
        "summary": "Get link for redirect",
        "operationId": "getLink",
        "parameters": [
          {
            "type": "string",
            "description": "Hash of short url",
            "name": "hash",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "301": {
            "$ref": "#/responses/Redirect"
          },
          "404": {
            "$ref": "#/responses/NotFound"
          }
        }
      }
    },
    "/users": {
      "post": {
        "tags": [
          "User"
        ],
        "summary": "User registrations",
        "operationId": "createUser",
        "parameters": [
          {
            "description": "User registration",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "required": [
                "username",
                "password",
                "email"
              ],
              "properties": {
                "email": {
                  "type": "string",
                  "maxLength": 100,
                  "example": "simple@example.com"
                },
                "password": {
                  "type": "string",
                  "maxLength": 60,
                  "minLength": 5,
                  "example": "simple_password"
                },
                "username": {
                  "type": "string",
                  "maxLength": 60,
                  "minLength": 5,
                  "example": "simple_username"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "$ref": "#/responses/User"
          },
          "400": {
            "$ref": "#/responses/BadRequest"
          }
        }
      }
    },
    "/users/login": {
      "post": {
        "tags": [
          "User"
        ],
        "summary": "Logs user into the system",
        "operationId": "loginUser",
        "parameters": [
          {
            "description": "User login",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "required": [
                "username",
                "password"
              ],
              "properties": {
                "password": {
                  "type": "string",
                  "example": "simple_password"
                },
                "username": {
                  "type": "string",
                  "example": "simple_username"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "$ref": "#/responses/User"
          },
          "400": {
            "$ref": "#/responses/BadRequest"
          }
        }
      }
    },
    "/users/logout": {
      "get": {
        "security": [
          {
            "BasicAuth": []
          }
        ],
        "tags": [
          "User"
        ],
        "summary": "Logs out current logged in user session",
        "operationId": "logoutUser",
        "responses": {
          "204": {
            "$ref": "#/responses/NoContent"
          },
          "401": {
            "$ref": "#/responses/Unauthorized"
          }
        }
      }
    },
    "/users/me": {
      "get": {
        "security": [
          {
            "BasicAuth": []
          }
        ],
        "tags": [
          "Statistic"
        ],
        "summary": "Info about current user",
        "operationId": "getCurrentUser",
        "responses": {
          "200": {
            "$ref": "#/responses/User"
          },
          "401": {
            "$ref": "#/responses/Unauthorized"
          }
        }
      }
    },
    "/users/me/shorten_urls": {
      "get": {
        "security": [
          {
            "BasicAuth": []
          }
        ],
        "tags": [
          "Link"
        ],
        "summary": "Get all short links for this user",
        "operationId": "getLinks",
        "responses": {
          "200": {
            "$ref": "#/responses/Links"
          },
          "401": {
            "$ref": "#/responses/Unauthorized"
          }
        }
      },
      "post": {
        "security": [
          {
            "BasicAuth": []
          }
        ],
        "tags": [
          "Link"
        ],
        "summary": "Creating a new short link",
        "operationId": "createShortLink",
        "parameters": [
          {
            "description": "Full link",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "required": [
                "url"
              ],
              "properties": {
                "url": {
                  "description": "Full link to some site",
                  "type": "string",
                  "example": "https://ya.ru"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "$ref": "#/responses/Link"
          },
          "400": {
            "$ref": "#/responses/BadRequest"
          },
          "401": {
            "$ref": "#/responses/Unauthorized"
          }
        }
      }
    },
    "/users/me/shorten_urls/{hash}": {
      "get": {
        "security": [
          {
            "BasicAuth": []
          }
        ],
        "tags": [
          "Statistic"
        ],
        "summary": "Get information about a specific short user link",
        "operationId": "getLinkInfo",
        "parameters": [
          {
            "type": "string",
            "description": "Hash of short url",
            "name": "hash",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "$ref": "#/responses/Transitions"
          },
          "400": {
            "$ref": "#/responses/BadRequest"
          },
          "404": {
            "$ref": "#/responses/NotFound"
          }
        }
      },
      "delete": {
        "security": [
          {
            "BasicAuth": []
          }
        ],
        "tags": [
          "Link"
        ],
        "summary": "Remove short link",
        "operationId": "removeLink",
        "parameters": [
          {
            "type": "string",
            "description": "Hash of short url",
            "name": "hash",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "$ref": "#/responses/NoContent"
          },
          "401": {
            "$ref": "#/responses/Unauthorized"
          },
          "404": {
            "$ref": "#/responses/NotFound"
          }
        }
      }
    },
    "/users/me/shorten_urls/{hash}/referers": {
      "get": {
        "security": [
          {
            "BasicAuth": []
          }
        ],
        "tags": [
          "Statistic"
        ],
        "summary": "Get top from 20 referring sites",
        "operationId": "getReferers",
        "parameters": [
          {
            "type": "string",
            "description": "Hash of short url",
            "name": "hash",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "$ref": "#/responses/TopReferers"
          },
          "401": {
            "$ref": "#/responses/Unauthorized"
          },
          "404": {
            "$ref": "#/responses/NotFound"
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "required": [
        "code"
      ],
      "properties": {
        "code": {
          "type": "integer"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "Link": {
      "type": "object",
      "required": [
        "id",
        "shortUrl",
        "fullUrl"
      ],
      "properties": {
        "clicks": {
          "description": "Clicks this link",
          "type": "integer",
          "format": "int64",
          "default": 0,
          "example": 532
        },
        "dateCreated": {
          "type": "string",
          "format": "string",
          "example": "2016-08-29 09:12:33"
        },
        "fullUrl": {
          "description": "Full link",
          "type": "string",
          "example": "https://ya.ru"
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "minimum": 1,
          "example": 2
        },
        "shortUrl": {
          "description": "Short url",
          "type": "string",
          "example": "gD4AyQ"
        },
        "userId": {
          "type": "integer",
          "format": "int64",
          "minimum": 1,
          "example": 1
        }
      }
    },
    "Transition": {
      "type": "object",
      "required": [
        "id",
        "link"
      ],
      "properties": {
        "dateCreated": {
          "type": "string",
          "format": "date-time",
          "example": "2016-08-29T09:12:33"
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "minimum": 1,
          "example": 3
        },
        "link": {
          "$ref": "#/definitions/Link"
        },
        "referer": {
          "description": "Where the transition came from",
          "type": "string",
          "example": "https://vk.com"
        }
      }
    },
    "User": {
      "type": "object",
      "required": [
        "id",
        "username",
        "hash",
        "email",
        "dateCreated"
      ],
      "properties": {
        "dateCreated": {
          "type": "string",
          "example": "2016-08-29 09:12:33"
        },
        "email": {
          "type": "string",
          "example": "simple@example.com"
        },
        "hash": {
          "type": "string",
          "example": "$2a$04$Nx.6xqukBA2T5oC3NeSxqOsvulltyjIyyLnd0TYGyvcNCFtxfsJl."
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "minimum": 1,
          "example": 1
        },
        "language": {
          "type": "string",
          "example": "ru_RU"
        },
        "timezone": {
          "type": "string",
          "example": "Europe/Moscow"
        },
        "username": {
          "type": "string",
          "example": "simple_username"
        }
      }
    }
  },
  "responses": {
    "BadRequest": {
      "description": "Bad Request",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    },
    "Link": {
      "description": "Get link",
      "schema": {
        "$ref": "#/definitions/Link"
      }
    },
    "Links": {
      "description": "Get array of links",
      "schema": {
        "type": "object",
        "required": [
          "links"
        ],
        "properties": {
          "links": {
            "description": "Array of links",
            "type": "array",
            "items": {
              "$ref": "#/definitions/Link"
            }
          }
        }
      }
    },
    "NoContent": {
      "description": "NoContent"
    },
    "NotFound": {
      "description": "The specified resource was not found",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    },
    "Redirect": {
      "description": "Link redirect"
    },
    "TopReferers": {
      "description": "Successful operation",
      "schema": {
        "type": "object",
        "required": [
          "top"
        ],
        "properties": {
          "top": {
            "type": "array",
            "items": {
              "type": "object",
              "required": [
                "id",
                "referer",
                "count"
              ],
              "properties": {
                "count": {
                  "type": "integer",
                  "minimum": 1,
                  "example": 539
                },
                "id": {
                  "type": "integer",
                  "minimum": 1,
                  "example": 1
                },
                "referer": {
                  "type": "string",
                  "example": "https://vk.com"
                }
              }
            }
          }
        }
      }
    },
    "Transitions": {
      "description": "Transitions for link",
      "schema": {
        "type": "object",
        "required": [
          "link",
          "transitions"
        ],
        "properties": {
          "link": {
            "$ref": "#/definitions/Link"
          },
          "transitions": {
            "description": "Array of transitions",
            "type": "array",
            "items": {
              "type": "object",
              "required": [
                "date",
                "count"
              ],
              "properties": {
                "count": {
                  "description": "Count of transitions for date",
                  "type": "integer",
                  "format": "Int64",
                  "example": 539
                },
                "date": {
                  "description": "Date of transition",
                  "type": "string",
                  "format": "date-time",
                  "example": "2016-08-29T09:12:39"
                }
              }
            }
          }
        }
      }
    },
    "Unauthorized": {
      "description": "Authentication information is missing or invalid",
      "schema": {
        "$ref": "#/definitions/Error"
      },
      "headers": {
        "WWW_Authenticate": {
          "type": "string"
        }
      }
    },
    "User": {
      "description": "Get user",
      "schema": {
        "$ref": "#/definitions/User"
      }
    }
  },
  "securityDefinitions": {
    "BasicAuth": {
      "type": "basic"
    }
  },
  "tags": [
    {
      "description": "Operations with users",
      "name": "User"
    },
    {
      "description": "Operations with user statistic",
      "name": "Statistic"
    },
    {
      "description": "Operations with links",
      "name": "Link"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is a simple API for Shortener URL",
    "title": "Shortener API",
    "contact": {
      "email": "syrchikov_max@mail.ru"
    },
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "host": "localhost:8080",
  "basePath": "/api/v1",
  "paths": {
    "/shorten_urls/{hash}": {
      "get": {
        "tags": [
          "Link"
        ],
        "summary": "Get link for redirect",
        "operationId": "getLink",
        "parameters": [
          {
            "type": "string",
            "description": "Hash of short url",
            "name": "hash",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "301": {
            "description": "Link redirect"
          },
          "404": {
            "description": "The specified resource was not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/users": {
      "post": {
        "tags": [
          "User"
        ],
        "summary": "User registrations",
        "operationId": "createUser",
        "parameters": [
          {
            "description": "User registration",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "required": [
                "username",
                "password",
                "email"
              ],
              "properties": {
                "email": {
                  "type": "string",
                  "maxLength": 100,
                  "example": "simple@example.com"
                },
                "password": {
                  "type": "string",
                  "maxLength": 60,
                  "minLength": 5,
                  "example": "simple_password"
                },
                "username": {
                  "type": "string",
                  "maxLength": 60,
                  "minLength": 5,
                  "example": "simple_username"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Get user",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/users/login": {
      "post": {
        "tags": [
          "User"
        ],
        "summary": "Logs user into the system",
        "operationId": "loginUser",
        "parameters": [
          {
            "description": "User login",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "required": [
                "username",
                "password"
              ],
              "properties": {
                "password": {
                  "type": "string",
                  "example": "simple_password"
                },
                "username": {
                  "type": "string",
                  "example": "simple_username"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Get user",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/users/logout": {
      "get": {
        "security": [
          {
            "BasicAuth": []
          }
        ],
        "tags": [
          "User"
        ],
        "summary": "Logs out current logged in user session",
        "operationId": "logoutUser",
        "responses": {
          "204": {
            "description": "NoContent"
          },
          "401": {
            "description": "Authentication information is missing or invalid",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "headers": {
              "WWW_Authenticate": {
                "type": "string"
              }
            }
          }
        }
      }
    },
    "/users/me": {
      "get": {
        "security": [
          {
            "BasicAuth": []
          }
        ],
        "tags": [
          "Statistic"
        ],
        "summary": "Info about current user",
        "operationId": "getCurrentUser",
        "responses": {
          "200": {
            "description": "Get user",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "401": {
            "description": "Authentication information is missing or invalid",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "headers": {
              "WWW_Authenticate": {
                "type": "string"
              }
            }
          }
        }
      }
    },
    "/users/me/shorten_urls": {
      "get": {
        "security": [
          {
            "BasicAuth": []
          }
        ],
        "tags": [
          "Link"
        ],
        "summary": "Get all short links for this user",
        "operationId": "getLinks",
        "responses": {
          "200": {
            "description": "Get array of links",
            "schema": {
              "type": "object",
              "required": [
                "links"
              ],
              "properties": {
                "links": {
                  "description": "Array of links",
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/Link"
                  }
                }
              }
            }
          },
          "401": {
            "description": "Authentication information is missing or invalid",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "headers": {
              "WWW_Authenticate": {
                "type": "string"
              }
            }
          }
        }
      },
      "post": {
        "security": [
          {
            "BasicAuth": []
          }
        ],
        "tags": [
          "Link"
        ],
        "summary": "Creating a new short link",
        "operationId": "createShortLink",
        "parameters": [
          {
            "description": "Full link",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "required": [
                "url"
              ],
              "properties": {
                "url": {
                  "description": "Full link to some site",
                  "type": "string",
                  "example": "https://ya.ru"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Get link",
            "schema": {
              "$ref": "#/definitions/Link"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "401": {
            "description": "Authentication information is missing or invalid",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "headers": {
              "WWW_Authenticate": {
                "type": "string"
              }
            }
          }
        }
      }
    },
    "/users/me/shorten_urls/{hash}": {
      "get": {
        "security": [
          {
            "BasicAuth": []
          }
        ],
        "tags": [
          "Statistic"
        ],
        "summary": "Get information about a specific short user link",
        "operationId": "getLinkInfo",
        "parameters": [
          {
            "type": "string",
            "description": "Hash of short url",
            "name": "hash",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Transitions for link",
            "schema": {
              "type": "object",
              "required": [
                "link",
                "transitions"
              ],
              "properties": {
                "link": {
                  "$ref": "#/definitions/Link"
                },
                "transitions": {
                  "description": "Array of transitions",
                  "type": "array",
                  "items": {
                    "type": "object",
                    "required": [
                      "date",
                      "count"
                    ],
                    "properties": {
                      "count": {
                        "description": "Count of transitions for date",
                        "type": "integer",
                        "format": "Int64",
                        "minimum": 0,
                        "example": 539
                      },
                      "date": {
                        "description": "Date of transition",
                        "type": "string",
                        "format": "date-time",
                        "example": "2016-08-29T09:12:39"
                      }
                    }
                  }
                }
              }
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "The specified resource was not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "delete": {
        "security": [
          {
            "BasicAuth": []
          }
        ],
        "tags": [
          "Link"
        ],
        "summary": "Remove short link",
        "operationId": "removeLink",
        "parameters": [
          {
            "type": "string",
            "description": "Hash of short url",
            "name": "hash",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "NoContent"
          },
          "401": {
            "description": "Authentication information is missing or invalid",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "headers": {
              "WWW_Authenticate": {
                "type": "string"
              }
            }
          },
          "404": {
            "description": "The specified resource was not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/users/me/shorten_urls/{hash}/referers": {
      "get": {
        "security": [
          {
            "BasicAuth": []
          }
        ],
        "tags": [
          "Statistic"
        ],
        "summary": "Get top from 20 referring sites",
        "operationId": "getReferers",
        "parameters": [
          {
            "type": "string",
            "description": "Hash of short url",
            "name": "hash",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "type": "object",
              "required": [
                "top"
              ],
              "properties": {
                "top": {
                  "type": "array",
                  "items": {
                    "type": "object",
                    "required": [
                      "id",
                      "referer",
                      "count"
                    ],
                    "properties": {
                      "count": {
                        "type": "integer",
                        "minimum": 1,
                        "example": 539
                      },
                      "id": {
                        "type": "integer",
                        "minimum": 1,
                        "example": 1
                      },
                      "referer": {
                        "type": "string",
                        "example": "https://vk.com"
                      }
                    }
                  }
                }
              }
            }
          },
          "401": {
            "description": "Authentication information is missing or invalid",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "headers": {
              "WWW_Authenticate": {
                "type": "string"
              }
            }
          },
          "404": {
            "description": "The specified resource was not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "required": [
        "code"
      ],
      "properties": {
        "code": {
          "type": "integer"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "Link": {
      "type": "object",
      "required": [
        "id",
        "shortUrl",
        "fullUrl"
      ],
      "properties": {
        "clicks": {
          "description": "Clicks this link",
          "type": "integer",
          "format": "int64",
          "default": 0,
          "minimum": 0,
          "example": 532
        },
        "dateCreated": {
          "type": "string",
          "format": "string",
          "example": "2016-08-29 09:12:33"
        },
        "fullUrl": {
          "description": "Full link",
          "type": "string",
          "example": "https://ya.ru"
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "minimum": 1,
          "example": 2
        },
        "shortUrl": {
          "description": "Short url",
          "type": "string",
          "example": "gD4AyQ"
        },
        "userId": {
          "type": "integer",
          "format": "int64",
          "minimum": 1,
          "example": 1
        }
      }
    },
    "Transition": {
      "type": "object",
      "required": [
        "id",
        "link"
      ],
      "properties": {
        "dateCreated": {
          "type": "string",
          "format": "date-time",
          "example": "2016-08-29T09:12:33"
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "minimum": 1,
          "example": 3
        },
        "link": {
          "$ref": "#/definitions/Link"
        },
        "referer": {
          "description": "Where the transition came from",
          "type": "string",
          "example": "https://vk.com"
        }
      }
    },
    "User": {
      "type": "object",
      "required": [
        "id",
        "username",
        "hash",
        "email",
        "dateCreated"
      ],
      "properties": {
        "dateCreated": {
          "type": "string",
          "example": "2016-08-29 09:12:33"
        },
        "email": {
          "type": "string",
          "example": "simple@example.com"
        },
        "hash": {
          "type": "string",
          "example": "$2a$04$Nx.6xqukBA2T5oC3NeSxqOsvulltyjIyyLnd0TYGyvcNCFtxfsJl."
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "minimum": 1,
          "example": 1
        },
        "language": {
          "type": "string",
          "example": "ru_RU"
        },
        "timezone": {
          "type": "string",
          "example": "Europe/Moscow"
        },
        "username": {
          "type": "string",
          "example": "simple_username"
        }
      }
    }
  },
  "responses": {
    "BadRequest": {
      "description": "Bad Request",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    },
    "Link": {
      "description": "Get link",
      "schema": {
        "$ref": "#/definitions/Link"
      }
    },
    "Links": {
      "description": "Get array of links",
      "schema": {
        "type": "object",
        "required": [
          "links"
        ],
        "properties": {
          "links": {
            "description": "Array of links",
            "type": "array",
            "items": {
              "$ref": "#/definitions/Link"
            }
          }
        }
      }
    },
    "NoContent": {
      "description": "NoContent"
    },
    "NotFound": {
      "description": "The specified resource was not found",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    },
    "Redirect": {
      "description": "Link redirect"
    },
    "TopReferers": {
      "description": "Successful operation",
      "schema": {
        "type": "object",
        "required": [
          "top"
        ],
        "properties": {
          "top": {
            "type": "array",
            "items": {
              "type": "object",
              "required": [
                "id",
                "referer",
                "count"
              ],
              "properties": {
                "count": {
                  "type": "integer",
                  "minimum": 1,
                  "example": 539
                },
                "id": {
                  "type": "integer",
                  "minimum": 1,
                  "example": 1
                },
                "referer": {
                  "type": "string",
                  "example": "https://vk.com"
                }
              }
            }
          }
        }
      }
    },
    "Transitions": {
      "description": "Transitions for link",
      "schema": {
        "type": "object",
        "required": [
          "link",
          "transitions"
        ],
        "properties": {
          "link": {
            "$ref": "#/definitions/Link"
          },
          "transitions": {
            "description": "Array of transitions",
            "type": "array",
            "items": {
              "type": "object",
              "required": [
                "date",
                "count"
              ],
              "properties": {
                "count": {
                  "description": "Count of transitions for date",
                  "type": "integer",
                  "format": "Int64",
                  "minimum": 0,
                  "example": 539
                },
                "date": {
                  "description": "Date of transition",
                  "type": "string",
                  "format": "date-time",
                  "example": "2016-08-29T09:12:39"
                }
              }
            }
          }
        }
      }
    },
    "Unauthorized": {
      "description": "Authentication information is missing or invalid",
      "schema": {
        "$ref": "#/definitions/Error"
      },
      "headers": {
        "WWW_Authenticate": {
          "type": "string"
        }
      }
    },
    "User": {
      "description": "Get user",
      "schema": {
        "$ref": "#/definitions/User"
      }
    }
  },
  "securityDefinitions": {
    "BasicAuth": {
      "type": "basic"
    }
  },
  "tags": [
    {
      "description": "Operations with users",
      "name": "User"
    },
    {
      "description": "Operations with user statistic",
      "name": "Statistic"
    },
    {
      "description": "Operations with links",
      "name": "Link"
    }
  ]
}`))
}
