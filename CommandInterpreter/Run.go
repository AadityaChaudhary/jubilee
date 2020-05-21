package CommandInterpreter

import (
	"Jubilee/Driver"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func RunCommands() {
	for x,command := range commands.Commands {
		fmt.Print("step: " + string(x) + " : ")
		res := ExecuteCommand(command)
		if !res {
			return
		}
	}
}

func RunCommandLine() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("quit", text) == 0 {
			return
		}

	}
}

func ExecuteCommand(com Command) bool {

	var msg 	string
	var result	bool

	switch com.CommandType {

	case "open":
		msg, result = Driver.OpenSite(com.Link)
	case "click":
		msg, result = Driver.Click(com.Id)
	case "type":
		msg, result = Driver.Type(com.Id,com.Input)
	case "waitFor":
		msg, result = Driver.WaitForElem(com.Id, com.Timeout)
	}

	fmt.Println(msg)

	return result
}
