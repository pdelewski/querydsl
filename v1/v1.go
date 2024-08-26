package v1

type tuple struct {
	field1 int
	field2 string
}

type table []tuple

func from(t ...table) table {
	result := make(table, 0)
	for _, tt := range t {
		result = append(result, tt...)
	}
	return result
}

func (t table) where(predicate func(tuple) bool) table {
	result := make(table, 0)
	for _, tuple := range t {
		if predicate(tuple) {
			result = append(result, tuple)
		}
	}
	return result
}

func (t table) project(selector func(tuple) tuple) table {
	result := make(table, 0)
	for _, tuple := range t {
		result = append(result, selector(tuple))
	}
	return result
}

func Program() {
	table1 := table{
		{1, "a"},
		{2, "b"},
		{3, "c"},
	}
	t1 := from(table1)

	t2 := t1.where(func(t tuple) bool {
		return t.field1 > 1
	})

	t3 := t2.project(func(t tuple) tuple {
		return tuple{t.field1, t.field2}
	})

	for _, t := range t3 {
		println(t.field1, t.field2)
	}
}
