package v3

import (
	"fmt"
)

type tuple map[string]interface{}
type table []map[string]interface{}

func from(slices ...table) table {
	if len(slices) == 0 {
		return nil
	}

	// Start with the first slice of maps
	result := slices[0]

	// Combine with each subsequent slice
	for i := 1; i < len(slices); i++ {
		var newResult table
		for _, r := range result {
			for _, t := range slices[i] {
				// Combine two maps into a new one
				combined := copyMap(r)
				for k, v := range t {
					combined[k] = v
				}
				newResult = append(newResult, combined)
			}
		}
		result = newResult
	}

	return result
}

// copyMap creates a deep copy of a map.
func copyMap(m tuple) tuple {
	copied := make(tuple)
	for k, v := range m {
		copied[k] = v
	}
	return copied
}

// where filters a slice of maps based on a predicate function.
func where(maps table, predicate func(key string, value interface{}) bool) table {
	var result table

	for _, m := range maps {
		newMap := make(tuple)
		for k, v := range m {
			if predicate(k, v) {
				newMap[k] = v
			}
		}
		if len(newMap) > 0 {
			result = append(result, newMap)
		}
	}

	return result
}

// projection applies a transformation function to each map in the slice.
func projection(maps table, selector func(tuple) tuple) table {
	var result table

	for _, m := range maps {
		result = append(result, selector(m))
	}

	return result
}

func Program() {
	// Example data
	table1 := table{
		{"field1": 1},
		{"field1": 2},
	}

	table2 := table{
		{"field2": "one"},
		{"field2": "two"},
	}

	table3 := table{
		{"field3": 3.14},
		{"field3": 2.71},
	}

	table4 := from(table1, table2, table3)

	table5 := where(table4, func(key string, value interface{}) bool {
		if v, ok := value.(int); ok && key == "field1" {
			return v > 1
		}
		return false
	})

	// Project maps to include only specific columns
	table6 := projection(table5, func(m tuple) tuple {
		result := make(tuple)
		if val, ok := m["field1"]; ok {
			result["field1"] = val
		}
		if val, ok := m["field2"]; ok {
			result["field2"] = val
		}
		return result
	})
	fmt.Println("projection:", table6)
}
