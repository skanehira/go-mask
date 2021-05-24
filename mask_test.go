package mask

import (
	"encoding/json"
	"os"
	"path"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var (
	str  string  = "string"
	i    int     = 32
	i8   int8    = int8(i)
	i16  int16   = int16(i)
	i32  int32   = int32(i)
	i64  int64   = int64(i)
	b    byte    = byte(32)
	r    rune    = 32
	ui   uint    = 32
	ui8  uint8   = uint8(ui)
	ui16 uint16  = uint16(i)
	ui32 uint32  = uint32(i)
	ui64 uint64  = uint64(i)
	f32  float32 = float32(32.32)
	f64  float64 = float64(32.32)
)

type Struct struct {
	String string `mask:"string"`
}

type MaskStruct struct {
	String      string   `mask:"string"`
	StringPtr   *string  `mask:"string_ptr"`
	Byte        byte     `mask:"byte"`
	BytePtr     *byte    `mask:"byte_ptr"`
	Rune        rune     `mask:"rune"`
	RunPtr      *rune    `mask:"run_ptr"`
	Uint        uint     `mask:"uint"`
	UintPtr     *uint    `mask:"uint_ptr"`
	Uint8       uint8    `mask:"uint_8"`
	Uint8Ptr    *uint8   `mask:"uint_8_ptr"`
	Uint16      uint16   `mask:"uint_16"`
	Uint16Ptr   *uint16  `mask:"uint_16_ptr"`
	Uint32      uint32   `mask:"uint_32"`
	Uint32Ptr   *uint32  `mask:"uint_32_ptr"`
	Uint64      uint64   `mask:"uint_64"`
	Uint64Ptr   *uint64  `mask:"uint_64_ptr"`
	Int         int      `mask:"int"`
	IntPtr      *int     `mask:"int_ptr"`
	Int8        int8     `mask:"int_8"`
	Int8Ptr     *int8    `mask:"int_8_ptr"`
	Int16       int16    `mask:"int_16"`
	Int16Ptr    *int16   `mask:"int_16_ptr"`
	Int32       int32    `mask:"int_32"`
	Int32Ptr    *int32   `mask:"int_32_ptr"`
	Int64       int64    `mask:"int_64"`
	Int64Ptr    *int64   `mask:"int_64_ptr"`
	Float32     float32  `mask:"float_32"`
	Float32Ptr  *float32 `mask:"float_32_ptr"`
	Float64     float64  `mask:"float_64"`
	Float64Ptr  *float64 `mask:"float_64_ptr"`
	StructNoTag Struct
	Struct      Struct  `mask:"struct"`
	StructPtr   *Struct `mask:"struct_ptr"`
	NilPtr      *int    `mask:"nil_ptr"`
}

type StructNoTagged struct {
	String string
}

type UnMaskStruct struct {
	String     string
	StringPtr  *string
	Byte       byte
	BytePtr    *byte
	Rune       rune
	RunPtr     *rune
	Uint       uint
	UintPtr    *uint
	Uint8      uint8
	Uint8Ptr   *uint8
	Uint16     uint16
	Uint16Ptr  *uint16
	Uint32     uint32
	Uint32Ptr  *uint32
	Uint64     uint64
	Uint64Ptr  *uint64
	Int        int
	IntPtr     *int
	Int8       int8
	Int8Ptr    *int8
	Int16      int16
	Int16Ptr   *int16
	Int32      int32
	Int32Ptr   *int32
	Int64      int64
	Int64Ptr   *int64
	Float32    float32
	Float32Ptr *float32
	Float64    float64
	Float64Ptr *float64
	Struct     StructNoTagged
	StructPtr  *StructNoTagged
}

func TestMaskStruct(t *testing.T) {
	s := Struct{
		String: str,
	}
	ms := MaskStruct{
		String:     str,
		StringPtr:  &str,
		Byte:       b,
		BytePtr:    &b,
		Rune:       r,
		RunPtr:     &r,
		Uint:       ui,
		UintPtr:    &ui,
		Uint8:      ui8,
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
		Int16:      i16,
		Int16Ptr:   &i16,
		Int32:      i32,
		Int32Ptr:   &i32,
		Int64:      i64,
		Int64Ptr:   &i64,
		Float32:    f32,
		Float32Ptr: &f32,
		Float64:    f64,
		Float64Ptr: &f64,
		Struct:     s,
		StructPtr:  &s,
		NilPtr:     nil,
	}

	got := Mask(ms)
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

func TestUnMaskStruct(t *testing.T) {
	s := StructNoTagged{
		String: str,
	}
	ms := UnMaskStruct{
		String:     str,
		StringPtr:  &str,
		Byte:       b,
		BytePtr:    &b,
		Rune:       r,
		RunPtr:     &r,
		Uint:       ui,
		UintPtr:    &ui,
		Uint8:      ui8,
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
		Int16:      i16,
		Int16Ptr:   &i16,
		Int32:      i32,
		Int32Ptr:   &i32,
		Int64:      i64,
		Int64Ptr:   &i64,
		Float32:    f32,
		Float32Ptr: &f32,
		Float64:    f64,
		Float64Ptr: &f64,
		Struct:     s,
		StructPtr:  &s,
	}

	got := Mask(ms)
	var want UnMaskStruct
	filename := path.Join("testdata", "unmask_struct.json.golden")
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

func TestMaskPtrStruct(t *testing.T) {
	s := Struct{
		String: str,
	}
	ms := MaskStruct{
		String:     str,
		StringPtr:  &str,
		Byte:       b,
		BytePtr:    &b,
		Rune:       r,
		RunPtr:     &r,
		Uint:       ui,
		UintPtr:    &ui,
		Uint8:      ui8,
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
		Int16:      i16,
		Int16Ptr:   &i16,
		Int32:      i32,
		Int32Ptr:   &i32,
		Int64:      i64,
		Int64Ptr:   &i64,
		Float32:    f32,
		Float32Ptr: &f32,
		Float64:    f64,
		Float64Ptr: &f64,
		Struct:     s,
		StructPtr:  &s,
		NilPtr:     nil,
	}

	got := Mask(&ms)
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
		t.Errorf("mask pointer struct mismatch (-want +got):\n%s", diff)
	}
}
