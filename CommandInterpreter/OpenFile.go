package CommandInterpreter

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func LoadFile(path string)  (string, string, int, string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", "", 0, "", err
	}

	byteVal, err := ioutil.ReadAll(file)
	if err != nil {
		return "", "", 0, "", err
	}

	//unloading the json into the commands variable
	err = json.Unmarshal(byteVal, &commands)
	if err != nil {
		return "", "", 0, "", err
	}

	return commands.SeleniumPath, commands.DriverPath, commands.Port, commands.Browser, nil

}
