package enigma

import (
	"github.com/RobinVerachtert/GoEnigma/rotor"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestEnigma(t *testing.T) {
	Convey("[Enigma] full test", t, func() {
		r := [3]rotor.Rotor{}
		r[0] = rotor.NewRotor("XANTIPESOKRWUDVBCFGHJLMQYZ")
		r[1] = rotor.NewRotor("AJDKSIRUXBLHWTMCQGZNPYFVOE")
		r[2] = rotor.NewRotor("EKMFLGDQVZNTOWYHXUSPAIBRCJ")

		refl := NewReflector("YRUHQSLDPXNGOKMIEBFZCWVJAT")

		plug := NewPlugboard("PBMEDFLHIZKGCNOAQRSWUVTXYJ")

		e := NewEnigma(r, refl, plug)

		ret := e.Encode("P", "XEB")
		So(ret, ShouldEqual, "W")

		ret2 := e.Encode(ret, "XEB")
		So(ret2, ShouldEqual, "P")

		ret3 := e.Encode("ZUWWYOWUDROSEGFXLRVLMRXCCJHMBZKXGGOUMZMT", "AAA")
		So(ret3, ShouldEqual, "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	})

}
