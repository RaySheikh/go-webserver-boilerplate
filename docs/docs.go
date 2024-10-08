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
        "/user/{id}": {
            "get": {
                "description": "Get details of a user by their ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Get user by ID",
                "operationId": "get-user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/server.UserResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "server.ErrorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "description": "Error message",
                    "type": "string"
                }
            }
        },
        "server.UserResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "description": "User's email",
                    "type": "string"
                },
                "id": {
                    "description": "User ID",
                    "type": "integer"
                },
                "name": {
                    "description": "User's name",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Go Webserver API",
	Description:      "This is a sample server for a Go web server boilerplate.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
