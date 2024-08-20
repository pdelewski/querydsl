package v3

import "fmt"

type tuple struct {
	field1 int
	field2 string
}

type table struct {
	tuples []tuple
	ir     []string
}

func (table) from(t ...table) table {
	tuples := make([]tuple, 0)
	result := table{tuples: tuples}
	for _, tt := range t {
		result.tuples = append(result.tuples, tt.tuples...)
		result.ir = append(tt.ir, "from")
	}

	return result
}

func (t table) where(predicate func(tuple) bool) table {
	tuples := make([]tuple, 0)
	result := table{tuples: tuples}
	for _, tuple := range t.tuples {
		if predicate(tuple) {
			result.tuples = append(result.tuples, tuple)
		}
	}
	result.ir = append(t.ir, "where")

	return result
}

func (t table) project(selector func(tuple) tuple) table {
	tuples := make([]tuple, 0)
	result := table{tuples: tuples}
	for _, tuple := range t.tuples {
		result.tuples = append(result.tuples, selector(tuple))
	}
	result.ir = append(t.ir, "project")

	return result
}

func (t table) emit() {
	for _, i := range t.ir {
		fmt.Println(i)
	}
}

type QueryDsl interface {
	from(t ...table) table
	where(predicate func(tuple) bool) table
	project(selector func(tuple) tuple) table
	emit()
}

func Program() {
	var t0 QueryDsl

	t0 = table{}

	t1 := t0.from(table{tuples: []tuple{
		{1, "a"},
		{2, "b"},
		{3, "c"},
	}})

	t2 := t1.where(func(t tuple) bool {
		return t.field1 > 1
	})

	t3 := t2.project(func(t tuple) tuple {
		return tuple{t.field1, t.field2 + "!"}
	})

	t3.emit()
}
