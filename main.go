package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/url"
	"path"
)

func main() {
	//create a new router
	router := mux.NewRouter()
	//specify endpoints
	router.HandleFunc("/pets", getPets).Methods("GET")
	router.HandleFunc("/pet/{id}", getPetById).Methods("GET")
	router.HandleFunc("/create", createPetResource).Methods("POST")
	http.Handle("/", router)
	fmt.Println("Server started")
	//start and listen to requests
	log.Fatal(http.ListenAndServe(":8081", router))
}

type Pet struct {
	Id    int     `json:"id"`
	Type  string  `json:"type"`
	Price float64 `json:"price"`
}

type CreatePet struct {
	Pets    Pet    `json:"pet"`
	Message string `json:"message"`
}

func getPets(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("inside get pets")
	url := "http://petstore-demo-endpoint.execute-api.com/petstore/pets"
	response, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	var pets []Pet
	json.NewDecoder(response.Body).Decode(&pets)
	fmt.Println("pets we got is ", pets)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(pets)
	if err != nil {
		fmt.Println("error is ", err)
		return
	}
	fmt.Println("json response is ", jsonResponse)
	rw.Write(jsonResponse)

	return
}

func getPetById(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("pet id functionality is called")
	vars := mux.Vars(req)
	petId := vars["id"]
	baseURL, _ := url.Parse("http://petstore-demo-endpoint.execute-api.com/petstore/pets")

	baseURL.Path = path.Join(baseURL.Path, "/"+petId)

	fmt.Println("URL formed is ", baseURL.String())
	response, err := http.Get(baseURL.String())
	if err != nil {
		log.Fatalln(err)
	}
	var pets Pet
	json.NewDecoder(response.Body).Decode(&pets)

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(pets)
	if err != nil {
		fmt.Println("error is ", err)
		return
	}
	rw.Write(jsonResponse)

	fmt.Println("json response is ", jsonResponse)
	return

}

func createPetResource(rw http.ResponseWriter, req *http.Request) {

	fmt.Println("id we have is ", req.FormValue("id"))

	decoder := json.NewDecoder(req.Body)
	var request Pet
	err := decoder.Decode(&request)
	if err != nil {
		panic(err)
	}

	/*if err := req.ParseForm(); err != nil {
		fmt.Fprintf(rw, "ParseForm() err: %v", err)
		return
	}

	input := map[string]string{"id": req.FormValue("id"),
		"type":  req.FormValue("type"),
		"price": req.FormValue("price"),
	}
	fmt.Println("input is ", input)*/

	fmt.Println("request is ", request)

	Marshalled_data, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("http://petstore-demo-endpoint.execute-api.com/petstore/pets", "application/json",
		bytes.NewBuffer(Marshalled_data))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("http response is ", resp.Body)

	var pet CreatePet
	json.NewDecoder(resp.Body).Decode(&pet)
	fmt.Println("pets we got is ", pet)

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(pet)
	if err != nil {
		fmt.Println("error is ", err)
		return
	}
	rw.Write(jsonResponse)
}
