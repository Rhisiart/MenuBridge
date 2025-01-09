package queries

import (
	"fmt"
	"reflect"
	"strings"
)

type UpdateQuery struct {
	Query
}

func NewUpdateQuery() *UpdateQuery {
	return &UpdateQuery{}
}

func (c *UpdateQuery) setStatement(obj interface{}) error {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	} else if t.Kind() != reflect.Struct {
		return fmt.Errorf("the variable must be a pointer to a struct")
	}

	var columnsAndPlaceholders string
	var values []interface{}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		tableName := fieldType.Tag.Get("db")

		if tableName == "" ||
			!field.CanInterface() ||
			reflect.DeepEqual(
				field.Interface(),
				reflect.Zero(field.Type()).Interface()) {
			continue
		}

		expression := strings.Join(
			[]string{tableName, fmt.Sprintf("$%d", i+1)},
			" = ")

		if columnsAndPlaceholders == "" {
			columnsAndPlaceholders = expression
		} else {
			columnsAndPlaceholders = strings.Join(
				[]string{columnsAndPlaceholders, expression},
				", ")
		}

		values = append(values, field.Interface())
	}

	c.Statement = fmt.Sprintf(
		"UPDATE %s SET %s",
		t.Name(),
		columnsAndPlaceholders)
	c.Values = values

	return nil
}

func (c *UpdateQuery) setFilter(filters ...map[string]interface{}) {
	conditions := []string{}
	for _, filter := range filters {
		for columnName, value := range filter {
			placeholder := fmt.Sprintf("$%d", len(c.Values)+1)
			conditions = append(conditions, fmt.Sprintf("%s = %s", columnName, placeholder))
			c.Values = append(c.Values, value)
		}
	}
	if len(conditions) > 0 {
		c.Filter = fmt.Sprintf("WHERE %s", strings.Join(conditions, " AND "))
	} else {
		c.Filter = ""
	}
}

func (c *UpdateQuery) getQuery() Query {
	return Query{
		Statement: c.Statement,
		Filter:    c.Filter,
		Values:    c.Values,
	}
}
