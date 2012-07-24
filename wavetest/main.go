package main

import (
	"math"
	"os"
	"wave"
)

const (
	PCM = 1
)

func AppendSine(file *wave.File, frequency, duration, volume uint32) {
	var i, datalen uint32
	var omega float64

	datalen = duration * file.Samplerate / 1000
	omega = 2 * math.Pi * float64(frequency)
	for i = 0; i < datalen; i++ {
		data := 127 + byte(float64(volume)*math.Sin(float64(i)*omega/float64(file.Samplerate)))
		file.Data = append(file.Data, data)
	}
}

func main() {
	fd, err := os.OpenFile("test.wav", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		println(err.Error())
		return
	}
	defer fd.Close()

	file := wave.CreateFile(PCM, 1, 44100, 8)
	AppendSine(file, 440, 5000, 20)
	err = file.Write(fd)
	if err != nil {
		println(err.Error())
		return
	}
}
