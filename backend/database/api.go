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

	columns, err := rows.Columns()
	if err != nil {
		return fmt.Errorf("error fetching columns: %v", err)
	}

	for rows.Next() {
		elem := reflect.New(elemType).Elem()
		fields := make([]interface{}, len(columns))

		// Temporary storage for nullable fields
		tempValues := make([]interface{}, len(columns))

		// Map each struct field to the correct pointer
		for i := 0; i < elem.NumField(); i++ {
			field := elem.Field(i)
			if field.Kind() == reflect.String {
				// Use a temporary interface{} for nullable strings
				tempValues[i] = new(sql.NullString)
				fields[i] = tempValues[i]
			} else {
				fields[i] = field.Addr().Interface()
			}
		}

		// Scan the row into the temporary storage
		err := rows.Scan(fields...)
		if err != nil {
			return fmt.Errorf("error scanning row: %v", err)
		}

		// Post-process nullable fields
		for i := 0; i < elem.NumField(); i++ {
			field := elem.Field(i)
			if field.Kind() == reflect.String {
				nullStr, ok := tempValues[i].(*sql.NullString)
				if ok && !nullStr.Valid {
					field.SetString("") // Set to empty string or `nil` equivalent
				} else if ok {
					field.SetString(nullStr.String)
				}
			}
		}

		slice = reflect.Append(slice, elem)
	}

	slicePtr.Elem().Set(slice)
	return nil
}

func convertStringToType(name string) interface{} {
	switch name {
	case "Vendor":
		return &[]models.Vendor{}
	case "Meal":
		return &[]models.Meal{}
	case "Rider":
		return &[]models.Rider{}
	case "Customer":
		return &[]models.Customer{}
	case "Discount":
		return &[]models.Discount{}
	case "Orders":
		return &[]models.Orders{}
	case "OrderDetail":
		return &[]models.OrderDetail{}
	case "CustomerSessions":
		return &[]models.CustomerSessions{}
	case "VendorSessions":
		return &[]models.VendorSessions{}
	default:
		return nil
	}
}

/* convert from sql to struct interface */
func GetSQL(db *sql.DB, tableName string) (result interface{}, err error) {
	query := fmt.Sprintf(" SELECT * FROM %s", tableName)
	fmt.Println(query)
	rows, err := db.Query(query)
	if err != nil {
		log.Println("GetTable: Error querying rows:", tableName, err)
		return nil, err
	}
	defer rows.Close()

	result = convertStringToType(tableName)
	if result == nil {
		fmt.Println("GetTable: Unknown table:", tableName)
	}
	if err := scanRowsIntoStruct(rows, result); err != nil {
		log.Println("GetTable: Error scanning rows:", tableName, err)
		return nil, err
	}
	return result, nil
}

/* convert from struct interface to sql table */
func PostSQL(db *sql.DB, tableName string, table interface{}) error {
	t := reflect.TypeOf(table)
	v := reflect.ValueOf(table)

	if t.Kind() != reflect.Struct {
		return fmt.Errorf("table must be a struct")
	}

	var idNames []string
	var placeholders []string
	var idValues []interface{}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldName := field.Tag.Get("db")
		if fieldName == "" {
			fieldName = field.Name
		}

		if !v.Field(i).CanInterface() {
			continue
		}

		fieldValue := v.Field(i).Interface()
		idValues = append(idValues, fieldValue)
		idNames = append(idNames, fieldName)

		// Format the value for SQL
		switch v.Field(i).Kind() {
		case reflect.String:
			placeholders = append(placeholders, fmt.Sprintf("'%s'", fieldValue))
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			placeholders = append(placeholders, fmt.Sprintf("%d", fieldValue))
		case reflect.Bool:
			placeholders = append(placeholders, fmt.Sprintf("%t", fieldValue))
		case reflect.Float32, reflect.Float64:
			placeholders = append(placeholders, fmt.Sprintf("%f", fieldValue))
		default:
			if fieldValue == nil {
				placeholders = append(placeholders, "NULL")
			} else {
				placeholders = append(placeholders, fmt.Sprintf("'%v'", fieldValue))
			}
		}
	}

	if len(idNames) == 0 {
		return fmt.Errorf("no fields available to insert into %s", tableName)
	}

	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		tableName,
		strings.Join(idNames, ", "),
		strings.Join(placeholders, ", "),
	)
	fmt.Println(query)
	_, err := db.Exec(query)
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
