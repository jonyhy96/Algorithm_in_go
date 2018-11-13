package quick

func Sort(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	base := data[0]
	left, right := 0, len(data)-1
	for i := 1; i <= right; {
		if data[i] > base {
			data[i], data[right] = data[right], data[i]
			right--
		} else {
			data[i], data[left] = data[left], data[i]
			left++
			i++
		}
	}
	Sort(data[:left])
	Sort(data[left+1:])
	return data
}
