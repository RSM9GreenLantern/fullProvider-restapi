package main

import (
	"encoding/json"
	"log"
	"net/http"

	//"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
	. "github.com/SunPriest/fullProvider-getrestapi/config"
	. "github.com/SunPriest/fullProvider-getrestapi/dao"
)

var config = Config{}
var dao = DAO{}

// GET a prac by its ID
func FindPracEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	Prac, err := dao.FindPracById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid prac ID")
		return
	}
	respondWithJson(w, http.StatusOK, Prac)
}

// GET a CED by its ID
func FindCEDEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	CED, err := dao.FindCEDById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid CED ID")
		return
	}
	respondWithJson(w, http.StatusOK, CED)
}

// GET a serviceloc by its ID
func FindServLocEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ServLoc, err := dao.FindServLocById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid service location ID")
		return
	}
	respondWithJson(w, http.StatusOK, ServLoc)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

// Define HTTP request routes
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/prac/{id}", FindPracEndpoint).Methods("GET")
	r.HandleFunc("/ced/{id}", FindCEDEndpoint).Methods("GET")
	r.HandleFunc("/servloc/{id}", FindServLocEndpoint).Methods("GET")
	
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
