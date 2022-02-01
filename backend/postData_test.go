package backend

import (
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"petStore/logger"
	"petStore/structs"
	"testing"
)

func TestCreatePetResource(t *testing.T) {
	// initializing logger
	logger.NewLogger()
	req, err := http.NewRequest("POST", "/create", nil)
	if err != nil {
		t.Fatal(err)
	}

	form := url.Values{}

	form.Add("id", "5")
	form.Add("type", "fish")
	form.Add("price", "0.99")
	req.PostForm = form
	rr := httptest.NewRecorder()
	mockResponse := `{
    		"pet": {
        		"id": 5,
        		"type": "fish",
        		"price": 0.99
			},
    	"message": "success"
	}`

	expected := `{"message":"success","pet":{"id":5,"price":0.99,"type":"fish"}}`
	Client = MockClient{
		mockResponse: mockResponse,
		statusCode:   http.StatusOK,
		httpError:    nil,
	}
	handler := http.HandlerFunc(GetPetById)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
	assert.Equal(t, expected, rr.Body.String())
}

func TestPostData(t *testing.T) {
	// initializing logger
	logger.NewLogger()

	t.Run("happyPath", func(t *testing.T) {
		mockResponse := `{
    		"pet": {
        		"id": 5,
        		"type": "fish",
        		"price": 0.99
			},
    	"message": "success"
	}`
		Client = MockClient{
			mockResponse: mockResponse,
			statusCode:   http.StatusOK,
			httpError:    nil,
		}

		mockRequest := structs.Pet{Id: 1, Type: "dog", Price: 23.5}
		marshalledData, err := json.Marshal(mockRequest)
		if err != nil {
			log.Fatal(err)
		}

		value, err := postData(marshalledData, "https://petStore/pets", []structs.CreatePet{})
		assert.NotNil(t, value)
		assert.NoError(t, err)
	})

	t.Run("failPath", func(t *testing.T) {
		mockResponse := ``
		Client = MockClient{
			mockResponse: mockResponse,
			statusCode:   http.StatusInternalServerError,
			httpError:    errors.New("connection failed"),
		}

		mockRequest := structs.Pet{Id: 1, Type: "dog", Price: 23.5}
		marshalledData, err := json.Marshal(mockRequest)
		if err != nil {
			log.Fatal(err)
		}
		value, err := postData(marshalledData, "https://petStore/pets", []structs.CreatePet{})
		assert.NotNil(t, value)
		assert.Error(t, err)
	})

}
