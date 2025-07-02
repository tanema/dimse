package encoding

import (
	"bytes"
	"encoding/binary"
)

type (
	Writer struct {
		out *bytes.Buffer
		bo  binary.ByteOrder
	}
)

func NewWriter(bo binary.ByteOrder) *Writer {
	return &Writer{out: bytes.NewBuffer(nil), bo: bo}
}

func (w *Writer) Bytes() []byte {
	return w.out.Bytes()
}

func (w *Writer) Len() int {
	return w.out.Len()
}

func (w *Writer) Write(data []byte) (int, error) {
	err := binary.Write(w.out, w.bo, data)
	return len(data), err
}

func (w *Writer) WriteParts(parts ...any) error {
	for _, data := range parts {
		if skip, isSkip := data.(Skip); isSkip {
			zeros := make([]byte, int(skip))
			if err := binary.Write(w.out, w.bo, zeros); err != nil {
				return err
			}
		} else if err := binary.Write(w.out, w.bo, data); err != nil {
			return err
		}
	}
	return nil
}
