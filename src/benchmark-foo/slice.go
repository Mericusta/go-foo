package benchmarkfoo

import "sort"

// ----------------------------------------------------------------

// 从 slice 中删除重复元素的效率对比

// 从 slice 中删除重复元素，借助 map
func removeDuplication_map(arr []string) []string {
	set := make(map[string]struct{}, len(arr))
	j := 0
	for _, v := range arr {
		_, ok := set[v]
		if ok {
			continue
		}
		set[v] = struct{}{}
		arr[j] = v
		j++
	}

	return arr[:j]
}

// 从 slice 中删除重复元素，排序
func removeDuplication_sort(arr []string) []string {
	sort.Strings(arr)

	length := len(arr)
	if length == 0 {
		return arr
	}

	j := 0
	for i := 1; i < length; i++ {
		if arr[i] != arr[j] {
			j++
			if j < i {
				swap(arr, i, j)
			}
		}
	}

	return arr[:j+1]
}

func swap(arr []string, a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}

// 从 slice 中删除重复元素，双层循环
func simple_removeDuplication_map(arr []string) []string {
	la := make([]string, 0, len(arr))
	for _, rs := range arr {
		for _, ls := range la {
			if ls == rs {
				goto NEXT
			}
		}
		la = append(la, rs)
	NEXT:
	}
	return la
}
