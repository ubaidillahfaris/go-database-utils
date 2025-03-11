package database

import "fmt"

// GetSQLType mengembalikan tipe data SQL berdasarkan database yang digunakan
func GetSQLType(dbType, columnType string) (string, error) {
	typeMap := map[string]map[string]string{
		"postgres": {
			"int":       "INTEGER",
			"bigint":    "BIGINT",
			"smallint":  "SMALLINT",
			"decimal":   "NUMERIC(10,2)",
			"float":     "REAL",
			"double":    "DOUBLE PRECISION",
			"string":    "VARCHAR(255)",
			"text":      "TEXT",
			"char":      "CHAR(255)",
			"date":      "DATE",
			"timestamp": "TIMESTAMP",
			"uuid":      "UUID",
			"boolean":   "BOOLEAN",
		},
		"mysql": {
			"int":       "INT",
			"bigint":    "BIGINT",
			"smallint":  "SMALLINT",
			"decimal":   "DECIMAL(10,2)",
			"float":     "FLOAT",
			"double":    "DOUBLE",
			"string":    "VARCHAR(255)",
			"text":      "TEXT",
			"char":      "CHAR(255)",
			"date":      "DATE",
			"timestamp": "DATETIME",
			"uuid":      "CHAR(36)",
			"boolean":   "TINYINT(1)",
		},
		"oracle": {
			"int":       "NUMBER(10,0)",
			"bigint":    "NUMBER(19,0)",
			"smallint":  "NUMBER(5,0)",
			"decimal":   "NUMBER(10,2)",
			"float":     "BINARY_FLOAT",
			"double":    "BINARY_DOUBLE",
			"string":    "VARCHAR2(255)",
			"text":      "CLOB",
			"char":      "CHAR(255)",
			"date":      "DATE",
			"timestamp": "TIMESTAMP",
			"uuid":      "RAW(16)",
			"boolean":   "NUMBER(1)",
		},
	}

	dbTypes, ok := typeMap[dbType]
	if !ok {
		return "", fmt.Errorf("database type '%s' is not supported", dbType)
	}

	sqlType, ok := dbTypes[columnType]
	if !ok {
		return "", fmt.Errorf("column type '%s' is not supported for database '%s'", columnType, dbType)
	}

	return sqlType, nil
}
