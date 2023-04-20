package ppts

type ChatgptPromptModel struct {
	Title       string   `json:"title"`
	UserId      int64    `json:"userId"`
	Content     string   `json:"content"`
	Description string   `json:"description"`
	Model       string   `json:"model"`
	Version     string   `json:"version"`
	Category    []string `json:"category"`
}
