package object

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/arianxx/aminer/internal"
	"strconv"

	"github.com/graphql-go/graphql"

	"github.com/arianxx/aminer/web/model"
)

var queryPapers = &graphql.Field{
	Description: "获取指定的 Paper 列表",
	Type: graphql.NewObject(graphql.ObjectConfig{
		Name: "PapersInfo",
		Fields: graphql.Fields{
			"data": &graphql.Field{
				Type: graphql.NewList(paperType),
			},
			"info": &graphql.Field{
				Type: graphql.NewNonNull(graphql.NewList(queryInfoType)),
			},
		},
	}),
	Args: graphql.FieldConfigArgument{
		"title": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.String),
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
		"lang": &graphql.ArgumentConfig{
			Type:        langEnumType,
			Description: "过滤语言",
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		vars := map[string]string{
			"$title":  p.Args["title"].(string),
			"$first":  strconv.Itoa(p.Args["first"].(int)),
			"$offset": strconv.Itoa(p.Args["offset"].(int)),
		}

		q := model.QueryPaperList.GetQuery()
		if lang, ok := p.Args["lang"]; ok {
			q.SetFilters([]string{
				fmt.Sprintf("eq(lang, %s)", lang.(string)),
			})
		}

		query, err := q.Text(model.ListTemplate)
		if err != nil {
			return nil, err
		}

		ctx := context.Background()
		resJson, err := Db.QueryWithVars(ctx, query, vars)
		if err != nil {
			return nil, err
		}
		var res internal.PaperList
		err = json.Unmarshal(resJson, &res)
		if err != nil {
			return nil, err
		}

		q.Name = "lang"
		q.Other = internal.LangMap
		query, err = q.Text(model.CountTemplate)
		if err != nil {
			return nil, err
		}

		resJson, err = Db.QueryWithVars(ctx, query, vars)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(resJson, &res.Info[0].Lang)
		if err != nil {
			return nil, err
		}

		return res, nil
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
		ctx := context.Background()
		resJson, err := Db.QueryWithVars(ctx, model.QueryPaper, vars)
		if err != nil {
			return nil, err
		}
		var res internal.PaperList
		err = json.Unmarshal(resJson, &res)
		if err != nil {
			return nil, err
		}
		if len(res.Data) > 0 {
			return res.Data[0], nil
		} else {
			return nil, nil
		}
	},
}
