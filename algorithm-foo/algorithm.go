package algorithmfoo

import (
	"fmt"
	"strings"
	"time"
)

func ConvertCamelCase2SnakeCaseWithPhraseTest() {
	fmt.Printf("%v\n", ConvertCamelCase2SnakeCaseWithPhrase("ElfAwakeExample", map[string]struct{}{"elf": {}}))
	fmt.Printf("%v\n", ConvertCamelCase2SnakeCaseWithPhrase("AwakeElfExample", map[string]struct{}{"elf": {}}))
	fmt.Printf("%v\n", ConvertCamelCase2SnakeCaseWithPhrase("AwakeExampleElf", map[string]struct{}{"elf": {}}))
	fmt.Printf("%v\n", ConvertCamelCase2SnakeCaseWithPhrase("ELFAwakeExample", map[string]struct{}{"elf": {}}))
	fmt.Printf("%v\n", ConvertCamelCase2SnakeCaseWithPhrase("AwakeELFExample", map[string]struct{}{"elf": {}}))
	fmt.Printf("%v\n", ConvertCamelCase2SnakeCaseWithPhrase("AwakeExampleELF", map[string]struct{}{"elf": {}}))

}

// Abc 开头 | 中间 DONE
// ABC 开头 | 中间
// ConvertCamelCase2SnakeCaseWithPhrase 将驼峰命名法转换为蛇形命名法：XxxYyyZzz -> xxx_yyy_zzz
func ConvertCamelCase2SnakeCaseWithPhrase(camelCase string, phraseMap map[string]struct{}) string {
	allPhraseSubString := make(map[string]struct{})
	for phrase := range phraseMap {
		for index := 0; index != len(phrase); index++ {
			allPhraseSubString[phrase[0:index]] = struct{}{}
		}
	}

	builder := strings.Builder{}
	phraseBuilder := strings.Builder{}
	isFirstPhrase := true
	for _, r := range camelCase {
		if 'a' <= r && r <= 'z' {
			phraseBuilder.WriteRune(r)
		} else {
			if phraseBuilder.Len() > 0 {
				if _, isPhrase := phraseMap[phraseBuilder.String()]; isPhrase {
					if isFirstPhrase {
						isFirstPhrase = false
					} else {
						builder.WriteRune('_')
					}
					builder.WriteString(phraseBuilder.String())
					phraseBuilder.Reset()
				} else {
					if _, maybePhrase := allPhraseSubString[phraseBuilder.String()]; !maybePhrase {
						if isFirstPhrase {
							isFirstPhrase = false
						} else {
							builder.WriteRune('_')
						}
						builder.WriteString(phraseBuilder.String())
						phraseBuilder.Reset()
					}
				}
			}
			phraseBuilder.WriteRune(r + 32)
		}
	}
	builder.WriteRune('_')
	builder.WriteString(phraseBuilder.String())
	return builder.String()
}

// CalculateYearsOld 根据出生时间戳计算当前年龄
func CalculateYearsOld(birthTimstamp int64) int {
	birthTime := time.Unix(birthTimstamp, 0)
	nowTime := time.Now()
	if nowTime.Month() < birthTime.Month() || (nowTime.Month() == birthTime.Month() && nowTime.Day() < birthTime.Day()) {
		return nowTime.Year() - birthTime.Year() - 1
	}
	return nowTime.Year() - birthTime.Year()
}

func CalculateYearsOldTest() {
	// 1995.6.5 15:00
	birthTimestamp := 802335600
	fmt.Printf("now from 1995.6.5 15:00, %v\n", CalculateYearsOld(int64(birthTimestamp)))

	// 1995.3.5 15:00
	birthTimestamp = 794386800
	fmt.Printf("now from 1995.3.5 15:00, %v\n", CalculateYearsOld(int64(birthTimestamp)))

	// 1995.2.5 15:00
	birthTimestamp = 791967600
	fmt.Printf("now from 1995.2.5 15:00, %v\n", CalculateYearsOld(int64(birthTimestamp)))
}

type AntiAddictionData struct {
	CurrentLoginTimestamp int64
	AntiAddictionType     int64 `json:"anti_addiction_type,omitempty"`
	LastLoginTimestamp    int64 `json:"last_login_timestamp,omitempty"`
	DailyOnlineSeconds    int32 `json:"daily_online_seconds,omitempty"`
	WeeklyOnlineSeconds   int32 `json:"weekly_online_seconds,omitempty"`
}

