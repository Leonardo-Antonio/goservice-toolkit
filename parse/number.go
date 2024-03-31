package parse

import "strconv"

func StringToInt(val string) int {
	num, err := strconv.Atoi(val)
	if err != nil {
		return 0
	}
	return num
}

func StringToUint(val string) uint {
	num, err := strconv.Atoi(val)
	if err != nil {
		return 0
	}
	return uint(num)
}

func StringToUint8(val string) uint8 {
	num, err := strconv.Atoi(val)
	if err != nil {
		return 0
	}
	return uint8(num)
}

func StringTouUint16(val string) uint16 {
	num, err := strconv.Atoi(val)
	if err != nil {
		return 0
	}
	return uint16(num)
}

func StringTouUint32(val string) uint32 {
	num, err := strconv.Atoi(val)
	if err != nil {
		return 0
	}
	return uint32(num)
}

func StringTouUint64(val string) uint64 {
	num, err := strconv.Atoi(val)
	if err != nil {
		return 0
	}
	return uint64(num)
}
