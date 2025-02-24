package services

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"regexp"
	"sort"
	"strings"
)

type DatabaseService struct {
	connection *pgxpool.Pool
}

func NewDatabaseService(conn *pgxpool.Pool) *DatabaseService {
	return &DatabaseService{
		connection: conn,
	}
}

func validateSQLObjectName(objName string) (string, error) {
	if !regexp.MustCompile(`^[a-zA-Z_\.]+$`).MatchString(objName) {
		return "", fmt.Errorf("invalid method name: %s", objName)
	}
	return objName, nil
}

func createSQLSignature(name string, parameters map[string]interface{}) (string, []interface{}, error) {
	name, err := validateSQLObjectName(name)
	if err != nil {
		return "", []interface{}{}, err
	}
	params := []string{}
	values := []interface{}{}
	paramNames := make([]string, 0, len(parameters))

	for name := range parameters {
		paramNames = append(paramNames, name)
	}
	sort.Strings(paramNames)

	for i, name := range paramNames {
		cleanName := strings.ReplaceAll(name, "*", "_")
		params = append(params, fmt.Sprintf("%s := $%d", cleanName, i+1))
		values = append(values, parameters[name])
	}

	paramsToString := strings.Join(params, ", ")
	return fmt.Sprintf("%s(%s)", name, paramsToString), values, nil
}

func (db *DatabaseService) RunProcedure(ctx context.Context, procedureName string, args map[string]interface{}) error {
	sqlSignature, values, err := createSQLSignature(procedureName, args)
	if err != nil {
		log.Printf("Error creating SQL signature: %s", err)
		return err
	}

	query := fmt.Sprintf("CALL %s", sqlSignature)
	log.Printf("Executing SQL Query: %s", query)
	log.Printf("With Arguments: %v", values)

	_, err = db.connection.Exec(ctx, query, values...)
	if err != nil {
		log.Printf("Error running procedure: %s", err)
		return err
	}
	return nil
}

func (db DatabaseService) RunFunction(ctx context.Context, functionName string, args map[string]interface{}) ([]map[string]interface{}, error) {
	sqlSignature, values, err := createSQLSignature(functionName, args)
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf("SELECT * FROM %s", sqlSignature)
	return db.execute(ctx, query, values...)
}

func (db *DatabaseService) execute(ctx context.Context, query string, args ...interface{}) ([]map[string]interface{}, error) {
	rows, err := db.connection.Query(ctx, query, args...)
	if err != nil {
		log.Printf("Error executing query: %s", err)
		return nil, err
	}
	defer rows.Close()

	var results []map[string]interface{}

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Printf("Error getting row values: %s", err)
			return nil, err
		}

		rowMap := make(map[string]interface{})
		fieldDescriptions := rows.FieldDescriptions()
		for i, field := range fieldDescriptions {
			rowMap[field.Name] = values[i]
		}
		results = append(results, rowMap)
	}

	return results, nil
}
