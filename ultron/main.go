package main

import (
	_ "./controllers"
	_ "./models"
	"github.com/gernest/utron"
)

func main() {
	utron.Run()
}
