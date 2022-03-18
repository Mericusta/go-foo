package algorithmfoo

import (
	"testing"
	"time"
)

func TestAntiAddictionData_GetReleaseOnlineTime(t *testing.T) {
	// 中国防沉迷：每周 5，6，7 晚上 20~21 点提供网络游戏服务
	chineseAntiAddictionNormalCfgMap := map[int]map[int]*antiAddictionNormalCfg{
		1: {
			1: &antiAddictionNormalCfg{
				onlineTimeLimit:          nil,
				onlineSecondsLimit:       0,
				weeklyOnlineSecondsLimit: 0,
			},
			2: &antiAddictionNormalCfg{
				onlineTimeLimit:          nil,
				onlineSecondsLimit:       0,
				weeklyOnlineSecondsLimit: 0,
			},
			3: &antiAddictionNormalCfg{
				onlineTimeLimit:          nil,
				onlineSecondsLimit:       0,
				weeklyOnlineSecondsLimit: 0,
			},
			4: &antiAddictionNormalCfg{
				onlineTimeLimit:          nil,
				onlineSecondsLimit:       0,
				weeklyOnlineSecondsLimit: 0,
			},
			5: &antiAddictionNormalCfg{
				onlineTimeLimit: map[int]int{
					20: 21,
				},
				onlineSecondsLimit:       0,
				weeklyOnlineSecondsLimit: 0,
			},
			6: &antiAddictionNormalCfg{
				onlineTimeLimit: map[int]int{
					20: 21,
				},
				onlineSecondsLimit:       0,
				weeklyOnlineSecondsLimit: 0,
			},
			0: &antiAddictionNormalCfg{
				onlineTimeLimit: map[int]int{
					20: 21,
				},
				onlineSecondsLimit:       0,
				weeklyOnlineSecondsLimit: 0,
			},
		},
	}

	// 中国防沉迷：节假日晚上 20~21 点提供网络游戏服务
	chineseAntiAddictionSpecialCfgMap := map[string]*antiAddictionSpecialCfg{
		"2022.04.05": {
			onlineTimeLimit: map[int]int{
				20: 21,
			},
		},
		"2022.05.01": {
			onlineTimeLimit: map[int]int{
				20: 21,
			},
		},
	}

	// 自定义防沉迷：
	// 每周 1，2，3，4，5 中午 13~14 点，晚上 20~21 点提供网络游戏服务，单日限制1小时
	// 每周 6，7 9-21 点提供网络游戏服务，单日限制3小时，周总限制6小时
	testAntiAddictionNormalCfgMap := map[int]map[int]*antiAddictionNormalCfg{
		1: {
			1: &antiAddictionNormalCfg{
				onlineTimeLimit: map[int]int{
					13: 14,
					20: 21,
				},
				onlineSecondsLimit:       3600,
				weeklyOnlineSecondsLimit: 21600,
			},
			2: &antiAddictionNormalCfg{
				onlineTimeLimit: map[int]int{
					13: 14,
					20: 21,
				},
				onlineSecondsLimit:       3600,
				weeklyOnlineSecondsLimit: 21600,
			},
			3: &antiAddictionNormalCfg{
				onlineTimeLimit: map[int]int{
					13: 14,
					20: 21,
				},
				onlineSecondsLimit:       3600,
				weeklyOnlineSecondsLimit: 21600,
			},
			4: &antiAddictionNormalCfg{
				onlineTimeLimit: map[int]int{
					13: 14,
					20: 21,
				},
				onlineSecondsLimit:       3600,
				weeklyOnlineSecondsLimit: 21600,
			},
			5: &antiAddictionNormalCfg{
				onlineTimeLimit: map[int]int{
					13: 14,
					20: 21,
				},
				onlineSecondsLimit:       3600,
				weeklyOnlineSecondsLimit: 21600,
			},
			6: &antiAddictionNormalCfg{
				onlineTimeLimit: map[int]int{
					9: 21,
				},
				onlineSecondsLimit:       10800,
				weeklyOnlineSecondsLimit: 21600,
			},
			0: &antiAddictionNormalCfg{
				onlineTimeLimit: map[int]int{
					9: 21,
				},
				onlineSecondsLimit:       10800,
				weeklyOnlineSecondsLimit: 21600,
			},
		},
	}

	testAntiAddictionSpecialCfgMap := make(map[string]*antiAddictionSpecialCfg)

	type args struct {
		ts                         int64
		antiAddictionNormalCfgMap  map[int]map[int]*antiAddictionNormalCfg
		antiAddictionSpecialCfgMap map[string]*antiAddictionSpecialCfg
	}
	tests := []struct {
		name string
		a    *AntiAddictionData
		args args
		want int64
	}{
		// TODO: Add test cases.
		{
			name: "中国防沉迷测试：非节假日周一晚上首次登录",
			a: &AntiAddictionData{
				AntiAddictionType:   1,
				LastLoginTimestamp:  0,
				DailyOnlineSeconds:  0,
				WeeklyOnlineSeconds: 0,
			},
			args: args{
				ts:                         time.Date(2022, time.March, 14, 20, 0, 0, 0, time.Local).Unix(),
				antiAddictionNormalCfgMap:  chineseAntiAddictionNormalCfgMap,
				antiAddictionSpecialCfgMap: chineseAntiAddictionSpecialCfgMap,
			},
			want: 0,
		},
		{
			name: "中国防沉迷测试：非节假日周五中午首次登录",
			a: &AntiAddictionData{
				AntiAddictionType:   1,
				LastLoginTimestamp:  0,
				DailyOnlineSeconds:  0,
				WeeklyOnlineSeconds: 0,
			},
			args: args{
				ts:                         time.Date(2022, time.March, 18, 12, 0, 0, 0, time.Local).Unix(),
				antiAddictionNormalCfgMap:  chineseAntiAddictionNormalCfgMap,
				antiAddictionSpecialCfgMap: chineseAntiAddictionSpecialCfgMap,
			},
			want: 0,
		},
		{
			name: "中国防沉迷测试：非节假日周五20点首次登录，可玩3600秒",
			a: &AntiAddictionData{
				AntiAddictionType:   1,
				LastLoginTimestamp:  0,
				DailyOnlineSeconds:  0,
				WeeklyOnlineSeconds: 0,
			},
			args: args{
				ts:                         time.Date(2022, time.March, 18, 20, 0, 0, 0, time.Local).Unix(),
				antiAddictionNormalCfgMap:  chineseAntiAddictionNormalCfgMap,
				antiAddictionSpecialCfgMap: chineseAntiAddictionSpecialCfgMap,
			},
			want: 3600,
		},
		{
			name: "中国防沉迷测试：非节假日周五21点30分首次登录，可玩1800秒",
			a: &AntiAddictionData{
				AntiAddictionType:   1,
				LastLoginTimestamp:  0,
				DailyOnlineSeconds:  0,
				WeeklyOnlineSeconds: 0,
			},
			args: args{
				ts:                         time.Date(2022, time.March, 18, 20, 30, 0, 0, time.Local).Unix(),
				antiAddictionNormalCfgMap:  chineseAntiAddictionNormalCfgMap,
				antiAddictionSpecialCfgMap: chineseAntiAddictionSpecialCfgMap,
			},
			want: 1800,
		},
		{
			name: "中国防沉迷测试：节假日21点30分首次登录，可玩1800秒",
			a: &AntiAddictionData{
				AntiAddictionType:   1,
				LastLoginTimestamp:  0,
				DailyOnlineSeconds:  0,
				WeeklyOnlineSeconds: 0,
			},
			args: args{
				ts:                         time.Date(2022, time.April, 5, 20, 30, 0, 0, time.Local).Unix(),
				antiAddictionNormalCfgMap:  chineseAntiAddictionNormalCfgMap,
				antiAddictionSpecialCfgMap: chineseAntiAddictionSpecialCfgMap,
			},
			want: 1800,
		},
		{
			name: "自定义防沉迷测试：周一12点00分首次登录，可玩0秒",
			a: &AntiAddictionData{
				AntiAddictionType:   1,
				LastLoginTimestamp:  0,
				DailyOnlineSeconds:  0,
				WeeklyOnlineSeconds: 0,
			},
			args: args{
				ts:                         time.Date(2022, time.March, 14, 12, 00, 0, 0, time.Local).Unix(),
				antiAddictionNormalCfgMap:  testAntiAddictionNormalCfgMap,
				antiAddictionSpecialCfgMap: testAntiAddictionSpecialCfgMap,
			},
			want: 0,
		},
		{
			name: "自定义防沉迷测试：周一13点30分首次登录，可玩1800秒",
			a: &AntiAddictionData{
				AntiAddictionType:   1,
				LastLoginTimestamp:  0,
				DailyOnlineSeconds:  0,
				WeeklyOnlineSeconds: 0,
			},
			args: args{
				ts:                         time.Date(2022, time.March, 14, 13, 30, 0, 0, time.Local).Unix(),
				antiAddictionNormalCfgMap:  testAntiAddictionNormalCfgMap,
				antiAddictionSpecialCfgMap: testAntiAddictionSpecialCfgMap,
			},
			want: 1800,
		},
		{
			name: "自定义防沉迷测试：周一20点00分非首次登录，已玩1800秒，可玩1800秒",
			a: &AntiAddictionData{
				AntiAddictionType:   1,
				LastLoginTimestamp:  time.Date(2022, time.March, 14, 13, 30, 0, 0, time.Local).Unix(),
				DailyOnlineSeconds:  1800,
				WeeklyOnlineSeconds: 1800,
			},
			args: args{
				ts:                         time.Date(2022, time.March, 14, 20, 0, 0, 0, time.Local).Unix(),
				antiAddictionNormalCfgMap:  testAntiAddictionNormalCfgMap,
				antiAddictionSpecialCfgMap: testAntiAddictionSpecialCfgMap,
			},
			want: 1800,
		},
		{
			name: "自定义防沉迷测试：周二20点00分首次登录，可玩3600秒",
			a: &AntiAddictionData{
				AntiAddictionType:   1,
				LastLoginTimestamp:  time.Date(2022, time.March, 14, 20, 0, 0, 0, time.Local).Unix(),
				DailyOnlineSeconds:  0,
				WeeklyOnlineSeconds: 3600,
			},
			args: args{
				ts:                         time.Date(2022, time.March, 15, 20, 0, 0, 0, time.Local).Unix(),
				antiAddictionNormalCfgMap:  testAntiAddictionNormalCfgMap,
				antiAddictionSpecialCfgMap: testAntiAddictionSpecialCfgMap,
			},
			want: 3600,
		},
		{
			name: "自定义防沉迷测试：周三13点00分首次登录，可玩3600秒",
			a: &AntiAddictionData{
				AntiAddictionType:   1,
				LastLoginTimestamp:  time.Date(2022, time.March, 15, 20, 0, 0, 0, time.Local).Unix(),
				DailyOnlineSeconds:  0,
				WeeklyOnlineSeconds: 7200,
			},
			args: args{
				ts:                         time.Date(2022, time.March, 16, 13, 0, 0, 0, time.Local).Unix(),
				antiAddictionNormalCfgMap:  testAntiAddictionNormalCfgMap,
				antiAddictionSpecialCfgMap: testAntiAddictionSpecialCfgMap,
			},
			want: 3600,
		},
		{
			name: "自定义防沉迷测试：周三20点30分非首次登录，已玩3600秒，可玩0秒",
			a: &AntiAddictionData{
				AntiAddictionType:   1,
				LastLoginTimestamp:  time.Date(2022, time.March, 16, 13, 0, 0, 0, time.Local).Unix(),
				DailyOnlineSeconds:  3600,
				WeeklyOnlineSeconds: 10800,
			},
			args: args{
				ts:                         time.Date(2022, time.March, 16, 20, 30, 0, 0, time.Local).Unix(),
				antiAddictionNormalCfgMap:  testAntiAddictionNormalCfgMap,
				antiAddictionSpecialCfgMap: testAntiAddictionSpecialCfgMap,
			},
			want: 0,
		},
		{
			name: "自定义防沉迷测试：周四20点30分首次登录，可玩1800秒",
			a: &AntiAddictionData{
				AntiAddictionType:   1,
				LastLoginTimestamp:  time.Date(2022, time.March, 16, 20, 30, 0, 0, time.Local).Unix(),
				DailyOnlineSeconds:  0,
				WeeklyOnlineSeconds: 10800,
			},
			args: args{
				ts:                         time.Date(2022, time.March, 17, 20, 30, 0, 0, time.Local).Unix(),
				antiAddictionNormalCfgMap:  testAntiAddictionNormalCfgMap,
				antiAddictionSpecialCfgMap: testAntiAddictionSpecialCfgMap,
			},
			want: 1800,
		},
		{
			name: "自定义防沉迷测试：周五20点00分首次登录，可玩3600秒",
			a: &AntiAddictionData{
				AntiAddictionType:   1,
				LastLoginTimestamp:  time.Date(2022, time.March, 17, 20, 30, 0, 0, time.Local).Unix(),
				DailyOnlineSeconds:  0,
				WeeklyOnlineSeconds: 12600,
			},
			args: args{
				ts:                         time.Date(2022, time.March, 18, 20, 0, 0, 0, time.Local).Unix(),
				antiAddictionNormalCfgMap:  testAntiAddictionNormalCfgMap,
				antiAddictionSpecialCfgMap: testAntiAddictionSpecialCfgMap,
			},
			want: 3600,
		},
		{
			name: "自定义防沉迷测试：周六8点00分首次登录，可玩0秒",
			a: &AntiAddictionData{
				AntiAddictionType:   1,
				LastLoginTimestamp:  time.Date(2022, time.March, 18, 20, 0, 0, 0, time.Local).Unix(),
				DailyOnlineSeconds:  0,
				WeeklyOnlineSeconds: 16200,
			},
			args: args{
				ts:                         time.Date(2022, time.March, 19, 8, 0, 0, 0, time.Local).Unix(),
				antiAddictionNormalCfgMap:  testAntiAddictionNormalCfgMap,
				antiAddictionSpecialCfgMap: testAntiAddictionSpecialCfgMap,
			},
			want: 0,
		},
		{
			name: "自定义防沉迷测试：周六9点00分首次登录，可玩5400秒",
			a: &AntiAddictionData{
				AntiAddictionType:   1,
				LastLoginTimestamp:  time.Date(2022, time.March, 18, 20, 0, 0, 0, time.Local).Unix(),
				DailyOnlineSeconds:  0,
				WeeklyOnlineSeconds: 16200,
			},
			args: args{
				ts:                         time.Date(2022, time.March, 19, 9, 0, 0, 0, time.Local).Unix(),
				antiAddictionNormalCfgMap:  testAntiAddictionNormalCfgMap,
				antiAddictionSpecialCfgMap: testAntiAddictionSpecialCfgMap,
			},
			want: 5400,
		},
		{
			name: "自定义防沉迷测试：周六15点00分非首次登录，已玩3600秒，可玩1800秒",
			a: &AntiAddictionData{
				AntiAddictionType:   1,
				LastLoginTimestamp:  time.Date(2022, time.March, 19, 9, 0, 0, 0, time.Local).Unix(),
				DailyOnlineSeconds:  3600,
				WeeklyOnlineSeconds: 19800,
			},
			args: args{
				ts:                         time.Date(2022, time.March, 19, 15, 0, 0, 0, time.Local).Unix(),
				antiAddictionNormalCfgMap:  testAntiAddictionNormalCfgMap,
				antiAddictionSpecialCfgMap: testAntiAddictionSpecialCfgMap,
			},
			want: 1800,
		},
		{
			name: "自定义防沉迷测试：周六20点00分非首次登录，已玩5400秒，可玩0秒",
			a: &AntiAddictionData{
				AntiAddictionType:   1,
				LastLoginTimestamp:  time.Date(2022, time.March, 19, 15, 0, 0, 0, time.Local).Unix(),
				DailyOnlineSeconds:  5400,
				WeeklyOnlineSeconds: 21600,
			},
			args: args{
				ts:                         time.Date(2022, time.March, 19, 20, 0, 0, 0, time.Local).Unix(),
				antiAddictionNormalCfgMap:  testAntiAddictionNormalCfgMap,
				antiAddictionSpecialCfgMap: testAntiAddictionSpecialCfgMap,
			},
			want: 0,
		},
		{
			name: "自定义防沉迷测试：周日9点00分首次登录，可玩0秒",
			a: &AntiAddictionData{
				AntiAddictionType:   1,
				LastLoginTimestamp:  time.Date(2022, time.March, 20, 12, 0, 0, 0, time.Local).Unix(),
				DailyOnlineSeconds:  0,
				WeeklyOnlineSeconds: 21600,
			},
			args: args{
				ts:                         time.Date(2022, time.March, 20, 0, 0, 0, 0, time.Local).Unix(),
				antiAddictionNormalCfgMap:  testAntiAddictionNormalCfgMap,
				antiAddictionSpecialCfgMap: testAntiAddictionSpecialCfgMap,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.GetReleaseOnlineTime(tt.args.ts, tt.args.antiAddictionNormalCfgMap, tt.args.antiAddictionSpecialCfgMap); got != tt.want {
				t.Errorf("AntiAddictionData.GetReleaseOnlineTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
