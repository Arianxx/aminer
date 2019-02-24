package model

import (
	"context"
	"encoding/json"
	"github.com/arianxx/aminer/internal"
	"github.com/dgraph-io/dgo"
	"io/ioutil"
	"strings"
)

type Db struct {
	Conn   *dgo.Dgraph
	Cancel func() error
}

func InitDgraphConn(host string) (*Db) {
	conn, cancel := internal.GetDgraphClient(host)
	return &Db{conn, cancel}
}

func (d *Db) QueryWithVars(ctx context.Context, query string, vars map[string]string) ([]byte, error) {
	resp, err := d.Conn.NewTxn().QueryWithVars(ctx, query, vars)
	if err != nil {
		return nil, err
	}
	return resp.Json, nil
}

func (d *Db) Query(ctx context.Context, query string) ([]byte, error) {
	resp, err := d.Conn.NewTxn().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	return resp.Json, nil
}

func (d *Db) GetDataWithVars(ctx context.Context, query string, vars map[string]string, res interface{}) error {
	resJson, err := d.QueryWithVars(ctx, query, vars)
	if err != nil {
		return err
	}
	err = json.Unmarshal(resJson, res)
	if err != nil {
		return err
	}
	return nil
}

func GetPaperFilePaths() ([]string, error) {
	paths := []string{internal.AMINER_PAPERS_DIRECTORY, internal.MAG_PAPERS_DIRECTORY}
	res := []string{}
	for _, path := range paths {
		files, err := ioutil.ReadDir(path)
		if err != nil {
			return nil, err
		}
		for _, f := range files {
			res = append(res, strings.Join([]string{path, f.Name()}, "/"))
		}
	}
	return res, nil
}
