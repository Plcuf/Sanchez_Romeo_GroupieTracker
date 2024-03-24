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

type Anime struct {
	Tag           string
	ID            string             `json:"id"`
	Type          string             `json:"type"`
	Links         AnimeLinks         `json:"links"`
	Attributes    AnimeAttributes    `json:"attributes"`
	Relationships AnimeRelationships `json:"relationships"`
}

type AnimeLinks struct {
	Self string `json:"self"`
}

type AnimeAttributes struct {
	CreatedAt           time.Time `json:"createdAt"`
	UpdatedAt           time.Time `json:"updatedAt"`
	Slug                string    `json:"slug"`
	Synopsis            string    `json:"synopsis"`
	Description         string    `json:"description"`
	CoverImageTopOffset int       `json:"coverImageTopOffset"`
	Titles              struct {
		En   string `json:"en"`
		EnJP string `json:"en_jp"`
		EnUS string `json:"en_us"`
		JaJP string `json:"ja_jp"`
	} `json:"titles"`
	CanonicalTitle    string            `json:"canonicalTitle"`
	AbbreviatedTitles []string          `json:"abbreviatedTitles"`
	AverageRating     string            `json:"averageRating"`
	RatingFrequencies map[string]string `json:"ratingFrequencies"`
	UserCount         int               `json:"userCount"`
	FavoritesCount    int               `json:"favoritesCount"`
	StartDate         string            `json:"startDate"`
	EndDate           string            `json:"endDate"`
	NextRelease       interface{}       `json:"nextRelease"`
	PopularityRank    int               `json:"popularityRank"`
	RatingRank        int               `json:"ratingRank"`
	AgeRating         string            `json:"ageRating"`
	AgeRatingGuide    string            `json:"ageRatingGuide"`
	Subtype           string            `json:"subtype"`
	Status            string            `json:"status"`
	TBA               interface{}       `json:"tba"`
	PosterImage       struct {
		Tiny     string `json:"tiny"`
		Large    string `json:"large"`
		Small    string `json:"small"`
		Medium   string `json:"medium"`
		Original string `json:"original"`
		Meta     struct {
			Dimensions struct {
				Tiny   Dimension `json:"tiny"`
				Large  Dimension `json:"large"`
				Small  Dimension `json:"small"`
				Medium Dimension `json:"medium"`
			} `json:"dimensions"`
		} `json:"meta"`
	} `json:"posterImage"`
	CoverImage struct {
		Tiny     string `json:"tiny"`
		Large    string `json:"large"`
		Small    string `json:"small"`
		Original string `json:"original"`
		Meta     struct {
			Dimensions struct {
				Tiny  Dimension `json:"tiny"`
				Large Dimension `json:"large"`
				Small Dimension `json:"small"`
			} `json:"dimensions"`
		} `json:"meta"`
	} `json:"coverImage"`
	EpisodeCount   int    `json:"episodeCount"`
	EpisodeLength  int    `json:"episodeLength"`
	TotalLength    int    `json:"totalLength"`
	YoutubeVideoID string `json:"youtubeVideoId"`
	ShowType       string `json:"showType"`
	NSFW           bool   `json:"nsfw"`
}

type AnimeRelationships struct {
	Genres             RelationshipLinks `json:"genres"`
	Categories         RelationshipLinks `json:"categories"`
	Castings           RelationshipLinks `json:"castings"`
	Installments       RelationshipLinks `json:"installments"`
	Mappings           RelationshipLinks `json:"mappings"`
	Reviews            RelationshipLinks `json:"reviews"`
	MediaRelationships RelationshipLinks `json:"mediaRelationships"`
	Characters         RelationshipLinks `json:"characters"`
	Staff              RelationshipLinks `json:"staff"`
	Productions        RelationshipLinks `json:"productions"`
	Quotes             RelationshipLinks `json:"quotes"`
	Episodes           RelationshipLinks `json:"episodes"`
	StreamingLinks     RelationshipLinks `json:"streamingLinks"`
	AnimeProductions   RelationshipLinks `json:"animeProductions"`
	AnimeCharacters    RelationshipLinks `json:"animeCharacters"`
	AnimeStaff         RelationshipLinks `json:"animeStaff"`
}

type RelationshipLinks struct {
	Links RelationshipURLs `json:"links"`
}

type RelationshipURLs struct {
	Self    string `json:"self"`
	Related string `json:"related"`
}

type Dimension struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}
