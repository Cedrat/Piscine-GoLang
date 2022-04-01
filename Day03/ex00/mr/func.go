package mr

import "strconv"

func SumInt(value1, value2 Any) Any {
	var ret1, ret2 int
	switch v1 := value1.(type) {
	case int:
		ret1 = int(v1)
	case string:
		ret1, _ = strconv.Atoi(v1)
	}
	switch v2 := value2.(type) {
	case int:
		ret2 = int(v2)
	case string:
		ret2, _ = strconv.Atoi(v2)
	}
	return ret1 + ret2
}
func AddOneInt(value Any) (ret Any) {
	switch v := value.(type) {
	case int:
		ret = v + 1
	case string:
		vi, _ := strconv.Atoi(v)
		ret = strconv.Itoa(vi + 1)
	}
	return
}
func AddOneString(value Any) (ret Any) {
	switch v := value.(type) {
	case int:
		ret, _ = strconv.Atoi(strconv.Itoa(v) + "1")
	case string:
		ret = v + "1"
	}
	return
}
