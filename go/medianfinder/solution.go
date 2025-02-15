package medianfinder

type MedianFinder struct {
	Data []int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	obj := MedianFinder{
		Data: make([]int, 0),
	}
	return obj
}

func (this *MedianFinder) AddNum(num int) {

	this.Data = append(this.Data, num)

	var i, val int
	for i, val = range this.Data {
		if val > num {
			break
		}
	}

	if i != len(this.Data)-1 {
		copy(this.Data[i+1:], this.Data[i:len(this.Data)-1])
		this.Data[i] = num
	}
}

func (this *MedianFinder) FindMedian() float64 {

	length := len(this.Data)
	m := length % 2
	if m == 1 {
		return float64(this.Data[(length-1)/2])
	}

	m = length / 2
	return float64((this.Data[m-1] + this.Data[m])) / 2
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
