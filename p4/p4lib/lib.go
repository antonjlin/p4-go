package p4lib

import (
	"fmt"
	p4Config "github.com/antonjlin/p4-go/p4/config/v1"
	p4 "github.com/antonjlin/p4-go/p4/v1"
)

type ConfigBuilder struct {
	Client p4.P4RuntimeClient
	Config p4Config.P4Info
}

func CreateConfigBuilder(client p4.P4RuntimeClient) ConfigBuilder {

	config, err := p4.GetPipelineConfigs(client)
	var info = p4Config.P4Info{}
	if err != nil {
		fmt.Printf("Could not get config from client: %+v - using blank config", err)
	} else {
		info = *config
	}

	builder := ConfigBuilder{
		Client: client,
		Config: info,
	}

	return builder
}

func (builder *ConfigBuilder) AddTableEntry(table p4Config.Table){
	builder.Config.Tables = append(builder.Config.Tables, &table)
}

