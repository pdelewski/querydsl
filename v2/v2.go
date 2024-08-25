package v2

type tuple struct {
	field1 int
	field2 string
}

type tuple2 struct {
	field1 int
	field2 string
	field3 string
}

func where[T any](data []T, predicate func(T) bool) []T {
	result := make([]T, 0)
	for _, item := range data {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

func project[T any, R any](data []T, mapper func(T) R) []R {
	result := make([]R, 0, len(data)) // Preallocate with the same length as the input slice
	for _, item := range data {
		result = append(result, mapper(item))
	}
	return result
}

func Program() {
	table := []tuple{
		{1, "a"},
		{2, "b"},
		{3, "c"},
	}

	t2 := where(table, func(t tuple) bool {
		return t.field1 > 1
	})

	t3 := project(t2, func(t tuple) tuple2 {
		return tuple2{t.field1, t.field2, "!"}
	})

	for _, t := range t3 {
		println(t.field1, t.field2, t.field3)
	}
}
