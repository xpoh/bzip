package atack

import (
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
	lc := len(a.chars) - 1
	ln := len(a.passN) - 1

	for i := 0; i < ln+1; i++ {

		if a.passN[i] < lc {
			a.passN[i]++
			return a.buildString(a.passN)
		} else {
			if i < ln {
				a.passN[i] = 0
				a.passN[i+1]++
				return a.buildString(a.passN)
			}
		}
	}
	return ""
}

/*
	buildPass - build strin pass from numbers
*/
func (a *Atack) buildString(i []int) string {
	s := ""
	for j := 0; j < len(i); j++ {
		s = s + string(a.chars[i[j]])
	}
	return s
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
			println("build=", s)
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
