package handler

import (
	"context"
	"fmt"

	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
	"google.golang.org/protobuf/proto"
)

func EventHandler(client *whatsmeow.Client) func(evt interface{}) {
	return func(evt interface{}) {
		switch v := evt.(type) {
		case *events.Message:
			if !v.Info.IsFromMe && v.Info.PushName == "Matíasdd" && v.Message.GetConversation() == "!help" {
				protoMsg := &waProto.Message{
					Conversation: proto.String("funcionando"),
				}

				test := &types.MessageInfo{
					MessageSource: types.MessageSource{
						Chat: v.Info.Sender,
					},
				}

				client.SendMessage(context.Background(), v.Info.Sender, protoMsg)

				old := client.BuildHistorySyncRequest(test, 6)
				prueba := whatsmeow.SendRequestExtra{Peer: true}
				test3, err := client.SendMessage(context.Background(), v.Info.Sender, old, prueba)

				if err != nil {
					panic(err)
				}

				fmt.Println(test3)

			}

		}

	}
}
