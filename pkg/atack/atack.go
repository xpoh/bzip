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
}

/*
	NewAtack is a constructor for Atack type
*/
func NewAtack(atack Atacker, pass string, maxLength int, chars []rune) *Atack {
	return &Atack{atack: atack, pass: pass, maxLength: maxLength, chars: chars}
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
	lc := len(a.chars) - 1
	ln := a.maxLength - 1

	passN := make([]int, a.maxLength)
	var i int
	for {
		select {
		case p := <-chIn:
			close(chOut)
			return p, nil
		default:
			s := a.buildString(passN)
			println("build=", s)
			chOut <- s
			if passN[i] < lc {
				passN[i]++
			} else {
				if i < ln {
					passN[i] = 0
					passN[i+1]++
				} else {
					close(chOut)
					return "", nil
				}
			}
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
