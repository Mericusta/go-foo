package algorithmfoo

import (
	"testing"
)

const TestTimes = 1

func TestGenerateRoomID(t *testing.T) {
	for testTime := 0; testTime != TestTimes; testTime++ {
		// NOTE: 初始房间组检查
		if len(roomGroup) != 0 {
			t.Errorf("RoomGroup init size error, current: %d, want: %d", len(roomGroup), 0)
		}

		if len(roomIDGroup) != 0 {
			t.Errorf("RoomIDGroup init size error, current: %d, want: %d", len(roomIDGroup), 0)
		}

		doneMap := make(map[int64]int)
		for groupIndex := groupSize - 1; groupIndex >= 0; groupIndex-- {
			// fmt.Printf("groupIndex = %d\n", groupIndex)
			var groupStartRoomID int64
			var groupEndRoomID int64

			// NOTE: RoomIDGroupSize 次数随机拿取房间号
			for roomIndex := 0; roomIndex < roomIDGroupSize; roomIndex++ {
				var beforeRoomGroup = currentRoomGroup
				var roomID = GenerateRoomID()

				// NOTE: 生成房间号边界检查
				if roomID < minRoomID || maxRoomID < roomID {
					t.Errorf("roomID range error, roomID == %d, want: [%d, %d]", roomID, minRoomID, maxRoomID)
				}

				// NOTE: 当房间分组发生变化时，更新起始结束组号
				if beforeRoomGroup != currentRoomGroup {
					groupStartRoomID = minRoomID + int64(roomIDGroupSize*currentRoomGroup)
					groupEndRoomID = groupStartRoomID + int64(roomIDGroupSize)

					// NOTE: 由于最后一组只有99个数字，所以检查的时候遇到的时候 roomIndex 需要加1, groupEndRoomID 需要减1
					if groupEndRoomID > maxRoomID {
						roomIndex++
						groupEndRoomID--
					}
				}

				if groupStartRoomID < minRoomID || maxRoomID < groupStartRoomID {
					t.Errorf("groupStartRoomID error, groupStartRoomID = %d, want: [%d, %d]", groupStartRoomID, minRoomID, maxRoomID)
				}

				if groupEndRoomID < minRoomID || maxRoomID < groupEndRoomID {
					t.Errorf("groupEndRoomID error, groupEndRoomID = %d, want: [%d, %d]", groupEndRoomID, minRoomID, maxRoomID)
				}

				// NOTE: 生成房间号属于当前组检查
				if !(groupStartRoomID <= roomID && roomID <= groupEndRoomID) {
					t.Errorf("roomID group range error, roomID == %d, want: [%d, %d]", roomID, groupStartRoomID, groupEndRoomID)
				}

				// NOTE: 房间组检查
				if (roomIDGroupSize - roomIndex - 1) != len(roomIDGroup) {
					t.Errorf("RoomIDGroup size error, current: %d, want: %d", len(roomIDGroup), roomIndex)
				}

				doneMap[roomID]++
				if doneMap[roomID] > 1 {
					panic(roomID)
				} else if len(doneMap) > 9000 {
					panic(doneMap)
				}
			}

			// NOTE: 房间组长度检查
			if groupIndex != len(roomGroup) {
				t.Errorf("RoomGroup size error, current: %d, want: %d", len(roomGroup), groupIndex)
			}

			t.Log(doneMap)
		}
	}
}
