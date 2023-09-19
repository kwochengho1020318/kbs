package api

import (
	"context"
	"fmt"
	"main/config"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	openai "github.com/sashabaranov/go-openai"
)

func Chat(w http.ResponseWriter, r *http.Request) {
	prompt := r.URL.Query().Get("prompt")
	history := r.URL.Query().Get("history")
	phase := mux.Vars(r)["phase"]
	if phase == "codegen" {
		response := codegen(prompt, history)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	} else if phase == "requirements" {
		response := meetrequirements(prompt, history)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	}
	// else if phase == "interface" {
	// 	response := interfacejson(prompt)
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.WriteHeader(http.StatusOK)
	// 	w.Write([]byte(response))
	// }
}

func OpenAIClient() *openai.Client {
	myconfig := config.NewConfig("appsettings.json")
	client := openai.NewClient(myconfig.OpenAI.APIKey)
	return client
}

func codegen(prompt string, history string) string {
	client := OpenAIClient()
	content, _ := os.ReadFile("files/Rules/logic.txt")
	importlogic := openai.ChatCompletionMessage{Role: openai.ChatMessageRoleSystem, Content: string(content)}

	logic := openai.ChatCompletionMessage{Role: openai.ChatMessageRoleSystem, Content: "This is a set of rules and example formats for JSON. For subsequent responses, please provide answers in the same format as I have given you and meet my custom requirements. Please adhere strictly to the format and rules or instruction."}
	hint := openai.ChatCompletionMessage{Role: openai.ChatMessageRoleSystem, Content: "If I directly ask you for an example, please first consider what tables and functional interfaces this system needs. Then, generate JSON based on the tables and functional interfaces you come up with."}
	ban_question := openai.ChatCompletionMessage{Role: openai.ChatMessageRoleSystem, Content: "Must:don't include any explanations and introduction in your responses "}
	rule := openai.ChatCompletionMessage{Role: openai.ChatMessageRoleSystem, Content: "this is history" + history}
	a := openai.ChatCompletionMessage{Role: openai.ChatMessageRoleUser, Content: prompt}

	conversation := []openai.ChatCompletionMessage{}
	conversation = append(conversation, importlogic, logic, hint, ban_question, rule, a)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo16K,
			Messages: conversation,
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return ""
	}
	return resp.Choices[0].Message.Content
}

func meetrequirements(prompt string, history string) string {
	client := OpenAIClient()
	content, _ := os.ReadFile("files/Rules/requirement.txt")
	hint := openai.ChatCompletionMessage{Role: openai.ChatMessageRoleSystem, Content: string(content)}
	a := openai.ChatCompletionMessage{Role: openai.ChatMessageRoleUser, Content: prompt}
	conversation := []openai.ChatCompletionMessage{}
	conversation = append(conversation, hint, a)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo16K,
			Messages: conversation,
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return ""
	}
	return resp.Choices[0].Message.Content
}
