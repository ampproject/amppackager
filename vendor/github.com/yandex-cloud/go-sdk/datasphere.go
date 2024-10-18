// Copyright (c) 2023 Yandex LLC. All rights reserved.
// Author: Ratbek Nurlanbekuulu <ratbek@yandex-team.ru>

package ycsdk

import (
	"github.com/yandex-cloud/go-sdk/gen/datasphere"
)

const (
	DatasphereServiceID Endpoint = "datasphere"
)

func (sdk *SDK) Datasphere() *datasphere.Datasphere {
	return datasphere.NewDatasphere(sdk.getConn(DatasphereServiceID))
}
