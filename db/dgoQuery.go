package main

import "github.com/arianxx/aminer/internal"

var queryAuthorUid = `
query Author($name: string, $org: string){
  data(func: eq(name, $name))@filter(eq(org, $org)){
    uid
  }
}
`

type authorRoot struct {
	data []internal.Paper `json:"data"`
}
