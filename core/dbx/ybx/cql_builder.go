package ybx

import (
	"encoding/json"
	"strconv"
)

func AnysToStrings(anys []any) (strings []string, err error) {
	for _, an := range anys {
		// type check
		switch v := an.(type) {
		case string:
			strings = append(strings, v)
		case int:
			strings = append(strings, strconv.Itoa(v))
		case int64:
			strings = append(strings, strconv.FormatInt(v, 10))
		case float64:
			strings = append(strings, strconv.FormatFloat(v, 'f', 6, 64))
		case bool:
			strings = append(strings, strconv.FormatBool(v))
		default:
			// convert to json string
			bytes, err := json.Marshal(v)
			if err != nil {
				return strings, err
			}
			strings = append(strings, string(bytes))
		}

	}
	return strings, err
}

func BuildInsertQuery(keyspace string, table string, columns []string, values []any) (string, error) {
	var vals []string
	var err error
	if vals, err = AnysToStrings(values); err != nil {
		return "", err
	}
	var insertStmt string = "INSERT INTO " + keyspace + "." + table + "("
	for i, column := range columns {
		if i == len(columns)-1 {
			insertStmt += column + ")"
		} else {
			insertStmt += column + ", "
		}
	}
	insertStmt += " VALUES ("
	for i, value := range vals {
		if i == len(vals)-1 {
			insertStmt += value + ")"
		} else {
			insertStmt += value + ", "
		}
	}
	return insertStmt, nil
}
