package atack

import (
	"errors"
	"runtime"
	"sync"
)

type Atacker interface {
	check(pass string) bool
}

type Atack struct {
	atack     Atacker
	pass      string
	maxLength int
	lenght    int
	chars     []rune
	passN     []int
}

/*
	NewAtack is a constructor for Atack type
*/
func NewAtack(atack Atacker, pass string, maxLength int, chars []rune) *Atack {
	a := &Atack{atack: atack, pass: pass, maxLength: maxLength, chars: chars}
	a.passN = make([]int, maxLength)
	a.lenght = 1

	return a
}

/*
	NewAtack is a constructor for Atack type
*/
func (a *Atack) GenNextPass(idx int) error {

	lc := len(a.chars) - 1
	ln := a.lenght - 1

	a.passN[idx]++

	if a.passN[idx] > lc {
		if idx == ln {
			if a.lenght < a.maxLength {
				a.lenght++
				a.passN[idx] = 0
				a.passN[idx+1] = 0
			} else {
				return errors.New("Overflow")
			}
		} else {
			a.passN[idx] = 0
			err := a.GenNextPass(idx + 1)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

/*
	buildPass - build string pass from numbers
*/
func (a *Atack) buildString(i []int) (string, error) {
	s := ""
	for j := 0; j < a.lenght; j++ {
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

	wg := sync.WaitGroup{}

	for i := 0; i < N; i++ {
		go a.bruteWorker(&wg, chOut, chIn)
		wg.Add(1)
	}

	var p string
	for {
		select {
		case p = <-chIn:
			close(chOut)
			return p, nil
		default:
			err := a.GenNextPass(0)
			if err != nil {
				close(chOut)
				wg.Wait()
				close(chIn)
				for p = range chIn {
				}
				return p, nil
			} else {
				s, _ := a.buildString(a.passN)
				chOut <- s
			}
		}
	}
}

func (a *Atack) bruteWorker(wg *sync.WaitGroup, chIn chan string, chOut chan string) {
	for s := range chIn {
		if a.atack.check(s) {
			chOut <- s
		}
	}
	wg.Done()
}
