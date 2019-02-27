package model

import (
	"bytes"
	"fmt"
	"github.com/graphql-go/graphql"
	"text/template"
)

type CondtionInput struct {
	Field *graphql.InputObjectFieldConfig
	Form  string
}

type ListQuery struct {
	Name      string
	Variables map[string]string
	Function  string
	Filters   []string
	Orderdesc string
	Orderasc  string
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

func (q *ListQuery) SetFilters(f string) {
	if !q.copied {
		return
	}

	q.Filters = append(q.Filters, f)
}

func (q *ListQuery) SetConditions(m map[string]interface{}, s map[string]CondtionInput) {
	if !q.copied {
		return
	}

	for k, c := range s {
		if v, ok := m[k]; ok {
			var f string
			switch v.(type) {
			case int:
				f = fmt.Sprintf(c.Form, v.(int))
			case string:
				f = fmt.Sprintf(c.Form, v.(string))
			}

			if len(q.Function) == 0 {
				q.Function = f
			} else {
				q.SetFilters(f)
			}
		}
	}
}

func (q *ListQuery) SetOrder(desc, asc string) {
	if !q.copied {
		return
	}

	q.Orderdesc, q.Orderasc = desc, asc
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
