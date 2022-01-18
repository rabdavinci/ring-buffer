package main

import (
	"fmt"
	"time"
)

func main() {
	b := NewCircleByteBuffer(1024)
	go func() {
		for i := 0; i < 10; i++ {
			bts := []byte{byte(i)}
			b.Write(bts)
			time.Sleep(time.Second)
		}
		fmt.Println("Close!")
		b.Close()
	}()

	buf := make([]byte, 512)
	for {
		n, err := b.Read(buf)
		if n > 0 {
			fmt.Println("Reads:", buf[0:n])
		}
		if err != nil {
			break
		}
	}
	fmt.Println("end!")
}
