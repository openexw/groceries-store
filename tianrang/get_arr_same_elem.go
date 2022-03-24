package tianrang

func arrSameElem(a1, a2 []int) (rst []int) {
	set1 := make(map[int]int)
	set2 := make(map[int]int)
	var target []int

	for _, v := range a1 {
		if _, ok := set1[v]; ok {
			set1[v] += 1
		} else {
			set1[v] = 1
		}

	}
	for _, v := range a2 {
		if _, ok := set2[v]; ok {
			set2[v] += 1
		} else {
			set2[v] = 1
		}
	}

	sameV := make(map[int]struct{})

	for k, v := range set1 {
		if val, has := set2[k]; has {
			sameV[k] = struct{}{}
			if val > v {
				target = a2
			} else {
				target = a1
			}
		}
	}

	// 时间复杂 O(n)
	// 空间复杂度 O(n)
	for _, v := range target {
		if _, ok := sameV[v]; ok {
			rst = append(rst, v)
		}
	}
	return rst
}
