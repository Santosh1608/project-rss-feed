package requests

type CreateFeed struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
