package queries

import (
	"log/slog"
)

type IQueryBuilder interface {
	getQuery() Query
	setConditions() string
	setStatement(obj interface{}) error
}

type Query struct {
	Statement string
	Values    []any
}

type QueryBuilder struct {
	builder IQueryBuilder
}

func GetBuilder(builderType string) IQueryBuilder {
	switch builderType {
	case "Create":
		return NewCreateQuery()

	default:
		return NewCreateQuery()
	}
}

func NewBuilder(b IQueryBuilder) *QueryBuilder {
	return &QueryBuilder{
		builder: b,
	}
}

func (b *QueryBuilder) BuildQuery(obj interface{}) Query {
	err := b.builder.setStatement(obj)

	if err != nil {
		slog.Error("Unable to build the query", "Error", err.Error())
	}

	return b.builder.getQuery()
}
