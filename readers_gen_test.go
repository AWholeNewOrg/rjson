// Code generated by gen/testgen. DO NOT EDIT.

package rjson

import "testing"

func TestReadUint64_gen(t *testing.T) {
	t.Parallel()
	testReadChecker(t, genReadUint64Checker, true)
}

func genReadUint64Checker(t *testing.T, data []byte) int {
	t.Helper()
	got, p, err := ReadUint64(data)
	return assertMatchesUnmarshal(t, data, got, p, err)
}

func TestReadUint32_gen(t *testing.T) {
	t.Parallel()
	testReadChecker(t, genReadUint32Checker, true)
}

func genReadUint32Checker(t *testing.T, data []byte) int {
	t.Helper()
	got, p, err := ReadUint32(data)
	return assertMatchesUnmarshal(t, data, got, p, err)
}

func TestReadInt64_gen(t *testing.T) {
	t.Parallel()
	testReadChecker(t, genReadInt64Checker, true)
}

func genReadInt64Checker(t *testing.T, data []byte) int {
	t.Helper()
	got, p, err := ReadInt64(data)
	return assertMatchesUnmarshal(t, data, got, p, err)
}

func TestReadInt32_gen(t *testing.T) {
	t.Parallel()
	testReadChecker(t, genReadInt32Checker, true)
}

func genReadInt32Checker(t *testing.T, data []byte) int {
	t.Helper()
	got, p, err := ReadInt32(data)
	return assertMatchesUnmarshal(t, data, got, p, err)
}

func TestReadInt_gen(t *testing.T) {
	t.Parallel()
	testReadChecker(t, genReadIntChecker, true)
}

func genReadIntChecker(t *testing.T, data []byte) int {
	t.Helper()
	got, p, err := ReadInt(data)
	return assertMatchesUnmarshal(t, data, got, p, err)
}

func TestReadUint_gen(t *testing.T) {
	t.Parallel()
	testReadChecker(t, genReadUintChecker, true)
}

func genReadUintChecker(t *testing.T, data []byte) int {
	t.Helper()
	got, p, err := ReadUint(data)
	return assertMatchesUnmarshal(t, data, got, p, err)
}

func TestReadFloat64_gen(t *testing.T) {
	t.Parallel()
	testReadChecker(t, genReadFloat64Checker, true)
}

func genReadFloat64Checker(t *testing.T, data []byte) int {
	t.Helper()
	got, p, err := ReadFloat64(data)
	return assertMatchesUnmarshal(t, data, got, p, err)
}

func TestReadStringBytes_gen(t *testing.T) {
	t.Parallel()
	testReadChecker(t, genReadStringBytesChecker, true)
}

func genReadStringBytesChecker(t *testing.T, data []byte) int {
	t.Helper()
	got, p, err := ReadStringBytes(data, nil)
	return assertMatchesUnmarshal(t, data, string(got), p, err)
}

func TestReadString_gen(t *testing.T) {
	t.Parallel()
	testReadChecker(t, genReadStringChecker, true)
}

func genReadStringChecker(t *testing.T, data []byte) int {
	t.Helper()
	got, p, err := ReadString(data, nil)
	return assertMatchesUnmarshal(t, data, string(got), p, err)
}

func TestReadBool_gen(t *testing.T) {
	t.Parallel()
	testReadChecker(t, genReadBoolChecker, true)
}

func genReadBoolChecker(t *testing.T, data []byte) int {
	t.Helper()
	got, p, err := ReadBool(data)
	return assertMatchesUnmarshal(t, data, got, p, err)
}

func TestReadValue_gen(t *testing.T) {
	t.Parallel()
	testReadChecker(t, genReadValueChecker, true)
}

func genReadValueChecker(t *testing.T, data []byte) int {
	t.Helper()
	got, p, err := ReadValue(data)
	return assertMatchesUnmarshal(t, data, got, p, err)
}

func TestReadObject_gen(t *testing.T) {
	t.Parallel()
	testReadChecker(t, genReadObjectChecker, false)
}

func genReadObjectChecker(t *testing.T, data []byte) int {
	t.Helper()
	got, p, err := ReadObject(data)
	return assertMatchesUnmarshal(t, data, got, p, err)
}

func TestReadArray_gen(t *testing.T) {
	t.Parallel()
	testReadChecker(t, genReadArrayChecker, false)
}

func genReadArrayChecker(t *testing.T, data []byte) int {
	t.Helper()
	got, p, err := ReadArray(data)
	return assertMatchesUnmarshal(t, data, got, p, err)
}