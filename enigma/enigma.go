package enigma

import (
	"GoEnigma/rotor"
	"strings"
)

type Enigma interface {
	Encode(text string, rotorPos string) string
}

type EnigmaClass struct {
	reflector Reflector
	plugboard Plugboard
	rotors    [3]rotor.Rotor
}

func (e *EnigmaClass) encode(text string, rotorPos string) string {
	// set the start position of the rotors:
	e.rotors[0].SetOfset(rune(rotorPos[0]))
	e.rotors[1].SetOfset(rune(rotorPos[1]))
	e.rotors[2].SetOfset(rune(rotorPos[2]))

	// clear the text of any whitespace
	r := strings.NewReplacer(" ", "")
	text = r.Replace(text)
	// set all characters to uppercase
	text = strings.ToUpper(text)

	//encode letter per letter
	newText := ""
	for i := 0; i < len(text); i++ {
		working := rune(text[i])

		working = e.plugboard.Plug(working)

		working = e.rotors[0].EncodeRtoL(working)
		working = e.rotors[1].EncodeRtoL(working)
		working = e.rotors[2].EncodeRtoL(working)

		working = e.reflector.Reflect(working)

		working = e.rotors[0].EncodeLtoR(working)
		working = e.rotors[1].EncodeLtoR(working)
		working = e.rotors[2].EncodeLtoR(working)

		working = e.plugboard.Plug(working)

		newText += string(working)

		// click the rotors
		click := true
		for j := 0; j < 3; j++ {
			if click {
				click = e.rotors[j].Click()
			}
		}
	}

	return newText
}
