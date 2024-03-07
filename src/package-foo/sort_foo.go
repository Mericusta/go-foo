package packagefoo

import (
	"slices"
	"sort"
)

type info struct {
	ID    int
	Value int
}

func SortFoo(s []int, infos []*info, reverse bool) []int {
	sort.Slice(s, func(i, j int) bool {
		iIndex := slices.IndexFunc(infos, func(_info *info) bool {
			return infos[i].ID == _info.ID
		})
		jIndex := slices.IndexFunc(infos, func(_info *info) bool {
			return infos[j].ID == _info.ID
		})
		if reverse {
			return infos[iIndex].Value > infos[jIndex].Value
		}
		return infos[iIndex].Value < infos[jIndex].Value
	})
	return s
}
