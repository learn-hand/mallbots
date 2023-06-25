package ddd

const (
	AggregateIdKey   = "aggregate_id"
	AggregateNameKey = "aggregate_name"
)

type Metadata map[string]any

func (m Metadata) configureEvent(e *Event) {
	for key, value := range m {
		e.metadata[key] = value
	}
}

func (m Metadata) Get(key string) any {
	return m[key]
}