const AntiAddictionSpecialDateFromat string = "2006.01.02"

type antiAddictionNormalCfg struct {
	onlineTimeLimit          map[int]int
	onlineSecondsLimit       int32
	weeklyOnlineSecondsLimit int32
}

type antiAddictionSpecialCfg struct {
	onlineTimeLimit    map[int]int
	onlineSecondsLimit int32
}

// GetReleaseOnlineTime 获得可在线时长，返回-1则无限制
func (a *AntiAddictionData) GetReleaseOnlineTime(ts int64, antiAddictionNormalCfgMap map[int]map[int]*antiAddictionNormalCfg, antiAddictionSpecialCfgMap map[string]*antiAddictionSpecialCfg) int64 {
	aancwm, ok := antiAddictionNormalCfgMap[int(a.AntiAddictionType)]
	if aancwm == nil || !ok {
		return -1
	}

	lastLoginTime := time.Unix(a.LastLoginTimestamp, 0)

	// nowTime := time.Now()
	nowTime := time.Unix(ts, 0)

	var timeReleaseSeconds int64 = -1
	var durationReleaseSeconds int64 = -1

	// 指定日期限制
	aasc, ok := antiAddictionSpecialCfgMap[nowTime.Format(AntiAddictionSpecialDateFromat)]
	if aasc != nil && ok {
		// 时间段限制
		has, endHour := false, 0
		for bh, eh := range aasc.onlineTimeLimit {
			if bh <= nowTime.Hour() && nowTime.Hour() < eh {
				has = true
				endHour = eh
			}
		}
		if !has {
			return 0
		}
		endTimestamp := nowTime.Unix() - int64(nowTime.Minute()*60) - int64(nowTime.Second()) + int64((int64(endHour)-int64(nowTime.Hour()))*3600)
		timeReleaseSeconds = endTimestamp - nowTime.Unix()

		// 时长限制
		if aasc.onlineSecondsLimit > 0 {
			durationReleaseSeconds = int64(aasc.onlineSecondsLimit)
			if lastLoginTime.Day() == nowTime.Day() {
				if a.DailyOnlineSeconds >= aasc.onlineSecondsLimit {
					return 0
				}
				durationReleaseSeconds -= int64(a.DailyOnlineSeconds)
			}
		}
	} else {
		// 常规日期
		aanc, ok := aancwm[int(nowTime.Weekday())]
		if aanc == nil || !ok {
			return 0
		}

		// 时间段限制
		has, endHour := false, 0
		for bh, eh := range aanc.onlineTimeLimit {
			if bh <= nowTime.Hour() && nowTime.Hour() < eh {
				has = true
				endHour = eh
			}
		}
		if !has {
			return 0
		}
		endTimestamp := nowTime.Unix() - int64(nowTime.Minute()*60) - int64(nowTime.Second()) + int64((int64(endHour)-int64(nowTime.Hour()))*3600)
		timeReleaseSeconds = endTimestamp - nowTime.Unix()

		// 时长限制
		if aanc.onlineSecondsLimit > 0 {
			durationReleaseSeconds = int64(aanc.onlineSecondsLimit)
			if lastLoginTime.Day() == nowTime.Day() {
				if a.DailyOnlineSeconds >= aanc.onlineSecondsLimit {
					return 0
				}
				durationReleaseSeconds -= int64(a.DailyOnlineSeconds)
			}
		}

		// 周时长限制
		if aanc.weeklyOnlineSecondsLimit > 0 {
			ly, lw := lastLoginTime.ISOWeek()
			ny, nw := nowTime.ISOWeek()
			if ly == ny && lw == nw {
				if a.WeeklyOnlineSeconds >= aanc.weeklyOnlineSecondsLimit {
					return 0
				}
				if v := aanc.weeklyOnlineSecondsLimit - a.WeeklyOnlineSeconds; v < int32(durationReleaseSeconds) {
					durationReleaseSeconds = int64(v)
				}
			}
		}
	}

	switch {
	case timeReleaseSeconds < 0 && durationReleaseSeconds < 0:
		return 0
	case timeReleaseSeconds < 0 && durationReleaseSeconds > 0:
		return durationReleaseSeconds
	case timeReleaseSeconds > 0 && durationReleaseSeconds < 0:
		return timeReleaseSeconds
	default:
		if timeReleaseSeconds < durationReleaseSeconds {
			return timeReleaseSeconds
		} else {
			return durationReleaseSeconds
		}
	}
}
