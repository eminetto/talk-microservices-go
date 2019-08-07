package feedback

type Feedback struct {
	Email string
	Title string `json:"title"`
	Body string `json:"body"`
}
