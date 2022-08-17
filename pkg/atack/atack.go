package atack

import (
	"errors"
	"runtime"
)

type Atacker interface {
	check(pass string) bool
}

type Atack struct {
	atack     Atacker
	pass      string
	maxLength int
	chars     []rune
	passN     []int
}

/*
	NewAtack is a constructor for Atack type
*/
func NewAtack(atack Atacker, pass string, maxLength int, chars []rune) *Atack {
	a := &Atack{atack: atack, pass: pass, maxLength: maxLength, chars: chars}
	a.passN = make([]int, maxLength)

	return a
}

/*
	NewAtack is a constructor for Atack type
*/
func (a *Atack) GenNextPass() string {
	var ans string
	var err error
	ans, err = a.buildString(a.passN)

	if err != nil {
		return "!!!"
	}

	lc := len(a.chars) - 1
	ln := len(a.passN) - 1

	var i int
	for (a.passN[i] == lc) && (i < ln) {
		i++
	}

	if a.passN[i] == 0 {
		for j := i - 1; j >= 0; j-- {
			a.passN[j] = 0
		}
	}
	if a.passN[i] <= lc {
		a.passN[i]++
	}

	return ans
}

/*
	buildPass - build string pass from numbers
*/
func (a *Atack) buildString(i []int) (string, error) {
	s := ""
	for j := 0; j < len(i); j++ {
		if i[j] < len(a.chars) {
			s = s + string(a.chars[i[j]])
		} else {
			return "", errors.New("Overflow")
		}
	}
	return s, nil
}

/*
	brute pass
*/
func (a *Atack) brute() (pass string, err error) {
	N := runtime.NumCPU()

	chOut := make(chan string, N)
	chIn := make(chan string)

	for i := 0; i < N; i++ {
		go a.brute_worker(chOut, chIn)
	}

	for {
		select {
		case p := <-chIn:
			close(chOut)
			return p, nil
		default:
			s := a.GenNextPass()
			//s := a.buildString(passN)
			if s != "!!!" {
				println("build=", s)
			}
			//chOut <- s

			//close(chOut)
			//return "", nil
		}
	}
}

func (a *Atack) brute_worker(ch_in chan string, ch_out chan string) {
	for s := range ch_in {
		//fmt.Println(s)
		if a.atack.check(s) {
			ch_out <- s
		}
	}
}
