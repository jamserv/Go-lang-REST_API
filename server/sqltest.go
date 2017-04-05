package main

// Notice in the import list there's one package prefaced by a ".",
// which allows referencing functions in that package without naming the library in
// the call (if using . "fmt", I can call Println as Println, not fmt.Println)
import (
	"database/sql"
	"flag"
	"fmt"
	"os"
)

const strVERSION string = "0.18 compiled on 8/11/2015"

// sqltest is a small application for demonstrating/testing/learning about SQL database connectivity from Go
func main() {

	// Flags
	ptrVersion := flag.Bool("version", false, "Display program version")
	//ptrDeleteIt := flag.Bool("deletedb", false, "Delete the database")
	ptrServer := flag.String("server", "192.168.2.194", "Server to connect to")
	ptrUser := flag.String("username", "usrbgjobs", "Username for authenticating to database; if you use a backslash, it must be escaped or in quotes")
	ptrPass := flag.String("password", "usrdev2", "Password for database connection")
	//ptrDBName := flag.String("dbname", "PlatformlaRelease", "Database name")

	flag.Parse()

	// Does the user just want the version of the application?
	if *ptrVersion == true {
		fmt.Println("Version " + strVERSION)
		os.Exit(0)
	}

	// Open connection to the database server; this doesn't verify anything until you
	// perform an operation (such as a ping).
	db, err := sql.Open("mssql", "server="+*ptrServer+";user id="+*ptrUser+";password="+*ptrPass)
	if err != nil {
		fmt.Println("From Open() attempt: " + err.Error())
	}

	// When main() is done, this should close the connections
	defer db.Close()

}
