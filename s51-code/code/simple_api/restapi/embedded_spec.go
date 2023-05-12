// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/network.golang.demoapi-measurement.v1+json"
  ],
  "produces": [
    "application/network.golang.demoapi-list.v1+json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Return sensor readings from database",
    "title": "Sensor readings",
    "version": "1.0.0"
  },
  "paths": {
    "/measurements": {
      "get": {
        "tags": [
          "measurements"
        ],
        "responses": {
          "200": {
            "description": "List of measurements",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/measurement"
              }
            }
          }
        }
      }
    }
  },
  "definitions": {
    "measurement": {
      "type": "object",
      "required": [
        "timeTaken",
        "sensor",
        "value"
      ],
      "properties": {
        "sensor": {
          "type": "string",
          "minLength": 8
        },
        "timeTaken": {
          "type": "string",
          "format": "date-time"
        },
        "value": {
          "type": "string"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/network.golang.demoapi-measurement.v1+json"
  ],
  "produces": [
    "application/network.golang.demoapi-list.v1+json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Return sensor readings from database",
    "title": "Sensor readings",
    "version": "1.0.0"
  },
  "paths": {
    "/measurements": {
      "get": {
        "tags": [
          "measurements"
        ],
        "responses": {
          "200": {
            "description": "List of measurements",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/measurement"
              }
            }
          }
        }
      }
    }
  },
  "definitions": {
    "measurement": {
      "type": "object",
      "required": [
        "timeTaken",
        "sensor",
        "value"
      ],
      "properties": {
        "sensor": {
          "type": "string",
          "minLength": 8
        },
        "timeTaken": {
          "type": "string",
          "format": "date-time"
        },
        "value": {
          "type": "string"
        }
      }
    }
  }
}`))
}
