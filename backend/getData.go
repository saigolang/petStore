package backend

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/url"
	"path"
	"petStore/constants"
	"petStore/structs"
)

func GetPets(rw http.ResponseWriter, req *http.Request) {
	response, err := httpGet(constants.URL)
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

	baseURL, _ := url.Parse(constants.URL)
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
