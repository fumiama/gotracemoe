package trmoe

type Result struct {
	FrameCount int     `json:"frameCount"`
	Error      string  `json:"error"`
	Result     []Anime `json:"result"`
}

type Anime struct {
	Anilist    Anilist `json:"anilist"`
	Filename   string  `json:"filename"`
	Episode    int     `json:"episode"`
	From       float32 `json:"from"`
	To         float32 `json:"to"`
	Similarity float64 `json:"similarity"`
	Video      string  `json:"video"`
	Image      string  `json:"image"`
}

type Anilist struct {
	Id       int      `json:"id"`
	IdMal    int      `json:"idMal"`
	Title    Title    `json:"title"`
	Synonyms []string `json:"synonyms"`
	IsAdult  bool     `json:"isAdult"`
}

type Title struct {
	Native  string `json:"native"`
	Romaji  string `json:"romaji"`
	English string `json:"english"`
}
