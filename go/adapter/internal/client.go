package plug

import "fmt"

type Client struct{}

type Computer interface {
	InsertIntoLightningPort()
}

func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}
