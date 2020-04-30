package collection

import "fmt"

type BaseCollection struct {
	data []interface{}
}

func (b *BaseCollection) ToString() string {
	s := "["
	for _, d := range b.data {
		s += fmt.Sprintf("%+v,", d)
	}
	return s[:len(s)-1] + "]"
}

func (b *BaseCollection) ForEach(l func(interface{})) {
	for _, d := range b.data {
		l(d)
	}
}

func (b *BaseCollection) Map(l ...func(element interface{}) interface{}) []interface{} {
	for _, l1 := range l {
		newData := []interface{}{}
		for _, d := range b.data {
			newData = append(newData, l1(d))
		}
		b.data = newData
	}

	return b.data
}
