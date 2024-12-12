package tools

import (
	"bytes"
	"encoding/json"
	"fmt"
	mcpgolang "github.com/metoro-io/mcp-golang"
	"github.com/metoro-io/metoro-mcp-server/model"
	"github.com/metoro-io/metoro-mcp-server/utils"
)

type GetK8sServiceInformationHandlerArgs struct {
	TimeConfig   utils.TimeConfig `json:"time_config" jsonschema:"required,description=The time to get state of the service. e.g. if you want to see the state of the service 5 minutes ago, you would set time_period=5 and time_window=Minutes"`
	ServiceName  string           `json:"serviceName" jsonschema:"required,description=The name of the service to get information for"`
	Environments []string         `json:"environments" jsonschema:"description=The environments to get information for. If empty, all environments will be used."`
}

func GetK8sServiceInformationHandler(arguments GetK8sServiceInformationHandlerArgs) (*mcpgolang.ToolResponse, error) {
	startTime, endTime := utils.CalculateTimeRange(arguments.TimeConfig)
	request := model.GetPodsRequest{
		StartTime:    startTime,
		EndTime:      endTime,
		ServiceName:  arguments.ServiceName,
		Environments: arguments.Environments,
	}
	jsonBody, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %v", err)
	}

	resp, err := utils.MakeMetoroAPIRequest("POST", "k8s/summary", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("error making Metoro call: %v", err)
	}

	return mcpgolang.NewToolReponse(mcpgolang.NewTextContent(fmt.Sprintf("%s", string(resp)))), nil
}
