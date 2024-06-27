package database

import "strings"

func stateQuery(typeName string, states []string) string {
	return `DO $$ BEGIN IF NOT EXISTS ` + "(SELECT 1 FROM pg_type WHERE typname ='" + typeName +
		"') THEN CREATE TYPE " + typeName + " AS ENUM ('" + strings.Join(states, "', '") + "');" +
		" END IF; END $$;"
}
