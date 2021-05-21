package extended_map



func CalculateAdd(arr []int) {		
	Map(add, arr)
}

func add(a int)int{
	return a +1
}

func Map(f func(int)int, arr []int) {
	for i, v := range arr {
		arr[i] = f(v)
	}	
}