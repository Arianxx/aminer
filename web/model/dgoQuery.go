package model

import "github.com/arianxx/aminer/internal"

type info struct {
	Num int `json:"num"`
}

type PaperList struct {
	Data []internal.Paper `json:"data"`
	Info []info           `json:"info"`
}

type AuthorList struct {
	Data []internal.Author `json:"data"`
	Info []info            `json:"info"`
}

var QueryPaperList = `
query Paper($title: string, $offset: int, $first: int){
  data(func: alloftext(title, $title), first: $first, offset: $offset){
    expand(_all_){
		name
		org
		papers: ~authors {
			id
		}
    }
  }
  info(func: alloftext(title, $title)){
	num: count(uid)
  }
}
`

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

var QueryAuthorsList = `
query Authors($name: string, $offset: int, $first: int){
  data(func: allofterms(name, $name), first: $first, offset: $offset){
    name,
	org,
	papers: ~authors {
		id
	}
  }
  info(func: allofterms(name, $name)){
	num: count(uid)
  }
}
`
