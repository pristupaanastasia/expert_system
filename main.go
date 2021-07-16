package main

import (
	"fmt"
	"io"
	"os"

)



func or(a bool, b bool) bool{
	return a || b
}

func and(a bool, b bool) bool{
	return a && b
}

func xor(a bool,b bool) bool{
	return (a || b) && !(a && b)
}
func ifOnlyIf(a bool,b bool) bool{
	if (a && b) || (!a && !b){
		return true
	}
	if (a && !b) || (!a && b){
		return false
	}
	return false
}

func implies(a bool,b bool) bool{
	if a && !b{
		return false
	}
	return true

}
func initOper()  map[byte]int {
	oper := make(map[byte]int)
	oper['('] = 4
	oper[')'] = 4
	oper['!'] = 1
	oper['+'] = 2
	oper['|'] = 2
	oper['^'] = 2
	oper['-'] = 3// <=>
	oper['='] = 3//=>
	return oper
}

func recurs(mass []byte, valueByte map[byte]bool) bool {
	st := make([]bool, 0)
	
	for _, val := range mass {
		if _,ok := valueByte[val]; ok && (val >= 'A') && (val <= 'Z') {
			st = append(st, valueByte[val])
			
		} else {
			if _,ok = valueByte[val]; !ok && (val >= 'A') && (val <= 'Z'){
				return false
			}
			if val == '!' && len(st) > 0 {

				st[len(st)-1] = !st[len(st)-1]
		
			} else {
				if len(st) > 1 {
					var buf bool
					if val == '-' {
						buf = ifOnlyIf(st[len(st)-2], st[len(st)-1])
					}
					if val == '=' {
						buf = implies(st[len(st)-2], st[len(st)-1])
					}
					if val == '+' {
						buf = and(st[len(st)-2], st[len(st)-1])
					}
					if val == '^' {
						buf = xor(st[len(st)-2], st[len(st)-1])
					}
					if val == '|' {
						buf = or(st[len(st)-2], st[len(st)-1])
					}
				
					st = st[:len(st)-2]
					st = append(st, buf)
				}
			}

		}
	}
	fmt.Println(st," len ", len(st))
	return st[0]
}

func nores(value byte, res []byte) bool{
	for _,val := range res{
		if val == value{
			return false
		}
	}
	return true
}

func brutforce(val map[byte]bool, mass []byte, res []byte) (map[byte]bool, []byte){
	var buf []byte
	for _, value :=range mass{
		if value >= 'A' && value <= 'Z' && nores(value,res){
			if _, ok := val[value]; !ok {
				val[value] = false
				buf = append(buf,value)

			}
		}
	}
	if len(buf)>0{
		return val,buf
	}
	for _, value :=range mass{
		if value >= 'A' && value <= 'Z' && nores(value,res){
			if _, ok := val[value]; ok{
				val[value] = true
				buf = append(buf,value)

			}
		}
	}
	if len(buf)>0{
		return val,buf
	}
	return nil, nil
}

func computation(mass [][]byte, val map[byte]bool, res []byte) {
	var buf []byte
	var newval map[byte]bool
	var valueByte map[byte]bool
	result := make([]byte, 0)
	for k, _ := range valueByte{
		fmt.Printf("valubyte %s \n",string(k))
		result = append(result,k)
	}
	i:=0

	strb := make(map[byte]int)

	valueByte = val
	metka = make(map[byte]int)
	for i < len(mass) {
		prev := newval
		buf = nil
		if len(mass[i]) >0 {
			newval = valueByte
			
			for !recurs(mass[i], valueByte) {
				fmt.Printf("RES %s \n",result)
				fmt.Printf("val %s \n",string(mass[i]))
				newval, buf = brutforce(newval,mass[i],result)
				for key, value := range newval{
					if _, ok:=strb[key]; !ok{
						strb[key] = i
					}
					fmt.Printf("key %s val %t \n", string(key), value)
					// if strb[key] <= i && metka[key]==0{
						
					// }
					valueByte[key] = value
				}
				fmt.Println("buf",buf)
				// if buf == nil{
				// 	break
				// }
				
			}
			
			if buf == nil{
				// i--
				// fmt.Println("new val nil")
				// if prev != nil{
				// 	fmt.Println("prev nen nil")
				// 	flag :=0
				// 	for k, _ := range prev{
				// 		_, exist:=val[k]
				// 		if _,ok :=valueByte[k];ok  && !exist && valueByte[k] {
				// 			fmt.Println(valueByte[k],"k",k,"k")
				// 			valueByte[k] = false
							
				// 			i = strb[k]
				// 			for _, l:= range mass[i]{
				// 				if _,ok:=strb[l];ok && strb[l] < i{
				// 					i = strb[l]
				// 					if !recurs(mass[i], valueByte){
				// 						valueByte[k] = true
				// 					}
				// 				}
				// 			}
				// 			fmt.Println(valueByte[k],"false")
				// 			flag = 1

				// 		}
				// 	}
				// 	if flag != 1{
				// 		for k, _ := range prev{
				// 			_, exist:=val[k]
				// 			if _,ok :=valueByte[k];ok  && !exist && !valueByte[k] {
				// 				fmt.Println(valueByte[k],"k",k,"k")
				// 				valueByte[k] = true
				// 				i = strb[k]
				// 				for _, l:= range mass[i]{
				// 					if _,ok:=strb[l];ok && strb[l] < i{
				// 						i = strb[l]
				// 					}
				// 				}
				// 				fmt.Println("true",valueByte[k])
				// 				flag = -1

				// 			}
				// 		}
				// 	}
					
				//}
				if i < 0{
					fmt.Printf("Sorry i didn't find solution")
					return
				}
			}
			if len(buf)>0 {
				i++
			}
		}
	}
	fmt.Printf("RESULT: ")
	for key, value := range valueByte{

		fmt.Printf("key %s val %t \n", string(key), value)
	}
	//fmt.Printf("res %s \n",string(res))
	for _,val := range res{
		if _, ok := valueByte[val]; ok{
			fmt.Printf("RESULT %s %+v \n",string(val),valueByte[val])
		}
	}
}

