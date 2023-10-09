package algorithmfoo

import (
	"math/rand"
)

const (
	groupSize       = 90   // 组数量
	roomIDGroupSize = 100  // 分组大小
	minRoomID       = 1000 // 最小房间号
	maxRoomID       = 9999 // 最大房间号
)

var (
	roomGroup        []int   // 还未分配的组
	currentRoomGroup int     // 当前房间组号
	roomIDGroup      []int64 // 当前已分组 ID
)

func init() {
	// 初始化房间组号
	currentRoomGroup = -1
}

// 简单防猜离散ID算法
// 按总数划分组 -> 取随机组 -> 按组大小划分组ID -> 取组内随机ID并删除 -> 组ID消耗尽删除组
func GenerateRoomID() int64 {
	// re-assign room ID group
	if len(roomIDGroup) == 0 {
		// re-assign group slice
		if len(roomGroup) == 0 {
			for index := 0; index != 90; index++ {
				roomGroup = append(roomGroup, index)
			}
		}

		// pick a group and remove it from group slice
		var roomGroupIndex = rand.Intn(len(roomGroup))
		currentRoomGroup = roomGroup[roomGroupIndex]
		roomGroup = append(roomGroup[:roomGroupIndex], roomGroup[roomGroupIndex+1:]...)

		// calculate room ID in group range
		var startRoomID int64 = minRoomID + int64(roomIDGroupSize*currentRoomGroup)
		var endRoomID int64 = startRoomID + int64(roomIDGroupSize)
		if endRoomID > maxRoomID {
			endRoomID = maxRoomID
		}

		// makeup room ID group
		for index := startRoomID; index != endRoomID; index++ {
			roomIDGroup = append(roomIDGroup, int64(index))
		}
	}

	// random pick room ID group index and remove it from group
	var randomIndex = rand.Int31n(int32(len(roomIDGroup)))
	var roomID = roomIDGroup[randomIndex]
	roomIDGroup = append(roomIDGroup[:randomIndex], roomIDGroup[randomIndex+1:]...)

	return roomID
}

func LocalUniqueID() {
	
}
