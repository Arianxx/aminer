package internal

import (
	"context"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
)

var schema = `
	id: string @index(exact) .
	title: string @index(fulltext, exact) .
	venue: string @index(term) .
	year: int @index(int) .
	n_citation: int @index(int) .
	page_start: string @index(exact) .
	page_end: string @index(exact) .
	volume: string @index(exact) .
	issue: string @index(exact) .
	issn: string @index(exact) .
	isbn: string @index(exact) .
	doi: string @index(exact) .
	url: [string] @index(exact) .
	abstract: string @index(term) .
	authors: uid @reverse .
	name: string @index(term) .
	org: string @index(term) .	
	venue: string @index(exact) .
	keyword: [string] @index(term) .
	fos: [string] @index(term) .
	references: [string] @index(exact) .
	doc_type: string @index(term) .
	lang: string @index(exact) .
	publisher: string @index(exact) .
`

func AlterSchema(dg *dgo.Dgraph) error {
	op := &api.Operation{}
	op.Schema = schema
	ctx := context.Background()
	err := dg.Alter(ctx, op)
	return err
}
