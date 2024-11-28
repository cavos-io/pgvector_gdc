package config

// GetConfigSchema returns the main configuration schema for the data connector
func GetConfigSchema() map[string]interface{} {
	return map[string]interface{}{
		"type":     "object",
		"nullable": false,
		"properties": map[string]interface{}{
			"db": map[string]interface{}{
				"description": "Name of the database. Omit to use the default database.",
				"type":        "string",
				"nullable":    true,
			},
			"distance_metric": map[string]interface{}{
				"description": "Distance metric for vector similarity (cosine, euclidean, or inner_product).",
				"type":        "string",
				"enum":        []string{"cosine", "euclidean", "inner_product"},
				"nullable":    false,
			},
			"index_type": map[string]interface{}{
				"description": "Indexing strategy for vectors (ivfflat, hnsw, or brute_force).",
				"type":        "string",
				"enum":        []string{"ivfflat", "hnsw", "brute_force"},
				"nullable":    false,
			},
		},
	}
}

// GetOtherSchemas returns an empty map for additional schemas
func GetOtherSchemas() map[string]interface{} {
	return map[string]interface{}{}
}
