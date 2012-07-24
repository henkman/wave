package main

import (
	"github.com/papplampe/wave"
	"math"
	"os"
)

const (
	PCM = 1
)

func main() {
	fd, err := os.OpenFile("test.wav", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		println(err.Error())
		return
	}
	defer fd.Close()

	file := wave.CreateFile(PCM, 1, 44100, 8)
	file.AppendSine(file, 440, 5000, 20)
	err = file.Write(fd)
	if err != nil {
		println(err.Error())
	}
}
