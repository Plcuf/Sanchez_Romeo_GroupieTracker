package temps

import (
	"fmt"
	"html/template"
)

func GetTemps() *template.Template {
	temp, err := template.ParseGlob("temps/*.html")
	if err != nil {
		fmt.Println("NoT cool : ", err.Error())
	}

	return temp
}
