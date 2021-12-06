//go:build gofuzzbeta
// +build gofuzzbeta

package rjson

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func FuzzReadFloat64(f *testing.F)        { fuzzFuzzer(f, fuzzReadFloat64) }
func FuzzReadUint64(f *testing.F)         { fuzzFuzzer(f, fuzzReadUint64) }
func FuzzReadUint32(f *testing.F)         { fuzzFuzzer(f, fuzzReadUint32) }
func FuzzReadUint(f *testing.F)           { fuzzFuzzer(f, fuzzReadUint) }
func FuzzReadInt64(f *testing.F)          { fuzzFuzzer(f, fuzzReadInt64) }
func FuzzReadInt32(f *testing.F)          { fuzzFuzzer(f, fuzzReadInt32) }
func FuzzReadInt(f *testing.F)            { fuzzFuzzer(f, fuzzReadInt) }
func FuzzReadString(f *testing.F)         { fuzzFuzzer(f, fuzzReadString) }
func FuzzReadStringBytes(f *testing.F)    { fuzzFuzzer(f, fuzzReadStringBytes) }
func FuzzReadBool(f *testing.F)           { fuzzFuzzer(f, fuzzReadBool) }
func FuzzReadNull(f *testing.F)           { fuzzFuzzer(f, fuzzReadNull) }
func FuzzSkipValue(f *testing.F)          { fuzzFuzzer(f, fuzzSkipValue) }
func FuzzValid(f *testing.F)              { fuzzFuzzer(f, fuzzValid) }
func FuzzNextToken(f *testing.F)          { fuzzFuzzer(f, fuzzNextToken) }
func FuzzReadArray(f *testing.F)          { fuzzFuzzer(f, fuzzReadArray) }
func FuzzReadObject(f *testing.F)         { fuzzFuzzer(f, fuzzReadObject) }
func FuzzReadValue(f *testing.F)          { fuzzFuzzer(f, fuzzReadValue) }
func FuzzDecodeFloat64(f *testing.F)      { fuzzFuzzer(f, fuzzDecodeFloat64) }
func FuzzDecodeUint64(f *testing.F)       { fuzzFuzzer(f, fuzzDecodeUint64) }
func FuzzDecodeUint32(f *testing.F)       { fuzzFuzzer(f, fuzzDecodeUint32) }
func FuzzDecodeUint(f *testing.F)         { fuzzFuzzer(f, fuzzDecodeUint) }
func FuzzDecodeInt64(f *testing.F)        { fuzzFuzzer(f, fuzzDecodeInt64) }
func FuzzDecodeInt32(f *testing.F)        { fuzzFuzzer(f, fuzzDecodeInt32) }
func FuzzDecodeInt(f *testing.F)          { fuzzFuzzer(f, fuzzDecodeInt) }
func FuzzDecodeString(f *testing.F)       { fuzzFuzzer(f, fuzzDecodeString) }
func FuzzDecodeBool(f *testing.F)         { fuzzFuzzer(f, fuzzDecodeBool) }
func FuzzHandleArrayValues(f *testing.F)  { fuzzFuzzer(f, fuzzHandleArrayValues) }
func FuzzHandleObjectValues(f *testing.F) { fuzzFuzzer(f, fuzzHandleObjectValues) }

func FuzzAll(f *testing.F) {
	f.SkipNow()
	addCorpus(f)
	f.Fuzz(func(t *testing.T, data string) {
		for _, fd := range fuzzers {
			d := make([]byte, len(data))
			copy(d, data)
			_, err := fd.fn(d)
			require.NoError(t, err)
		}
	})
}

func addCorpus(f *testing.F) {
	f.Helper()
	corpusDir := filepath.FromSlash(`testdata/fuzz/corpus`)
	dir, err := ioutil.ReadDir(corpusDir)
	require.NoError(f, err)
	for _, info := range dir {
		if info.IsDir() {
			continue
		}
		dd, err := ioutil.ReadFile(filepath.Join(corpusDir, info.Name()))
		require.NoError(f, err)
		f.Add(string(dd))
	}
}

func fuzzFuzzer(f *testing.F, fn func(data []byte) (int, error)) {
	f.Helper()
	addCorpus(f)
	f.Fuzz(func(t *testing.T, data string) {
		_, err := fn([]byte(data))
		require.NoError(t, err)
	})
}

