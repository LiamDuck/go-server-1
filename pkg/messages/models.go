package messages

type Message struct {
	ID      string   `json:"id"`
	Content Contents `json:"content"`
}
type Contents struct {
	User string `json:"user"`
	Text string `json:"text"`
}
