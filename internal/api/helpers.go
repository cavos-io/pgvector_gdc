package api

func mapPostgresTypeToHasuraType(pgType string) string {
	switch pgType {
	case "character varying", "text", "uuid":
		return "string"
	case "integer", "bigint", "numeric", "real", "double precision":
		return "number"
	case "boolean":
		return "boolean"
	case "vector":
		return "string"
	case "timestamp without time zone", "timestamp with time zone":
		return "timestamp"
	default:
		return "unknown"
	}
}
