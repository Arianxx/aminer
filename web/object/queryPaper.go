package object

import (
	"encoding/json"
	"strconv"

	"github.com/graphql-go/graphql"

	"github.com/arianxx/aminer/web/model"
)

var queryPapers = &graphql.Field{
	Description: "获取指定的 Paper 列表",
	Type:        graphql.NewList(paperType),
	Args: graphql.FieldConfigArgument{
		"title": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "待搜索的 Paper 题目",
		},
		"offset": &graphql.ArgumentConfig{
			Type:         graphql.Int,
			DefaultValue: 0,
			Description:  "从何处开始查询",
		},
		"first": &graphql.ArgumentConfig{
			Type:         graphql.Int,
			DefaultValue: 10,
			Description:  "查询数目",
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		vars := map[string]string{
			"$title":  p.Args["title"].(string),
			"$first":  strconv.Itoa(p.Args["first"].(int)),
			"$offset": strconv.Itoa(p.Args["offset"].(int)),
		}
		resJson, err := Db.QueryWithVars(model.QueryPaperList, vars)
		if err != nil {
			return nil, err
		}
		var res model.PaperList
		json.Unmarshal(resJson, &res)
		return res.Data, nil
	},
}

var queryPaper = &graphql.Field{
	Description: "获取指定的 Paper 信息",
	Type:        paperType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Description: "要获取的 Paper id",
			Type:        graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		vars := map[string]string{
			"$id": p.Args["id"].(string),
		}
		resJson, err := Db.QueryWithVars(model.QueryPaper, vars)
		if err != nil {
			return nil, err
		}
		var res model.PaperList
		json.Unmarshal(resJson, &res)
		if len(res.Data) > 0 {
			return res.Data[0], nil
		} else {
			return nil, nil
		}
	},
}
