package main

import "fmt"
//排序  左边整数 右边负数，空间复杂度为o(n)

var arr = []int{6, 4, -3, 0, 5, -2, -1, 0, 1, -9,321,-9392,321,-923,0}
func main(){
	fmt.Println(sort(arr))
}

func sort(arr []int)[]int{

	for index,item:=range arr{
		if item>0{
			continue
		}else{
			var flag =false
			for ind2,tt:=range arr[index:]{
				if tt>0{
					arr[index] = tt
					arr[index+ind2] = item
					flag=true
				}
			}
			if flag==false{
				for indnew,itnew:=range arr[index:]{
					if itnew==0{
						continue
					}else{
						var flag =false
						for ind2,tt:=range arr[index+indnew:]{
							if tt==0{
								arr[index+indnew] = tt
								arr[index+indnew+ind2] = itnew
								flag=true
							}
						}
						if flag==false{
							return arr
						}
					}
				}
			}
		}
	}

	return nil

}

