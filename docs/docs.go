// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/test/add": {
            "post": {
                "description": "add test history",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "test"
                ],
                "summary": "add history",
                "operationId": "add",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.addRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.sResponse"
                        }
                    }
                }
            }
        },
        "/test/history": {
            "get": {
                "description": "Show all test history",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "test"
                ],
                "summary": "Show history",
                "operationId": "history",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.historyResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.eResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.Test": {
            "type": "object",
            "properties": {
                "test": {
                    "type": "string",
                    "example": "test"
                }
            }
        },
        "v1.addRequest": {
            "type": "object",
            "properties": {
                "test": {
                    "type": "string",
                    "example": "test"
                }
            }
        },
        "v1.eResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "message"
                }
            }
        },
        "v1.historyResponse": {
            "type": "object",
            "properties": {
                "history": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Test"
                    }
                }
            }
        },
        "v1.sResponse": {
            "type": "object",
            "properties": {
                "success": {
                    "type": "string",
                    "example": "message"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/v1",
	Schemes:          []string{"http", "https"},
	Title:            "swagger test",
	Description:      "swagger test example",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
