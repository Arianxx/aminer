package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"

	"github.com/arianxx/aminer/internal"
)

func storeMultiFile(conn *dgo.Dgraph, paths []string, concurrency int) error {
	sem := getSem(concurrency)
	for _, path := range paths {
		sem.acquire()
		go func(path string) {
			defer sem.release()
			if err := storeDataFromFile(conn, path); err != nil {
				fmt.Println(err)
			}
		}(path)
	}
	sem.wait()
	return nil
}

// storeDataFromFile 将目录为 filepath 的文件里面的 json 数据存入 dgraph
func storeDataFromFile(conn *dgo.Dgraph, filepath string) error {
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	dec := json.NewDecoder(f)

	s := getSem(50)
	for dec.More() {
		var p internal.Paper
		err := dec.Decode(&p)
		if err != nil {
			fmt.Println(err)
			return err
		}

		s.acquire()
		go func(p *internal.Paper) {
			defer s.release()

			err = storePaper(conn, p)
			if err != nil {
				fmt.Println(err)
			}
		}(&p)
	}
	s.wait()

	fmt.Printf("\n-------------")
	fmt.Printf("A file %s has been processed!", filepath)
	fmt.Printf("\n-------------")

	return nil
}

// replaceAuthor 查询 author 是否已经存在，如果存在就获取他的 uid 并存入结构体
func replaceAuthor(conn *dgo.Dgraph, author *internal.Author) error {
	ctx := context.Background()
	var authorVars map[string]string
	var r authorRoot

	authorVars = make(map[string]string)
	authorVars["$name"], authorVars["$org"] = author.Name, author.Org
	resp, err := conn.NewTxn().QueryWithVars(ctx, queryAuthorUid, authorVars)
	if err != nil {
		return err
	}
	err = json.Unmarshal(resp.Json, &r)
	if err != nil {
		return err
	}

	if len(r.data) > 0 {
		author.Uid = r.data[0].Uid
	}

	return nil
}

// storePaper 将 paper 存入 dgraph
func storePaper(conn *dgo.Dgraph, paper *internal.Paper) error {
	for _, author := range paper.Authors {
		if err := replaceAuthor(conn, &author); err != nil {
			return err
		}
	}

	mu := &api.Mutation{}
	paperJson, err := json.Marshal(paper)
	if err != nil {
		return err
	}

	ctx := context.Background()
	mu.SetJson = paperJson
	mu.CommitNow = true

	resp, err := conn.NewTxn().Mutate(ctx, mu)
	if err != nil {
		return err
	}

	fmt.Println(resp)

	return nil
}
