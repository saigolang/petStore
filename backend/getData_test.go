package backend

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"petStore/logger"
	"petStore/structs"
	"testing"
)

func TestGetPets(t *testing.T) {
	// initializing logger
	logger.NewLogger()

	req, err := http.NewRequest("GET", "/pets", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	mockResponse := `[
    {
        "id": 1,
        "type": "Hamster",
        "price": 249.99
    },
    {
        "id": 2,
        "type": "cat",
        "price": 124.99
    },
    {
        "id": 3,
        "type": "fish",
        "price": 0.99
    }
]`
	Client = MockClient{
		mockResponse: mockResponse,
		statusCode:   http.StatusOK,
		httpError:    nil,
	}
	handler := http.HandlerFunc(GetPets)
	expected := `[{"id":1,"price":249.99,"type":"Hamster"},{"id":2,"price":124.99,"type":"cat"},{"id":3,"price":0.99,"type":"fish"}]`

	handler.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
	assert.Equal(t, expected, rr.Body.String())
}

func TestGetPetById(t *testing.T) {
	// initializing logger
	logger.NewLogger()

	req, err := http.NewRequest("GET", "/pet/id", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	mockResponse := `
    {
        "id": 1,
        "type": "fish",
        "price": 0.99
    }`
	Client = MockClient{
		mockResponse: mockResponse,
		statusCode:   http.StatusOK,
		httpError:    nil,
	}
	handler := http.HandlerFunc(GetPetById)
	expected := `{"id":1,"price":0.99,"type":"fish"}`

	handler.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
	assert.Equal(t, expected, rr.Body.String())
}
func TestGetHTTPData(t *testing.T) {
	// initializing logger
	logger.NewLogger()

	t.Run("happyPath", func(t *testing.T) {
		mockResponse := `[
    {
        "id": 1,
        "type": "dog",
        "price": 249.99
    },
    {
        "id": 2,
        "type": "cat",
        "price": 124.99
    },
    {
        "id": 3,
        "type": "fish",
        "price": 0.99
    }
]`
		Client = MockClient{
			mockResponse: mockResponse,
			statusCode:   http.StatusOK,
			httpError:    nil,
		}
		value, err := getHTTPData("https://petStore/pets", []structs.Pet{})
		assert.NotNil(t, value)
		assert.NoError(t, err)
	})

	t.Run("failPath", func(t *testing.T) {
		mockResponse := ``
		Client = MockClient{
			mockResponse: mockResponse,
			statusCode:   http.StatusInternalServerError,
			httpError:    errors.New("connectivity failed"),
		}
		value, err := getHTTPData("https://petStore/pets", []structs.Pet{})
		assert.Nil(t, value)
		assert.Error(t, err)
	})

}
