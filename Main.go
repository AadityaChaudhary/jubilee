package main

import (
	"Jubilee/CommandInterpreter"
	"Jubilee/Driver"
	"bufio"
	"fmt"
	"os"
	"strings"
)


func main() {

	var options string
	var path string

	if len(os.Args) < 2 {
		fmt.Println("Not enough Arguments!")
		return
	} else if len(os.Args) == 2 {
		OneArg()
		return
	} else if len(os.Args) == 3 {
		options, path = TwoArgs()
	} else {
		fmt.Println("Too many Args!")
		return
	}

	Begin(options, path)


}

func pause() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("continue", text) == 0 {
			return
		}

	}

}

func OneArg() {
	if strings.HasPrefix(os.Args[1], "-")  {
		if strings.Compare(os.Args[1], "-h") == 0 {
			Help()
		} else {
			fmt.Println("you need to provide a path as well")
		}

	} else {
		seleniumPath, driverPath, port, browser, err := CommandInterpreter.LoadFile(os.Args[1])
		if err != nil {
			panic(err)
		}
		Driver.Start(seleniumPath, driverPath, port, browser)
		CommandInterpreter.RunCommands()
	}
}

func TwoArgs() (string,string) {
	var options string
	if strings.HasPrefix(os.Args[1], "-") {
		options = os.Args[1]
	} else {
		options = "-"
	}

	path := os.Args[2]

	return options, path
}

func Begin(options string, path string) {
	if strings.Contains(options, "d") {
				//put folder code here
	} else {
		StartFile(options, path)
	}

}

func StartFile(options string, path string) {
	seleniumPath, driverPath, port, browser, err := CommandInterpreter.LoadFile(path)
	if err != nil {
		panic(err)
	}
	Driver.Start(seleniumPath, driverPath, port, browser)
	CommandInterpreter.RunCommands()

	if strings.Contains(options, "p") {
		pause()
	}
}

func Help() {
	fmt.Println("Format: [options] [path]")
	fmt.Println("d:		pass a directory path, and run all of the scripts in that directory")
	fmt.Println("p:		pause after running the script, instead of closing the browser")
}
