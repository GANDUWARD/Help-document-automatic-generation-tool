package types

type Request struct {
	Name string `path:"name"`
}
type PostRequest struct {
	Path string `json:"path"`
}
type SaveRequest struct {
	Name    string `json:"name"`
	Path    string `json:"path"`
	Content string `json:"content"`
}
type Response struct {
	Message string `json:"message"`
	Data    string `json:"data"`
}
