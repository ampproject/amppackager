package v2

type (
	List[T Zone | RRSet] struct {
		Count      int  `json:"count"`
		NextOffset int  `json:"next_offset"`
		Items      []*T `json:"result"` //nolint: tagliatelle
	}

	APIError struct {
		Error string `json:"error,omitempty"`
	}
)

func (l List[T]) GetCount() int {
	return l.Count
}

func (l List[T]) GetNextOffset() int {
	return l.NextOffset
}

func (l List[T]) GetItems() []*T {
	return l.Items
}
