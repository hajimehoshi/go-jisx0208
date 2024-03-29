// Copyright 2013 Hajime Hoshi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package jisx0208_test

import (
	"testing"

	. "github.com/hajimehoshi/go-jisx0208"
)

func TestConvert(t *testing.T) {
	testCases := []struct {
		input       int
		expectRune  rune
		expectError bool
	}{
		{0x2121, '　', false},
		{0x2422, 'あ', false},
		{0x306C, '一', false},
		{0x1000, 0, true},
	}
	for _, testCase := range testCases {
		input := testCase.input
		wantedRune := testCase.expectRune
		wantedError := testCase.expectError
		gotRune, gotError := Rune(input)
		if wantedRune != gotRune {
			t.Errorf("Rune(%d) = '%c', want '%c'",
				input, gotRune, wantedRune)
		}
		if !wantedError && gotError != nil {
			t.Errorf("Rune(%d) raises error \"%v\", want nil",
				input, gotError.Error())
		}
	}
}

func TestRevConvert(t *testing.T) {
	testCases := []struct {
		input       rune
		expectCode  int
		expectError bool
	}{
		{'　', 0x2121, false},
		{'あ', 0x2422, false},
		{'一', 0x306C, false},
		{0xE0000, 0, true}, // Private Use Area in BMP
	}
	for _, testCase := range testCases {
		input := testCase.input
		wantedCode := testCase.expectCode
		wantedError := testCase.expectError
		gotCode, gotError := Code(input)
		if wantedCode != gotCode {
			t.Errorf("Code(%c) = '%d', want '%d'",
				input, gotCode, wantedCode)
		}
		if !wantedError && gotError != nil {
			t.Errorf("Rune(%d) raises error \"%v\", want nil",
				input, gotError.Error())
		}
	}
}
