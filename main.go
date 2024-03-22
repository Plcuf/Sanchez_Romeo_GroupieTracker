package main

import (
	r "main/routeur"
	t "main/temps"
)

func main() {
	// Initialize routes
	t.GetTemps()
	r.InitServer()
}
