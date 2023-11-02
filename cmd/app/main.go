package main

import "fmt"

func main(){
	fmt.Println("run now")
}	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	host := os.Getenv("PROJECT_HOST")
	port := os.Getenv("PROJECT_PORT")

	serverAddress := fmt.Sprintf("%s:%s", host, port)

