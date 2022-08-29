package atack

import (
	"context"
	"errors"
	"log"
	"runtime"
	"sync"
	"time"
)

type Atacker interface {
	check(pass string) bool
}

type Atack struct {
	atack     Atacker
	maxLength int
	lenght    int
	chars     []rune
	passN     []int
	pass      string
	count     int64
}

/*
NewAtack is a constructor for Atack type
*/
func NewAtack(atack Atacker, maxLength int, chars []rune) *Atack {
	a := &Atack{atack: atack, maxLength: maxLength, chars: chars}
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
	a.pass = s
	a.count++
	return s, nil
}

/*
Brute pass
*/
func (a *Atack) Brute() (pass string, err error) {
	N := runtime.NumCPU()
	log.Printf("Use %v CPU\n", N)

	chOut := make(chan string, N*1e8)
	chIn := make(chan string, 1)

	wg := sync.WaitGroup{}
	ctx := context.Background()
	defer ctx.Done()

	for i := 0; i < N; i++ {
		go a.bruteWorker(&wg, chOut, chIn)
		wg.Add(1)
	}

	go a.startStatusLogger(ctx)

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
	defer wg.Done()

	for s := range chIn {
		if a.atack.check(s) {
			chOut <- s
		}
	}
}

func (a *Atack) startStatusLogger(ctx context.Context) {
	var t int
	totalTime := 0 * time.Second
	for {
		select {
		case <-ctx.Done():
			return
		default:
			time.Sleep(time.Second)
			log.Printf("Current pass %v, Speed %v op/sec, TotalCount %v, time %v", a.pass, a.count, a.count, totalTime)
			a.count = 0
			t++
			totalTime = totalTime + 1*time.Second
		}
	}
}
