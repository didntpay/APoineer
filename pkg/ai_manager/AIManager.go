package aimanager

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

/*
	AIManager is a struct that contains all the models that are available to the user.
	functions it should have:
		- CreateConnectionWithModel
		- open
*/

type AIManager interface {
	createConnectionWithModel(model string) error
	sendMessage(msg string) error
}

type AIManagerImpl struct {
	model string
}

type GPTManagerImpl struct {
	AIManagerImpl
	client *openai.Client
}

func (gpt *GPTManagerImpl) createConnectionWithModel(apikey string) error {
	gpt.client = openai.NewClient(apikey)
	return nil
}

func (gpt *GPTManagerImpl) sendMessage(msg string) (openai.ChatCompletionResponse, error) {
	resp, err := gpt.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: gpt.model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: msg,
				},
			},
		},
	)

	return resp, err

}

// client := openai.NewClient(os.Getenv("OPENAI_KEY"))
// 	fmt.Println(os.Getenv("OPENAI_KEY"))
// 	resp, err := client.CreateChatCompletion(
// 		context.Background(),
// 		openai.ChatCompletionRequest{
// 			Model: openai.GPT3Dot5Turbo,
// 			Messages: []openai.ChatCompletionMessage{
// 				{
// 					Role:    openai.ChatMessageRoleUser,
// 					Content: "Hello!",
// 				},
// 			},
// 		},
// 	)

// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(resp)
// 	}
