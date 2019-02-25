package object

import (
	"context"
	"github.com/arianxx/aminer/internal"
	"strconv"

	"github.com/graphql-go/graphql"

	"github.com/arianxx/aminer/web/model"
)

var searchConditionsInputMap = map[string]model.CondtionInput{
	"title": {
		Field: &graphql.InputObjectFieldConfig{
			Description: "Paper 题目",
			Type:        graphql.String,
		},
		Form: "anyoftext(title, \"%s\")",
	},
	"lang": {
		Field: &graphql.InputObjectFieldConfig{
			Description: "Paper 语言",
			Type:        langEnumType,
		},
		Form: "eq(lang, \"%s\")",
	},
	"keywords": {
		Field: &graphql.InputObjectFieldConfig{
			Description: "Paper 关键词",
			Type:        graphql.String,
		},
		Form: "anyofterms(keywords, \"%s\")",
	},
	"fos": {
		Field: &graphql.InputObjectFieldConfig{
			Description: "Paper fos",
			Type:        graphql.String,
		},
		Form: "anyofterms(fos, \"%s\")",
	},
	"publisher": {
		Field: &graphql.InputObjectFieldConfig{
			Description: "Paper 出版机构",
			Type:        graphql.String,
		},
		Form: "eq(publisher, \"%s\")",
	},
	"nCitationGe": {
		Field: &graphql.InputObjectFieldConfig{
			Description: "Paper 引用数量，返回大于等于此数的 paper",
			Type:        graphql.Int,
		},
		Form: "ge(n_citation, %d)",
	},
	"nCitationLe": {
		Field: &graphql.InputObjectFieldConfig{
			Description: "Paper 引用数量，返回小于等于此数的 paper",
			Type:        graphql.Int,
		},
		Form: "le(n_citation, %d)",
	},
}

var searchInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name:        "SearchCondtion",
	Description: "搜索条件",
	Fields: func() graphql.InputObjectConfigFieldMap {
		res := make(graphql.InputObjectConfigFieldMap)
		for k, v := range searchConditionsInputMap {
			res[k] = v.Field
		}
		return res
	}(),
})

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
		"SearchCondtion": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(searchInput),
			Description: "搜索条件",
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
		"orderdesc": &graphql.ArgumentConfig{
			Type:         orderFieldEnumType,
			DefaultValue: "",
			Description:  "按某一字段降序排序",
		},
		"orderasc": &graphql.ArgumentConfig{
			Type:         orderFieldEnumType,
			DefaultValue: "",
			Description:  "按某一字段升序排序",
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		searchCondition := p.Args["SearchCondtion"].(map[string]interface{})
		vars := map[string]string{
			"$first":  strconv.Itoa(p.Args["first"].(int)),
			"$offset": strconv.Itoa(p.Args["offset"].(int)),
		}

		q := model.QueryPaperList.GetQuery()
		q.SetConditions(searchCondition, searchConditionsInputMap)
		q.SetOrder(
			p.Args["orderdesc"].(string),
			p.Args["orderasc"].(string),
		)

		query, err := q.Text(model.ListTemplate)
		if err != nil {
			return nil, err
		}
		ctx := context.Background()
		var res internal.PaperList
		err = Db.GetDataWithVars(ctx, query, vars, &res)
		if err != nil {
			return nil, err
		}

		q.Name = "lang"
		q.Other = internal.LangMap
		query, err = q.Text(model.CountTemplate)
		if err != nil {
			return nil, err
		}
		err = Db.GetDataWithVars(ctx, query, vars, &res.Info[0].Lang)

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
		var res internal.PaperList
		err := Db.GetDataWithVars(ctx, model.QueryPaper, vars, &res)
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
