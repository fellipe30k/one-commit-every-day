package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/robfig/cron/v3"
)

const ollamaURL = "http://ollama:11434/api/generate"

// generateMessage communicates with Ollama and parses JSON lines (streaming response).
func generateMessage() string {
	payload := map[string]interface{}{
		"model":  "llama3.2:1b",
		"prompt": "I'm going to write a simple readme.md that how write a simple hello word in python:",
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshaling payload:", err)
		return ""
	}

	resp, err := http.Post(ollamaURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error contacting Ollama:", err)
		return ""
	}
	defer resp.Body.Close()

	// Process the streaming response
	var fullMessage string
	reader := bufio.NewReader(resp.Body)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading response:", err)
			return ""
		}

		// Parse each JSON line
		var response map[string]interface{}
		if err := json.Unmarshal(line, &response); err != nil {
			fmt.Println("Error unmarshaling line:", string(line), err)
			continue
		}

		// Append the "response" field to the full message
		if part, ok := response["response"].(string); ok {
			fullMessage += part
		}

		// Stop if "done" is true
		if done, ok := response["done"].(bool); ok && done {
			break
		}
	}

	return fullMessage
}

// commitFile adds and commits a file to the Git repository.
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

// generateAndCommitMessage generates a message, saves it to a file, and commits it.
func generateAndCommitMessage() {
	message := generateMessage()
	if message == "" {
		fmt.Println("Failed to generate message.")
		return
	}

	filename := fmt.Sprintf("message_%s.md", time.Now().Format("20060102T150405"))
	err := os.WriteFile(filename, []byte(message), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	fmt.Printf("Generated message saved to %s\n", filename)
	commitFile(filename, message)
}

// startCron initializes the cron scheduler.
func startCron() {
	c := cron.New()
	_, err := c.AddFunc("0 8 * * *", generateAndCommitMessage)
	if err != nil {
		fmt.Println("Error scheduling cron job:", err)
		return
	}

	fmt.Println("Cron job initialized.")
	c.Start()
	select {}
}

// testApp verifies that the system is creating files, committing, and communicating with Ollama.
func testApp() {
	fmt.Println("Starting test...")

	message := generateMessage()
	if message == "" {
		fmt.Println("Test failed: unable to generate message.")
		return
	}
	fmt.Println("Message generated:", message)

	filename := fmt.Sprintf("test_message_%s.md", time.Now().Format("20060102T150405"))
	err := os.WriteFile(filename, []byte(message), 0644)
	if err != nil {
		fmt.Println("Test failed: unable to write file.")
		return
	}
	fmt.Printf("Test file created: %s\n", filename)

	commitFile(filename, message)
	fmt.Println("Test completed successfully.")
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "test" {
		testApp()
		return
	}

	fmt.Println("System initialized.")
	startCron()
}
