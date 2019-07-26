package main

import (
	"context"
	"fmt"
	p4Config "github.com/antonjlin/p4-go/p4/config/v1"
	p4 "github.com/antonjlin/p4-go/p4/v1"
	"log"
	"time"
)

func main() {

	client := p4.GetClient("localhost:50001")
	stream, sErr := client.StreamChannel(context.Background())
	if sErr != nil {
		fmt.Println(sErr)
		log.Fatalf("cannot open stream channel with the server")
	}

	listener := p4.OpenStreamListener(stream)

	p4.SetMastership(stream)
	p4.SetPipelineConfigFromFile(client, "resources/p4info.p4")
	p4.PrintTables(client)

	table := p4Config.CreateTable()
	table.AddPreamble(
		p4Config.CreatePreamble(
			33596298,
			"FabricIngress.filtering.fwd_classifier",
			"fwd_classifier",
		),
	)

	table.AddMatchField(
		p4Config.CreateMatchField(
			1,
			"ig_port",
			nil,
			9,
			p4Config.MatchField_EXACT,
		),
	)

	table.AddMatchField(
		p4Config.CreateMatchField(
			2,
			"eth_dst",
			nil,
			48,
			p4Config.MatchField_TERNARY,
		),
	)

	table.AddMatchField(
		p4Config.CreateMatchField(
			3,
			"eth_type",
			nil,
			16,
			p4Config.MatchField_EXACT,
		),
	)

	table.AddActionRef(
		p4Config.CreateActionRef(
			16840921,
			p4Config.ActionRef_TABLE_AND_DEFAULT,
			nil,
		),
	)

	info := p4Config.P4Info{
		Tables: []*p4Config.Table{
			&table,
		},
	}

	fmt.Printf("%+v\n", info)

	p4.SetPipelineConfig(client, &info)
	time.Sleep(1000 * time.Millisecond)

	newConfig , err := p4.GetPipelineConfigs(client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", newConfig)

	listener.Wait()

}
