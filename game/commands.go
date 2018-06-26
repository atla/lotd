package game

import (
	"log"
	"strings"
)

// CommandProcessor ... global user struct to control logins
type CommandProcessor struct {
	commands map[string]Command
}

// RegisterCommand ... register
func (commandProcessor *CommandProcessor) RegisterCommand(key string, command Command) {
	commandProcessor.commands[key] = command
}

// Process ...asd
func (commandProcessor *CommandProcessor) Process(game *Game, message *Message) bool {

	parts := strings.Fields(message.Data)

	if len(parts) > 0 {
		var key = parts[0]
		if val, ok := commandProcessor.commands[key]; ok {

			log.Println("Found command " + key + " executing...")
			return val.Execute(game, message)
		}
	}

	return false

}

// ScreamCommand ... foo
type ScreamCommand struct {
}

// Execute ... executes scream command
func (screamCommand *ScreamCommand) Execute(game *Game, message *Message) bool {

	parts := strings.Fields(message.Data)
	newMsg := strings.Join(parts[1:len(parts)], " ")

	var newMessage = "-- " + message.FromUser.ID + " screams " + strings.ToUpper(newMsg) + "!!!!!"
	game.OnMessageReceived <- NewMessage(nil, newMessage)
	return true

}

func (commandProcessor *CommandProcessor) registerCommands() {

	commandProcessor.RegisterCommand("scream", &ScreamCommand{})

}

// NewCommandProcessor .. creates a new command processor
func NewCommandProcessor() *CommandProcessor {
	var commandProcessor = &CommandProcessor{
		commands: make(map[string]Command),
	}
	// only once?
	commandProcessor.registerCommands()
	return commandProcessor
}

// Command ... commands
type Command interface {
	Execute(game *Game, message *Message) bool
}
