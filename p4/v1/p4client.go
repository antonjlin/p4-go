package p4_v1

import (
	"context"
	"fmt"
	p4config "github.com/antonjlin/p4-go/p4/config/v1"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"io"
	"io/ioutil"
	"log"
	"sync"
)

//protoc --go_out=paths=source_relative:. --proto_path=proto/ proto/p4-go/status.proto
//protoc --go_out=plugins=grpc:.  --proto_path=proto/ proto/p4-go/v1/p4runtime.proto

func GetClient(host string) P4RuntimeClient {
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return NewP4RuntimeClient(conn)
}

func GetPipelineConfigs(client P4RuntimeClient) (*p4config.P4Info, error) {
	getReq := &GetForwardingPipelineConfigRequest{
		DeviceId:     1,
		ResponseType: GetForwardingPipelineConfigRequest_ALL,
	}
	reply, err := client.GetForwardingPipelineConfig(context.Background(), getReq)
	if err != nil {
		fmt.Println(err)
		return nil, errors.Wrap(err, "error getting pipeline config")
	}
	return reply.GetConfig().P4Info, nil
}



func PrintTables(client P4RuntimeClient) {
	res, err := GetPipelineConfigs(client)
	if err != nil {
		log.Fatal(err)
	}
	for i, table := range res.Tables {
		fmt.Printf("%v - %+v \n", i, table.Preamble.GetName()) //with name, value and type
	}
}

func SetPipelineConfigFromFile(client P4RuntimeClient, filename string) {

	config, err := GetConfigFromFile(filename)
	if err != nil {
		log.Fatalf("Could not get config from file")
	}
	SetPipelineConfig(client, config)
}

func SetPipelineConfig(client P4RuntimeClient, config *p4config.P4Info) {

	req := &SetForwardingPipelineConfigRequest{
		Action:   SetForwardingPipelineConfigRequest_VERIFY_AND_COMMIT,
		DeviceId: 1, RoleId: 2,
		ElectionId: &Uint128{High: 10000, Low: 9999},
		Config:     &ForwardingPipelineConfig{P4Info: config},
	}

	_, err1 := client.SetForwardingPipelineConfig(context.Background(), req)

	if err1 != nil {
		fmt.Println("\nError with pipeline req")
		fmt.Println(err1)
	}
}

func OpenStreamListener(stream P4Runtime_StreamChannelClient) sync.WaitGroup {

	var waitg sync.WaitGroup
	waitg.Add(1)

	go func() {
		for {
			inData, err := stream.Recv()
			if err == io.EOF {
				waitg.Done()
				return
			}
			if err != nil {
				fmt.Printf("[STREAM-ERROR] (%T) : %+v\n", err, err)
			}
			if inData != nil {
				fmt.Printf("[STREAM-INCOMING] (%T) : %+v\n", inData, inData)

			}
			// Act on the received message
		}
	}()

	return waitg
}

func SetMastership(stream P4Runtime_StreamChannelClient) {
	req := StreamMessageRequest{
		Update: &StreamMessageRequest_Arbitration{
			Arbitration: &MasterArbitrationUpdate{
				DeviceId: 1,
				Role: &Role{
					Id: 2,
				},
				ElectionId: &Uint128{High: 10000, Low: 9999},
			},
		},
	}

	err := stream.Send(&req)
	if err != nil {
		fmt.Println("ERROR SENDING STREAM REQUEST:")
		fmt.Println(err)
	}
}

func GetConfigFromFile(filename string) (*p4config.P4Info, error) {

	i := p4config.P4Info{}

	if d, err := ioutil.ReadFile(filename); err != nil {
		return nil, errors.Wrapf(err, "error reading p4info file: %s", filename)
	} else if err := proto.UnmarshalText(string(d), &i); err != nil {
		return nil, errors.Wrapf(err, "error parsing p4info file: %s", filename)
	}
	return &i, nil
}
