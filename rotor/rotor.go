package rotor

import ()

type Rotor interface {
	Click() bool
	SetOffset(rune)
	EncodeRtoL(rune) rune
	EncodeLtoR(rune) rune
}

type EnigmaRotor struct {
	mapping        map[rune]rune
	inverseMapping map[rune]rune
	offset         int
}

func (e *EnigmaRotor) Click() bool {
	e.offset = (e.offset + 1) % 26

	if e.offset == 0 {
		return true
	} else {
		return false
	}
}

func (e *EnigmaRotor) SetOffset(r rune) {
	e.offset = int(r - 65)
}

func (e *EnigmaRotor) EncodeRtoL(r rune) rune {
	after_offset := rune(Modulo(int(r-65)+e.offset, 26) + 65)
	char := e.mapping[after_offset]
	undo_offset := rune(Modulo(int(char-65)-e.offset, 26) + 65)
	return undo_offset
}

func (e *EnigmaRotor) EncodeLtoR(r rune) rune {
	after_offset := rune(Modulo(int(r-65)+e.offset, 26) + 65)
	char := e.inverseMapping[after_offset]
	undo_offset := rune(Modulo(int(char-65)-e.offset, 26) + 65)
	return undo_offset
}

func NewRotor(config string) Rotor {
	offset1 := 0
	mapping1 := make(map[rune]rune)
	inverseMapping1 := make(map[rune]rune)
	for i := 0; i < 26; i++ {
		mapping1[rune(i+65)] = rune(config[i])
		inverseMapping1[rune(config[i])] = rune(i + 65)
	}
	r := &EnigmaRotor{mapping: mapping1, inverseMapping: inverseMapping1, offset: offset1}
	return r
}

func Modulo(a, b int) int {
	temp := a % b
	if temp < 0 {
		return (temp + b)
	}
	return temp
}
