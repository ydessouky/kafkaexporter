// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     stream_data_record_message_schema.avsc
 */
package kafkaexporter

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

func writeArrayDateTimeColumnFormatDTO(r []DateTimeColumnFormatDTO, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writeDateTimeColumnFormatDTO(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayDateTimeColumnFormatDTOWrapper struct {
	Target *[]DateTimeColumnFormatDTO
}

func (_ ArrayDateTimeColumnFormatDTOWrapper) SetBoolean(v bool)     { panic("Unsupported operation") }
func (_ ArrayDateTimeColumnFormatDTOWrapper) SetInt(v int32)        { panic("Unsupported operation") }
func (_ ArrayDateTimeColumnFormatDTOWrapper) SetLong(v int64)       { panic("Unsupported operation") }
func (_ ArrayDateTimeColumnFormatDTOWrapper) SetFloat(v float32)    { panic("Unsupported operation") }
func (_ ArrayDateTimeColumnFormatDTOWrapper) SetDouble(v float64)   { panic("Unsupported operation") }
func (_ ArrayDateTimeColumnFormatDTOWrapper) SetBytes(v []byte)     { panic("Unsupported operation") }
func (_ ArrayDateTimeColumnFormatDTOWrapper) SetString(v string)    { panic("Unsupported operation") }
func (_ ArrayDateTimeColumnFormatDTOWrapper) SetUnionElem(v int64)  { panic("Unsupported operation") }
func (_ ArrayDateTimeColumnFormatDTOWrapper) Get(i int) types.Field { panic("Unsupported operation") }
func (_ ArrayDateTimeColumnFormatDTOWrapper) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ ArrayDateTimeColumnFormatDTOWrapper) Finalize()        {}
func (_ ArrayDateTimeColumnFormatDTOWrapper) SetDefault(i int) { panic("Unsupported operation") }
func (r ArrayDateTimeColumnFormatDTOWrapper) HintSize(s int) {
	if len(*r.Target) == 0 {
		*r.Target = make([]DateTimeColumnFormatDTO, 0, s)
	}
}
func (r ArrayDateTimeColumnFormatDTOWrapper) NullField(i int) {
	panic("Unsupported operation")
}

func (r ArrayDateTimeColumnFormatDTOWrapper) AppendArray() types.Field {
	var v DateTimeColumnFormatDTO
	v = NewDateTimeColumnFormatDTO()

	*r.Target = append(*r.Target, v)
	return &types.Record{Target: &(*r.Target)[len(*r.Target)-1]}
}
