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
        "/create_session": {
            "post": {
                "description": "Создание новой игровой сессии",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json",
                    "application/json"
                ],
                "tags": [
                    "Session"
                ],
                "summary": "Создание игровой сессии",
                "responses": {
                    "200": {
                        "description": "Successful response with session ID",
                        "schema": {
                            "$ref": "#/definitions/main.SuccessResponse"
                        }
                    }
                }
            }
        },
        "/current_games": {
            "get": {
                "description": "Возвращает список текущих игровых сессий, включая информацию о каждой сессии.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Session"
                ],
                "summary": "Получить список текущих игровых сессий",
                "responses": {
                    "200": {
                        "description": "Массив сессий",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.Session"
                            }
                        }
                    }
                }
            }
        },
        "/join_session": {
            "post": {
                "description": "Присоединение к игровой сессии",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Session"
                ],
                "summary": "Присоединение к игровой сессии",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID игровой сессии",
                        "name": "session_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Имя игрока",
                        "name": "player_name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Игрок успешно присоединился к сессии",
                        "schema": {
                            "$ref": "#/definitions/main.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Некорректный запрос или неверный ID сессии",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/leaderboard": {
            "get": {
                "description": "Возвращает рейтинг игроков в текущих игровых сессиях.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Session"
                ],
                "summary": "Получить текущий рейтинг игроков",
                "responses": {
                    "200": {
                        "description": "map[имя_игрока:количество_побед]",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "integer"
                            }
                        }
                    }
                }
            }
        },
        "/play": {
            "post": {
                "description": "Производит игровой ход в рамках сессии, включая выбор игрока и определение победителя в раунде.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Session"
                ],
                "summary": "Играть в рок-ножницы-бумага",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Идентификатор игровой сессии",
                        "name": "session_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Имя игрока",
                        "name": "player_name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Выбор игрока (rock, paper, scissors)",
                        "name": "choice",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Результат игры",
                        "schema": {
                            "$ref": "#/definitions/main.PlayResponse"
                        }
                    },
                    "400": {
                        "description": "Некорректный запрос или неверный идентификатор сессии",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "main.PlayResponse": {
            "type": "object",
            "properties": {
                "game_result": {
                    "type": "string",
                    "example": "Win 1 player"
                }
            }
        },
        "main.Player": {
            "type": "object",
            "properties": {
                "choice": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "main.Scores": {
            "type": "object"
        },
        "main.Session": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "players": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.Player"
                    }
                },
                "round": {
                    "type": "integer"
                },
                "score": {
                    "$ref": "#/definitions/main.Scores"
                }
            }
        },
        "main.SuccessResponse": {
            "type": "object",
            "properties": {
                "session_id": {
                    "type": "integer"
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
