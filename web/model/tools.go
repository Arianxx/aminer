package model

import (
	"github.com/arianxx/aminer/internal"
	"context"
	"github.com/dgraph-io/dgo"
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
