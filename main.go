package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"time"
	"github.com/robfig/cron/v3"
)

const ollamaURL = "http://ollama:11435/api/v1/completions"

func generateMessage() string {
	prompt := "Good morning! Please generate a positive message for the day."
	data := map[string]string{"prompt": prompt}
	jsonData, _ := json.Marshal(data)

	resp, err := http.Post(ollamaURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error contacting Ollama:", err)
		return ""
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var response map[string]interface{}
	json.Unmarshal(body, &response)

	if message, ok := response["completion"].(string); ok {
		return message
	}
	return "Have a great day!"
}

func commitFile(filename, message string) {
	cmd := exec.Command("git", "add", filename)
	if err := cmd.Run(); err != nil {
		fmt.Println("Error adding file to git:", err)
		return
	}

	commitMessage := fmt.Sprintf("✨ %s", message)
	cmd = exec.Command("git", "commit", "-m", commitMessage)
	if err := cmd.Run(); err != nil {
		fmt.Println("Error committing file:", err)
		return
	}
	fmt.Println("File committed with message:", commitMessage)
}

func main() {
	c := cron.New()
	_, err := c.AddFunc("0 8 * * *", func() {
		message := generateMessage()
		if message == "" {
			fmt.Println("Failed to generate message.")
			return
		}
		filename := fmt.Sprintf("message_%s.md", time.Now().Format("20060102T080000"))
		err := os.WriteFile(filename, []byte(message), 0644)
		if err != nil {
			fmt.Println("Error writing file:", err)
			return
		}
		fmt.Println("Generated message saved to", filename)
		commitFile(filename, message)
	})

	if err != nil {
		fmt.Println("Error scheduling cron job:", err)
		return
	}

	c.Start()
	select {}
}
