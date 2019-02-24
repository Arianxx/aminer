package main

import (
	"github.com/arianxx/aminer/web/model"
	"log"

	"github.com/arianxx/aminer/internal"
)

func main() {
	conn, cancel := internal.GetDgraphClient("127.0.0.1:9080")
	defer func() {
		if err := cancel(); err != nil {
			log.Fatal(err)
		}
	}()

	paths, err := model.GetPaperFilePaths()
	if err != nil {
		log.Fatal(err)
	}

	err = storeMultiFile(conn, paths, 10)
	if err != nil {
		log.Fatal(err)
	}
}
