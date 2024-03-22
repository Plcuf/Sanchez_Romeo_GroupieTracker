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
	fmt.Println(r.FormFile("image"))
	file, _, err := r.FormFile("image")
	if err != nil {
		fmt.Println("NoT cool : ", err.Error())
	}
	defer file.Close()

	var buffer bytes.Buffer
	write := multipart.NewWriter(&buffer)
	part, err := write.CreateFormFile("image", "image.jpg")
	io.Copy(part, file)
	write.Close()

	url := "https://api.trace.moe/searchhttps://api.trace.moe/search"

	req, err := http.NewRequest("POST", url, &buffer)
	if err != nil {
		fmt.Println("NoT cool : ", err.Error())
	}

	req.Header.Set("Content-Type", write.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("NoT cool : ", err.Error())
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)

	temps := temps.GetTemps()
	temps.ExecuteTemplate(w, "tracedmoe", nil)
}

func HandleError(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "404 Not Found", http.StatusNotFound)
}
