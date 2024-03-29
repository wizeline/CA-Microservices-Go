package controller

type httpRequestTest struct {
	params  map[string]string
	payload []byte
}

type httpResponseTest struct {
	code int
	body string
}
