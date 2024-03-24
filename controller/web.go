package controller

import (
	"bytes"
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
		fmt.Println("NoT cool : ", err.Error())
	}
	defer file.Close()

	var buffer bytes.Buffer
	write := multipart.NewWriter(&buffer)

	part, err := write.CreateFormFile("image", "image.jpg")
	if err != nil {
		fmt.Println("NoT cool : ", err.Error())
	}

	io.Copy(part, file)

	url := "https://api.trace.moe/search"

	req, err := http.Post(url, "multipart/form-data", &buffer)
	if err != nil {
		fmt.Println("NoT cool : ", err.Error())
	}

	defer req.Body.Close()

	fmt.Println("response Status:", req.Status)

	http.Redirect(w, r, "/tracemoe/result", http.StatusSeeOther)
}

func Tracemoeresult(w http.ResponseWriter, r *http.Request) {

}

func Search(w http.ResponseWriter, r *http.Request) {
}

func Searchtreatment(w http.ResponseWriter, r *http.Request) {
}

func Anime(w http.ResponseWriter, r *http.Request) {
}

func HandleError(w http.ResponseWriter, r *http.Request) {
	temps := temps.GetTemps()
	temps.ExecuteTemplate(w, "error", nil)
}
