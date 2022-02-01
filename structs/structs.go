package structs

type Pet struct {
	Id    int     `json:"id"`
	Type  string  `json:"type"`
	Price float64 `json:"price"`
}

type CreatePet struct {
	Pets    Pet    `json:"pet"`
	Message string `json:"message"`
}

type ErrorMessage struct {
	Message string
}