func polsky(valOper [][]byte, oper map[byte]int) [][]byte{

	mass := make([][]byte,len(valOper))
	//fmt.Printf("val oper %s",valOper)
	for i, initoper :=range valOper{
	
		buf := make([]byte,0)
		stackOper := make([]byte,0)
		for _, value := range initoper{
			if value >= 'A' && value <= 'Z'{
				buf = append(buf,value)
		
			}else{
				if oper[value] > 0{
					if value == ')' && len(stackOper) > 1{
						j := len(stackOper) -1
						for stackOper[j] != '(' && len(stackOper) > 0{
							buf = append(buf,stackOper[j])
					
							stackOper = stackOper[:j]
							j--
						}
					}else {
						if len(stackOper) > 0 && oper[value] >= oper[stackOper[len(stackOper)-1]] {
							for len(stackOper) > 0 && oper[value] >= oper[stackOper[len(stackOper)-1]] {
					
								buf = append(buf, stackOper[len(stackOper)-1])
								stackOper = stackOper[:len(stackOper)-1]
							}
								stackOper = append(stackOper, value)

						} else {
					
							stackOper = append(stackOper, value)
						}
					}
					fmt.Printf("%s \n",string(buf))
				}
			}
		}
		for len(stackOper) > 0{
			buf = append(buf,stackOper[len(stackOper) -1])
			stackOper = stackOper[:len(stackOper) -1]
		}
		mass[i] = append(mass[i],buf...)
	}
	return mass

}

func sortCalc(mass[][]byte, val map[byte]bool) [][]byte{

	newSlice := make([][]byte,0)
	i:=0
	for i < len(mass) {
		fmt.Printf("mass i %s %d\n",mass[i],i)
		for _,value := range mass[i]{
			if _,ok := val[value]; ok && len(mass) > 2  {
				newSlice = append(newSlice, mass[i])
				buf := mass[i+1:]
				mass = mass[0:i]
				mass = append(mass,buf...)

			}
		}
		i++
	}
	fmt.Printf("mass %s \n",mass)
	fmt.Printf("newsLICW %s \n",newSlice)

	for len(mass) >0{
		fmt.Printf("mass %s \n",mass)
		fmt.Printf("newsLICW %s \n",newSlice)
		newSlice = append(newSlice, mass[0])
		mass = mass[1:]

	}
	fmt.Printf("mass %s \n",mass)
	fmt.Printf("newsLICW %s \n",newSlice)

	return newSlice
}


func calcul(data []byte, val map[byte]bool, res []byte){
	mass := make([][]byte,0)
	buf := make([]byte,0)
	j :=0
	oper := initOper()
	for i := range data{
		if data[i] == '\n' && len(data) > i + 1 && data[i+1] >= 'A' && data[i+1] <= 'Z'{
			buf = make([]byte,0)
			i++
			for data[i] != '\n' && data[i] != '#'{
				if (data[i] >= 'A' && data[i] <= 'Z') || (data[i] == '=' && data[i+1] == '>') || oper[data[i]] > 0  {
					if data[i] == '=' && data[i+1] == '>' && data[i-1] == '<'{
						buf = append(buf, '-')
					}else{
						buf = append(buf, data[i])
					}
				}
				i++
			}
			mass = append(mass, []byte{})
			mass[j] = append(mass[j], buf...)
			j++
		}
	}
	for _,value := range mass{
		fmt.Println("SORT:", string(value))
	}

	if len(mass)>2 {
		mass = sortCalc(mass, val)
	}


	mass = polsky(mass, oper)
	computation(mass,val,res)
}

func parseData(data []byte){
	val := make(map[byte]bool)
	res := make([]byte, 0)
	for i  := range data{
		if data[i] >= 'A' && data[i] <= 'Z'{
			//val[data[i]] = false
		}
		if data[i] == '=' && len(data) > i+1 && data[i+1] >= 'A' && data[i+1] <= 'Z'{
			i++
			for i < len(data) && data[i] != ' ' && data[i] != '\n'{
				//_, ok := val[data[i]]; if ok{
					val[data[i]] = true
				//}else{
				//	fmt.Printf("error parse %s int %d", string(data[i]), i)
				//	os.Exit(1)
				//}
				i++
			}
		}
		if len(data) > i+1 && data[i] == '?' {
			i++
			for  i < len(data) && data[i] != ' ' &&  data[i] != '\n'{
				//_, ok := val[data[i]]; if ok{
					res = append(res, data[i])
				//}else{
				//	fmt.Printf("error parse true value")
				//	os.Exit(1)
				//}
				i++
			}
		}

	}
	fmt.Printf(string(res))
	calcul(data,val,res)
}

func main(){
	if len(os.Args) > 1{

		data := make([]byte,0)
		file, err := os.Open(os.Args[1])
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {

			}
		}(file)
		if err != nil{
			fmt.Printf(err.Error())
			os.Exit(1)
		}
		buf := make([]byte, 64)
		for{
			n, err := file.Read(buf)
			if err == io.EOF{
				break
			}
			fmt.Print(string(buf[:n]))
			buf = buf[:n]
			data = append(data,buf...)
		}
		parseData(data)
	}else{
		fmt.Printf("input file doesnt exist")
		os.Exit(1)
	}
}
