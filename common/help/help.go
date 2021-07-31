package help

import "strconv"

func StringToInt32(a string) int32 {
	d, _ := strconv.ParseInt(a, 10, 32)
	return int32(d)
}

func IntToString(a int) string {
	str := strconv.Itoa(a)
	return str
}

func Int64ToString(a int64) string {
	str := strconv.FormatInt(a,10)
	return str
}