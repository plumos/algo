package main

import (
	"fmt"
	"strings"
)

//解析反向json
var ss =map[string]interface{}{"1": "bar", "2": "foo.bar", "3": "foo.foo", "4": "baz.cloudmall.com", "5": "baz.cloudmall.ai"}

func main(){
	var result = []map[string]interface{}{}
	for key,item:= range ss{
		result = append(result,techlist(item.(string),key))
	}

	var res  map[string]interface{}
	for _,item:=range result{
		res=allmap(res,item)

	}
	fmt.Println(res)

}

func techlist(str,value string)map[string]interface{}{
	kv := reverse(strings.Split(str, "."))

	var tmp = map[string]interface{}{}
	for key,item:=range kv{
		var newtmp = map[string]interface{}{}
		if key==0{
			tmp[item] = value
		}else{
			newtmp[item] = tmp
			tmp = newtmp
		}
	}
	return tmp
}

func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func allmap(m1,m2 map[string]interface{})map[string]interface{}{
	if m1==nil{
		return m2
	}
	var flag = false
	for k,v :=range m1{
		if m2[k]!=nil{
			v1,flag1:=v.(map[string]interface{})
			v2,flag2:=m2[k].(map[string]interface{})
			if flag1 &&flag2{
				flag =true
				m1[k] = allmap(v1,v2)
			}
		}
	}

	if flag==false{
		for k,v :=range m2{
			m1[k] = v
			return m1
		}
	}

	return m1
}