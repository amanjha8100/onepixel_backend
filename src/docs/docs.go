// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://github.com/championswimmer/onepixel_backend",
        "contact": {
            "name": "Arnav Gupta",
            "email": "dev@championswimmer.in"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/urls": {
            "get": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Get all urls",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "urls"
                ],
                "summary": "Get all urls",
                "responses": {
                    "200": {
                        "description": "GetAllUsers",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Create random short url",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "urls"
                ],
                "summary": "Create random short url",
                "operationId": "create-random-url",
                "parameters": [
                    {
                        "description": "Url",
                        "name": "url",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.CreateUrlRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dtos.UrlResponse"
                        }
                    },
                    "400": {
                        "description": "The request body is not valid",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "long_url is required to create url",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/urls/{shortcode}": {
            "put": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Create specific short url",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "urls"
                ],
                "summary": "Create specific short url",
                "operationId": "create-specific-url",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Shortcode",
                        "name": "shortcode",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Url",
                        "name": "url",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.CreateUrlRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dtos.UrlResponse"
                        }
                    },
                    "400": {
                        "description": "The request body is not valid",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Shortcode is not allowed",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorResponse"
                        }
                    },
                    "409": {
                        "description": "Shortcode already exists",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "long_url is required to create url",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/users": {
            "post": {
                "security": [
                    {
                        "APIKeyAuth": []
                    }
                ],
                "description": "Register new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Register new user",
                "operationId": "register-user",
                "parameters": [
                    {
                        "description": "User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dtos.UserResponse"
                        }
                    },
                    "400": {
                        "description": "The request body is not valid",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorResponse"
                        }
                    },
                    "409": {
                        "description": "User with this email already exists",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "email and password are required to create user",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/users/:userid": {
            "get": {
                "description": "Get user info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get user info",
                "operationId": "get-user-info",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "patch": {
                "description": "Update user info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Update user info",
                "operationId": "update-user-info",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/users/login": {
            "post": {
                "security": [
                    {
                        "APIKeyAuth": []
                    }
                ],
                "description": "Login user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Login user",
                "operationId": "login-user",
                "parameters": [
                    {
                        "description": "User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.LoginUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.UserResponse"
                        }
                    },
                    "401": {
                        "description": "Invalid email or password",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dtos.CreateUrlRequest": {
            "type": "object",
            "properties": {
                "long_url": {
                    "type": "string"
                }
            }
        },
        "dtos.CreateUserRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "dtos.ErrorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Something went wrong"
                },
                "status": {
                    "type": "integer",
                    "example": 400
                }
            }
        },
        "dtos.LoginUserRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "dtos.UrlResponse": {
            "type": "object",
            "properties": {
                "long_url": {
                    "type": "string",
                    "example": "https://www.google.com"
                },
                "short_url": {
                    "type": "string",
                    "example": "nhg145"
                }
            }
        },
        "dtos.UserResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "user@test.com"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "token": {
                    "type": "string",
                    "example": "\u003cJWT_TOKEN\u003e"
                }
            }
        }
    },
    "securityDefinitions": {
        "APIKeyAuth": {
            "type": "apiKey",
            "name": "X-API-Key",
            "in": "header"
        },
        "BearerToken": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1",
	Host:             "api.onepixel.link",
	BasePath:         "/api/v1",
	Schemes:          []string{"http", "https"},
	Title:            "onepixel API",
	Description:      "1px.li URL Shortner API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
