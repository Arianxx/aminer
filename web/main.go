package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"

	"github.com/arianxx/aminer/web/model"
	"github.com/arianxx/aminer/web/object"
)

func init() {
	object.Db = model.InitDgraphConn("127.0.0.1:9080")
}

func main() {
	defer func() {
		if err := object.Db.Cancel(); err != nil {
			log.Fatal(err)
		}
	}()

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: object.QueryType,
	})
	if err != nil {
		log.Fatal(err)
	}

	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: true,
	})

	fmt.Println("Web Server Running....")

	http.Handle("/graphql", h)
	err = http.ListenAndServe("0.0.0.0:1234", nil)
	if err != nil {
		log.Fatal(err)
	}
}
