// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     stream_data_record_message_schema.avsc
 */
package kafkaexporter

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

type UnionNullMapStringTypeEnum int

const (
	UnionNullMapStringTypeEnumMapString UnionNullMapStringTypeEnum = 1
)

type UnionNullMapString struct {
	Null      *types.NullVal
	MapString map[string]string
	UnionType UnionNullMapStringTypeEnum
}

func writeUnionNullMapString(r *UnionNullMapString, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullMapStringTypeEnumMapString:
		return writeMapString(r.MapString, w)
	}
	return fmt.Errorf("invalid value for *UnionNullMapString")
}

func NewUnionNullMapString() *UnionNullMapString {
	return &UnionNullMapString{}
}

func (r *UnionNullMapString) Serialize(w io.Writer) error {
	return writeUnionNullMapString(r, w)
}

func DeserializeUnionNullMapString(r io.Reader) (*UnionNullMapString, error) {
	t := NewUnionNullMapString()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, t)

	if err != nil {
		return t, err
	}
	return t, err
}

func DeserializeUnionNullMapStringFromSchema(r io.Reader, schema string) (*UnionNullMapString, error) {
	t := NewUnionNullMapString()
	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, t)

	if err != nil {
		return t, err
	}
	return t, err
}

func (r *UnionNullMapString) Schema() string {
	return "[\"null\",{\"type\":\"map\",\"values\":\"string\"}]"
}

func (_ *UnionNullMapString) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullMapString) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullMapString) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullMapString) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullMapString) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullMapString) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullMapString) SetLong(v int64) {

	r.UnionType = (UnionNullMapStringTypeEnum)(v)
}

func (r *UnionNullMapString) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.MapString = make(map[string]string)
		return &MapStringWrapper{Target: (&r.MapString)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullMapString) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullMapString) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullMapString) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullMapString) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullMapString) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullMapString) Finalize()                        {}

func (r *UnionNullMapString) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullMapStringTypeEnumMapString:
		return json.Marshal(map[string]interface{}{"map": r.MapString})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullMapString")
}

func (r *UnionNullMapString) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["map"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.MapString)
	}
	return fmt.Errorf("invalid value for *UnionNullMapString")
}
