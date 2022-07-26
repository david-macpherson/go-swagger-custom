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
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "go-swagger custom extension scratch pad",
    "title": "go-swagger custom extension scratch pad",
    "contact": {
      "name": "David Macpherson",
      "email": "david@tensorworks.com.au"
    },
    "version": "0.0.0"
  },
  "basePath": "/",
  "paths": {
    "/{firstName}_{lastName}": {
      "post": {
        "description": "A Place where a user can upload a big file",
        "consumes": [
          "multipart/form-data"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "A Place where a user can upload a big file",
        "operationId": "UploadBigFile",
        "parameters": [
          {
            "type": "string",
            "description": "First Name",
            "name": "firstName",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "Last Name",
            "name": "lastName",
            "in": "path",
            "required": true
          },
          {
            "type": "file",
            "name": "bigFile",
            "in": "formData",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful"
          }
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "go-swagger custom extension scratch pad",
    "title": "go-swagger custom extension scratch pad",
    "contact": {
      "name": "David Macpherson",
      "email": "david@tensorworks.com.au"
    },
    "version": "0.0.0"
  },
  "basePath": "/",
  "paths": {
    "/{firstName}_{lastName}": {
      "post": {
        "description": "A Place where a user can upload a big file",
        "consumes": [
          "multipart/form-data"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "A Place where a user can upload a big file",
        "operationId": "UploadBigFile",
        "parameters": [
          {
            "type": "string",
            "description": "First Name",
            "name": "firstName",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "Last Name",
            "name": "lastName",
            "in": "path",
            "required": true
          },
          {
            "type": "file",
            "name": "bigFile",
            "in": "formData",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful"
          }
        }
      }
    }
  }
}`))
}
