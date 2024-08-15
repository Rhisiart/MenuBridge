package protocol

import (
	"encoding"
	"fmt"
	"io"
)

type FrameWriter struct {
	Writer io.Writer
}

func NewFrameWriter(writer io.Writer) *FrameWriter {
	return &FrameWriter{
		Writer: writer,
	}
}

func (fw *FrameWriter) Write(bytes encoding.BinaryMarshaler) error {
	data, err := bytes.MarshalBinary()

	if err != nil {
		return fmt.Errorf("")
	}

	length := len(data)

	for length > 0 {
		bytesConverted, err := fw.Writer.Write(data)

		if err != nil {
			return err
		}

		data = data[bytesConverted:]
		length -= bytesConverted
	}

	return nil
}
