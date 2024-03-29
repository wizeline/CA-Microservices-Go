package controller

const dateFormat = "2006-01-02"

// basicMessage is the representation of a basic http JSON response.
type basicMessage struct {
	Message string `json:"message"`
}
