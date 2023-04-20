package ppts

type ArtPromptModel struct {
	Title            string   `json:"title"`
	UserId           int64    `json:"userId"`
	Description      string   `json:"description"`
	Model            string   `json:"model"`
	Version          string   `json:"version"`
	Category         []string `json:"category"`
	Seed             string   `json:"seed"`
	Type             string   `json:"type"`
	Steps            string   `json:"steps"`
	SubModel         string   `json:"subModel"`
	GuidanceScale    string   `json:"guidanceScale"`
	Sampler          string   `json:"sampler"`
	Content          string   `json:"content"`
	NegativeContent  string   `json:"negativeContent"`
	ExtraInstruction string   `json:"extraInstruction"`
	Images           []string `json:"images"`
	Medias           []int64  `json:"medias"`
}
