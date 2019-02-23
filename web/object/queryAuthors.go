package object

import (
	"context"
	"encoding/json"
	"github.com/arianxx/aminer/internal"
	"strconv"

	"github.com/graphql-go/graphql"

	"github.com/arianxx/aminer/web/model"
)

var queryAuthors = &graphql.Field{
	Description: "获取指定的 Authors 列表",
	Type: graphql.NewObject(graphql.ObjectConfig{
		Name: "AuthorsInfo",
		Fields: graphql.Fields{
			"data": &graphql.Field{
				Type: graphql.NewList(authorType),
			},
			"info": &graphql.Field{
				Type: graphql.NewNonNull(graphql.NewList(queryInfoType)),
			},
		},
	}),
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "待搜索的 Authors 名称",
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
			"$name":   p.Args["name"].(string),
			"$first":  strconv.Itoa(p.Args["first"].(int)),
			"$offset": strconv.Itoa(p.Args["offset"].(int)),
		}
		query, err := model.QueryAuthorsList.Text(model.ListTemplate)
		if err != nil {
			return nil, err
		}

		ctx := context.Background()
		resJson, err := Db.QueryWithVars(ctx, query, vars)
		if err != nil {
			return nil, err
		}
		var res internal.AuthorList
		err = json.Unmarshal(resJson, &res)
		if err != nil {
			return nil, err
		}
		return res, nil
	},
}
