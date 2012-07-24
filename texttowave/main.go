package main

import (
	"bufio"
	"flag"
	"github.com/papplampe/wave"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	PCM = 1
)

type Tone struct {
	Frequency uint32
	Duration  uint32
}

var (
	_in  string
	_out string
	_vol uint
)

func init() {
	flag.StringVar(&_in, "f", "", "textfile")
	flag.StringVar(&_out, "o", "", "wavefile")
	flag.UintVar(&_vol, "v", 20, "volume")
}

func parseLine(line []byte) *Tone {
	s := strings.Split(string(line), ":")
	if len(s) != 2 {
		return nil
	}

	freq, err := strconv.Atoi(s[0])
	if err != nil {
		return nil
	}

	dur, err := strconv.Atoi(s[1])
	if err != nil {
		return nil
	}

	return &Tone{uint32(freq), uint32(dur)}
}

func readSong(in io.Reader) ([]*Tone, error) {
	song := make([]*Tone, 0)
	bufin := bufio.NewReader(in)

	for {
		line, err := bufin.ReadBytes('\n')
		if len(line) > 3 {
			tone := parseLine(line[:len(line)-1])
			if tone != nil {
				song = append(song, tone)
			}
		}

		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}

	return song, nil
}

func textToWave(textfile, wavefile string, volume uint32) error {
	in, err := os.OpenFile(textfile, os.O_RDONLY, 0600)
	if err != nil {
		return err
	}
	defer in.Close()

	song, err := readSong(in)
	if err != nil {
		return err
	}

	out, err := os.OpenFile(wavefile, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer out.Close()

	file := wave.CreateFile(PCM, 1, 44100, 8)
	for _, tone := range song {
		file.AppendSine(tone.Frequency, tone.Duration, volume)
	}
	return file.Write(out)
}

func main() {
	flag.Parse()

	if _in == "" || _out == "" {
		flag.Usage()
		return
	}

	err := textToWave(_in, _out, uint32(_vol))
	if err != nil {
		println(err.Error())
	}
}
