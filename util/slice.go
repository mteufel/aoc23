package util

func RemoveValueFromIntSlice(s []int64, value int64) []int64 {
	var new []int64
	for _, v := range s {
		if v != value {
			new = append(new, v)
		}
	}
	return new
}
