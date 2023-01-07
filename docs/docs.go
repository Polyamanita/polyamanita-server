// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/auths": {
            "post": {
                "description": "Sends an email to the address passed in with a verification code to verify their email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Send a Verification Code",
                "parameters": [
                    {
                        "description": "Email address to send code to",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/routes.PostAuths.AuthEmailInputStruct"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Expiry time of code",
                        "schema": {
                            "$ref": "#/definitions/routes.PostAuths.AuthEmailOutputStruct"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "routes.PostAuths.AuthEmailInputStruct": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
        "routes.PostAuths.AuthEmailOutputStruct": {
            "type": "object",
            "properties": {
                "codeExpiry": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "<some url>",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Polyamanita API",
	Description:      "API for Polyamanita server functions",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
