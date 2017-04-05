
package main

import (
	"flag"
	"fmt"
	"database/sql"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
)

var debug = flag.Bool("debug", false, "enable debugging")
var password = flag.String("password", "usrdev2", "usrdev2")
var port *int = flag.Int("port", 1433, "the database port")
var server = flag.String("server", "192.168.2.194", "192.168.2.194")
var user = flag.String("user", "usrbgjobs", "usrbgjobs")
var dbname = flag.String("dbname", "PlatformlaRelease\\VMSQL2005QAM", "PlatformlaRelease")
//var connStr = fmt.Sprintf("server=%s;Initial Catalog=MySchema;userid=%s;password=%s;port=%d", *server, *user, *password, *port)

func main() {
	simplemain()
}

func simplemain() {

		flag.Parse() // parse the command line args

	if *debug {
		fmt.Printf(" password:%s\n", *password)
		fmt.Printf(" port:%d\n", *port)
		fmt.Printf(" server:%s\n", *server)
		fmt.Printf(" user:%s\n", *user)
	}

	connString := fmt.Sprintf("server=%s;Initial Catalog=PlatformlaRelease\\VMSQL2005QAM;user id=%s;password=%s;port=%d", *server, *user, *password, *port)
	if *debug {
		fmt.Printf(" connString:%s\n", connString)
	}

	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}else {
		_, err = conn.Exec("USE PlatformlaRelease")
	}

	defer conn.Close()

	//stmt, err := conn.Prepare("select 1, 'abc'")
	stmt, err := conn.Query("SELECT id FROM dbo.BulkPaymentDetail")
	if err != nil {
		log.Fatal("Query failed:", err.Error())
	}
	defer stmt.Close()

	/*
	row := stmt.QueryRow()

	var somenumber int64
	var somechars string
	err = row.Scan(&somenumber, &somechars)
	if err != nil {
		log.Fatal("Scan failed:", err.Error())
	}
	fmt.Printf("somenumber:%d\n", somenumber)
	fmt.Printf("somechars:%s\n", somechars)
	*/

	fmt.Printf("bye\n")


}
