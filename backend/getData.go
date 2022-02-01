package backend

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/url"
	"path"
	"petStore/structs"
)

func GetPets(rw http.ResponseWriter, req *http.Request) {
	url := "http://petstore-demo-endpoint.execute-api.com/petstore/pets"
	response, err := httpGet(url)
	if err != nil {
		log.Fatalln(err)
	}
	var pets []structs.Pet
	json.NewDecoder(response.Body).Decode(&pets)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(pets)
	if err != nil {
		fmt.Println("error is ", err)
		return
	}
	rw.Write(jsonResponse)
	return
}

func GetPetById(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	petId := vars["id"]
	baseURL, _ := url.Parse("http://petstore-demo-endpoint.execute-api.com/petstore/pets")

	baseURL.Path = path.Join(baseURL.Path, "/"+petId)

	response, err := httpGet(baseURL.String())
	if err != nil {
		log.Fatalln(err)
	}
	var pets structs.Pet
	json.NewDecoder(response.Body).Decode(&pets)

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(pets)
	if err != nil {
		fmt.Println("error is ", err)
		return
	}
	rw.Write(jsonResponse)

	return

}

func httpGet(url string) (*http.Response, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return response, nil
}
