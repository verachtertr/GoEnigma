package rotor

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestRotor(t *testing.T) {
	Convey("[Rotor] full test", t, func() {

		r1 := NewRotor("XANTIPESOKRWUDVBCFGHJLMQYZ")
		r2 := NewRotor("AJDKSIRUXBLHWTMCQGZNPYFVOE")
		r3 := NewRotor("EKMFLGDQVZNTOWYHXUSPAIBRCJ")

		_ = r3 // to prevent compile error: r3 declared and not used

		char1 := r1.EncodeRtoL('A')
		char2 := r1.EncodeLtoR('A')
		char3 := r2.EncodeRtoL('A')
		char4 := r2.EncodeRtoL('B')
		char5 := r2.EncodeLtoR('J')

		So(char1, ShouldEqual, 'X')
		So(char2, ShouldEqual, 'B')
		So(char3, ShouldEqual, 'A')
		So(char4, ShouldEqual, 'J')
		So(char5, ShouldEqual, 'B')

	})

}
