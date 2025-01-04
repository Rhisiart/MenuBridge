package queries

import (
	"fmt"
	"reflect"
	"strings"
)

type CreateQuery struct {
	Query
}

func NewCreateQuery() *CreateQuery {
	return &CreateQuery{}
}

func (c *CreateQuery) setStatement(obj interface{}) error {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	if t.Kind() != reflect.Struct {
		return fmt.Errorf("the variable is not a struct")
	}

	columns := t.Field(0).Tag.Get("db")
	placeholders := fmt.Sprintf("$%d", 1)
	var values []interface{}

	for i := 1; i < v.NumField(); i++ {
		columns = strings.Join([]string{columns, t.Field(i).Tag.Get("db")}, ", ")
		placeholders = strings.Join([]string{placeholders, fmt.Sprintf("$%d", i+1)}, ", ")
		values = append(values, v.Field(i).Interface())
	}

	c.Statement = fmt.Sprintf(
		"INSERT INTO %s (%s) Values (%s)",
		t.Name(),
		columns,
		placeholders)
	c.Values = values

	return nil
}

func (c *CreateQuery) setConditions() string {
	return ""
}

func (c *CreateQuery) getQuery() Query {
	return Query{
		Statement: c.Statement,
		Values:    c.Values,
	}
}
