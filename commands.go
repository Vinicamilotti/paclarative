package main

type Commands struct {
	command *CommandParser
	config  *ConfigFiles
	cmdList map[string]Command
}

func CreateCommands(cmd *CommandParser, config *ConfigFiles) *Commands {
	var cmdList = make(map[string]Command)

	cmdList["sync"] = Sync
	cmdList["ensure"] = Ensure
	cmdList["remove"] = Remove
	cmdList["list"] = List

	var c = Commands{
		command: cmd,
		config:  config,
		cmdList: cmdList,
	}

	return &c
}

func (cmd *Commands) ExecCommand() {
	cmd.cmdList[cmd.command.command](cmd.command.args, cmd.config)
}

func Sync(args []string, conf *ConfigFiles)   {}
func Ensure(args []string, conf *ConfigFiles) {}
func Remove(args []string, conf *ConfigFiles) {}
func List(args []string, conf *ConfigFiles)   {}
