package enigma

import (
	"GoEnigma/rotor"
	"encoding/json"
	"errors"
	"os"
)

type RotorConfig struct {
	Id    int
	Setup string
}

type EnigmaConfig struct {
	ReflectorSetup  string
	PlugboardConfig string
	Rotors          [3]RotorConfig
}

// unmarshall enigma object
func ReadConfig(file_name string) (*EnigmaConfig, error) {

	//Open the config file
	config_file, err := os.Open(file_name)

	if err != nil {
		return nil, errors.New("Could not open config file: " + err.Error())
	}

	//Parse the config file
	enigma := EnigmaConfig{}

	jsonParser := json.NewDecoder(config_file)
	err = jsonParser.Decode(&enigma)

	if err != nil {
		return nil, errors.New("Could not parse config file: " + err.Error())
	} else {
		return &enigma, nil
	}

}

func (e EnigmaConfig) CreateEnigma() Enigma {

	// create the rotor
	r := [3]rotor.Rotor{}
	r[0] = rotor.NewRotor(e.Rotors[0].Setup)
	r[1] = rotor.NewRotor(e.Rotors[1].Setup)
	r[2] = rotor.NewRotor(e.Rotors[2].Setup)

	p := NewPlugboard(e.PlugboardConfig)
	refl := NewReflector(e.ReflectorSetup)

	return NewEnigma(r, refl, p)
}
