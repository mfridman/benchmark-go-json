package benchmarkgojson

import (
	stdjson "encoding/json"
	"io"
	"log"
	"os"
	"testing"

	stdjsonv2 "github.com/go-json-experiment/json"
	goccyjson "github.com/goccy/go-json"
	segmentjson "github.com/segmentio/encoding/json"
)

var testData *AbbreviatedMetadata

func init() {
	fileContent, err := os.ReadFile("input.json")
	if err != nil {
		log.Fatal(err)
	}
	testData = new(AbbreviatedMetadata)
	if err := stdjson.Unmarshal(fileContent, &testData); err != nil {
		log.Fatal(err)
	}
}

func BenchmarkStandardLibrary(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for range b.N {
		if err := stdjson.NewEncoder(io.Discard).Encode(testData); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkStandardLibraryV2(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for range b.N {
		// if _, err := stdjsonv2.Marshal(testData); err != nil {
		// 	b.Error(err)
		// }
		if err := stdjsonv2.MarshalWrite(io.Discard, testData); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkGoccyJSON(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for range b.N {
		if err := goccyjson.NewEncoder(io.Discard).Encode(testData); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkSegmentJSON(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for range b.N {
		if err := segmentjson.NewEncoder(io.Discard).Encode(testData); err != nil {
			b.Error(err)
		}
	}
}
