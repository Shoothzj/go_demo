package demo_leetcode

func canPlaceFlowers(flowerbed []int, n int) bool {
	var freeFields []int
	var aux = 0
	for _, val := range flowerbed {
		if aux == 0 && val == 1 {
			continue
		}
		if val == 0 {
			aux = aux + 1
		}
		if aux != 0 && val == 1 {
			freeFields = append(freeFields, aux)
			aux = 0
		}
	}
	if aux != 0 {
		freeFields = append(freeFields, aux)
	}
	if len(freeFields) == 0 {
		if n <= 0 {
			return true
		} else {
			return false
		}
	}
	// 如果两边都是临界值
	if len(freeFields) == 1 {
		if flowerbed[0] == 0 && flowerbed[len(flowerbed)-1] == 0 {
			return (freeFields[0]+1)/2 >= n
		}
		if flowerbed[0] == 0 || flowerbed[len(flowerbed)-1] == 0 {
			return (freeFields[0])/2 >= n
		}
		return (freeFields[0]-1)/2 >= n
	}
	// 主要看左右是否临边
	var temp = 0
	if flowerbed[0] == 0 {
		temp = temp + (freeFields[0])/2
	} else {
		temp = temp + (freeFields[0]-1)/2
	}

	if flowerbed[len(flowerbed)-1] == 0 {
		temp = temp + (freeFields[len(freeFields)-1])/2
	} else {
		temp = temp + (freeFields[len(freeFields)-1]-1)/2
	}
	for i := 1; i < len(freeFields)-1; i++ {
		temp = temp + (freeFields[i]-1)/2
	}
	return temp >= n
}
