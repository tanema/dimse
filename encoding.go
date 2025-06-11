package dimse

import (
	"encoding/binary"
	"io"
)

type (
	pduDecoder struct{ in []io.Reader }
	pduEncoder struct{ out io.Writer }
	encSkip    int64
)

func newDecoder(in io.Reader) *pduDecoder  { return &pduDecoder{in: []io.Reader{in}} }
func newEncoder(out io.Writer) *pduEncoder { return &pduEncoder{out: out} }

func (d *pduDecoder) Read(parts ...any) error {
	for _, data := range parts {
		if skip, isSkip := data.(encSkip); isSkip {
			zeros := make([]byte, int(skip))
			if err := binary.Read(d.in[len(d.in)-1], binary.BigEndian, zeros); err != nil {
				return err
			}
		} else if err := binary.Read(d.in[len(d.in)-1], binary.BigEndian, data); err != nil {
			return err
		}
	}
	return nil
}

func (d *pduDecoder) PushLimit(l int) {
	d.in = append(d.in, &io.LimitedReader{R: d.in[len(d.in)-1], N: int64(l)})
}

func (d *pduDecoder) PopLimit() {
	if len(d.in) <= 1 {
		return
	}
	x := len(d.in) - 1
	d.in = d.in[:x:x]
}

func (d *pduDecoder) String(length int) (string, error) {
	data := make([]byte, length)
	return string(data), d.Read(&data)
}

func (d *pduEncoder) Write(parts ...any) error {
	for _, data := range parts {
		if skip, isSkip := data.(encSkip); isSkip {
			zeros := make([]byte, int(skip))
			if err := binary.Write(d.out, binary.BigEndian, zeros); err != nil {
				return err
			}
		} else if err := binary.Write(d.out, binary.BigEndian, data); err != nil {
			return err
		}
	}
	return nil
}
