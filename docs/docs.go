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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "跑路了",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/comments/orders/{order_id}": {
            "post": {
                "description": "用户为指定id的订单添加评价",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "评价管理"
                ],
                "summary": "用户评价订单",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/orders/{order_id}": {
            "get": {
                "description": "获取指定id的订单详情",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "订单管理"
                ],
                "summary": "获取订单详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "订单id",
                        "name": "order_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OrderDetailResponse",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "ErrorResponse",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "取消指定id的订单",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "订单管理"
                ],
                "summary": "取消订单",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "订单id",
                        "name": "order_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "succeed",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "ErrorResponse",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/riders/orders/{order_id}/accept": {
            "post": {
                "description": "骑手接收指定id的订单",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "骑手管理"
                ],
                "summary": "骑手接单",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "订单id",
                        "name": "order_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "接单成功",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "ErrorResponse",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/riders/orders/{order_id}/complete": {
            "post": {
                "description": "骑手完成指定id的订单",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "骑手管理"
                ],
                "summary": "骑手完成订单",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "订单id",
                        "name": "order_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "订单完成",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "ErrorResponse",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/dish/add": {
            "post": {
                "description": "添加菜品",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "菜品管理"
                ],
                "summary": "添加菜品",
                "parameters": [
                    {
                        "type": "string",
                        "description": "菜品名称",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "菜品描述",
                        "name": "description",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "菜品图片",
                        "name": "picture",
                        "in": "formData"
                    },
                    {
                        "type": "number",
                        "description": "菜品价格",
                        "name": "price",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "所属餐厅id",
                        "name": "merchantID",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "添加成功",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "添加失败",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "401": {
                        "description": "输入非法",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/dish/fuzzy": {
            "post": {
                "description": "根据菜品名称模糊搜索菜品信息",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "菜品管理"
                ],
                "summary": "模糊搜索菜品信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "菜品名称",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Dish",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/dish/list": {
            "get": {
                "description": "获取所有菜品列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "菜品管理"
                ],
                "summary": "列出所有菜品",
                "responses": {
                    "200": {
                        "description": "获取成功",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "500": {
                        "description": "ErrorResponse",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/dish/perfect": {
            "post": {
                "description": "根据菜品名称准确获取菜品信息",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "菜品管理"
                ],
                "summary": "准确获取菜品信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "菜品名称",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "所属餐厅id",
                        "name": "merchantID",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Dish",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "ErrorResponse",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "500": {
                        "description": "ErrorResponse",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/merchant/add": {
            "post": {
                "description": "添加商家",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "商家管理"
                ],
                "summary": "添加商家",
                "parameters": [
                    {
                        "type": "string",
                        "description": "餐厅名称",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "餐厅地址",
                        "name": "address",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "餐厅电话",
                        "name": "phone",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "添加成功",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "添加失败",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "401": {
                        "description": "输入非法",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/merchant/fuzzy": {
            "get": {
                "description": "根据商家名称模糊搜索商家信息",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "商家管理"
                ],
                "summary": "模糊搜索商家信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "商家名称",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Merchant",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/merchant/list": {
            "get": {
                "description": "获取所有商家列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "商家管理"
                ],
                "summary": "列出所有商家",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "500": {
                        "description": "ErrorResponse",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/merchant/perfect": {
            "get": {
                "description": "根据商家名称获取商家信息",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "商家管理"
                ],
                "summary": "准确获取商家信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "商家名称",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Merchant",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "ErrorResponse",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "500": {
                        "description": "ErrorResponse",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/orders": {
            "get": {
                "description": "获取所有订单列表",
                "produces": [
                    "application/json"
                ],
                "summary": "列出所有订单",
                "responses": {
                    "200": {
                        "description": "succeed",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "500": {
                        "description": "ErrorResponse",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "用户创建新订单",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "创建订单",
                "parameters": [
                    {
                        "description": "新订单信息",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Order",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "ErrorResponse",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "500": {
                        "description": "ErrorResponse",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/users": {
            "post": {
                "description": "创建新用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "创建用户",
                "parameters": [
                    {
                        "description": "新用户信息",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "ErrorResponse",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "500": {
                        "description": "ErrorResponse",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/users/login": {
            "post": {
                "description": "用户使用用户名和密码登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "用户登录信息",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "ErrorResponse",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "500": {
                        "description": "ErrorResponse",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/users/logout": {
            "post": {
                "description": "用户登出",
                "produces": [
                    "application/json"
                ],
                "summary": "用户登出",
                "responses": {
                    "204": {
                        "description": "退出成功",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "500": {
                        "description": "ErrorResponse",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Response": {
            "type": "object",
            "properties": {
                "data": {
                    "description": "响应数据"
                },
                "msg": {
                    "description": "响应消息",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:9090",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "饿了么项目复刻",
	Description:      "在东软的教学包基础上去做拓展",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
