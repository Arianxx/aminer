package internal

import (
	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"google.golang.org/grpc"
	"log"
)

const (
	AMINER_PAPERS_DIRECTORY = "./aminer_papers"
	MAG_PAPERS_DIRECTORY    = "./mag_papers"
)

func GetDgraphClient(host string) (*dgo.Dgraph, func() error) {
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	dc := api.NewDgraphClient(conn)
	dg := dgo.NewDgraphClient(dc)

	return dg, conn.Close
}
