package main

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

func main() {
	llm, _ := ollama.New(ollama.WithModel("llama3"))

	cmd := exec.Command("git", "status")

	output, _ := cmd.CombinedOutput()

	res, _ := llms.GenerateFromSinglePrompt(context.TODO(), llm, "It is your job to write a git commit summary for this commit YOU MUST ONLY RESPOND WITH THE MESSAGE NO MORE TEXT also please make very detailed message but try to avoid talking about the commit in the message like don't say 'for this staged commit' use the context of the message here is the results of git status: "+string(output))

	cmd = exec.Command("git", "add", ".")
	cmd.Run()
	cmd = exec.Command("git", "commit", "-m", res)
	cmd.Run()

	fmt.Println("Commited", res)
}
