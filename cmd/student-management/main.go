package main

import (
	"example.com/student-management/initializers"
	"example.com/student-management/internal/router"
)

func main() {
	initializers.ConnectDB()
	router.IntializeRouter()

}
