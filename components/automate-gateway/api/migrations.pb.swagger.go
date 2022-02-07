package api

func init() {
	Swagger.Add("migrations", `{
  "swagger": "2.0",
  "info": {
    "title": "external/infra_proxy/migrations/migrations.proto",
    "version": "version not set"
  },
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/api/v0/infra/servers/migrations/status/{migration_id}": {
      "get": {
        "operationId": "InfraProxyMigration_GetMigrationStatus",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.migrations.response.GetMigrationStatus"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "migration_id",
            "description": "Migration ID.",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "InfraProxyMigration"
        ]
      }
    },
    "/api/v0/infra/servers/{server_id}/migrations/cancel_migration/{migration_id}": {
      "get": {
        "operationId": "InfraProxyMigration_CancelMigration",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.migrations.response.CancelMigrationResponce"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Server ID",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "migration_id",
            "description": "Migration ID",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "InfraProxyMigration"
        ]
      }
    }
  },
  "definitions": {
    "chef.automate.api.infra_proxy.migrations.response.CancelMigrationResponce": {
      "type": "object",
      "properties": {
        "success": {
          "type": "boolean",
          "format": "boolean",
          "title": "migration ID"
        },
        "errors": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "title": "Error message"
        }
      }
    },
    "chef.automate.api.infra_proxy.migrations.response.GetMigrationStatus": {
      "type": "object",
      "properties": {
        "migration_id": {
          "type": "string",
          "description": "Migration ID."
        },
        "migration_type": {
          "type": "string",
          "title": "Migration type"
        },
        "migration_status": {
          "type": "string",
          "title": "Migration status"
        }
      }
    },
    "google.protobuf.Any": {
      "type": "object",
      "properties": {
        "type_url": {
          "type": "string"
        },
        "value": {
          "type": "string",
          "format": "byte"
        }
      }
    },
    "grpc.gateway.runtime.Error": {
      "type": "object",
      "properties": {
        "error": {
          "type": "string"
        },
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        },
        "details": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/google.protobuf.Any"
          }
        }
      }
    }
  }
}
`)
}