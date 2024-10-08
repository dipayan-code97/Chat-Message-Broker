package components

import (
	"fmt"

	"github.com/ASV44/chat-message-broker/client/models/receiver"
	"github.com/ASV44/chat-message-broker/client/models/sender"
	"github.com/ASV44/chat-message-broker/common"
)

/* CommunicationService represents abstraction of for performing message communication with broker */
type CommunicationService interface {
	sendMessage(sender.Message)
	getMessage() (receiver.Message, error)
}

/* CommunicationManager represents service for performing message communication with broker */
type CommunicationManager struct {
	connection common.Connection
}

/* NewCommunicationManager creates new instance of CommunicationManager */
func NewCommunicationManager(connection common.Connection) CommunicationManager {
	return CommunicationManager{connection: connection}
}

/* sendMessage send message to broker and handle error in case when sending of message fails */
func (managerRef *CommunicationManager) sendMessage(message sender.Message) {
	err := managerRef.connection.SendMessage(message)
	if err != nil {
		fmt.Printf(
			"Failed to send message.\nType:%s\nTarget: %s\nTime: %s\nErr: %s\n",
			message.Type,
			message.Target,
			message.Time,
			err,
		)
		// Handle error of sending message here.
		// Here is possible to implement various scenarios of handling communication error.
		// Resending of message, saving in DB, adding in queue of not delivered messages.
	}
}

/* GetMessage get from broker connection and handle error */
func (managerRef *CommunicationManager) getMessage() (receiver.Message, error) {
	var messageReceiver receiver.Message
	err := managerRef.connection.GetMessage(&messageReceiver)
	return messageReceiver, err
}
