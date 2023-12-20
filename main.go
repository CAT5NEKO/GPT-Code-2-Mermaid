package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

const (
	openaiApiKey = "hogehogefugafuga"
)

func main() {
	code := readCodeFromFile("inputHoge.go")
	extractedProcess := extractProcess(code)
	mermaidDiagram := convertToMermaid(extractedProcess)
	constructFlow(mermaidDiagram)
	generateMarkdownFile(mermaidDiagram, "output.md")
}

func readCodeFromFile(filename string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func extractProcess(code string) string {
	return code
}

func convertToMermaid(process string) string {
	client := openai.NewClient(openaiApiKey)
	ctx := context.Background()
	resp, err := client.CreateCompletion(
		ctx,
		openai.CompletionRequest{
			Model:     openai.GPT3Dot5Turbo,
			MaxTokens: 20,
			Prompt:    process,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	return resp.Choices[0].Text
}

func constructFlow(mermaidDiagram string) {
	cmd := exec.Command("mmdc", "--input", "-", "--output", "output.png")
	cmd.Stdin = strings.NewReader(mermaidDiagram)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("フローダイアグラムが正常に作成されました。")
}

func generateMarkdownFile(mermaidDiagram, filename string) {
	mdContent := fmt.Sprintf("# Process Flow\n\n```mermaid\n%s\n```\n", mermaidDiagram)

	err := ioutil.WriteFile(filename, []byte(mdContent), 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Mdファイル '%s' が正常に生成されました。\n", filename)
}
