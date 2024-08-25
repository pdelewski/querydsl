package v3

import (
	"fmt"
)

func from(slices ...[]map[string]interface{}) []map[string]interface{} {
	if len(slices) == 0 {
		return nil
	}

	// Start with the first slice of maps
	result := slices[0]

	// Combine with each subsequent slice
	for i := 1; i < len(slices); i++ {
		var newResult []map[string]interface{}
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
func copyMap(m map[string]interface{}) map[string]interface{} {
	copied := make(map[string]interface{})
	for k, v := range m {
		copied[k] = v
	}
	return copied
}

// where filters a slice of maps based on a predicate function.
func where(maps []map[string]interface{}, predicate func(key string, value interface{}) bool) []map[string]interface{} {
	var result []map[string]interface{}

	for _, m := range maps {
		newMap := make(map[string]interface{})
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
func projection(maps []map[string]interface{}, selector func(map[string]interface{}) map[string]interface{}) []map[string]interface{} {
	var result []map[string]interface{}

	for _, m := range maps {
		result = append(result, selector(m))
	}

	return result
}

func Program() {
	// Example data
	table1 := []map[string]interface{}{
		{"field1": 1},
		{"field1": 2},
	}

	table2 := []map[string]interface{}{
		{"field2": "one"},
		{"field2": "two"},
	}

	table3 := []map[string]interface{}{
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
	table6 := projection(table5, func(m map[string]interface{}) map[string]interface{} {
		result := make(map[string]interface{})
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
