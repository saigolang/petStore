package backend

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"petStore/constants"
	"petStore/logger"
	"petStore/structs"
)

func CreatePetResource(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var request structs.Pet
	err := decoder.Decode(&request)
	if err != nil {
		logger.Log.WithField("unmarshalling error", err.Error())
		handleError(rw, err)
		return
	}

	marshalledData, err := json.Marshal(request)
	if err != nil {
		logger.Log.WithField("marshalling error", err.Error())
		handleError(rw, err)
		return
	}
	pet, postErr := postData(marshalledData, constants.URL, structs.CreatePet{})
	if postErr != nil {
		logger.Log.WithField("connection error", postErr)
		handleError(rw, postErr)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)

	jsonResponse, err := json.Marshal(pet)
	if err != nil {
		logger.Log.WithField("marshalling error", err.Error())
		handleError(rw, err)
		return
	}
	rw.Write(jsonResponse)
}

func postData(marshalledData []byte, url string, input interface{}) (interface{}, error) {

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(marshalledData))
	if err != nil {
		logger.Log.WithField("connection error", err.Error())
		return structs.CreatePet{}, err
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")
	response, err := Client.Do(request)
	if err != nil {
		logger.Log.WithField("connection error", err.Error())
		return structs.CreatePet{}, err
	}
	if response.StatusCode != http.StatusOK {
		logger.Log.Error("connection Failed")
		return nil, errors.New("connection Failed")
	}
	json.NewDecoder(response.Body).Decode(&input)
	return input, nil
}
