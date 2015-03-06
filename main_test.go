package main

import (
	"GoEnigma/enigma"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMain(t *testing.T) {
	Convey("[Main] full test with config.json", t, func() {

		enigma_config, err := enigma.ReadConfig("jsonConfigFiles/config.json")
		noError := true
		if err != nil {
			noError = false
		}

		So(noError, ShouldEqual, true)

		e := enigma_config.CreateEnigma()

		encoded := e.Encode("TEST", "AAA")
		So(encoded, ShouldEqual, "ETOE")

		encoded = e.Encode("Lorem ipsum dolor sit amet consectetur adipiscing elit sed do eiusmod tempor incididunt ut labore et dolore magna aliqua Ut enim ad minim veniam quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur Excepteur sint occaecat cupidatat non proident sunt in culpa qui officia deserunt mollit anim id est laborum", "XGT")
		So(encoded, ShouldEqual, "VJPRSGXMCSHFNFEMGVYSRVUFLMRUVUYEXWOLHLMQLJZUILYMUOODULEVYZJETYKZQBLFBJBJWLEWENXIZVDDOETZTVDQYCIYYZNMFYFODINYXEYJCJYWDCJXYSLJQCKQZGLEDSKQWZYMYZJFLUUMAWJUMGJQZXFJRJDNYFJZDJVOHOYWEBBEUEWETFGALQMCLPFQLMGPJLJGCXUXJPYAGYAGVGFCGAXSXFHBZJYSRSGHGHDAOKKOJDHHWFIEHEGOOWMWRDTRBQGGJOJTHJRQTDCXDORDQTMHBIHZZLBZLIZNJTVLILIUHUJFHTVOYVGWYVRYHWKQJPWREDDRHRJFOJKMNUQRHXXVQDUVRVAKJQXDWHMVS")
	})

}
