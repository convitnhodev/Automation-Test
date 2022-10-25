package nodeModel

import (
	"errors"
	"strings"

)

type Node struct {
	SerialNumber string `bson:"serial_number,omitempty"`
	NodeName     string `bson:"node_name,omitempty"`
	UserName     string `bson:"user_name,omitempty"`
}

func (node *Node) Validate() error {

	// check validate of node name

	node.NodeName = strings.TrimSpace(node.NodeName)

	if node.NodeName == "" {
		return errors.New("username name can not be blank")
	}

	return nil
}
