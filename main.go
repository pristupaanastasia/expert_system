package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type point struct {
}

func or(a bool, b bool) bool {
	return a || b
}

func and(a bool, b bool) bool {
	return a && b
}

func xor(a bool, b bool) bool {
	return (a || b) && !(a && b)
}
func ifOnlyIf(a bool, b bool) bool {
	if (a && b) || (!a && !b) {
		return true
	}
	if (a && !b) || (!a && b) {
		return false
	}
	return false
}

func implies(a bool, b bool) bool {
	if a && !b {
		return false
	}
	return true

}
func initOper() map[byte]int {
	oper := make(map[byte]int)
	oper['('] = 4
	oper[')'] = 4
	oper['!'] = 1
	oper['+'] = 2
	oper['|'] = 2
	oper['^'] = 2
	oper['-'] = 3 // <=>
	oper['='] = 3 //=>
	return oper
}

func recurs(mass []byte, valueByte map[byte]bool) bool {
	st := make([]bool, 0)

	for _, val := range mass {
		if _, ok := valueByte[val]; ok && (val >= 'A') && (val <= 'Z') {
			st = append(st, valueByte[val])

		} else {
			if _, ok = valueByte[val]; !ok && (val >= 'A') && (val <= 'Z') {
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
	fmt.Println(st, " len ", len(st))
	return st[0]
}

func nores(value byte, res []byte) bool {
	for _, val := range res {
		if val == value {
			return false
		}
	}
	return true
}

func brutforce(val map[byte]bool, mass []byte, res []byte, valRes map[byte]bool, buf []byte) (map[byte]bool, []byte) {
	flag := 0
	var newV []byte
	for _, value := range mass {

		if value >= 'A' && value <= 'Z' && nores(value, res) {

			if _, ok := val[value]; !ok {
				val[value] = false
				buf = append(buf, value)
				log.Println(buf)
				flag = 1
			}
			newV = append(newV, value)
		}
	}
	if flag == 1 {
		return val, buf
	}
	i := 0

	j := 1
	for j <= len(newV) {
		i = 0
		for i < len(newV) {

			if _, ok := valRes[newV[i]]; !ok {
				val[newV[i]] = true
				log.Printf(" %s %v \n", string(newV[i]), val[newV[i]])
				if recurs(mass, val) {
					return val, buf
				}
				if i%j == 0 {
					val[newV[i]] = false
					log.Printf(" %s %v \n", string(newV[i]), val[newV[i]])
				}
			}
			i++
		}
		j = j * 2
	}
	return val, nil

}

func binapp(add map[byte]bool, history []byte, addbin int) map[byte]bool {
	i := 0
	buf := addbin
	log.Printf("%b \n", addbin)
	log.Println("bin ", len(history), string(history))
	for i < len(history) {

		log.Printf("add %b \n", addbin)
		buf = buf & 000000001
		log.Printf("buf posle %b \n", buf)
		if buf == 1 {
			add[history[i]] = true
		} else {
			add[history[i]] = false
		}
		i++
		addbin = addbin >> 1
		buf = addbin
	}

	return add
}

func computation(mass [][]byte, val map[byte]bool, res []byte) {

	var valueByte map[byte]bool
	//result := make([]byte, 0)

	var history []byte
	valueByte = make(map[byte]bool)

	for k, v := range val {
		valueByte[k] = v
		if v == false {
			history = append(history, k)
		}
	}

	flag := 0
	addbin := 0
	for {
		valueByte = binapp(valueByte, history, addbin)
		for i := range mass {
			if !recurs(mass[i], valueByte) {
				log.Println(recurs(mass[i], valueByte))
				flag = 1
			}
		}
		if flag == 0 {
			log.Printf("Done")
			break
		}
		flag = 0
		addbin++
	}

	fmt.Printf("RESULT: ")
	for key, value := range valueByte {

		fmt.Printf("key %s val %t \n", string(key), value)
	}
	//fmt.Printf("res %s \n",string(res))
	for _, va := range res {
		if _, ok := valueByte[va]; ok {
			fmt.Printf("RESULT %s %+v \n", string(va), valueByte[va])
		}
	}
}

func polsky(valOper [][]byte, oper map[byte]int) [][]byte {

	mass := make([][]byte, len(valOper))
	//fmt.Printf("val oper %s",valOper)
	for i, initoper := range valOper {

		buf := make([]byte, 0)
		stackOper := make([]byte, 0)
		for _, value := range initoper {
			if value >= 'A' && value <= 'Z' {
				buf = append(buf, value)

			} else {
				if oper[value] > 0 {
					if value == ')' && len(stackOper) > 1 {
						j := len(stackOper) - 1
						for stackOper[j] != '(' && len(stackOper) > 0 {
							buf = append(buf, stackOper[j])

							stackOper = stackOper[:j]
							j--
						}
					} else {
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
					fmt.Printf("%s \n", string(buf))
				}
			}
		}
		for len(stackOper) > 0 {
			buf = append(buf, stackOper[len(stackOper)-1])
			stackOper = stackOper[:len(stackOper)-1]
		}
		mass[i] = append(mass[i], buf...)
	}
	return mass

}

func sortCalc(mass [][]byte, val map[byte]bool) [][]byte {

	newSlice := make([][]byte, 0)
	i := 0
	for i < len(mass) {
		fmt.Printf("mass i %s %d\n", mass[i], i)
		for _, value := range mass[i] {
			if _, ok := val[value]; ok && len(mass) > 2 {
				newSlice = append(newSlice, mass[i])
				buf := mass[i+1:]
				mass = mass[0:i]
				mass = append(mass, buf...)

			}
		}
		i++
	}
	fmt.Printf("mass %s \n", mass)
	fmt.Printf("newsLICW %s \n", newSlice)

	for len(mass) > 0 {
		fmt.Printf("mass %s \n", mass)
		fmt.Printf("newsLICW %s \n", newSlice)
		newSlice = append(newSlice, mass[0])
		mass = mass[1:]

	}
	fmt.Printf("mass %s \n", mass)
	fmt.Printf("newsLICW %s \n", newSlice)

	return newSlice
}

func calcul(data []byte, val map[byte]bool, res []byte) {
	mass := make([][]byte, 0)
	buf := make([]byte, 0)
	j := 0
	oper := initOper()
	for i := range data {
		if data[i] == '\n' && len(data) > i+1 && data[i+1] >= 'A' && data[i+1] <= 'Z' {
			buf = make([]byte, 0)
			i++
			for data[i] != '\n' && data[i] != '#' {
				if (data[i] >= 'A' && data[i] <= 'Z') || (data[i] == '=' && data[i+1] == '>') || oper[data[i]] > 0 {
					if data[i] == '=' && data[i+1] == '>' && data[i-1] == '<' {
						buf = append(buf, '-')
					} else {
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
	for _, value := range mass {
		fmt.Println("SORT:", string(value))
	}

	if len(mass) > 2 {
		mass = sortCalc(mass, val)
	}

	mass = polsky(mass, oper)
	computation(mass, val, res)
}

func parseData(data []byte) {
	val := make(map[byte]bool)
	res := make([]byte, 0)
	i := 0
	for i < len(data) {
		if data[i] == '#' {
			for data[i] != '\n' && i < len(data)-1 {
				i++
			}
		}
		if data[i] == '=' && len(data) > i+1 && data[i+1] >= 'A' && data[i+1] <= 'Z' {
			i++
			for i < len(data) && data[i] != ' ' && data[i] != '\n' {
				val[data[i]] = true
				i++
			}
		} else {
			if len(data) > i+1 && data[i] == '?' {
				i++
				for i < len(data) && data[i] != ' ' && data[i] != '\n' {

					res = append(res, data[i])

					i++
				}
			} else {
				if data[i] >= 'A' && data[i] <= 'Z' {
					val[data[i]] = false
				}
				i++
			}
		}

	}
	fmt.Printf(string(res))
	calcul(data, val, res)
}

func main() {
	if len(os.Args) > 1 {

		data := make([]byte, 0)
		file, err := os.Open(os.Args[1])
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {

			}
		}(file)
		if err != nil {
			fmt.Printf(err.Error())
			os.Exit(1)
		}
		buf := make([]byte, 64)
		for {
			n, err := file.Read(buf)
			if err == io.EOF {
				break
			}
			fmt.Print(string(buf[:n]))
			buf = buf[:n]
			data = append(data, buf...)
		}
		parseData(data)
	} else {
		fmt.Printf("input file doesnt exist")
		os.Exit(1)
	}
}
