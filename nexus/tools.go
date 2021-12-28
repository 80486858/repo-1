package nexus

import (
	"os"
)

func interfaceSliceToStringSlice(data []interface{}) []string {
	result := make([]string, len(data))
	for i, v := range data {
		result[i] = v.(string)
	}
	return result
}

func stringSliceToInterfaceSlice(strings []string) []interface{} {
	s := make([]interface{}, len(strings))
	for i, v := range strings {
		s[i] = string(v)
	}
	return s
}

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}

// Copied from https://siongui.github.io/2018/03/09/go-match-common-element-in-two-array/
func intersection(a, b []int) (c []int) {
	m := make(map[int]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}
