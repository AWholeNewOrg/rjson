// Code generated by gen/fuzzgen. DO NOT EDIT.

package rjson

import (
	"encoding/json"
	"bytes"
)

var generatedFuzzers = []fuzzer{
	{name: "fuzzReadUint64Gen", fn: fuzzReadUint64Gen},
	{name: "fuzzReadUint32Gen", fn: fuzzReadUint32Gen},
	{name: "fuzzReadInt64Gen", fn: fuzzReadInt64Gen},
	{name: "fuzzReadInt32Gen", fn: fuzzReadInt32Gen},
	{name: "fuzzReadIntGen", fn: fuzzReadIntGen},
	{name: "fuzzReadUintGen", fn: fuzzReadUintGen},
	{name: "fuzzReadFloat64Gen", fn: fuzzReadFloat64Gen},
	{name: "fuzzReadStringBytesGen", fn: fuzzReadStringBytesGen},
	{name: "fuzzReadStringGen", fn: fuzzReadStringGen},
	{name: "fuzzReadBoolGen", fn: fuzzReadBoolGen},
	{name: "fuzzReadValueGen", fn: fuzzReadValueGen},
	{name: "fuzzReadObjectGen", fn: fuzzReadObjectGen},
	{name: "fuzzReadArrayGen", fn: fuzzReadArrayGen},
}

func fuzzReadUint64Gen(data []byte) (int, error) {
	nullP, nullErr := ReadNull(data)
	if nullErr == nil {
		return nullP, nullErr
	}
	var want uint64
	wantErr := json.NewDecoder(bytes.NewReader(data)).Decode(&want)
	got, _, gotErr := ReadUint64(data)
	err := checkFuzzErrors(wantErr, gotErr)
	if err != nil || wantErr != nil {
		return 0, err
	}
	gotVal, wantVal := removeJSONRuneError(got, want)
	err = fuzzCompare(wantVal, gotVal)
	return 0, err
}

func fuzzReadUint32Gen(data []byte) (int, error) {
	nullP, nullErr := ReadNull(data)
	if nullErr == nil {
		return nullP, nullErr
	}
	var want uint32
	wantErr := json.NewDecoder(bytes.NewReader(data)).Decode(&want)
	got, _, gotErr := ReadUint32(data)
	err := checkFuzzErrors(wantErr, gotErr)
	if err != nil || wantErr != nil {
		return 0, err
	}
	gotVal, wantVal := removeJSONRuneError(got, want)
	err = fuzzCompare(wantVal, gotVal)
	return 0, err
}

func fuzzReadInt64Gen(data []byte) (int, error) {
	nullP, nullErr := ReadNull(data)
	if nullErr == nil {
		return nullP, nullErr
	}
	var want int64
	wantErr := json.NewDecoder(bytes.NewReader(data)).Decode(&want)
	got, _, gotErr := ReadInt64(data)
	err := checkFuzzErrors(wantErr, gotErr)
	if err != nil || wantErr != nil {
		return 0, err
	}
	gotVal, wantVal := removeJSONRuneError(got, want)
	err = fuzzCompare(wantVal, gotVal)
	return 0, err
}

func fuzzReadInt32Gen(data []byte) (int, error) {
	nullP, nullErr := ReadNull(data)
	if nullErr == nil {
		return nullP, nullErr
	}
	var want int32
	wantErr := json.NewDecoder(bytes.NewReader(data)).Decode(&want)
	got, _, gotErr := ReadInt32(data)
	err := checkFuzzErrors(wantErr, gotErr)
	if err != nil || wantErr != nil {
		return 0, err
	}
	gotVal, wantVal := removeJSONRuneError(got, want)
	err = fuzzCompare(wantVal, gotVal)
	return 0, err
}

func fuzzReadIntGen(data []byte) (int, error) {
	nullP, nullErr := ReadNull(data)
	if nullErr == nil {
		return nullP, nullErr
	}
	var want int
	wantErr := json.NewDecoder(bytes.NewReader(data)).Decode(&want)
	got, _, gotErr := ReadInt(data)
	err := checkFuzzErrors(wantErr, gotErr)
	if err != nil || wantErr != nil {
		return 0, err
	}
	gotVal, wantVal := removeJSONRuneError(got, want)
	err = fuzzCompare(wantVal, gotVal)
	return 0, err
}

