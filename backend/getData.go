package backend

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"net/url"
	"path"
	"petStore/constants"
	"petStore/logger"
	"petStore/structs"
)

func GetPets(rw http.ResponseWriter, req *http.Request) {
	var pets []structs.Pet
	response, err := getHTTPData(constants.URL, pets)
	rw.Header().Set("Content-Type", "application/json")
	if err != nil {
		handleError(rw, err)
		return
	}
	rw.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		logger.Log.WithField("marshalling error", err.Error())
		handleError(rw, err)
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

	response, err := getHTTPData(baseURL.String(), structs.Pet{})
	if err != nil {
		handleError(rw, err)
		logger.Log.WithField("connection error", err.Error())
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		logger.Log.WithField("marshalling error", err.Error())
		handleError(rw, err)
		return
	}
	rw.Write(jsonResponse)
	return
}

func getHTTPData(url string, input interface{}) (interface{}, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		logger.Log.WithField("connection error", err.Error())
		return nil, err
	}
	request.Header.Add("Accept", "application/json")
	response, err := Client.Do(request)
	if err != nil {
		logger.Log.WithField("connection error", err.Error())
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		logger.Log.Error("connection Failed")
		return nil, errors.New("connection Failed")
	}
	json.NewDecoder(response.Body).Decode(&input)
	return input, nil
}
