package routeur

import (
	"fmt"
	ctrl "main/controller"
	"net/http"
	"os"
)

func InitServer() {
	//Initialisation des routes
	http.HandleFunc("/index", ctrl.Index)
	http.HandleFunc("/tracemoe", ctrl.Tracemoe)
	http.HandleFunc("/tracemoe/result", ctrl.Tracemoetreatment)
	http.HandleFunc("/search", ctrl.Search)
	http.HandleFunc("/anime/treatment", ctrl.AnimeTreatment)
	http.HandleFunc("/animedisplay", ctrl.Animedisplay)

	// Handle 404
	http.HandleFunc("/", ctrl.HandleError)

	// Serve static files
	RootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(RootDoc + "/assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fileserver))

	//Le lien d'ou est lancé le serveur
	fmt.Println("(http://localhost:8282/index) - Server started on port:8282")
	http.ListenAndServe("localhost:8282", nil)
	fmt.Println("Server closed")
}
