package enigma

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestPlugboard(t *testing.T) {
	Convey("[Plugboard] full test", t, func() {
		plug := NewPlugboard("PBMEDFLHIZKGCNOAQRSWUVTXYJ")

		char1 := plug.Plug('A')
		char2 := plug.Plug('H')
		char3 := plug.Plug('Z')

		So(char1, ShouldEqual, 'P')
		So(char2, ShouldEqual, 'H')
		So(char3, ShouldEqual, 'J')

		So('A', ShouldEqual, plug.Plug(plug.Plug('A')))
		So('H', ShouldEqual, plug.Plug(plug.Plug('H')))
		So('Z', ShouldEqual, plug.Plug(plug.Plug('Z')))
	})

}
