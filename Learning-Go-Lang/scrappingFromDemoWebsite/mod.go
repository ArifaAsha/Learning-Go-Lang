
package main

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver"
	"io"
	"reflect"
)

func main() {
	results, err := runQuery("neo4j+s://demo.neo4jlabs.com:7687", "movies", "mUser", "s3cr3t")
	if err != nil {
		panic(err)
	}
	for _, result := range results {
		fmt.Println(result)
	}
}

func runQuery(uri, database, username, password string) (result []string, err error) {
	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		return nil, err
	}
	defer func() {err = handleClose(driver, err)}()
	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead, DatabaseName: database})
	defer func() {err = handleClose(session, err)}()
	results, err := session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"MATCH (m:Movie {title:$movieTitle})<-[:ACTED_IN]-(a:Person) RETURN a.name as actorName",
			map[string]interface{}{"movieTitle": "The Matrix"}
		)
		if err != nil {
			return nil, err
		}
		var arr []string
		for result.Next() {
			value, found := result.Record().Get("actorName")
			if found {
				arr = append(arr, value.(string))
			}
		}
		if err = result.Err(); err != nil {
			return nil, err
		}
		return arr, nil
	})
	if err != nil {
		return nil, err
	}
	result = results.([]string)
	return result, err
}

func handleClose(closer io.Closer, previousError error) error {
	err := closer.Close()
	if err == nil {
		return previousError
	}
	if previousError == nil {
		return err
	}
	return fmt.Errorf("%v closure error occurred:\n%s\ninitial error was:\n%w", reflect.TypeOf(closer), err.Error(), previousError)
}

