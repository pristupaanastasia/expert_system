package main

import (
	"bufio"
	"fmt"
	"strings"

	//"io"
	"log"
	"os"
)

/*
commits := map[string]int{
    "rsc": 3711,
    "r":   2138,
    "gri": 1908,
    "adg": 912,
}
*/
var (
	allowedSymbols string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ=><+|!^?()"
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
func ifOnlyIf(a bool, b bool, res bool) (bool, bool) {
	if a && b {
		return true, b
	}
	if !a && !b {
		return true, b
	}
	if res {
		if b == false {
			b = true
		} else {
			b = false
		}
		if a && b {
			return true, b
		}
		if !a && !b {
			return true, b
		}
	}

	return false, false
}

func implies(a bool, b bool, res bool) (bool, bool) {
	if a && !b && !res {
		return false, false
	}
	if res {
		if b == false {
			b = true
		} else {
			b = false
		}
		if a && !b && !res {
			return false, b
		}
	}
	return true, b

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

func recurs(mass []byte, valueByte map[byte]int) bool {
	st := make([]bool, 0)

	res := false
	log.Println(string(mass))
	for i, val := range mass {
		if (val >= 'A') && (val <= 'Z') {
			if _, ok := valueByte[val]; !ok {
				return false
			}
			if valueByte[val] == 1 {
				st = append(st, true)
			}
			if valueByte[val] == -1 {
				st = append(st, false)
			}
			if valueByte[val] == 0 {
				if i < len(mass) && (mass[i+1] == '=' || mass[i+1] == '-') { //а вдруг тут =>a&b  ab&=
					res = true
					st = append(st, false)
				} else {
					res = true
					st = append(st, false)
				}
			}

		} else {

			if val == '!' && len(st) > 0 {
				st[len(st)-1] = !st[len(st)-1]
			} else {
				var b bool
				if len(st) > 1 {
					var buf bool
					if val == '-' {
						buf, b = ifOnlyIf(st[len(st)-2], st[len(st)-1], res)
						if buf && res {
							if b {
								valueByte[mass[i-1]] = 1
							} else {
								valueByte[mass[i-1]] = -1
							}
						}
					}
					if val == '=' {
						buf, b = implies(st[len(st)-2], st[len(st)-1], res)
						if buf && res {
							if b {
								valueByte[mass[i-1]] = 1
							} else {
								valueByte[mass[i-1]] = -1
							}
						}
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
					log.Println(st)
				}
			}

		}
	}
	fmt.Println(st)
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

func binapp(add map[byte]bool, history []byte, addbin int) map[byte]bool {

	buf := addbin

	i := len(history) - 1
	for i > -1 {

		log.Printf("add %b \n", addbin)
		buf = buf & 000000000000001
		log.Printf("buf posle %b %s \n", buf, string(history[i]))
		if buf == 1 {
			add[history[i]] = true
		} else {
			add[history[i]] = false
		}
		i--
		addbin = addbin >> 1
		buf = addbin
	}

	return add
}

func cancel(val map[byte]bool, history []byte) bool {

	for _, v := range history {
		if val[v] == false {
			return false
		}
	}
	return true
}

func computation(mass [][]byte, val map[byte]int, res []byte) {

	flag := 0
	for {
		for i := range mass {
			if !recurs(mass[i], val) {
				flag = 1
			}
		}
		for key, v := range val {
			log.Println(string(key), v)
		}
		if flag == 0 {
			break
		}
		flag = 0
	}
	//var valueByte map[byte]bool
	////result := make([]byte, 0)
	//
	//var history []byte
	//valueByte = make(map[byte]bool)
	//
	//for k, v := range val {
	//	valueByte[k] = v
	//	if v == false && nores(k, res) {
	//		history = append(history, k)
	//	}
	//}
	//for _ ,v := range res{
	//	if valueByte[v] == false{
	//	history = append(history, v)
	//	}
	//}
	//
	//flag := 0
	//addbin := 0
	//log.Println(string(history))
	//minbaf := make(map[byte]bool)
	//min := -1
	//for {
	//	for {
	//		flag = 0
	//		valueByte = binapp(valueByte, history, addbin)
	//		for i := range mass {
	//			if recurs(mass[i], valueByte)  == false{
	//				log.Println(recurs(mass[i], valueByte),string(mass[i]) ,"lox")
	//				flag = 1
	//			}
	//		}
	//		if flag == 0 || cancel(valueByte, history) {
	//			log.Printf("Done")
	//			break
	//		}
	//
	//		addbin++
	//	}
	//	buf := 0
	//	for _, v := range valueByte {
	//		if v == true {
	//			buf++
	//		}
	//	}
	//	if (buf < min || min == -1) && flag == 0 {
	//		min = buf
	//		for k, v := range valueByte {
	//			minbaf[k] = v
	//		}
	//	}
	//	if cancel(valueByte, history) { //алгоритм окончания другой (когда все тру)
	//		break
	//	}
	//	addbin++
	//}
	//
	fmt.Printf("RESULT: ")
	for key, value := range val {

		fmt.Printf("key %s val %t \n", string(key), value)
	}
	//fmt.Printf("res %s \n",string(res))
	for _, va := range res {
		if _, ok := val[va]; ok {
			fmt.Printf("RESULT %s %+v \n", string(va), val[va])
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

func sortCalc(mass [][]byte, val map[byte]int) [][]byte {

	newSlice := make([][]byte, 0)
	i := 0
	for i < len(mass) {
		//		fmt.Printf("mass i %s %d\n", mass[i], i)
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

	for len(mass) > 0 {

		newSlice = append(newSlice, mass[0])
		mass = mass[1:]

	}

	return newSlice
}

func getOpsLines(lines []string) []string {
	i := 0
	res := make([]string, 0)
	for i < len(lines) {
		if lines[i][0] != '?' && lines[i][0] != '=' {
			res = append(res, lines[i])
		}
		i++
	}
	return res
}

func calculv2(lines []string, val map[byte]int, res []byte) {
	mass := make([][]byte, 0)
	//	buf := make([]byte, 0)
	//j := 0
	i := 0
	oper := initOper() // TODO: it's a global var
	ops := getOpsLines(lines)
	for i < len(ops) {
		k := 0
		buf := make([]byte, 0)
		for k < len(ops[i]) {
			if ops[i][k] == '<' && k < len(ops[i])-2 && ops[i][k]+1 == '=' && ops[i][k+2] == '>' {
				buf = append(buf, '-')
				k += 2

			} else {
				if ops[i][k] == '=' && k < len(ops[i])-2 && ops[i][k]+1 == '>' {
					buf = append(buf, '=')
					k += 1
				} else {
					buf = append(buf, ops[i][k])
				}
			}
			k++
		}
		mass = append(mass, buf)
		i++
	}

	if len(mass) > 2 {
		mass = sortCalc(mass, val)
	}

	mass = polsky(mass, oper)
	computation(mass, val, res)
}

func calcul(data []byte, val map[byte]int, res []byte) {
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

func parseData(lines []string, val map[byte]int, res []byte) {
	calculv2(lines, val, res)
}

func removeComment(line string) string {
	pepega := strings.Split(line, "#")
	return pepega[0]
}

func isComment(line string) bool {
	i := 0
	for i < len(line) {
		if line[i] == ' ' {
			i++
		} else {
			break
		}
	}
	return line[i] == '#'
}

func isEmpty(line string) bool {
	i := 0
	for i < len(line) {
		if line[i] != ' ' {
			return false
		}
		i++
	}
	return true
}

func removeInvalid(line string) string {
	i := 0
	for i < len(line) {
		if strings.Contains(allowedSymbols, string(line[i])) {
			i++
		} else {
			line = strings.ReplaceAll(line, string(line[i]), "")
			i = 0
		}
	}
	return line
}

func parserv2(file *os.File) ([]string, error) {
	var result = make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if err := scanner.Err(); err != nil {
			return nil, err
		}
		lineLen := len(text)
		if lineLen < 1 || isEmpty(text) || isComment(text) {
			continue
		}
		if strings.Contains(text, "#") {
			text = removeComment(text)
		}
		text = removeInvalid(text)
		if len(text) > 0 {
			result = append(result, text)
		}
	}
	return result, nil
}

func findFacts(lines []string) map[byte]int {
	val := make(map[byte]int)
	i := 0
	for i < len(lines) {
		k := 0
		if lines[i][k] != '?' && lines[i][k] != '=' {
			for k < len(lines[i]) {
				if lines[i][k] >= 'A' && lines[i][k] <= 'Z' {
					val[lines[i][k]] = 0
				}
				k++
			}
		}
		i++
	}
	return val
}

func updateFacts(lines []string, val map[byte]int) {
	i := 0
	for i < len(lines) {
		if lines[i][0] == '=' && len(lines[i]) > 1 {
			k := 0
			for k < len(lines[i]) {
				if lines[i][k] >= 'A' && lines[i][k] <= 'Z' {
					val[lines[i][k]] = 1
				}
				k++
			}
		}
		i++
	}
}

func getUnknown(lines []string, val map[byte]int) []byte {
	res := make([]byte, 0)
	i := 0
	for i < len(lines) {
		k := 0
		if lines[i][k] == '?' {
			for k < len(lines[i]) {
				if lines[i][k] >= 'A' && lines[i][k] <= 'Z' {
					res = append(res, lines[i][k])
					val[lines[i][k]] = 0
				}
				k++
			}
		}
		i++
	}
	return res
}

func parseDatav2(lines []string) (map[byte]int, []byte) {

	val := findFacts(lines)
	if len(val) < 1 {
		fmt.Printf("Failed to find Facts")
		os.Exit(0)
	}
	updateFacts(lines, val)
	res := getUnknown(lines, val)
	if len(res) < 1 {
		fmt.Printf("Failed to find Unknown Values")
		os.Exit(0)
	}
	return val, res
}

func validateData(lines []string) int {
	i := 0
	res := 0
	if len(lines) < 3 {
		return 0
	}
	for i < len(lines) {
		if lines[i][0] == '=' {
			res++
		}
		if lines[i][0] == '?' {
			res++
		}
		i++
	}
	return res
}

func main() {

	if len(os.Args) > 1 {
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
		if data, err := parserv2(file); err != nil {
			fmt.Println(err.Error())
			return
		} else {
			if validateData(data) != 2 {
				fmt.Printf("ivalid input")
				return
			}
			val, res := parseDatav2(data)
			parseData(data, val, res)
		}

	} else {
		fmt.Printf("input file doesnt exist")
		os.Exit(1)
	}
}
