package apis

import (
	"fmt"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/utils"
)

type Params interface {
	SetPathParams(...interface{}) error
}

// for ctl.
func SetPathParams(args []interface{}, ids ...interface{}) error {
	if len(args) == 0 {
		return nil
	}
	if len(args) != len(ids) {
		return fmt.Errorf("SetPathParams: args need %d items, but args len is %d", len(ids), len(args))
	}
	for i := range ids {
		switch v := ids[i].(type) {
		case *int64:
			val, err := utils.ToInt64(args[i])
			if err != nil {
				return fmt.Errorf("failed to cast to int64 `%s`: %w", args[i], err)
			}
			*v = val
		case *string:
			val, err := utils.ToString(args[i])
			if err != nil {
				return fmt.Errorf("failed to cast to string `%s`: %w", args[i], err)
			}
			*v = val
		default:
			panic(fmt.Sprintf("ids[%d] is not int64 or string", i))
		}
	}
	return nil
}

type Spec interface {
	api.Spec
	Params
}

type ListSpec interface {
	api.ListSpec
	Spec
}

type CountableListSpec interface {
	api.CountableListSpec
	Spec
}
