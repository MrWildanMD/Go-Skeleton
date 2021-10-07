package main

import (
	"fmt"
	"github.com/MrWildanMD/Go-Skeleton/routers"
	"github.com/joho/godotenv"
	"os"
)

func main() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	r := Routers.SetupRouter()

	port := os.Getenv("port")

	//for running on other port
	if len(os.Args) > 1 {
		resPort := os.Args[1]
		if reqPort != "" {
			port = reqPort
		}
	}

	if port == "" {
		port = "8080" //default localhost
	}

	type Job interface {
		Run()
	}

	r.Run(":" + port)
}
