package creational_patterns

import (
	"fmt"
	"sync"
)

/* The singleton pattern ensures that only one object of a particular class is ever created.
All further references to objects of the singleton class refer to the same underlying instance
*/

type databaseConnector struct{}

var dbConnector *databaseConnector

var once sync.Once

func GetDBConnector() *databaseConnector {
	once.Do(func() {
		fmt.Println("Init the database connector... (this will be printed only one time)")
		dbConnector = new(databaseConnector)
	})
	return dbConnector
}

func spamGetDBConnector() {
	_ = GetDBConnector()
}

func SingletonExample() {
	for i := 0; i < 10; i++ {
		go spamGetDBConnector()
	}
}
