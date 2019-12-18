package main

import (
	"fmt"
	"strconv"
)

type Signal struct {
	digits []int8
}

var (
	patValues = [4]int8{0, 1, 0, -1}
)

func main() {
	part2()
}

func part1() {
	s := makeSignal("59702216318401831752516109671812909117759516365269440231257788008453756734827826476239905226493589006960132456488870290862893703535753691507244120156137802864317330938106688973624124594371608170692569855778498105517439068022388323566624069202753437742801981883473729701426171077277920013824894757938493999640593305172570727136129712787668811072014245905885251704882055908305407719142264325661477825898619802777868961439647723408833957843810111456367464611239017733042717293598871566304020426484700071315257217011872240492395451028872856605576492864646118292500813545747868096046577484535223887886476125746077660705155595199557168004672030769602168262", 1)
	//s := makeSignal("69317163492948606335995924319873")
	for i := 0; i < 100; i++ {
		s.phase2()
	}
	for _, b := range s.digits[0:8] {
		fmt.Print(b)
	}
	fmt.Println("")
}

func part2() {
	s := makeSignal("59702216318401831752516109671812909117759516365269440231257788008453756734827826476239905226493589006960132456488870290862893703535753691507244120156137802864317330938106688973624124594371608170692569855778498105517439068022388323566624069202753437742801981883473729701426171077277920013824894757938493999640593305172570727136129712787668811072014245905885251704882055908305407719142264325661477825898619802777868961439647723408833957843810111456367464611239017733042717293598871566304020426484700071315257217011872240492395451028872856605576492864646118292500813545747868096046577484535223887886476125746077660705155595199557168004672030769602168262", 1)
	//s := makeSignal("69317163492948606335995924319873")
	for i := 0; i < 100; i++ {
		s.phase2()
		fmt.Printf("Phase %d\r", i+1)
	}
	fmt.Println("")
	temp := make([]byte, 0, 7)
	for _, b := range s.digits[0:7] {
		temp = append(temp, byte(b + '0'))
	}
	fmt.Println("Offset ", string(temp))
	offset, _ := strconv.Atoi(string(temp))
	offset = offset % len(s.digits)
	for _, b := range s.digits[offset:offset+8] {
		fmt.Print(b)
	}
	fmt.Println("")
}

func makeSignal(input string, repeat int) *Signal {
	s := new(Signal)
	s.digits = make([]int8, len(input) * repeat)
	for r := 0; r < repeat; r++ {
		for i, c := range input {
			s.digits[i] = int8(c) - '0'
		}
	}
	return s
}

func (s *Signal) phase1() {
	newDigits := make([]int8, len(s.digits))
	for i := range s.digits {
		pattern := makePattern(i+1, len(s.digits))
		newDigits[i] = x(pattern, s.digits)
	}
	s.digits = newDigits
}

func makePattern(position int, len int) []int8 {
	p := make([]int8, len+1)
	s := 0
	repeat := position - 1
	for i := 0; i <= len; i++{
		p[i] = patValues[s]
		if repeat <= 0 {
			s++
			if s == 4 {
				s = 0
			}
			repeat = position - 1
		} else {
			repeat--
		}
	}
	return p[1:]
}

func x(pattern []int8, digits []int8) int8 {
	acc := 0
	for i, b := range digits {
		j := b * pattern[i]
		acc += int(j)
	}
	if acc < 0 {
		acc = -acc
	}
	return int8(acc % 10)
}

func (s *Signal) phase2() {
	newDigits := make([]int8, len(s.digits))
	for i := range s.digits {
		newDigits[i] = x2(i, s.digits)
	}
	s.digits = newDigits
}

func x2(repeatCount int, digits []int8) int8 {
	acc := 0
	for i, b := range digits {
		//j := b * pb(i, repeatCount)
		j := b * patValues[((i+1)/(repeatCount+1))%4]
		acc += int(j)
	}
	if acc < 0 {
		acc = -acc
	}
	return int8(acc % 10)
}

func pb(pos, repeat int) int8 {
	grp := ((pos + 1) / (repeat+1)) % 4
	return patValues[grp]
}
