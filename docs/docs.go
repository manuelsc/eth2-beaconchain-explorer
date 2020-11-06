// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/healthz": {
            "get": {
                "description": "Health endpoint for montitoring if the explorer is in sync",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Health"
                ],
                "summary": "Health of the explorer",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/block/{slotOrHash}": {
            "get": {
                "description": "Returns a block by its slot or root hash",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Block"
                ],
                "summary": "Get block",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Block slot or root hash or the string latest",
                        "name": "slotOrHash",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/block/{slot}/attestations": {
            "get": {
                "description": "Returns the attestations included in a specific block",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Block"
                ],
                "summary": "Get the attestations included in a specific block",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Block slot",
                        "name": "slot",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/block/{slot}/attesterslashings": {
            "get": {
                "description": "Returns the attester slashings included in a specific block",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Block"
                ],
                "summary": "Get the attester slashings included in a specific block",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Block slot",
                        "name": "slot",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/block/{slot}/deposits": {
            "get": {
                "description": "Returns the deposits included in a specific block",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Block"
                ],
                "summary": "Get the deposits included in a specific block",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Block slot",
                        "name": "slot",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/block/{slot}/proposerslashings": {
            "get": {
                "description": "Returns the proposer slashings included in a specific block",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Block"
                ],
                "summary": "Get the proposer slashings included in a specific block",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Block slot",
                        "name": "slot",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/block/{slot}/voluntaryexits": {
            "get": {
                "description": "Returns the voluntary exits included in a specific block",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Block"
                ],
                "summary": "Get the voluntary exits included in a specific block",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Block slot",
                        "name": "slot",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/chart/{chart}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Charts"
                ],
                "summary": "Returns charts from the page https://beaconcha.in/charts as PNG",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Chart name (see https://github.com/gobitfly/eth2-beaconchain-explorer/blob/master/services/charts_updater.go#L20 for all available names)",
                        "name": "chart",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/epoch/{epoch}": {
            "get": {
                "description": "Returns information for a specified epoch by the epoch number or the latest epoch",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Epoch"
                ],
                "summary": "Get epoch by number",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Epoch number or the string latest",
                        "name": "epoch",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/epoch/{epoch}/blocks": {
            "get": {
                "description": "Returns all blocks for a specified epoch",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Epoch"
                ],
                "summary": "Get epoch blocks by epoch number",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Epoch number or the string latest",
                        "name": "epoch",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/eth1deposit/{txhash}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Eth1"
                ],
                "summary": "Get an eth1 deposit by its eth1 transaction hash",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Eth1 transaction hash",
                        "name": "txhash",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/user/dashboard/save": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Register or update your mobile notification token",
                "parameters": [
                    {
                        "description": "Your device` + "`" + `s firebase notification token",
                        "name": "token",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.ApiResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/types.ApiResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/types.ApiResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/user/mobile/settings": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get your device settings, currently only whether to enable mobile notifcations or not",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/types.ApiResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/types.MobileSettingsData"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/types.ApiResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/types.ApiResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Changing your devices mobile settings",
                "parameters": [
                    {
                        "description": "Whether to enable mobile notifications for this device or not",
                        "name": "notify_enabled",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "boolean"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/types.ApiResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/types.MobileSettingsData"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/types.ApiResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/types.ApiResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/user/token": {
            "post": {
                "security": [
                    {
                        "OAuthAccessCode": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Exchange your oauth code for an access token or refresh your access token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "grant_type use authorization_code for oauth code or refresh_token if you wish to refresh an token",
                        "name": "grant_type",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Only required when using authorization_code grant type. Code received via oauth redirect_uri",
                        "name": "code",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Only required when using authorization_code grant type. Must match the redirect_uri from your oauth flow.",
                        "name": "redirect_uri",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Only required when using refresh_token grant type. The refresh_token you received during authorization_code flow.",
                        "name": "refresh_token",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.OAuthResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.OAuthErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.OAuthErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/user/validator/saved": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get all your tagged validators",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/types.ApiResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/types.MinimalTaggedValidators"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/types.ApiResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/types.ApiResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/user/validator/{pubkey}/add": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "subscribes a user to get notifications from a specific validator",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Public Key of validator you want to subscribe to",
                        "name": "pubKey",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "Submit \\",
                        "name": "balance_decreases",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Submit \\",
                        "name": "validator_slashed",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.ApiResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/types.ApiResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/types.ApiResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/user/validator/{pubkey}/remove": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "unsubscribes a user from a specific validator",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Public Key of validator you want to subscribe to",
                        "name": "pubKey",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.ApiResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/types.ApiResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/types.ApiResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/validator/eth1/{address}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Validator"
                ],
                "summary": "Get all validators that belong to an eth1 address",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Eth1 address from which the validator deposits were sent",
                        "name": "eth1address",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/validator/leaderboard": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Validator"
                ],
                "summary": "Get the current top 100 performing validators (using the income over the last 7 days)",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/validator/{indexOrPubkey}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Validator"
                ],
                "summary": "Get up to 100 validators by their index",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Up to 100 validator indicesOrPubkeys, comma separated",
                        "name": "indexOrPubkey",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/validator/{indexOrPubkey}/attestations": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Validator"
                ],
                "summary": "Get all attestations during the last 100 epochs for up to 100 validators",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Up to 100 validator indicesOrPubkeys, comma separated",
                        "name": "indexOrPubkey",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/validator/{indexOrPubkey}/balancehistory": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Validator"
                ],
                "summary": "Get the balance history (last 100 epochs) of up to 100 validators",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Up to 100 validator indicesOrPubkeys, comma separated",
                        "name": "indexOrPubkey",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/validator/{indexOrPubkey}/deposits": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Validator"
                ],
                "summary": "Get all eth1 deposits for up to 100 validators",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Up to 100 validator indicesOrPubkeys, comma separated",
                        "name": "indexOrPubkey",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/validator/{indexOrPubkey}/performance": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Validator"
                ],
                "summary": "Get the current performance of up to 100 validators",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Up to 100 validator indicesOrPubkeys, comma separated",
                        "name": "indexOrPubkey",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/validator/{indexOrPubkey}/proposals": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Validator"
                ],
                "summary": "Get all proposed blocks during the last 100 epochs for up to 100 validators",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Up to 100 validator indicesOrPubkeys, comma separated",
                        "name": "indexOrPubkey",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "types.ApiResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "types.MinimalTaggedValidators": {
            "type": "object",
            "properties": {
                "index": {
                    "type": "integer"
                },
                "pubKey": {
                    "type": "string"
                }
            }
        },
        "types.MobileSettingsData": {
            "type": "object",
            "properties": {
                "notify_token": {
                    "type": "string"
                }
            }
        },
        "utils.OAuthErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "error_description": {
                    "type": "string"
                }
            }
        },
        "utils.OAuthResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "expires_in": {
                    "type": "integer"
                },
                "refresh_token": {
                    "type": "string"
                },
                "token_type": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        },
        "OAuthAccessCode": {
            "type": "oauth2",
            "flow": "accessCode",
            "authorizationUrl": "https://beaconcha.in/user/authorize",
            "tokenUrl": "https://beaconcha.in/user/token"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Beaconcha.in ETH2 API",
	Description: "High performance API for querying information from the Ethereum 2.0 beacon chain\nThe API is currently free to use. A fair use policy applies. Calls are rate limited to\n10 requests / 1 minute / IP. All API results are cached for 1 minute.\nIf you required a higher usage plan please checkout https://beaconcha.in/api/pricing.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
