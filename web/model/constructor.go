package model

import (
	"bytes"
	"fmt"
	"text/template"
)

type ListQuery struct {
	Name      string
	Variables map[string]string
	Function  string
	Filters   []string
	Want      string
	Other     map[string]string

	copied bool
}

func (q *ListQuery) GetQuery() ListQuery {
	res := *q
	res.Variables, res.Filters, res.Other = make(map[string]string), []string{}, make(map[string]string)
	for k, v := range q.Variables {
		res.Variables[k] = v
	}
	for k, v := range q.Other {
		res.Other[k] = v
	}
	res.Filters = append(res.Filters, q.Filters...)
	res.copied = true
	return res
}

func (q *ListQuery) SetVariables(m map[string]string) {
	if !q.copied {
		return
	}

	for k, v := range m {
		q.Variables[k] = v
	}
}

func (q *ListQuery) SetFilters(f []string) {
	if !q.copied {
		return
	}

	q.Filters = append(q.Filters, f...)
}

func (q *ListQuery) Text(t string) (string, error) {
	tpl, err := template.New("listTemplate").Parse(t)
	if err != nil {
		return "", fmt.Errorf("Error when parsing query text: %s", err)
	}

	var res bytes.Buffer
	err = tpl.Execute(&res, q)
	if err != nil {
		return "", fmt.Errorf("Error when formating query text: %s", err)
	}

	return res.String(), nil
}
