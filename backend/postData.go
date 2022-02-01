package backend

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"petStore/structs"
)

func CreatePetResource(rw http.ResponseWriter, req *http.Request) {

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

	resp, err := http.Post("http://petstore-demo-endpoint.execute-api.com/petstore/pets", "application/json",
		bytes.NewBuffer(marshalledData))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("http response is ", resp.Body)

	var pet structs.CreatePet
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
