package vote

type Vote struct {
	Email string
	TalkName string `json:"talk_name"`
	Score int `json:"score"`
}

