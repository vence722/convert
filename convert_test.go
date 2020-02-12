package convert

import "testing"

func TestConvertSignedIntegers(t *testing.T) {
	// Test int
	originInt := int(-300)
	converted := Int2Str(originInt)
	if converted != "-300" {
		t.Logf("failed Test int")
		t.Fail()
	}

	// Test int64
	originInt64 := int64(-100)
	converted = Int2Str(originInt64)
	if converted != "-100" {
		t.Logf("failed Test int64")
		t.Fail()
	}

	// Test int32
	originInt32 := int32(-100)
	converted = Int2Str(originInt32)
	if converted != "-100" {
		t.Logf("failed Test int32")
		t.Fail()
	}

	// Test int16
	originInt16 := int16(-30)
	converted = Int2Str(originInt16)
	if converted != "-30" {
		t.Logf("failed Test int16")
		t.Fail()
	}

	// Test int8
	originInt8 := int8(-7)
	converted = Int2Str(originInt8)
	if converted != "-7" {
		t.Logf("failed Test int8")
		t.Fail()
	}
}

func TestConvertUnsignedIntegers(t *testing.T) {
	// Test uint
	originUint := uint(300)
	converted := Int2Str(originUint)
	if converted != "300" {
		t.Logf("failed Test uint")
		t.Fail()
	}

	// Test uint64
	originUint64 := uint(300)
	converted = Int2Str(originUint64)
	if converted != "300" {
		t.Logf("failed Test uint64")
		t.Fail()
	}

	// Test uint32
	originUint32 := uint32(300)
	converted = Int2Str(originUint32)
	if converted != "300" {
		t.Logf("failed Test uint32")
		t.Fail()
	}

	// Test uint16
	originUint16 := uint16(300)
	converted = Int2Str(originUint16)
	if converted != "300" {
		t.Logf("failed Test uint16")
		t.Fail()
	}

	// Test uint8
	originUint8 := uint8(255)
	converted = Int2Str(originUint8)
	if converted != "255" {
		t.Logf("failed Test uint8")
		t.Fail()
	}
}

func TestBytesSlice2StrNonCopy(t *testing.T) {
	origin := "Vence is smart"
	originBytes := []byte(origin)
	converted := BytesSlice2StrNonCopy(originBytes)
	if origin != converted {
		t.Fail()
	}
}