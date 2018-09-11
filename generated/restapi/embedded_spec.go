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
  "swagger": "2.0",
  "info": {
    "title": "Rialto Entity Resolver",
    "version": "0.0.1"
  },
  "paths": {
    "/healthcheck": {
      "get": {
        "produces": [
          "application/json"
        ],
        "operationId": "healthCheck",
        "responses": {
          "200": {
            "description": "The service is working properly",
            "schema": {
              "$ref": "#/definitions/HealthCheckResponse"
            }
          },
          "503": {
            "description": "The service is not working properly",
            "schema": {
              "$ref": "#/definitions/HealthCheckResponse"
            }
          }
        }
      }
    },
    "/person": {
      "get": {
        "produces": [
          "text/plain"
        ],
        "operationId": "findOrCreatePerson",
        "parameters": [
          {
            "type": "string",
            "description": "First name of the person",
            "name": "first_name",
            "in": "query"
          },
          {
            "type": "string",
            "description": "Last name of the person",
            "name": "last_name",
            "in": "query"
          },
          {
            "type": "string",
            "description": "ORCID of the person",
            "name": "orcid",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "returns a URI for the person in Rialto",
            "schema": {
              "description": "contains the actual URI",
              "type": "string"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "HealthCheckResponse": {
      "type": "object",
      "properties": {
        "status": {
          "description": "The status of the service",
          "type": "string"
        }
      },
      "example": {
        "status": "OK"
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "swagger": "2.0",
  "info": {
    "title": "Rialto Entity Resolver",
    "version": "0.0.1"
  },
  "paths": {
    "/healthcheck": {
      "get": {
        "produces": [
          "application/json"
        ],
        "operationId": "healthCheck",
        "responses": {
          "200": {
            "description": "The service is working properly",
            "schema": {
              "$ref": "#/definitions/HealthCheckResponse"
            }
          },
          "503": {
            "description": "The service is not working properly",
            "schema": {
              "$ref": "#/definitions/HealthCheckResponse"
            }
          }
        }
      }
    },
    "/person": {
      "get": {
        "produces": [
          "text/plain"
        ],
        "operationId": "findOrCreatePerson",
        "parameters": [
          {
            "type": "string",
            "description": "First name of the person",
            "name": "first_name",
            "in": "query"
          },
          {
            "type": "string",
            "description": "Last name of the person",
            "name": "last_name",
            "in": "query"
          },
          {
            "type": "string",
            "description": "ORCID of the person",
            "name": "orcid",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "returns a URI for the person in Rialto",
            "schema": {
              "description": "contains the actual URI",
              "type": "string"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "HealthCheckResponse": {
      "type": "object",
      "properties": {
        "status": {
          "description": "The status of the service",
          "type": "string"
        }
      },
      "example": {
        "status": "OK"
      }
    }
  }
}`))
}
