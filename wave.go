package wave

import (
	"encoding/binary"
	"io"
	"errors"
)

const (
	HEADER_SIZE = 12
	FORMAT_SIZE = 24
	DATA_SIZE   = 8
	FILE_SIZE   = HEADER_SIZE + FORMAT_SIZE + DATA_SIZE
)

type File struct {
	// Chunkid [4]byte
	Filesize uint32
	// Rifftype [4]byte

	// Chunkid [4]byte
	// Fmtlen uint32
	Format        uint16
	Channels      uint16
	Samplerate    uint32
	BytesPerSec   uint32
	Blockalign    uint16
	BitsPerSample uint16

	// Chunkid [4]byte
	// Length uint32
	Data []byte
}

func ReadFile(in io.Reader) (*File, error) {
	return nil, errors.New("not implemented")
}

/*
	Intializes a new wave file
*/
func CreateFile(format uint16, channels uint16, samplerate uint32, bitspersample uint16) *File {
	var blockalign uint16 = channels * bitspersample / 8
	return &File{
		0,
		format,
		channels,
		samplerate,
		samplerate * uint32(blockalign),
		blockalign,
		bitspersample,
		make([]byte, 0),
	}
}

func (this *File) writeMetaData(out io.Writer) error {
	var datalen uint32 = uint32(len(this.Data))
	var filesize uint32 = FILE_SIZE + datalen - 8
	var fmtlen uint32 = 16

	out.Write([]byte("RIFF"))
	binary.Write(out, binary.LittleEndian, filesize)
	out.Write([]byte("WAVE"))

	out.Write([]byte("fmt "))
	binary.Write(out, binary.LittleEndian, fmtlen)
	binary.Write(out, binary.LittleEndian, this.Format)
	binary.Write(out, binary.LittleEndian, this.Channels)
	binary.Write(out, binary.LittleEndian, this.Samplerate)
	binary.Write(out, binary.LittleEndian, this.BytesPerSec)
	binary.Write(out, binary.LittleEndian, this.Blockalign)
	binary.Write(out, binary.LittleEndian, this.BitsPerSample)

	out.Write([]byte("data"))
	binary.Write(out, binary.LittleEndian, datalen)

	return nil
}

/*
	Writes the file to out
*/
func (this *File) Write(out io.Writer) error {
	this.writeMetaData(out)
	out.Write(this.Data)

	return nil
}
