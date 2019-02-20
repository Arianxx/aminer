package object

import (
	"encoding/json"
	"strconv"

	"github.com/graphql-go/graphql"

	"github.com/arianxx/aminer/web/model"
)

var queryAuthors = &graphql.Field{
	Description: "获取指定的 Authors 列表",
	Type:        graphql.NewList(authorType),
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type:        graphql.String,
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
		resJson, err := Db.QueryWithVars(model.QueryAuthorsList, vars)
		if err != nil {
			return nil, err
		}
		var res model.AuthorList
		json.Unmarshal(resJson, &res)
		return res.Data, nil
	},
}
