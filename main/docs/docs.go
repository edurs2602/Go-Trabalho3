// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/clients": {
            "get": {
                "description": "Retorna a lista de clientes.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Clients"
                ],
                "summary": "Lista os clientes.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.Client"
                            }
                        }
                    }
                }
            }
        },
        "/clients/add": {
            "post": {
                "description": "Adiciona um novo cliente.",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Clients"
                ],
                "summary": "Adiciona um novo cliente.",
                "parameters": [
                    {
                        "description": "Novo cliente a ser adicionado",
                        "name": "client",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.Client"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Cliente adicionado com sucesso"
                    }
                }
            }
        },
        "/clients/update/{name}": {
            "put": {
                "description": "Atualiza os detalhes de um cliente.",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Clients"
                ],
                "summary": "Atualiza os detalhes de um cliente.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Nome do cliente a ser atualizado",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Detalhes atualizados do cliente",
                        "name": "client",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.Client"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Cliente atualizado com sucesso"
                    }
                }
            }
        },
        "/pets": {
            "get": {
                "description": "Retorna a lista de pets.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pets"
                ],
                "summary": "Lista os pets.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.Pet"
                            }
                        }
                    }
                }
            }
        },
        "/pets/add": {
            "post": {
                "description": "Adiciona um novo pet.",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Pets"
                ],
                "summary": "Adiciona um novo pet.",
                "parameters": [
                    {
                        "description": "Novo pet a ser adicionado",
                        "name": "pet",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.Pet"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Pet adicionado com sucesso"
                    }
                }
            }
        },
        "/pets/update/{name}": {
            "put": {
                "description": "Atualiza os detalhes de um pet.",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Pets"
                ],
                "summary": "Atualiza os detalhes de um pet.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Nome do pet a ser atualizado",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Detalhes atualizados do pet",
                        "name": "pet",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.Pet"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Pet atualizado com sucesso"
                    }
                }
            }
        }
    },
    "definitions": {
        "main.Client": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "pets": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "main.Pet": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "owner": {
                    "type": "string"
                },
                "weight": {
                    "type": "number"
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
