

    package main
     
    import (
    	"encoding/json"
    	"fmt"
    	"io"
    	"net/http"
    	"os"
    	"strings"
    )
     
    const (
    	apiKey      = "add-your-own-api-key"
    	apiEndpoint = "https://api.openai.com/v1/engines/davinci-codex/completions"
    )
     
    type GPT3Request struct {
    	Prompt      string  `json:"prompt"`
    	MaxTokens   int     `json:"max_tokens"`
    	N           int     `json:"n"`
    	Stop        string  `json:"stop"`
    	Temperature float64 `json:"temperature"`
    }
     
    type GPT3Response struct {
    	Choices []struct {
    		Text string `json:"text"`
    	} `json:"choices"`
    }
     
    func main() {
    	if len(os.Args) != 2 {
    		fmt.Println("Usage: chatgpt \"question\"")
    		os.Exit(1)
    	}
     
    	question := os.Args[1]
    	prompt := fmt.Sprintf("I'm a chatbot trained by OpenAI, and I'm here to answer your questions.\n\nUser: %s\n\nAI:", question)
     
    	reqBody, _ := json.Marshal(GPT3Request{
    		Prompt:      prompt,
    		MaxTokens:   100,
    		N:           1,
    		Stop:        "\n",
    		Temperature: 0.5,
    	})
     
    	client := &http.Client{}
    	req, _ := http.NewRequest("POST", apiEndpoint, strings.NewReader(string(reqBody)))
    	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
    	req.Header.Set("Content-Type", "application/json")
     
    	resp, err := client.Do(req)
    	if err != nil {
    		fmt.Println("Error making request:", err)
    		os.Exit(1)
    	}
     
    	defer resp.Body.Close()
    	respBody, _ := io.ReadAll(resp.Body)
     
    	var gpt3Resp GPT3Response
    	err = json.Unmarshal(respBody, &gpt3Resp)
    	if err != nil {
    		return 
    	}
     
    	if len(gpt3Resp.Choices) > 0 {
    		answer := strings.TrimSpace(gpt3Resp.Choices[0].Text)
    		fmt.Println(answer)
    	} else {
    		fmt.Println("No answer found")
    	}
    }

