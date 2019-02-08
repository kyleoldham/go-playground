package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 4, 2, 8}
	fmt.Println(SliceUniqMap(arr))
}

func SliceUniqMap(s []int) []int {
	seen := make(map[int]struct{}, len(s))
	j := 0
	for _, v := range s {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		s[j] = v
		j++
	}
	return s[:j]
}
