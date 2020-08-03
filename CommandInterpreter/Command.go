package CommandInterpreter

type Commands struct {
	SeleniumPath	string		`json:"selenium-path"`
	DriverPath		string		`json:"driver-path"`
	Port			int			`json:"port"`
	Browser 		string		`json:"browser"`
	Commands 		[]Command	`json:"commands"`

}

type Command struct {
	OutputPath	string	`json:"output"`
	Input 		string 	`json:"input"`
	Timeout		int		`json:"timeout"`
}

var commands Commands

// this code was not approved by C+ gnag