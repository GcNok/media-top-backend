package response

type (
	GetArticlesResponse struct {
		Title        string `json:"title"`
		ArticleURL   string `json:"articleURL"`
		MainVisual   string `json:"mainVisual"`
		WriterImage  string `json:"writerImage"`
		WriterName   string `json:"writerName"`
		Last30daysPv string `json:"last30daysPv"`
		Updated      string `json:"updated"`
	}
)
