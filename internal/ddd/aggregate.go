package ddd

type Aggregate struct {
	aggregateId   string
	aggregateName string
	events        []Event
}

func NewAggregate(id string, name string) *Aggregate {
	return &Aggregate{
		aggregateId:   id,
		aggregateName: name,
		events:        make([]Event, 0),
	}
}

func (a *Aggregate) AggregateId() string {
	return a.aggregateId
}

func (a *Aggregate) AggregateName() string {
	return a.aggregateName
}

func (a *Aggregate) Events() []Event {
	return a.events
}

func (a *Aggregate) ClearEvents() {
	a.events = []Event{}
}

func (a *Aggregate) AddEvent(name string, payload any, options ...EventOption) {
	options = append(options, Metadata{
		AggregateIdKey:   a.aggregateId,
		AggregateNameKey: a.aggregateName,
	})
	a.events = append(a.events, NewEvent(name, payload, options...))
}
