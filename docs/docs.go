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
        "/auth/authorize": {
            "get": {
                "description": "Authorize user access token after login",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Authorize user",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/users.SuccessAuthorizeToken"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/users.UnauthorizedToken"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "description": "Login a new user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "Login user",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/users.LoginDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/users.SuccessLoginResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/users.BadRequestLoginResponse"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "Regist a new user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Register user",
                "parameters": [
                    {
                        "description": "Register user",
                        "name": "register",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/users.CreateUserDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/users.SuccessRegisterResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/users.BadRequestRegisterResponse"
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "description": "Responds if the APIs success.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "health"
                ],
                "summary": "Check health APIs",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/health.HealthSuccess"
                        }
                    }
                }
            }
        },
        "/product-category": {
            "get": {
                "description": "Get all product category",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product-category"
                ],
                "summary": "Get all product category",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Number of page to load",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Number page size or limit to display data",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/product_category.SuccessGetAllProductCategoryResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create new product category",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product-category"
                ],
                "summary": "Create product category",
                "parameters": [
                    {
                        "description": "Create product category",
                        "name": "product-category",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/product_category.CreateProductCategoryDTO"
                        }
                    },
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/product_category.SuccessCreateProductCategoryResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/product_category.BadRequestCreateProductCategoryResponse"
                        }
                    }
                }
            }
        },
        "/role": {
            "get": {
                "description": "Get all role that exist",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "role"
                ],
                "summary": "Get all role",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Number of page to load",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Number page size or limit to display role data",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Role"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create new role",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "role"
                ],
                "summary": "Create role",
                "parameters": [
                    {
                        "description": "Create role",
                        "name": "role",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/role.CreateRoleDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Role"
                        }
                    }
                }
            }
        },
        "/role/{id}": {
            "put": {
                "description": "Update role by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "role"
                ],
                "summary": "Update role by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Role id to update",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update role",
                        "name": "role",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/role.UpdateRoleDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Role"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete role by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "role"
                ],
                "summary": "Delete role by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Role id to delete",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Role"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "health.HealthSuccess": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "models.Role": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string",
                    "example": "2024-07-16T23:13:03.115483+07:00"
                },
                "deleted_at": {
                    "type": "string",
                    "example": "null"
                },
                "id": {
                    "type": "string",
                    "example": "ab7ac1cb-17c6-4e9a-8cd8-d51d8988c5ec"
                },
                "name": {
                    "type": "string",
                    "example": "Admin"
                },
                "updated_at": {
                    "type": "string",
                    "example": "2024-07-16T23:13:03.115483+07:00"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string",
                    "example": "2024-07-16T23:13:03.115483+07:00"
                },
                "deleted_at": {
                    "type": "string",
                    "example": "null"
                },
                "id": {
                    "type": "string",
                    "example": "ab7ac1cb-17c6-4e9a-8cd8-d51d8988c5ec"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "role": {
                    "$ref": "#/definitions/models.Role"
                },
                "role_id": {
                    "type": "string"
                },
                "salt": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string",
                    "example": "2024-07-16T23:13:03.115483+07:00"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "product_category.BadRequestCreateProductCategoryResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "error message"
                }
            }
        },
        "product_category.CreateProductCategoryDTO": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "example": "fashion"
                }
            }
        },
        "product_category.ProductCategories": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string",
                    "example": "2024-07-16T23:13:03.115483+07:00"
                },
                "deleted_at": {
                    "type": "string",
                    "example": "null"
                },
                "id": {
                    "type": "string",
                    "example": "ab7ac1cb-17c6-4e9a-8cd8-d51d8988c5ec"
                },
                "name": {
                    "type": "string",
                    "example": "electronic"
                },
                "updated_at": {
                    "type": "string",
                    "example": "2024-07-16T23:13:03.115483+07:00"
                }
            }
        },
        "product_category.SuccessCreateProductCategoryResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/product_category.ProductCategories"
                },
                "message": {
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "product_category.SuccessGetAllProductCategoryResponse": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer",
                    "example": 1
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/product_category.ProductCategories"
                    }
                },
                "message": {
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "role.CreateRoleDTO": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string",
                    "example": "Admin"
                }
            }
        },
        "role.UpdateRoleDTO": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string",
                    "example": "Shop"
                }
            }
        },
        "users.BadRequestLoginResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "username not exist"
                }
            }
        },
        "users.BadRequestRegisterResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "username is exist"
                }
            }
        },
        "users.CreateUserDTO": {
            "type": "object",
            "required": [
                "name",
                "password",
                "role_id",
                "username"
            ],
            "properties": {
                "name": {
                    "type": "string",
                    "example": "Admin Shop"
                },
                "password": {
                    "type": "string",
                    "example": "adminshop123"
                },
                "role_id": {
                    "type": "string",
                    "example": "8cb17fe3-3552-4bfa-852f-a9a7d91fd00d"
                },
                "username": {
                    "type": "string",
                    "example": "adminshop"
                }
            }
        },
        "users.LoginDTO": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "example": "adminshop123"
                },
                "username": {
                    "type": "string",
                    "example": "adminshop"
                }
            }
        },
        "users.SuccessAuthorizeToken": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "token is valid"
                }
            }
        },
        "users.SuccessLoginResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/users.TokenResponse"
                },
                "message": {
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "users.SuccessRegisterResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/models.User"
                },
                "message": {
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "users.TokenResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
                }
            }
        },
        "users.UnauthorizedToken": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "unauthorized"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:5000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Online Shop",
	Description:      "A Online Shop Backend Service",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
