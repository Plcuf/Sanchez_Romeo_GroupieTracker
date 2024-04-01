package controller

import "time"

type TraceMoeResult struct {
	FrameCount int    `json:"frameCount"`
	Error      string `json:"error"`
	Result     []struct {
		Anilist    int     `json:"anilist"`
		Filename   string  `json:"filename"`
		Episode    *int    `json:"episode"` // episode could be null, so use pointer type
		From       float64 `json:"from"`
		To         float64 `json:"to"`
		Similarity float64 `json:"similarity"`
		Video      string  `json:"video"`
		Image      string  `json:"image"`
	} `json:"result"`
}

type AnimeResponse struct {
	Offset   string
	Tag      string
	Year     string
	Rating   string
	Finished string
	Amount   string
	Data     []Anime `json:"data"`
	Meta     Meta    `json:"meta"`
	Links    Links   `json:"links"`
}

type Anime struct {
	ID            string        `json:"id"`
	Type          string        `json:"type"`
	Links         Links         `json:"links"`
	Attributes    Attributes    `json:"attributes"`
	Relationships Relationships `json:"relationships"`
}

type Links struct {
	Self string `json:"self"`
}

type Meta struct {
	Count int `json:"count"`
}

type Attributes struct {
	CreatedAt           time.Time         `json:"createdAt"`
	UpdatedAt           time.Time         `json:"updatedAt"`
	Slug                string            `json:"slug"`
	Synopsis            string            `json:"synopsis"`
	Description         string            `json:"description"`
	CoverImageTopOffset int               `json:"coverImageTopOffset"`
	Titles              Titles            `json:"titles"`
	CanonicalTitle      string            `json:"canonicalTitle"`
	AbbreviatedTitles   []string          `json:"abbreviatedTitles"`
	AverageRating       string            `json:"averageRating"`
	RatingFrequencies   map[string]string `json:"ratingFrequencies"`
	UserCount           int               `json:"userCount"`
	FavoritesCount      int               `json:"favoritesCount"`
	StartDate           string            `json:"startDate"`
	EndDate             string            `json:"endDate"`
	NextRelease         interface{}       `json:"nextRelease"`
	PopularityRank      int               `json:"popularityRank"`
	RatingRank          int               `json:"ratingRank"`
	AgeRating           string            `json:"ageRating"`
	AgeRatingGuide      string            `json:"ageRatingGuide"`
	Subtype             string            `json:"subtype"`
	Status              string            `json:"status"`
	TBA                 interface{}       `json:"tba"`
	PosterImage         Image             `json:"posterImage"`
	CoverImage          Image             `json:"coverImage"`
	EpisodeCount        int               `json:"episodeCount"`
	EpisodeLength       int               `json:"episodeLength"`
	TotalLength         int               `json:"totalLength"`
	YoutubeVideoID      string            `json:"youtubeVideoId"`
	ShowType            string            `json:"showType"`
	NSFW                bool              `json:"nsfw"`
}

type Titles struct {
	En   string `json:"en"`
	EnJP string `json:"en_jp"`
	EnUS string `json:"en_us"`
	JaJP string `json:"ja_jp"`
}

type Image struct {
	Tiny     string `json:"tiny"`
	Large    string `json:"large"`
	Small    string `json:"small"`
	Medium   string `json:"medium"`
	Original string `json:"original"`
	Meta     Meta   `json:"meta"`
}

type Relationships struct {
	Genres             Relationship `json:"genres"`
	Categories         Relationship `json:"categories"`
	Castings           Relationship `json:"castings"`
	Installments       Relationship `json:"installments"`
	Mappings           Relationship `json:"mappings"`
	Reviews            Relationship `json:"reviews"`
	MediaRelationships Relationship `json:"mediaRelationships"`
	Characters         Relationship `json:"characters"`
	Staff              Relationship `json:"staff"`
	Productions        Relationship `json:"productions"`
	Quotes             Relationship `json:"quotes"`
	Episodes           Relationship `json:"episodes"`
	StreamingLinks     Relationship `json:"streamingLinks"`
	AnimeProductions   Relationship `json:"animeProductions"`
	AnimeCharacters    Relationship `json:"animeCharacters"`
	AnimeStaff         Relationship `json:"animeStaff"`
}

type Relationship struct {
	Links Links `json:"links"`
}
