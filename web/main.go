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

type corsHandler struct {
	h func(w http.ResponseWriter, r *http.Request)
}

func (c *corsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")

	c.h(w, r)
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

	graphqlHandler := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: true,
	})
	h := &corsHandler{graphqlHandler.ServeHTTP}

	fmt.Println("Web Server Running....")

	http.Handle("/graphql", h)
	err = http.ListenAndServe("0.0.0.0:1234", nil)
	if err != nil {
		log.Fatal(err)
	}
}
