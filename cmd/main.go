package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/isubhampadhi56/remote-management/router"
)

func main() {
	router := router.MainRouter()
	fmt.Println("Staring API Server on 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalln(err)
	}
}
