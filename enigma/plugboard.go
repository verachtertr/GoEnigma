package enigma

import(
  "fmt"
)

type Plugboard interface {
  Plug(rune) rune
}

type EnigmaPlugboard struct {
  plugboard map[rune]rune
}

func (e EnigmaPlugboard) Plug(r rune) rune {
  return e.plugboard[r]
}

func NewPlugboard(config string) Plugboard {
  p := make(map[rune]rune)
  // set the plugboard according to the config string
  for i:=0; i<26;i++ {
    p[rune(i+65)] = rune(config[i])
  }

  // check that it is a 1 on 1 mapping
  for i:=int('A'); i<=int('Z');i++ {
    val1 := p[rune(i)]
    val2 := p[val1]

    if val2 != rune(i) {
      fmt.Printf("Plugboard: You did not give a correct config string, %q and %q should be a 1 on 1 mapping ", val1, val2 )
      // TODO exit the [program]
    }
  }

  pb := EnigmaPlugboard{plugboard: p}
  return pb

}
