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

	onelinelogcmd := exec.Command("git", "log", "--oneline")

	onelinelogoutput, _ := onelinelogcmd.CombinedOutput()

	res, _ := llms.GenerateFromSinglePrompt(context.TODO(), llm, fmt.Sprintf("It is your job to write a git commit summary for this commit YOU MUST ONLY RESPOND WITH THE MESSAGE NO MORE TEXT also please make it very short make sure to use the context that you have, the results of git status: %s and here is the output of git log --oneline: %s rember to make your commit message short I perfer under 50 chars, YOU CANNOT INCLUDE THE GIT DIFF IN THE COMMIT MESSAGE YOU MUST ONLY RETURN THE COMMIT MESSAGE AND NOTHING ELSE", string(output), string(onelinelogoutput)))

	cmd = exec.Command("git", "add", ".")
	cmd.Run()
	cmd = exec.Command("git", "commit", "-m", res)
	cmd.Run()

	fmt.Println("Commited", res)
}
