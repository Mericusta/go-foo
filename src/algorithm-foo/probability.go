package algorithmfoo

import "math/rand"

// PutBackRandom 放回随机
func PutBackRandom[T comparable](count int, rangeFunc func(func(T, int) bool)) []T {
	result := make([]T, 0, count)

	var totalWeight int
	rangeFunc(func(item T, weight int) bool {
		totalWeight += weight
		return true
	})
	if totalWeight == 0 {
		return result
	}

	for index := 0; index != count; index++ {
		randomNum := rand.Intn(totalWeight) + 1 // [0, totalWeight) -> [1, totalWeight]
		rangeFunc(func(item T, weight int) bool {
			if randomNum <= weight {
				result = append(result, item)
				return false
			} else {
				randomNum -= weight
			}
			return true
		})
	}

	return result
}

// PutAwayRandom 不放回随机
func PutAwayRandom[T comparable](count int, rangeFunc func(func(T, int) bool)) []T {
	result := make([]T, 0, count)

	// 计算总权重
	var totalWeight int
	idWeightMap := make(map[T]int)
	rangeFunc(func(ID T, weight int) bool {
		totalWeight += weight
		idWeightMap[ID] = weight
		return true
	})

	for index := 0; index != count; index++ {
		// 总权重消耗光，不再抽取
		if totalWeight <= 0 {
			continue
		}

		// 总权重随机
		n := rand.Intn(totalWeight)
		var selected bool
		var selectedID T
		var selectedWeight int
		// 减法得出随机值所在区间
		for ID, weight := range idWeightMap {
			if weight <= 0 {
				continue
			}
			if n >= weight {
				n -= weight
			} else {
				selected = true
				selectedID = ID
				selectedWeight = weight
				break
			}
		}

		// 没有随出则报错
		if !selected {
			return nil
		}

		// 添加结果
		result = append(result, selectedID)
		// 减少总权重
		totalWeight -= selectedWeight
		// 不放回
		delete(idWeightMap, selectedID)
	}

	return result
}
