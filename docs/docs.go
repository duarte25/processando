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
        "/climate": {
            "get": {
                "description": "Retorna uma lista com os valores de clima e na base do parâmetro e com os valores de total de acidentes, total de óbitos, total de envolvidos e feridos.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "climate"
                ],
                "summary": "Listar tabelas de clima (Climate)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Identificador dos dados a serem recuperados (ex.: data_climate_2022)",
                        "name": "data",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Lista de guarda corpo",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/day_week": {
            "get": {
                "description": "Retorna uma lista com os valores de dia da semana e na base do parâmetro e com os valores de total de acidentes, total de óbitos, total de envolvidos e feridos.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "day_week"
                ],
                "summary": "Listar tabelas de dia da semana (Day Week)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Identificador dos dados a serem recuperados (ex.: data_day_week_2022)",
                        "name": "data",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Lista de guarda corpo",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/guardrail": {
            "get": {
                "description": "Retorna uma lista com os valores de guarda-corpo e na base do parâmetro e com os valores de total de acidentes, total de óbitos, total de envolvidos e feridos.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "guardrail"
                ],
                "summary": "Listar tabelas de guarda-corpo (Guardrail)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Identificador dos dados a serem recuperados (ex.: data_guardrail_2022)",
                        "name": "data",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Lista de guarda corpo",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/highway": {
            "get": {
                "description": "Retorna uma lista com os valores de via rodoviária e na base do parâmetro e com os valores de total de acidentes, total de óbitos, total de envolvidos e feridos.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "highway"
                ],
                "summary": "Listar tabelas de via rodoviária (Highway)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Identificador dos dados a serem recuperados (ex.: data_highway_2022)",
                        "name": "data",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Lista de canteiro central",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/median": {
            "get": {
                "description": "Retorna uma lista com os valores de canteiro central e na base do parâmetro e com os valores de total de acidentes, total de óbitos, total de envolvidos e feridos.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "median"
                ],
                "summary": "Listar tabelas de canteiro central (Median)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Identificador dos dados a serem recuperados (ex.: data_median_2022)",
                        "name": "data",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Lista de canteiro central",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/month": {
            "get": {
                "description": "Retorna uma lista com os meses do ano e na base do parâmetro e com os valores de total de acidentes, total de óbitos, total de envolvidos e feridos.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "month"
                ],
                "summary": "Listar tabelas de mês do ano (Month)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Identificador dos dados a serem recuperados (ex.: data_month_2022)",
                        "name": "data",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Lista de mês do ano",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/phase_day": {
            "get": {
                "description": "Retorna uma lista com as fase do dia e na base do parâmetro e com os valores de total de acidentes, total de óbitos, total de envolvidos e feridos.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "phase_day"
                ],
                "summary": "Listar tabelas fase do dia (Phase Day)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Identificador dos dados a serem recuperados (ex.: data_phase_day_2022)",
                        "name": "data",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Lista de fase do dia",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/shoulder": {
            "get": {
                "description": "Retorna uma lista de acostamento com base no parâmetro e com os valores de total de acidentes, total de óbitos, total de envolvidos e feridos.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "shoulder"
                ],
                "summary": "Listar tabelas de acostamento (Shoulder)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Identificador dos dados a serem recuperados (ex.: data_shoulder_2022)",
                        "name": "data",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Lista de acostamento",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/speed": {
            "get": {
                "description": "Retorna uma lista de limite de velocidade com base no parâmetro e com os valores de total de acidentes, total de óbitos, total de envolvidos e feridos.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "speed"
                ],
                "summary": "Listar tabelas de velocidade (Speed)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Identificador dos dados a serem recuperados (ex.: data_speed_2022)",
                        "name": "data",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Lista de limite de velocidade da via",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/susp_alcohol": {
            "get": {
                "description": "Retorna uma lista de suspeito de álcool com base no parâmetro e com os valores de total de acidentes, total de óbitos, total de envolvidos e feridos.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "susp_alcohol"
                ],
                "summary": "Listar tabelas suspeito de álcool",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Identificador dos dados a serem recuperados (ex.: data_susp_alcohol_2022)",
                        "name": "data",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Lista de suspeito de álcool",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/uf": {
            "get": {
                "description": "Retorna uma lista de UFs com base no parâmetro e com os valores de total de acidentes, total de óbitos, total de envolvidos e feridos.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UF"
                ],
                "summary": "Listar unidades federativas (UFs)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Identificador dos dados a serem recuperados (ex.: data_uf_2022)",
                        "name": "data",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Lista de UFs",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
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
