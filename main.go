package main

import (
	"assignment2/router"
)

func main() {
	r := router.StartApp()

	r.Run(("localhost:8000"))
}
