package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"main/temps"
	"mime/multipart"
	"net/http"
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
	tag := r.FormValue("tag")
	year := r.FormValue("year")
	rating := r.FormValue("rating")
	finished := r.FormValue("finished")

	url := "https://kitsu.io/api/edge/anime?filter[categories]=" + tag

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

	fmt.Println(string(body))

	//Decode JSON response
	var response Anime
	if err := json.Unmarshal(body, &response); err != nil {
		http.Error(w, "Failed to decode JSON response: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response.Tag = tag

	temps := temps.GetTemps()
	temps.ExecuteTemplate(w, "search", response)
}

func Animedisplay(w http.ResponseWriter, r *http.Request) {
	slug := r.FormValue("slug")
	url := "https://kitsu.io/api/edge/anime?filter[slug]=" + slug

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

	fmt.Println(string(body))

	//Decode JSON response
	var response Anime
	if err := json.Unmarshal(body, &response); err != nil {
		http.Error(w, "Failed to decode JSON response: "+err.Error(), http.StatusInternalServerError)
		return
	}

	temps := temps.GetTemps()
	temps.ExecuteTemplate(w, "anime", response)
}
