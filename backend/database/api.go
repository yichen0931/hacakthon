package database

import (
	"database/sql"
	"fmt"
	"hackathon/models"
	"log"
	"reflect"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func scanRowsIntoStruct(rows *sql.Rows, result interface{}) error {
	slicePtr := reflect.ValueOf(result)
	if slicePtr.Kind() != reflect.Ptr || slicePtr.Elem().Kind() != reflect.Slice {
		return fmt.Errorf("result argument must be a pointer to a slice")
	}

	elemType := slicePtr.Elem().Type().Elem()
	slice := reflect.MakeSlice(slicePtr.Elem().Type(), 0, 0)

	for rows.Next() {
		elem := reflect.New(elemType).Elem()
		fields := make([]interface{}, elem.NumField())
		for i := 0; i < elem.NumField(); i++ {
			fields[i] = elem.Field(i).Addr().Interface()
		}

		err := rows.Scan(fields...)
		if err != nil {
			return fmt.Errorf("error scanning row: %v", err)
		}

		slice = reflect.Append(slice, elem)
	}

	slicePtr.Elem().Set(slice)
	return nil
}

/* convert from sql to struct interface */
func GetSQL(db *sql.DB, tableName string) (result interface{}, err error) {
	query := fmt.Sprintf(" SELECT * FROM %s", tableName)
	rows, err := db.Query(query)
	if err != nil {
		log.Println("GetTable: Error querying rows:", tableName, err)
		return nil, err
	}
	defer rows.Close()

	switch tableName {
	case "Vendor":
		var vendors []models.Vendor
		if err := scanRowsIntoStruct(rows, &vendors); err != nil {
			log.Println("GetTable: Error scanning rows:", tableName, err)
			return nil, err
		}
		result = vendors
	default:
		log.Println("GetTable: Invalid table:", tableName)
		return nil, err
	}
	return result, nil
}

/* convert from struct interface to sql table */
func PostSQL(db *sql.DB, tableName string, table interface{}) error {
	// Get the type and value of the struct
	t := reflect.TypeOf(table)
	v := reflect.ValueOf(table)
	if t.Kind() != reflect.Struct {
		return fmt.Errorf("table must be a struct")
	}

	var fieldNames []string
	var placeholders []string
	var fieldValues []interface{}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldNames = append(fieldNames, field.Name)
		placeholders = append(placeholders, "?")
		fieldValues = append(fieldValues, v.Field(i).Interface())
	}

	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		tableName,
		strings.Join(fieldNames, ", "),
		strings.Join(placeholders, ", "),
	)
	_, err := db.Exec(query, fieldValues...)
	if err != nil {
		return fmt.Errorf("error executing query: %w", err)
	}

	return nil
}

func PutSQL(db *sql.DB, tableName string, table interface{}, idField string, idValue interface{}) error {
	t := reflect.TypeOf(table)
	v := reflect.ValueOf(table)
	if t.Kind() != reflect.Struct {
		return fmt.Errorf("table must be a struct")
	}

	var updates []string
	var fieldValues []interface{}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldName := field.Name
		fieldValue := v.Field(i).Interface()

		// Skip the ID field (it shouldn't be updated)
		if strings.EqualFold(fieldName, idField) {
			continue
		}
		updates = append(updates, fmt.Sprintf("%s = ?", fieldName))
		fieldValues = append(fieldValues, fieldValue)
	}
	fieldValues = append(fieldValues, idValue)

	query := fmt.Sprintf(
		"UPDATE %s SET %s WHERE %s = ?",
		tableName,
		strings.Join(updates, ", "),
		idField,
	)
	_, err := db.Exec(query, fieldValues...)
	if err != nil {
		return fmt.Errorf("error executing query: %w", err)
	}

	return nil
}

func DeleteSQL(db *sql.DB, tableName string) error {
	_, err := db.Exec(fmt.Sprintf("DELETE FROM %s", tableName))
	if err != nil {
		return fmt.Errorf("error executing query: %w", err)
	}
	return nil
}
