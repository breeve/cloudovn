package main

import (
	"context"
	"fmt"
	"os"

	"github.com/breeve/cloudovn/pkg/gateway/schema"
	"github.com/ovn-kubernetes/libovsdb/client"
	"github.com/ovn-kubernetes/libovsdb/ovsdb"
)

func main() {
	dbModel, _ := schema.FullDatabaseModel()
	var options []client.Option
	options = append(options, client.WithEndpoint("tcp://127.0.0.1:2222"))
	dbClient, _ := client.NewOVSDBClient(dbModel, options...)
	dbClient.Connect(context.Background())

	// case
	routeTable := &schema.RouteTable{}
	ops, _ := dbClient.Create(routeTable)
	results, _ := dbClient.Transact(context.Background(), ops...)
	opErrors, err := ovsdb.CheckOperationResults(results, ops)
	if err != nil {
		fmt.Printf("%+v %s", opErrors, err)
		os.Exit(-1)
	}
}
