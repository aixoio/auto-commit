package main

import (
	"context"
	"fmt"
	"os/exec"
	"strings"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

func main() {
  llm, _ := ollama.New(ollama.WithModel("llama3.2"))

	cmd := exec.Command("git", "status")

	output, _ := cmd.CombinedOutput()

  res, _ := llms.GenerateFromSinglePrompt(context.TODO(), llm, "It is your job to write a git commit summary for this commit YOU MUST ONLY RESPOND WITH THE MESSAGE NO MORE TEXT also please make detailed but short make sure to use the context that you have, the results of git status you CANNOT use 'feat' to prefix your comiit, don't talk about per proeject problems only what was changed YOU CANNOT TALK ABOUT THE PROJECT IT SELF AND YOU CAN ONLY TALK ABOUT WHAT WAS CHANGED your message should look like \"Deleted main.go\" but change it to match the changes made to the project: "+string(output))

	cmd = exec.Command("git", "add", ".")
	cmd.Run()
	cmd = exec.Command("git", "commit", "-m", strings.TrimSuffix(strings.TrimSpace(res), "\n"))
	cmd.Run()

	fmt.Println("Commited", res)
}
