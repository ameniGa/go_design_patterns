package command

import "fmt"

type Appliance interface {
	Open()
	Close()
}

// Window represents the receiver of Command
type Window struct {
}

func (w *Window) Open() {
	fmt.Println("Window is opened")
}

func (w *Window) Close() {
	fmt.Println("Window is closed")
}

/*****************************************************************/

// Command is the contract to implement by any type of Command
type Command interface {
	Execute()
	Unexecute()
}

/*****************************************************************/

// closeCmd is a concrete Command
// implements the Command interface
type closeCmd struct {
	receiver Appliance
}

func (c *closeCmd) Execute() {
	c.receiver.Close()
}

func (c *closeCmd) Unexecute() {
	c.receiver.Open()
}

/*****************************************************************/

// openCmd is a concrete Command
// implements the Command interface
type openCmd struct {
	receiver Appliance
}

func (o *openCmd) Execute() {
	o.receiver.Open()
}

func (o *openCmd) Unexecute() {
	o.receiver.Close()
}

/*****************************************************************/

// RemoteControl represents the invoker
type RemoteControl struct {
	openCmd  Command
	closeCmd Command
}

func NewRemoteControl(openCmd, closeCmd Command) RemoteControl {
	return RemoteControl{
		openCmd:  openCmd,
		closeCmd: closeCmd,
	}
}

func (u RemoteControl) Open() {
	u.openCmd.Execute()
}

func (u RemoteControl) Close() {
	u.closeCmd.Execute()
}