func fuzzReadUintGen(data []byte) (int, error) {
	nullP, nullErr := ReadNull(data)
	if nullErr == nil {
		return nullP, nullErr
	}
	var want uint
	wantErr := json.NewDecoder(bytes.NewReader(data)).Decode(&want)
	got, _, gotErr := ReadUint(data)
	err := checkFuzzErrors(wantErr, gotErr)
	if err != nil || wantErr != nil {
		return 0, err
	}
	gotVal, wantVal := removeJSONRuneError(got, want)
	err = fuzzCompare(wantVal, gotVal)
	return 0, err
}

func fuzzReadFloat64Gen(data []byte) (int, error) {
	nullP, nullErr := ReadNull(data)
	if nullErr == nil {
		return nullP, nullErr
	}
	var want float64
	wantErr := json.NewDecoder(bytes.NewReader(data)).Decode(&want)
	got, _, gotErr := ReadFloat64(data)
	err := checkFuzzErrors(wantErr, gotErr)
	if err != nil || wantErr != nil {
		return 0, err
	}
	gotVal, wantVal := removeJSONRuneError(got, want)
	err = fuzzCompare(wantVal, gotVal)
	return 0, err
}

func fuzzReadStringBytesGen(data []byte) (int, error) {
	nullP, nullErr := ReadNull(data)
	if nullErr == nil {
		return nullP, nullErr
	}
	var want string
	wantErr := json.NewDecoder(bytes.NewReader(data)).Decode(&want)
	got, _, gotErr := ReadStringBytes(data, nil)
	err := checkFuzzErrors(wantErr, gotErr)
	if err != nil || wantErr != nil {
		return 0, err
	}
	gotVal, wantVal := removeJSONRuneError(string(got), want)
	err = fuzzCompare(wantVal, gotVal)
	return 0, err
}

func fuzzReadStringGen(data []byte) (int, error) {
	nullP, nullErr := ReadNull(data)
	if nullErr == nil {
		return nullP, nullErr
	}
	var want string
	wantErr := json.NewDecoder(bytes.NewReader(data)).Decode(&want)
	got, _, gotErr := ReadString(data, nil)
	err := checkFuzzErrors(wantErr, gotErr)
	if err != nil || wantErr != nil {
		return 0, err
	}
	gotVal, wantVal := removeJSONRuneError(got, want)
	err = fuzzCompare(wantVal, gotVal)
	return 0, err
}

func fuzzReadBoolGen(data []byte) (int, error) {
	nullP, nullErr := ReadNull(data)
	if nullErr == nil {
		return nullP, nullErr
	}
	var want bool
	wantErr := json.NewDecoder(bytes.NewReader(data)).Decode(&want)
	got, _, gotErr := ReadBool(data)
	err := checkFuzzErrors(wantErr, gotErr)
	if err != nil || wantErr != nil {
		return 0, err
	}
	gotVal, wantVal := removeJSONRuneError(got, want)
	err = fuzzCompare(wantVal, gotVal)
	return 0, err
}

func fuzzReadValueGen(data []byte) (int, error) {
	nullP, nullErr := ReadNull(data)
	if nullErr == nil {
		return nullP, nullErr
	}
	var want interface{}
	wantErr := json.NewDecoder(bytes.NewReader(data)).Decode(&want)
	got, _, gotErr := ReadValue(data)
	err := checkFuzzErrors(wantErr, gotErr)
	if err != nil || wantErr != nil {
		return 0, err
	}
	gotVal, wantVal := removeJSONRuneError(got, want)
	err = fuzzCompare(wantVal, gotVal)
	return 0, err
}

func fuzzReadObjectGen(data []byte) (int, error) {
	nullP, nullErr := ReadNull(data)
	if nullErr == nil {
		return nullP, nullErr
	}
	var want map[string]interface{}
	wantErr := json.NewDecoder(bytes.NewReader(data)).Decode(&want)
	got, _, gotErr := ReadObject(data)
	err := checkFuzzErrors(wantErr, gotErr)
	if err != nil || wantErr != nil {
		return 0, err
	}
	gotVal, wantVal := removeJSONRuneError(got, want)
	err = fuzzCompare(wantVal, gotVal)
	return 0, err
}

func fuzzReadArrayGen(data []byte) (int, error) {
	nullP, nullErr := ReadNull(data)
	if nullErr == nil {
		return nullP, nullErr
	}
	var want []interface{}
	wantErr := json.NewDecoder(bytes.NewReader(data)).Decode(&want)
	got, _, gotErr := ReadArray(data)
	err := checkFuzzErrors(wantErr, gotErr)
	if err != nil || wantErr != nil {
		return 0, err
	}
	gotVal, wantVal := removeJSONRuneError(got, want)
	err = fuzzCompare(wantVal, gotVal)
	return 0, err
}
