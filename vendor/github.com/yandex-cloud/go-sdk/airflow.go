package ycsdk

import (
	"github.com/yandex-cloud/go-sdk/gen/airflow"
)

const AirflowServiceID = "managed-airflow"

func (sdk *SDK) Airflow() *airflow.Airflow {
	return airflow.NewAirflow(sdk.getConn(AirflowServiceID))
}
