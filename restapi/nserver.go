package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"fmt"
	"os"

	"github.com/gorilla/mux"
	"flag"
	"database/sql"
	_"github.com/denisenkom/go-mssqldb"
	"github.com/BurntSushi/toml"
)

type Person struct {
	ID        string   "json:'id,omitempty'"
	Firstname string   "json:'firstname,omitempty'"
	Lastname  string   "json:'lastname,omitempty'"
	Address   *Address "json:'address,omitempty'"
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people = make(map[string]Person)

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	person, _ := people[params["id"]]

	// person might be empty if it wasn't found in the map
	json.NewEncoder(w).Encode(person)
}

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	//json.NewEncoder(w).Encode(people)

	js, err := json.Marshal(people)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("Content-Type", "application/xml")
	w.Write(js)
}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	//fmt.Fprintf(os.Stderr, "Params: %v\n", params);
	var person Person
	result := json.NewDecoder(req.Body).Decode(&person)
	if result != nil {
		fmt.Fprintf(os.Stderr, "result=%v\n", result)
		return
	}

	//fmt.Fprintf(os.Stderr, "Person: %v\n", person);
	person.ID, _ = params["id"]
	people[person.ID] = person
	json.NewEncoder(w).Encode(people)
}

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, _ := params["id"]

	delete(people, id)
	json.NewEncoder(w).Encode(people)
}

var usuPush = make(map[string] Users)

func insertRecords(db *sql.DB, user *Users)  {
	stms, err := db.Prepare("INSERT INTO dbo.Users (name, address) VALUES " +
		"('" + user.NAME + "', '" + user.ADDRESS +"')")
	_, err = stms.Exec()
	CheckError(err)
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	js, err := json.Marshal(usuPush)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

type Users struct {
	UUID 		string	"json:'uuid,omitempty'"
	NAME		string	"json:'name,omitempty'"
	ADDRESS 	string	"json:'address,omitempty'"
	AGE		int64	"json:'age,omitempty'"
	CREATEDT	string "json:'createDt,omitempty'"
}

func main() {
	/*
	router := mux.NewRouter()
	db, error := getConectionsDB()
	if error != nil {
		log.Panic(error)
	}

	//bulkAddUsr(db)

	buildQuery(db)
	*/

	/*	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
		router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
		router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
		router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")
	*/

	config := ReadConfig()
	log.Println("db_is::"+config.database)

	/*
	router.HandleFunc("/users", getAllUsers).Methods("GET")

	log.Fatal(http.ListenAndServe(":9984", router))

	fmt.Println("listenig on... http://localhost:9984/")
	*/
}

func ReadConfig() Configuration  {
	var configfile = "properties.toml"
	_, err := os.Stat(configfile)
	if err != nil {
		log.Fatal("Config file is missing: ", configfile)
	}

	var config Configuration
	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		log.Fatal(err)
	}
	//log.Print(config.Index)
	return config
}

func getConectionsDB()(*sql.DB, error) {
	config := ReadConfig()
	println("host___is::"+ config.host)

	var debug = flag.Bool("debug", true, "enable debugging")
	var password = flag.String("password", "admin9984", "the database password")
	var port *int = flag.Int("port", 1433, "the database port")
	var server = flag.String("server", "localhost", "the database server")
	var user = flag.String("user", "sa", "the database user")
	var database = flag.String("database", "janezdev", "the database name")
	flag.Parse() // parse the command line args

	if *debug {
		fmt.Printf(" password:%s\n", *password)
		fmt.Printf(" port:%d\n", *port)
		fmt.Printf(" server:%s\n", *server)
		fmt.Printf(" user:%s\n", *user)
		fmt.Printf(" database:%s\n", *database)
	}
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s", *server, *user, *password, *port, *database)
	if *debug {
		fmt.Printf("connString:%s\n", connString)
	}
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	return db, err
}

func bulkAddUsr(db *sql.DB)  {
	for index := 0; index < 100000; index++ {
		insertRecords(db, &Users{
			NAME: "name__" + strconv.Itoa(index),
			ADDRESS: "Address__" + strconv.Itoa(index),
			AGE: 45,
		})
	}
}

type Configuration struct {
	host		string
	port		string
	username	string
	password	string
	database	string
}

func buildQuery(db *sql.DB)  {
	rows, err := db.Query("select cast(uuid as char(36)), name, address, createDt from dbo.Users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	index := 0
	for rows.Next() {
		index++
		var r Users
		err = rows.Scan(&r.UUID, &r.NAME, &r.ADDRESS, &r.CREATEDT)
		if err != nil {
			log.Fatalf("Scan: %v", err)
		}
		usuPush["index__" + strconv.Itoa(index)] = r
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}