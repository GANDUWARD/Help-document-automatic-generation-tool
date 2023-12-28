package types

type Request struct {
	Name string `path:"name"`
}

type Response struct {
	Message string `json:"message"`
	Data    string `json:"data"`
}
