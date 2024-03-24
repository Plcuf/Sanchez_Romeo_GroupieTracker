package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"main/temps"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	temps := temps.GetTemps()
	temps.ExecuteTemplate(w, "index", nil)
}

func Tracemoe(w http.ResponseWriter, r *http.Request) {
	temps := temps.GetTemps()
	temps.ExecuteTemplate(w, "tracemoe", nil)
}

func Tracemoetreatment(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)
	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Failed to read image: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	var buffer bytes.Buffer
	write := multipart.NewWriter(&buffer)

	part, err := write.CreateFormFile("image", "image.jpg")
	if err != nil {
		http.Error(w, "Failed to create form file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := io.Copy(part, file); err != nil {
		http.Error(w, "Failed to copy image data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := write.Close(); err != nil {
		http.Error(w, "Failed to close multipart writer: "+err.Error(), http.StatusInternalServerError)
		return
	}

	url := "https://api.trace.moe/search"

	req, err := http.Post(url, write.FormDataContentType(), &buffer)
	if err != nil {
		http.Error(w, "Failed to make POST request to Trace.moe API: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer req.Body.Close()

	fmt.Println("response Status:", req.Status)
	// Read response body
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "Failed to read response body: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Decode JSON response
	var response TraceMoeResult
	if err := json.Unmarshal(body, &response); err != nil {
		http.Error(w, "Failed to decode JSON response: "+err.Error(), http.StatusInternalServerError)
		return
	}

	temps := temps.GetTemps()
	temps.ExecuteTemplate(w, "tracemoeresult", response)
}

func HandleError(w http.ResponseWriter, r *http.Request) {
	temps := temps.GetTemps()
	temps.ExecuteTemplate(w, "error", nil)
}

func Search(w http.ResponseWriter, r *http.Request) {
	OffSet := "0"
	tag := r.FormValue("category")
	year := r.FormValue("year")
	ratingStr := r.FormValue("rating")
	finished := r.FormValue("finished")
	if len(year) == 0 {
		year = ""
	}
	if len(ratingStr) == 0 {
		ratingStr = ""
	}
	if len(finished) == 0 {
		finished = ""
	}
	if len(year) != 0 {
		year = year + "-01-01"
	}
	var rating float64
	if ratingStr != "" {
		rating, _ = strconv.ParseFloat(ratingStr, 64)
	}

	url := "https://kitsu.io/api/edge/anime?page[limit]=20&page[offset]=" + string(OffSet) + "&filter[categories]=" + tag

	req, err := http.Get(url)
	if err != nil {
		http.Error(w, "Failed to make GET request to Kitsu API: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer req.Body.Close()

	fmt.Println("Response Status:", req.Status)

	// Read response body
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "Failed to read response body: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Decode JSON response
	var response AnimeResponse
	if err := json.Unmarshal(body, &response); err != nil {
		http.Error(w, "Failed to decode JSON response: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Filter the response based on the provided filters
	filteredAnime := AnimeResponse{}
	// Create a map to store unique anime IDs
	uniqueAnimeIDs := make(map[string]struct{})

	for _, anime := range response.Data {
		// Check if the anime ID has already been added
		if _, ok := uniqueAnimeIDs[anime.ID]; ok {
			continue // Skip if the anime ID has already been added
		}

		// Apply filters if they are present
		passedFilters := true

		if year != "" {
			// Filter by year
			StartDate, err := time.Parse("2006-01-02", anime.Attributes.StartDate)
			if err != nil {
				fmt.Println(err)
			}
			YearChek, err := time.Parse("2006-01-02", year)
			if err != nil {
				fmt.Println(err)
			}
			if StartDate.Before(YearChek) {
				passedFilters = false
			}
		}

		// Filter by rating
		if ratingStr != "" {
			avgRating, err := strconv.ParseFloat(anime.Attributes.AverageRating, 64)
			if err != nil || avgRating < rating {
				passedFilters = false
			}
		}

		// Filter by finished status
		if finished != "" && anime.Attributes.Status != finished {
			passedFilters = false
		}

		// If anime passes all filters, add it to filteredAnime and update the map
		if passedFilters {
			filteredAnime.Data = append(filteredAnime.Data, anime)
			uniqueAnimeIDs[anime.ID] = struct{}{}
		}
	}

	// Pass the filtered anime to your page
	temps := temps.GetTemps()
	if year != "" && len(year) != 0 || ratingStr != "" && len(ratingStr) != 0 || finished != "" {
		temps.ExecuteTemplate(w, "search", filteredAnime)
	} else {
		temps.ExecuteTemplate(w, "search", response)
	}
}

func Animedisplay(w http.ResponseWriter, r *http.Request) {
	anime_name := r.FormValue("name")
	url := "https://kitsu.io/api/edge/anime?filter[slug]=" + anime_name

	req, err := http.Get(url)
	if err != nil {
		http.Error(w, "Failed to make POST request to Trace.moe API: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer req.Body.Close()

	fmt.Println("response Status:", req.Status)
	// Read response body
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "Failed to read response body: "+err.Error(), http.StatusInternalServerError)
		return
	}

	//Decode JSON response
	var response AnimeResponse
	if err := json.Unmarshal(body, &response); err != nil {
		http.Error(w, "Failed to decode JSON response: "+err.Error(), http.StatusInternalServerError)
		return
	}

	temps := temps.GetTemps()
	temps.ExecuteTemplate(w, "animedisplay", response)
}
