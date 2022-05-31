package types

import (
	"encoding/json"
	"time"
)

func ParseTime(layout, value string) (Time, error) {
	var err error
	res := Time{}
	res.Time, err = time.Parse(layout, value)
	if err != nil {
		return res, err
	}
	return res, nil
}

// +k8s:deepcopy-gen=false
type Time struct {
	time.Time
}

func (i Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.Time)
}

func (i *Time) UnmarshalJSON(data []byte) error {
	if string(data) == `""` {
		return nil
	}
	var err error
	i.Time, err = time.Parse(`"`+time.RFC3339Nano+`"`, string(data))
	return err
}

func (i *Time) DeepCopyInto(in *Time) {
	i.Time = in.Time
}
