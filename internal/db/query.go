package db

import "github.com/pgvector/pgvector-go"

// FindNearestVectors queries for nearest vectors using pgvector
func FindNearestVectors(vector pgvector.Vector, topK int) ([]map[string]interface{}, error) {
	query := `
		SELECT id, embedding
		FROM vectors
		ORDER BY embedding <-> $1
		LIMIT $2;
	`

	rows, err := DB.Query(query, vector, topK)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []map[string]interface{}
	for rows.Next() {
		var id int
		var embedding pgvector.Vector
		if err := rows.Scan(&id, &embedding); err != nil {
			return nil, err
		}

		results = append(results, map[string]interface{}{
			"id":        id,
			"embedding": embedding,
		})
	}

	return results, nil
}
