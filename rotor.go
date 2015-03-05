package Enigma


type Rotor interface {
  Click() bool
  SetOfset(rune)
  EncodeRtoL(rune) rune
  EncodeLtoR(rune) rune
}

type EnigmaRotor struct {
  mapping map[rune]rune
  inverseMapping map[rune]rune
  ofset int
}

func (e EnigmaRotor) Click() bool {
    e.ofset = (e.ofset + 1) % 26

    if (e.ofset == 0) {
      return true
    } else {
      return false
    }
}

func (e EnigmaRotor) SetOfset(r rune) {
  e.ofset = int(r-65)
}

func (e EnigmaRotor) EncodeRtoL(r rune) rune {
  after_ofset := rune(modulo(int(r - 65) + e.ofset, 26) + 65)
  char := e.mapping[after_ofset]
  undo_ofset := rune(modulo(int(char - 65) - e.ofset, 26) + 65)
  return undo_ofset
}

func (e EnigmaRotor) EncodeLtoR(r rune) rune {
  after_ofset := rune(modulo(int(r - 65) + e.ofset, 26) + 65)
  char := e.inverseMapping[after_ofset]
  undo_ofset := rune(modulo(int(char - 65) - e.ofset, 26) + 65)
  return undo_ofset
}

func NewRotor(config string) Rotor {
  ofset1 := 0
  mapping1 := make(map[rune]rune)
  inverseMapping1 := make(map[rune]rune)
  for i:=0; i<26;i++ {
    inverseMapping1[rune(i+65)] = rune(config[i])
    mapping1[rune(config[i])] = rune(i+65)
  }
  r := EnigmaRotor{mapping: mapping1, inverseMapping: inverseMapping1, ofset: ofset1}
  return r
}
