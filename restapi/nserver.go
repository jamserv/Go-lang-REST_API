package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"fmt"
	"os"

	_ "github.com/denisenkom/go-mssqldb"

	"github.com/gorilla/mux"
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

var gpush = make(map[string]Client)

func getAllComits(w http.ResponseWriter, r *http.Request) {
	/*commits := map[string]int{
		"rsc": 3711,
		"r":   2138,
		"gri": 1908,
		"adg": 912,
	}*/
	/*gpush[1] = Client{
		9984,
		"janez",
	}*/

	js, err := json.Marshal(gpush)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

/*
type Producto struct {
	Id      int       "json:'id,omitempty'"
	Name    string    "json:'name, omitempty'"
	ProvInf Proveedor "json:'provinf'"
}*/

type Client struct {
	Id        int        "json:'id,omitempty'"
	Name      string     "json:'name,omitempty'"
	Proveedor *Proveedor "json:'proveedor,omitempty'"
}

type Proveedor struct {
	IdProv   int    "json:'idprov,omitempty'"
	NameProv string "json:'nameprov,omitempty'"
}

var debug = flag.Bool("debug", false, "enable debugging")
var password = flag.String("password", "usrdev2", "the database password")
var port *int = flag.Int("port", 1433, "the database port")
var server = flag.String("server", "192.168.2.194", "the database server")
var user = flag.String("user", "usrbgjobs", "the database user")

//PlatformlaRelease;instance=VMSQL2005QAM
func simplemain() {

	query := url.Values{}
	//query.Add("connection timeout", fmt.Sprintf("%d", 3000))

	u := &url.URL{
		Scheme:   "PlatformlaRelease",
		User:     url.UserPassword("usrbgjobs", "usrdev2"),
		Host:     fmt.Sprintf("%s:%d", "192.168.2.194", 1433),
		Path:     "VMSQL2005QAM", // if connecting to an instance instead of a port
		RawQuery: query.Encode(),
	}

	connectionString := u.String()

	println(connectionString)

	db, err := sql.Open("sqlserver", connectionString)

	if err != nil {
		fmt.Println("Cannot connect: ", err.Error())
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Cannot connect: ", err.Error())
		return
	}
	// or
	//db, err := sql.Open("mssql", connectionString)
	/*var (
		userid   = flag.String("U", "usrbgjobs", "usrbgjobs")
		password = flag.String("P", "usrdev2", "usrdev2")
		server   = flag.String("S", "192.168.2.194", "192.168.2.194[\\VMSQL2005QAM]")
		database = flag.String("d", "PlatformlaRelease", "PlatformlaRelease")
	)*/
	//flag.Parse()
	//dsn := "server=192.168.2.194;user id=usrbgjobs;password=usrdev2;database=PlatformlaRelease;instance=VMSQL2005QAM"
	/*dsn := "sqlserver://usrbgjobs:usrdev2@192.168.2.194:1433?database=PlatformlaRelease&instance=VMSQL2005QAM"

	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		fmt.Println("Cannot connect: ", err.Error())
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Cannot connect: ", err.Error())
		return
	}*/
}

func main() {
	router := mux.NewRouter()
	/*
		people["1"] = Person{ID: "1", Firstname: "Nic", Lastname: "Raboy", Address: &Address{City: "Dublin", State: "CA"}}
		people["2"] = Person{ID: "2", Firstname: "Maria", Lastname: "Raboy"}
	*/

	simplemain()

	for index := 0; index < 1000; index++ {
		gpush["inde_"+strconv.Itoa(index)] = Client{Id: index, Name: "janez___" + strconv.Itoa(index), Proveedor: &Proveedor{IdProv: index, NameProv: "prov_ca"}}
	}

	/*	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
		router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
		router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
		router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")
	*/
	router.HandleFunc("/commits", getAllComits).Methods("GET")

	log.Fatal(http.ListenAndServe(":9984", router))
}
