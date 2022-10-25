package commandModel

import "time"

type Node struct {
	NodeId string `json:"node_id" bson:"node_id"`
}

type Command struct {
	Command string `json:"command" bson:"command"`
}

type CommandNode struct {
	Nodes   []Node    `json:"node_id_list" bson:"node_id_list"`
	Command string    `json:"command"bson:"command,omitempty"`
	Date    time.Time `bson:"sampleDate" json:"sampleDate, omitempty"`
}

type EachNodeCommand struct {
	NodeId     string `json:"node_id" bson:"node_id"`
	Command    string `json:"command"bson:"command,omitempty"`
	CountStick int    `json:"count_stick" bson:"count_stick,omitempty"`
}
