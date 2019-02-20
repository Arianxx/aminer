package object

import (
	"github.com/graphql-go/graphql"

	"github.com/arianxx/aminer/web/model"
)

var Db *model.Db

var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"papers":  queryPapers,
		"paper":   queryPaper,
		"authors": queryAuthors,
	},
})
