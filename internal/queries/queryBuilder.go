package queries

import (
	"log/slog"
)

type IQueryBuilder interface {
	getQuery() Query
	setStatement(obj interface{}) error
	setFilter(filters ...map[string]interface{})
}

type Query struct {
	Statement string
	Filter    string
	Values    []any
}

type QueryBuilder struct {
	builder IQueryBuilder
	filters []map[string]interface{}
}

func NewBuilder(builderType string, f []map[string]interface{}) *QueryBuilder {
	return &QueryBuilder{
		builder: GetBuilder(builderType),
		filters: f,
	}
}

func GetBuilder(builderType string) IQueryBuilder {
	switch builderType {
	case "Create":
		return NewCreateQuery()
	case "Update":
		return NewUpdateQuery()
	default:
		return NewCreateQuery()
	}
}

func (b *QueryBuilder) SetBuilder(builderType string) {
	b.builder = GetBuilder(builderType)
}

func (b *QueryBuilder) BuildQuery(obj interface{}) Query {
	err := b.builder.setStatement(obj)
	b.builder.setFilter(b.filters...)

	if err != nil {
		slog.Error("Unable to build the query", "Error", err.Error())
	}

	return b.builder.getQuery()
}
