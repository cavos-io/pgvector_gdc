package api

import (
	"encoding/json"
	"net/http"

	"github.com/cavos-io/pgvector_gdc/config"
)

// CapabilitiesHandler handles the `/capabilities` endpoint.
func CapabilitiesHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"version": "1.0",
		"capabilities": map[string]interface{}{
			"queries": map[string]interface{}{
				"supports_primary_key":  true,
				"supports_foreign_keys": true,
				"supports_order_by":     true,
				"supports_filters":      true,
				"supports_limits":       true,
				"supports_offset":       true,
			},
			"mutations": map[string]interface{}{
				"supports_insert": false,
				"supports_update": false,
				"supports_delete": false,
			},
			"subscriptions": map[string]interface{}{},
			"scalar_types": map[string]interface{}{
				"vector": map[string]interface{}{
					"comparison_operators": map[string]string{
						"vector_similarity": "Float",
					},
					"aggregate_functions": map[string]string{
						"avg_vector": "String",
						"sum_vector": "String",
						"norm":       "Float",
					},
					"graphql_type": "String",
				},
			},
			"data_schema": map[string]interface{}{
				"supports_primary_keys": true,
				"supports_foreign_keys": true,
				"column_nullability":    "nullable_and_non_nullable",
			},
			"user_defined_functions": map[string]interface{}{},
		},
		"config_schemas": map[string]interface{}{
			"config_schema": config.GetConfigSchema(),
			"other_schemas": config.GetOtherSchemas(),
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
