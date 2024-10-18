package ycsdk

import (
	loadtestingapi "github.com/yandex-cloud/go-sdk/gen/loadtesting"
	loadtestingagent "github.com/yandex-cloud/go-sdk/gen/loadtesting/agent-api"
)

const (
	LoadtestingServiceID      = "loadtesting"
	LoadtestingAgentServiceID = "loadtesting/agent-api"
)

func (sdk *SDK) Loadtesting() *loadtestingapi.Loadtesting {
	return loadtestingapi.NewLoadtesting(sdk.getConn(LoadtestingServiceID))
}

func (sdk *SDK) LoadtestingAgent() *loadtestingagent.LoadtestingAgentAPI {
	return loadtestingagent.NewLoadtestingAgentAPI(sdk.getConn(LoadtestingAgentServiceID))
}
