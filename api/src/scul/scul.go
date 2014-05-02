//see https://code.google.com/p/go-wiki/wiki/SQLInterface
//see http://godoc.org/github.com/lib/pq
package scul

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"strconv"
)

// A Database wrapper
type DataB struct {
	Database     string
	Host         string
	User         string
	Password     string
	Databasename string
	Wrappeddb    *sql.DB
}

//Get Database URL
func (config *DataB) Url() string {
	return fmt.Sprintf("%v://%v:%v@%v/%v \n", config.Database, config.User, config.Password, config.Host, config.Databasename)
}

//Query executes a query that returns a single row, typically a SELECT. The args are for any placeholder parameters in the query and scan. queryArgsLen serves as a delimiter to separate query parameters from scan parameters in args.
func (db *DataB) QueryRow(query string, queryArgsLen int, args ...interface{}) error {
	log.Printf("QueryRow of %q", query)

	err := db.Wrappeddb.QueryRow(query, args[0:queryArgsLen]...).Scan(args[queryArgsLen:]...)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No rows. Query of %q", query)
	case err != nil:
		log.Fatal("Query of %q: %v", query, err)
	}
	return err
}

// Exec executes a query without returning any rows. The args are for any placeholder parameters in the query.
func (db *DataB) Exec(query string, args ...interface{}) (sql.Result, error) {
	log.Printf("Exec of %q", query)
	res, err := db.Wrappeddb.Exec(query, args...)
	if err != nil {
		log.Fatal("Exec of %q: %v", query, err)
	}
	return res, err
}

//Query executes a query that returns rows, typically a SELECT. The args are for any placeholder parameters in the query.
func (db *DataB) Query(query string, args ...interface{}) *sql.Rows {
	log.Printf("Query of %q", query)
	rows, err := db.Wrappeddb.Query(query, args...)
	if err != nil {
		log.Fatal("Exec of %q: %v", query, err)
	}
	return rows
}

//Opendb opens a database specified by its database driver name and a driver-specific data source name, usually consisting of at least a database name and connection information.
func (db *DataB) Opendb() *sql.DB {
	log.Printf("Open DB %v", db)
	adb, err := sql.Open("postgres", db.Url())
	if err != nil {
		log.Fatal("Could not open db %v: %v", adb, err)
		return nil
	}
	db.Wrappeddb = adb
	return db.Wrappeddb
}


//PostgreSQL Goodies
// array of HStore is built from an array of Map. The args is an array of map where each map represents an HStore element
func (db *DataB) HStoresToString(hstores []map[string]interface{}) string {
	arraylen := len(hstores)
	array := make([]string, arraylen, arraylen) 
	for i := 0; i < len(hstores); i++ {
		hstore := []string {}
		for key, value := range hstores[i] {
    		hstore = append(hstore, fmt.Sprintf("%s => %v", key, value))
		}
		array[i] = fmt.Sprintf("{ %s }", strings.Join(hstore, ", "))
	}
	return fmt.Sprintf("{ %s }", strings.Join(array, ", ")) 
}

func (db *DataB) StringToHStores(hstores string) []map[string]interface{} {
	hstoresAsMap := []map[string]interface{} {}
	arrayContent := strings.Trim(hstores, "{}")
	//fmt.Printf("arrayAsString : %v \n",arrayAsString)
	for _, hstore := range strings.Split(strings.Trim(arrayContent, "{}"),"},{") {
		//fmt.Printf("Hstore : %v \n",hstore)
		hstoreValue := map[string]interface{} {}
		for _, hstoreEl := range strings.Split(hstore,",") {
			keyValue := strings.SplitN(hstoreEl,"=>",2)
			hstoreValue[CleanHStoreString(keyValue[0])] = CleanHStoreAny(keyValue[1])
		}
		hstoresAsMap = append(hstoresAsMap,hstoreValue)
	} 
	return hstoresAsMap
}


func CleanHStoreString(s string) string {
	var v string
	v = strings.Replace(s, "\"", "", -1)
	v = strings.Replace(v, "\\", "", -1)
	return v
}

func CleanHStoreAny(s string) interface{} {
	v := CleanHStoreString(s)
	if vInt, err := strconv.ParseInt(v,10, 64); err == nil {
		return vInt
	}
	if vFloat, err := strconv.ParseFloat(v,64); err == nil {
		return vFloat
	}	
	return v
}






