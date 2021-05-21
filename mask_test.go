package mask

import (
	"encoding/json"
	"os"
	"path"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMaskStruct(t *testing.T) {
	type MaskStruct struct {
		String     string   `mask:"string"`
		StringPtr  *string  `mask:"string_ptr"`
		Byte       byte     `mask:"byte"`
		BytePtr    *byte    `mask:"byte_ptr"`
		Rune       rune     `mask:"rune"`
		RunPtr     *rune    `mask:"run_ptr"`
		Uint       uint     `mask:"uint"`
		UintPtr    *uint    `mask:"uint_ptr"`
		Uint8      uint8    `mask:"uint_8"`
		Uint8Ptr   *uint8   `mask:"uint_8_ptr"`
		Uint16     uint16   `mask:"uint_16"`
		Uint16Ptr  *uint16  `mask:"uint_16_ptr"`
		Uint32     uint32   `mask:"uint_32"`
		Uint32Ptr  *uint32  `mask:"uint_32_ptr"`
		Uint64     uint64   `mask:"uint_64"`
		Uint64Ptr  *uint64  `mask:"uint_64_ptr"`
		Int        int      `mask:"int"`
		IntPtr     *int     `mask:"int_ptr"`
		Int8       int8     `mask:"int_8"`
		Int8Ptr    *int8    `mask:"int_8_ptr"`
		Int32      int32    `mask:"int_32"`
		Int32Ptr   *int32   `mask:"int_32_ptr"`
		Int64      int64    `mask:"int_64"`
		Int64Ptr   *int64   `mask:"int_64_ptr"`
		Float32    float32  `mask:"float_32"`
		Float32Ptr *float32 `mask:"float_32_ptr"`
		Float64    float64  `mask:"float_64"`
		Float64Ptr *float64 `mask:"float_64_ptr"`
	}

	var str string = "hello"
	var i int = 32
	var i8 int8 = int8(i)
	var i32 int32 = int32(i)
	var i64 int64 = int64(i)
	var b byte = byte(32)
	var r rune = 32
	var ui uint = 32
	var ui8 uint8 = uint8(ui)
	var ui16 uint16 = uint16(i)
	var ui32 uint32 = uint32(i)
	var ui64 uint64 = uint64(i)
	var f32 float32 = float32(32.0)
	var f64 float64 = float64(32.0)

	s := MaskStruct{
		String:     "string",
		StringPtr:  &str,
		Byte:       b,
		BytePtr:    &b,
		Rune:       r,
		RunPtr:     &r,
		Uint:       ui,
		UintPtr:    &ui,
		Uint8:      uint8(ui),
		Uint8Ptr:   &ui8,
		Uint16:     ui16,
		Uint16Ptr:  &ui16,
		Uint32:     ui32,
		Uint32Ptr:  &ui32,
		Uint64:     ui64,
		Uint64Ptr:  &ui64,
		Int:        i,
		IntPtr:     &i,
		Int8:       i8,
		Int8Ptr:    &i8,
		Int32:      i32,
		Int32Ptr:   &i32,
		Int64:      i64,
		Int64Ptr:   &i64,
		Float32:    f32,
		Float32Ptr: &f32,
		Float64:    f64,
		Float64Ptr: &f64,
	}

	got := Mask(s)
	var want MaskStruct
	filename := path.Join("testdata", "mask_struct.json.golden")
	f, err := os.Open(filename)
	if err != nil {
		t.Fatalf("cannot open file %s: %s", filename, err)
	}
	defer f.Close()

	if err := json.NewDecoder(f).Decode(&want); err != nil {
		t.Fatalf("cannot decode %s to %T: %s", f.Name(), want, err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mask struct mismatch (-want +got):\n%s", diff)
	}
}
