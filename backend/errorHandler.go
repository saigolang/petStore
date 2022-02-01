package backend

import (
	"encoding/json"
	"net/http"
	"petStore/structs"
)

func handleError(rw http.ResponseWriter, err error) {
	var errorMessage structs.ErrorMessage
	errorMessage.Message = "Error in getting the response:" + err.Error()
	errorInfo, err := json.Marshal(errorMessage)
	if err != nil {
		return
	}
	rw.WriteHeader(http.StatusInternalServerError)
	rw.Write(errorInfo)
}
