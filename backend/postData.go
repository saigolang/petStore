package backend

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"petStore/constants"
	"petStore/structs"
)

func CreatePetResource(rw http.ResponseWriter, req *http.Request) {
	// validate input request
	decoder := json.NewDecoder(req.Body)
	var request structs.Pet
	err := decoder.Decode(&request)
	if err != nil {
		panic(err)
	}
	marshalledData, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	pet, postErr := postData(marshalledData)
	// todo handle error
	if postErr != nil {
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(pet)
	if err != nil {
		return
	}
	rw.Write(jsonResponse)
}

func postData(marshalledData []byte) (structs.CreatePet, error) {
	resp, err := http.Post(constants.URL, "application/json",
		bytes.NewBuffer(marshalledData))

	if err != nil {
		return structs.CreatePet{}, errors.New("error in posting the request: " + err.Error())
	}
	var pet structs.CreatePet
	json.NewDecoder(resp.Body).Decode(&pet)
	return pet, nil
}

// todo
func validatePostRequest() {

}
