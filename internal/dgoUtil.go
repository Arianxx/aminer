package internal

import (
	"io/ioutil"
	"log"
	"strings"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"google.golang.org/grpc"
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

func GetPaperFilePaths() ([]string, error) {
	paths := []string{AMINER_PAPERS_DIRECTORY, MAG_PAPERS_DIRECTORY}
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
