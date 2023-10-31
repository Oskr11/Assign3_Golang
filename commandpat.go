package main

import (
	"fmt"
)

func main() {
	turnOnLightCom := &ConcreteCommand{
		receiver: &ConcreteReceiver{},
	}

	invoker := &Invoker{}

	invoker.AddCommand(turnOnLightCom)

	invoker.ExecuteCommands()
}

type Command interface {
	Execute()
}

type Receiver interface {
	DoSomething()
}

type ConcreteCommand struct {
	receiver Receiver
}

func (c *ConcreteCommand) Execute() {
	c.receiver.DoSomething()
}

type ConcreteReceiver struct{}

func (r *ConcreteReceiver) DoSomething() {
	fmt.Println("Turning on the light...")
}

type Invoker struct {
	commands []Command
}

func (i *Invoker) AddCommand(command Command) {
	i.commands = append(i.commands, command)
}

func (i *Invoker) ExecuteCommands() {
	for _, command := range i.commands {
		command.Execute()
	}
}
