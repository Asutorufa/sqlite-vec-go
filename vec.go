package sqlitevecgo

import (
	"bytes"
	_ "embed"
	eb "encoding/binary"

	"github.com/ncruces/go-sqlite3"
)

//go:embed sqlite-vec.wasm
var binary []byte

func init() {
	sqlite3.Binary = binary
}

func SerializeFloat32(vector []float32) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := eb.Write(buf, eb.LittleEndian, vector)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
