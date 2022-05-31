package types

type Count struct {
	Count int32 `read:"count"`
}

func (c *Count) SetCount(v int32) { c.Count = v }
func (c *Count) GetCount() int32  { return c.Count }
