package collection

type List struct {
	BaseCollection
}

func NewList() *List {
	return &List{
		BaseCollection{
			data: []interface{}{},
		},
	}
}

func (l *List) Add(e interface{}) {
	l.data = append(l.data, e)
}

func ListOf(e ...interface{}) *List {
	l := NewList()
	for _, e1 := range e {
		l.Add(e1)
	}
	return l
}

func (l *List) ForEach(lambda func(e interface{})) {
	l.BaseCollection.ForEach(lambda)
}

func (l *List) Map(lambda ...func(e interface{}) interface{}) *List {
	l.BaseCollection.Map(lambda...)
	return l
}
