package internal

type AuthorsPaper struct {
	Id string `json:"id"`
}

type Author struct {
	Uid   string         `json:"uid,omitempty"`
	Name  string         `json:"name,omitempty"`
	Org   string         `json:"org,omitempty"`
	Paper []AuthorsPaper `json:"papers,omitempty"`
}

type Paper struct {
	Uid        string   `json:"uid,omitempty"`
	Id         string   `json:"id,omitempty"`
	Title      string   `json:"title,omitempty"`
	Authors    []Author `json:"authors,omitempty"`
	Venue      string   `json:"venue,omitempty"`
	Year       int      `json:"year,omitempty"`
	Keywords   []string `json:"keywords,omitempty"`
	Fos        []string `json:"fos,omitempty"`
	NCitation  int      `json:"n_citation,omitempty"`
	References []string `json:"references,omitempty"`
	PageStart  string   `json:"paper_start,omitempty"`
	PageEnd    string   `json:"page_end,omitempty"`
	DocType    string   `json:"doc_type,omitempty"`
	Lang       string   `json:"lang,omitempty"`
	Publisher  string   `json:"publisher,omitempty"`
	Volume     string   `json:"volume,omitempty"`
	Issue      string   `json:"issue,omitempty"`
	Issn       string   `json:"issn,omitempty"`
	ISBN       string   `json:"isbn,omitempty"`
	Doi        string   `json:"doi,omitempty"`
	Pdf        string   `json:"pdf,omitempty"`
	Url        []string `json:"url,omitempty"`
	Abstract   string   `json:"abstract,omitempty"`
}
