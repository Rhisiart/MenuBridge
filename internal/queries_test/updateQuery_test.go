package queriestest

import (
	"encoding/json"
	"testing"

	"github.com/Rhisiart/MenuBridge/internal/database"
	"github.com/Rhisiart/MenuBridge/internal/queries"
	"github.com/stretchr/testify/assert"
)

func TestUpdateQuery(t *testing.T) {
	menu := database.NewMenu(1, "Hamburger")
	builder := queries.NewBuilder("Update", []map[string]interface{}{{"Id": 1}})

	query := builder.BuildQuery(menu)
	assert.Equal(t, query.Statement, "UPDATE Menu SET id = $1, name = $2")
	assert.Equal(t, query.Filter, "WHERE Id = $3")
}

func TestUpdateQueryWithJsonMarshal(t *testing.T) {
	menuUnmarshal := &database.Menu{}
	_ = json.Unmarshal([]uint8{123, 34, 105, 100, 34, 58, 49, 44, 34, 110, 97, 109, 101, 34, 58, 34, 72, 97, 109, 98, 117, 114, 103, 101, 114, 34, 125}, menuUnmarshal)

	builder := queries.NewBuilder("Update", []map[string]interface{}{{"Id": 1}})

	query := builder.BuildQuery(menuUnmarshal)
	assert.Equal(t, query.Statement, "UPDATE Menu SET id = $1, name = $2")
	assert.Equal(t, query.Filter, "WHERE Id = $3")
}
