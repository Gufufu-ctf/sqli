package sqli


func GetHalf(min int,max int)int{
	if (min+max)%2==0 {
		return (min+max)/2
	}else{
		return (min+max-1)/2
	}

}


func FindByDichotomy(f func(int)bool)int{			 //f : 目标大于burp
	min := 0
	max := 300
	burp := 0
	for ;; {
		burp = GetHalf(min,max)
		if f(burp)  {
			if max-burp == 1{
				burp = max
				break
			}
			min = burp

		}else{
			if burp == min{
				burp = min
				break
			}
			max = burp
		}
	}
	return burp
}


func judge(burp int)bool{    //f eg
	target := 123
	if target > burp {
		return true
	}else{
		return false
	}
}