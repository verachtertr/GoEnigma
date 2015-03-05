package Enigma

import(
  "fmt"
  )

type Plugboard interface {
  plug(rune) rune
}

type EnigmaPlugboard struct {
  plugboard map[rune]rune
}

func (e EnigmaPlugboard) plug(r rune) rune {
  return e.plugboard[r]
}

func createPlugboard(config string) Plugboard {
  p := make(map[rune]rune)
  // set the plugboard according to the config string
  for i:=0; i<26;i++ {
    p[rune(i+65)] = rune(config[i])
  }

  // check that it is a 1 on 1 mapping
  for i:=0; i<26;i++ {
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
