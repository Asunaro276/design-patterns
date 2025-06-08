package main

import (
	"github.com/asunaro276/design-patterns/go/adapter/internal"
)

func main() {
	client := plug.Client{}
	mac := plug.Mac{}

	client.InsertLightningConnectorIntoComputer(&mac)

	windowsMachine := plug.Windows{}
	windowsMachineAdapter := plug.WindowsAdapter{
		WindowsMachine: &windowsMachine,
	}
	client.InsertLightningConnectorIntoComputer(&windowsMachineAdapter)
}
