package main

type CommandParser struct {
	command string
	args    []string
}

func (cm *CommandParser) Parse(cmd []string) {
	cm.command = cmd[0]
	cm.args = cmd[:1]
}

func CreateCommandParser(args []string) CommandParser {
	if len(args) < 2 {
		panic("Command not found")
	}


	var cmd = CommandParser{}
	cmd.Parse(args)

	return cmd
}
