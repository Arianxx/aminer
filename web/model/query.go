package model

var QueryPaperList = ListQuery{
	Name: "Paper",
	Want: `
	id
    title
    authors{
      name
      org
      papers: ~authors{
      id
    }
    }
  	venue
  	year
  	keywords
  	fos
  	n_citation
  	references
  	page_start
  	page_end
  	doc_type
  	lang
  	publisher
  	volume
  	issue
  	issn
  	isbn
  	doi
  	pdf
  	url
  	abstract
	`,
}

var QueryPaper = `
query Paper($id: string){
  data(func: eq(id, $id)){
    id
    title
    authors{
      name
      org
      papers: ~authors{
      id
    }
    }
  	venue
  	year
  	keywords
  	fos
  	n_citation
  	references
  	page_start
  	page_end
  	doc_type
  	lang
  	publisher
  	volume
  	issue
  	issn
  	isbn
  	doi
  	pdf
  	url
  	abstract
  }
}
`

var QueryAuthorsList = ListQuery{
	Name: "Authors",
	Variables: map[string]string{
		"$name": "string",
	},
	Function: "allofterms(name, $name)",
	Want: `
		name,
		org,
		papers: ~authors {
			id
		}
	`,
}
