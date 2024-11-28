package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/cavos-io/pgvector_gdc/internal/db"
)

// SchemaHandler handles the `/schema` endpoint.
func SchemaHandler(w http.ResponseWriter, r *http.Request) {
	client := db.DB

	tables, err := fetchTables(client)
	if err != nil {
		http.Error(w, "Failed to fetch tables", http.StatusInternalServerError)
		return
	}

	// Build the response
	response := map[string]interface{}{
		"tables": tables,
	}

	// Return the schema response as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func fetchTables(client *sql.DB) ([]map[string]interface{}, error) {
	tablesQuery := `
		SELECT table_name
		FROM information_schema.tables
		WHERE table_schema = 'public';
	`

	tablesRows, err := client.Query(tablesQuery)
	if err != nil {
		return nil, err
	}
	defer tablesRows.Close()

	tables := []map[string]interface{}{}
	for tablesRows.Next() {
		var tableName string
		if err := tablesRows.Scan(&tableName); err != nil {
			return nil, err
		}

		columns, err := fetchColumns(client, tableName)
		if err != nil {
			return nil, err
		}

		tables = append(tables, map[string]interface{}{
			"name":        []string{"public", tableName},
			"type":        "table",
			"columns":     columns,
			"primary_key": []string{},
		})
	}

	return tables, nil
}

func fetchColumns(client *sql.DB, tableName string) ([]map[string]interface{}, error) {
	columnsQuery := `
		SELECT column_name, data_type, is_nullable
		FROM information_schema.columns
		WHERE table_name = $1;
	`

	columnsRows, err := client.Query(columnsQuery, tableName)
	if err != nil {
		return nil, err
	}
	defer columnsRows.Close()

	columns := []map[string]interface{}{}
	for columnsRows.Next() {
		var columnName, dataType, isNullable string
		if err := columnsRows.Scan(&columnName, &dataType, &isNullable); err != nil {
			return nil, err
		}

		columns = append(columns, map[string]interface{}{
			"name":     columnName,
			"type":     mapPostgresTypeToHasuraType(dataType),
			"nullable": isNullable == "YES",
		})
	}

	return columns, nil
}
