// Code generated by swaggo/swag. DO NOT EDIT.

package auto_generate

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
        "/api/linting_rule/{project_id}": {
            "post": {
                "description": "If there is admin grant, allow creating or updating linting rules, otherwise not allowed",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "LintingRule"
                ],
                "summary": "Update linting rules by project id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "unique ID for a project",
                        "name": "project_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "session token",
                        "name": "token",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "custom linting rule in yaml",
                        "name": "{object}",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.RequestForPostByProjectID"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.ResponseProtocol"
                        }
                    }
                }
            }
        },
        "/api/linting_rule/{project_id}/validate": {
            "post": {
                "description": "Validate spec with linting rule for the project",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "LintingRule"
                ],
                "summary": "Validate spec",
                "parameters": [
                    {
                        "type": "string",
                        "description": "unique ID for a project",
                        "name": "project_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "session token",
                        "name": "token",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "Specification to be validated ",
                        "name": "{object}",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.RequestForValidate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "validate result",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.ResponseProtocol"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/rule_engine.ValidationResult"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "errors.ErrorCode": {
            "type": "integer",
            "enum": [
                0,
                400,
                401,
                402,
                500
            ],
            "x-enum-varnames": [
                "OK",
                "BadRequest",
                "Unauthorized",
                "Forbidden",
                "InternalServerError"
            ]
        },
        "handler.RequestForPostByProjectID": {
            "type": "object",
            "properties": {
                "rule": {
                    "type": "string"
                }
            }
        },
        "handler.RequestForValidate": {
            "type": "object",
            "properties": {
                "spec": {
                    "type": "string"
                }
            }
        },
        "handler.ResponseProtocol": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "business code, 0 means ok",
                    "allOf": [
                        {
                            "$ref": "#/definitions/errors.ErrorCode"
                        }
                    ]
                },
                "data": {
                    "description": "result"
                },
                "message": {
                    "description": "error message",
                    "type": "string"
                }
            }
        },
        "rule_engine.ValidationResult": {
            "type": "object",
            "properties": {
                "errors": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "isValid": {
                    "type": "boolean"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}