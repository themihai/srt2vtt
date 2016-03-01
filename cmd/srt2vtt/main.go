package main

import (
	"fmt"
	"github.com/ricksancho/srt2vtt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Open(os.Args[1])
	check(err)
	r, _ := srt2vtt.NewReader(f)

	for {
		b := make([]byte, 32768)
		n, _ := r.Read(b)
		if n == 0 {
			break
		}
		fmt.Printf("%s", string(b[:n]))
		//fmt.Printf("Len: %d Str: %q", n, string(b))
	}
}
