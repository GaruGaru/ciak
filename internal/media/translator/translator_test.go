package translator

import (
	"testing"
)

func TestDecoderNameWithTorrentNotation(t *testing.T) {

	input := "Godzilla.2014.720p.BluRay.x264-LEONARDO_[scarabey.org].mkv"

	output := Translate(input)

	expectedOutput := "Godzilla"

	if output != expectedOutput {
		t.Fatalf("expected %s as decoded output but got %s", expectedOutput, output)
	}

}

func TestDecoderSimpleName(t *testing.T) {

	input := "Godzilla 2014"

	output := Translate(input)

	expectedOutput := "Godzilla"

	if output != expectedOutput {
		t.Fatalf("expected %s as decoded output but got %s", expectedOutput, output)
	}

}
