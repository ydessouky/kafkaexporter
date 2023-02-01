// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     stream_data_record_message_schema.avsc
 */
package kafkaexporter

import (
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
	"io"
)

func writeMapInt(r map[string]int32, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for k, e := range r {
		err = vm.WriteString(k, w)
		if err != nil {
			return err
		}
		err = vm.WriteInt(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type MapIntWrapper struct {
	Target *map[string]int32
	keys   []string
	values []int32
}

func (_ *MapIntWrapper) SetBoolean(v bool)     { panic("Unsupported operation") }
func (_ *MapIntWrapper) SetInt(v int32)        { panic("Unsupported operation") }
func (_ *MapIntWrapper) SetLong(v int64)       { panic("Unsupported operation") }
func (_ *MapIntWrapper) SetFloat(v float32)    { panic("Unsupported operation") }
func (_ *MapIntWrapper) SetDouble(v float64)   { panic("Unsupported operation") }
func (_ *MapIntWrapper) SetBytes(v []byte)     { panic("Unsupported operation") }
func (_ *MapIntWrapper) SetString(v string)    { panic("Unsupported operation") }
func (_ *MapIntWrapper) SetUnionElem(v int64)  { panic("Unsupported operation") }
func (_ *MapIntWrapper) Get(i int) types.Field { panic("Unsupported operation") }
func (_ *MapIntWrapper) SetDefault(i int)      { panic("Unsupported operation") }

func (r *MapIntWrapper) HintSize(s int) {
	if r.keys == nil {
		r.keys = make([]string, 0, s)
		r.values = make([]int32, 0, s)
	}
}

func (r *MapIntWrapper) NullField(_ int) {
	panic("Unsupported operation")
}

func (r *MapIntWrapper) Finalize() {
	for i := range r.keys {
		(*r.Target)[r.keys[i]] = r.values[i]
	}
}

func (r *MapIntWrapper) AppendMap(key string) types.Field {
	r.keys = append(r.keys, key)
	var v int32
	r.values = append(r.values, v)
	return &types.Int{Target: &r.values[len(r.values)-1]}
}

func (_ *MapIntWrapper) AppendArray() types.Field { panic("Unsupported operation") }
