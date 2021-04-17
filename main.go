package main

import (
	"fmt"
	"io"
	"os"
	"sort"
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
}
func implies(a bool,b bool) bool{
	if a && !b{
		return false
	}
	return true

}
func initOper()  map[byte]int{
	oper := make(map[byte]int)
	oper['('] = 3
	oper[')'] = 3
	oper['!'] = 1
	oper['+'] = 2
	oper['|'] = 1
	oper['^'] = 2
	oper['-'] = 4// <=>
	oper['='] = 4//=>

	return oper
}
func computation(mass [][]byte, val map[byte]bool, res []byte){
	back := make(map[byte]bool)
	result := make(map[byte]bool)

	for
}
func polsky(val [][]byte, oper map[byte]int) [][]byte{

	mass := make([][]byte,len(val))
	for i, val :=range val{
		buf := make([]byte,0)
		stackOper := make([]byte,0)
		for _, v := range val{
			if v >= 'A' && v <= 'Z'{
				buf = append(buf,v)

			}else{
				if oper[v] > 0{
					if v == ')' && len(stackOper) > 1{
						j := len(stackOper) -1
						for stackOper[j] != '(' && len(stackOper) > 0{
							buf = append(buf,stackOper[j])
							stackOper = stackOper[:j]
							j--
						}
					}else {
						if len(stackOper) > 0 && oper[v] > oper[stackOper[len(stackOper)-1]] {
							buf = append(buf, stackOper[len(stackOper)-1])
							stackOper = stackOper[:len(stackOper)-1]
							stackOper = append(stackOper, v)
						} else {
							stackOper = append(stackOper, v)
						}
					}
					//fmt.Printf("%s \n",string(buf))
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

func calculation(data []byte, val map[byte]bool, res []byte){
	mass := make([][]byte,len(val))
	j :=0
	oper := initOper()
	for i, _ := range data{
		if data[i] == '\n' && data[i+1] >= 'A' && data[i+1] <= 'Z'{
			buf := make([]byte,0)
			for data[i + 1] != '\n' && data[i + 1] != '#'{
				if (data[i] >= 'A' && data[i] <= 'Z') || (data[i] == '=' && data[i+1] == '>') || oper[data[i]] > 0  {
					if data[i] == '=' && data[i+1] == '>' && data[i-1] == '<'{
						buf = append(buf, '-')
					}else{
						buf = append(buf, data[i])
					}
				}
				i++
			}
			mass[j] = append(mass[j], buf...)
			j++
		}
	}
	for _,val := range mass{
		fmt.Println("b:", string(val))
	}
	sort.Slice(mass,func(i,j int) bool {
		return len(mass[i]) < len(mass[j])
	})
	for _,val := range mass{
		fmt.Println("SORT:", string(val))
	}

	mass = polsky(mass, oper)
	for _,val := range mass{
		fmt.Println("polsky:", string(val))
	}
	computation(mass,val,res)
}

func parseData(data []byte){
	val := make(map[byte]bool)
	res := make([]byte, 0)
	for i, _ := range data{
		if data[i] >= 'A' && data[i] <= 'Z'{
			val[data[i]] = false
		}
		if data[i] == '=' && len(data) > i+1 && data[i+1] >= 'A' && data[i+1] <= 'Z'{
			i++
			for data[i] != ' '{
				_, ok := val[data[i]]; if ok{
					val[data[i]] = true
				}else{
					fmt.Printf("error parse %s int %d", string(data[i]), i)
					os.Exit(1)
				}
				i++
			}
		}
		if data[i] == '?' && len(data) > i+1{
			i++
			for data[i] != ' ' {
				_, ok := val[data[i]]; if ok{
					res = append(res, data[i])
				}else{
					fmt.Printf("error parse true value")
					os.Exit(1)
				}
				i++
			}
		}

	}
	fmt.Printf(string(res))
	calculation(data,val,res)

}

func main(){
	if len(os.Args) > 1{

		data := make([]byte,0)
		file, err := os.Open(os.Args[1])
		defer file.Close()
		if err != nil{
			fmt.Printf(err.Error())
			os.Exit(1)
		}
		buf := make([]byte, 64)
		for{
			n, err := file.Read(buf)
			if err == io.EOF{   // если конец файла
				break           // выходим из цикла
			}
			//fmt.Print(string(buf[:n]))
			buf = buf[:n]
			data = append(data,buf...)
		}
		parseData(data)
	}else{
		fmt.Printf("input file doesnt exist")
		os.Exit(1)
	}
}
