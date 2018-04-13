package jeniusunique

import (
	"fmt"
	"testing"
	"time"
)

func Threaded(cc chan string) {
	for x := 0; x < 100; x++ {
		cc <- GetUniqueGenInstance().NewXReferenceNo(24)
	}
}

func Check(cc chan string, m map[string]interface{}, t *testing.T) {
	var str string
	for {
		select {
		case str = <-cc:
			fmt.Println(str)
			if str == "DONE" {
				return
			}
			if _, ok := m[str]; ok {
				fmt.Printf("Duplicated %s", str)
				t.Fatal("Duplicated")
			} else {
				m[str] = nil
			}
		}
	}
}

func TestUniqueGen_NewXReferenceNo(t *testing.T) {
	mm := make(map[string]interface{})
	var cc chan string
	cc = make(chan string)
	go Threaded(cc)
	go Threaded(cc)
	go Threaded(cc)
	go Threaded(cc)
	go Threaded(cc)
	go Threaded(cc)
	go Threaded(cc)
	go Threaded(cc)
	go Check(cc, mm, t)
	time.Sleep(3 * time.Second)
	cc <- "DONE"
}
