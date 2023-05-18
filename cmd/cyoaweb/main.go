package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bocdagla/gostory/package/cyoa"
)

func main() {
	port := flag.Int("port", 3000, "The port to start CYOA web application on")
	filename := flag.String("file", "gopher.json", "the JSON file with the CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s. \n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		log.Fatalf("could not open the file %s \n Error Description: \n %v", *filename, err)
		return
	}

	story, err := cyoa.JsonStory(f)
	if err != nil {
		log.Fatalf("could not parse the file %s \n Error Description: \n %v", *filename, err)
		return
	}

	h := cyoa.NewHandler(story)
	fmt.Printf("Starting the server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
