package v3

import "fmt"

func from(maps ...map[string]interface{}) map[string]interface{} {
	combined := make(map[string]interface{})

	for _, m := range maps {
		for k, v := range m {
			combined[k] = v
		}
	}

	return combined
}

func where(data map[string]interface{}, predicate func(string, interface{}) bool) map[string]interface{} {
	result := make(map[string]interface{})

	for key, value := range data {
		if predicate(key, value) {
			result[key] = value
		}
	}

	return result
}

func projection(m map[string]interface{}, predicate func(key string, value interface{}) bool) map[string]interface{} {
	result := make(map[string]interface{})

	for k, v := range m {
		if predicate(k, v) {
			result[k] = v
		}
	}

	return result
}

func Program() {
	table1 := map[string]interface{}{
		"field1": 1,
		"field2": 2,
	}

	table2 := map[string]interface{}{
		"field3": "one",
		"field4": "two",
	}

	// Generate the Cartesian product of multiple slices
	table3 := from(table1, table2)

	table4 := where(table3, func(key string, value interface{}) bool {
		if number, ok := value.(int); ok {
			return number == 1 && key == "field1"
		}
		return false
	})

	table5 := projection(table4, func(key string, value interface{}) bool {
		return key == "field1" || key == "field4"
	})

	// Print the result
	for key, value := range table5 {
		fmt.Printf("%s: %v\n", key, value)
	}
}
