package model

import (
	"github.com/arianxx/aminer/internal"
	"context"
	"github.com/dgraph-io/dgo"
)

type Db struct {
	conn   *dgo.Dgraph
	Cancel func() error
}

func InitDgraphConn(host string) (*Db) {
	conn, cancel := internal.GetDgraphClient(host)
	return &Db{conn, cancel}
}

func (d *Db) QueryWithVars(query string, vars map[string]string) ([]byte, error) {
	resp, err := d.conn.NewTxn().QueryWithVars(context.Background(), query, vars)
	if err != nil {
		return nil, err
	}
	return resp.Json, nil
}