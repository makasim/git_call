package git_call

import (
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"io"
)

// todo: stdin
// todo: arrays and objects
// todo path validation
//reader := strings.NewReader(`
//{"user": {
//	"balance": 0}}
//`)

type Data struct {
	json string
}

func NewData(json string) *Data {
	return &Data{
		json: json,
	}
}

func (d *Data) Set(path string, v string) {
	d.json, _ = sjson.Set(d.json, path, v)
}

func (d *Data) SetInt64(path string, v int64) {
	d.json, _ = sjson.Set(d.json, path, v)
}

func (d Data) GetInt64(path string) int64 {
	return gjson.Get(d.json, path).Int()
}

func (d Data) Get(path string) string {
	return gjson.Get(d.json, path).String()
}

func (d Data) Fprint(writer io.Writer) {
	fmt.Fprint(writer, d.json)
}
