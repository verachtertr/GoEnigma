package enigma

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestReflector(t *testing.T) {
	Convey("[Reflector] full test", t, func() {
		refl := NewReflector("YRUHQSLDPXNGOKMIEBFZCWVJAT")

		char1 := refl.Reflect('A')
		char2 := refl.Reflect('H')
		char3 := refl.Reflect('Z')

		So(char1, ShouldEqual, 'Y')
		So(char2, ShouldEqual, 'D')
		So(char3, ShouldEqual, 'T')
	})

}
