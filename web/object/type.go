package object

import (
	"github.com/arianxx/aminer/internal"
	"github.com/graphql-go/graphql"
)

var langEnumType = graphql.NewEnum(graphql.EnumConfig{
	Name: "Lang",
	Values: func() graphql.EnumValueConfigMap {
		res := make(graphql.EnumValueConfigMap)
		for k, v := range internal.LangMap {
			res[k] = &graphql.EnumValueConfig{
				Value: v,
			}
		}
		return res
	}(),
})

var langInfoType = graphql.NewObject(graphql.ObjectConfig{
	Name: "LangInfo",
	Fields: graphql.Fields{
		"num": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
})

var langCountType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "LangCountType",
	Description: "语言统计类型。每个列表只有一个值，及使用该语言的 paper 数量。",
	Fields: func() graphql.Fields {
		res := make(graphql.Fields)
		for _, v := range internal.LangMap {
			res[v] = &graphql.Field{
				Type: graphql.NewNonNull(graphql.NewList(langInfoType)),
			}
		}
		return res
	}(),
})

var queryInfoType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "QueryInfoType",
	Description: "此次 query 相关信息",
	Fields: graphql.Fields{
		"num": &graphql.Field{
			Description: "结果总数",
			Type:        graphql.Int,
		},
		"lang": &graphql.Field{
			Description: "各语言数量。只有查 paper 才有此字段。",
			Type:        langCountType,
		},
	},
})

var authorsPaperType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "AuthorsPaper",
	Description: "作者发表的文章id",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Description: "Paper Id",
			Type:        graphql.String,
		},
	},
})

var authorType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Author",
	Description: "作者信息集合",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Description: "名称",
			Type:        graphql.String,
		},
		"org": &graphql.Field{
			Description: "附属机构",
			Type:        graphql.String,
		},
		"papers": &graphql.Field{
			Description: "发表的 Paper",
			Type:        graphql.NewNonNull(graphql.NewList(authorsPaperType)),
		},
	},
})

var paperType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Paper",
	Description: "Paper 信息集合",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Description: "Paper Id",
			Type:        graphql.String,
		},
		"title": &graphql.Field{
			Description: "标题",
			Type:        graphql.String,
		},
		"authors": &graphql.Field{
			Description: "作者",
			Type:        graphql.NewNonNull(graphql.NewList(authorType)),
		},
		"venue": &graphql.Field{
			Description: "发表地址",
			Type:        graphql.String,
		},
		"year": &graphql.Field{
			Description: "发表年份",
			Type:        graphql.Int,
		},
		"keywords": &graphql.Field{
			Description: "关键词",
			Type:        graphql.NewList(graphql.String),
		},
		"fos": &graphql.Field{
			Description: "研究领域",
			Type:        graphql.NewList(graphql.String),
		},
		"n_citation": &graphql.Field{
			Description: "被引数量",
			Type:        graphql.Int,
		},
		"references": &graphql.Field{
			Description: "引用资源",
			Type:        graphql.NewList(graphql.String),
		},
		"paper_start": &graphql.Field{
			Description: "开始页",
			Type:        graphql.String,
		},
		"page_end": &graphql.Field{
			Description: "结束页",
			Type:        graphql.String,
		},
		"doc_type": &graphql.Field{
			Description: "文档类型",
			Type:        graphql.String,
		},
		"lang": &graphql.Field{
			Description: "语言",
			Type:        graphql.String,
		},
		"publisher": &graphql.Field{
			Description: "出版商",
			Type:        graphql.String,
		},
		"volume": &graphql.Field{
			Description: "容量",
			Type:        graphql.String,
		},
		"issue": &graphql.Field{
			Description: "期号",
			Type:        graphql.String,
		},
		"issn": &graphql.Field{
			Description: "ISSN",
			Type:        graphql.String,
		},
		"isbn": &graphql.Field{
			Description: "ISBN",
			Type:        graphql.String,
		},
		"doi": &graphql.Field{
			Description: "Digital Object Identifier",
			Type:        graphql.String,
		},
		"pdf": &graphql.Field{
			Description: "PDF地址",
			Type:        graphql.String,
		},
		"url": &graphql.Field{
			Description: "URL",
			Type:        graphql.NewList(graphql.String),
		},
		"abstract": &graphql.Field{
			Description: "摘要",
			Type:        graphql.String,
		},
	},
})
