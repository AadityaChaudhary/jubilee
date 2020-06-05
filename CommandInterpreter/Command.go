package CommandInterpreter

type Commands struct {
	SeleniumPath	string		`json:"selenium-path"`
	DriverPath		string		`json:"driver-path"`
	Port			int			`json:"port"`
	Browser 		string		`json:"browser"`
	Commands 		[]Command	`json:"commands"`

}

type Command struct {
	CommandType string 	`json:"command"`
	Link 		string 	`json:"link"`
	Id 			string 	`json:"id"`
	Comment 	string 	`json:"comment"`
	OutputPath	string	`json:"output"`
	Input 		string 	`json:"input"`
	Timeout		int		`json:"timeout"`
}

var commands Commands

