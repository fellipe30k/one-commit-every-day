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
	"path/filepath"
	"time"

	"github.com/robfig/cron/v3"
)

const (
	ollamaURL     = "http://ollama:11434/api/generate"
	outputDir     = "generated_files"
	promptOllama  = "I'm going to write a simple readme.md that explains how to write a simple read CSV in Python:"
	cronSchedule  = "0 10 * * 1-5" // At 10:00 AM, Monday to Friday
	fileExtension = "md"
)

// generateMessage communicates with Ollama and parses JSON lines (streaming response).
func generateMessage() string {
	payload := map[string]interface{}{
		"model":  "llama3.2:1b",
		"prompt": promptOllama,
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

		var response map[string]interface{}
		if err := json.Unmarshal(line, &response); err != nil {
			fmt.Println("Error unmarshaling line:", string(line), err)
			continue
		}

		if part, ok := response["response"].(string); ok {
			fullMessage += part
		}

		if done, ok := response["done"].(bool); ok && done {
			break
		}
	}

	return fullMessage
}

// commitFile adds and commits a file to the Git repository.
func commitFile(filename, message string) {
	cmd := exec.Command("git", "add", "./generated_files/", filename)
	if err := cmd.Run(); err != nil {
		fmt.Println("Error adding file to git:", err)
		return
	}

	shortMessage := message
	if len(message) > 50 {
		shortMessage = message[:150] + "..."
	}

	commitMessage := fmt.Sprintf("✨ %s", shortMessage)
	cmd = exec.Command("git", "commit", "-m", commitMessage)
	if err := cmd.Run(); err != nil {
		fmt.Println("Error committing file:", err)
		return
	}
	fmt.Println("File committed with message:", commitMessage)
	
	cmd = exec.Command("git", "push")
	if err := cmd.Run(); err != nil {
		fmt.Println("Error pushing file:", err)
		return
	}
	fmt.Println("File pushed:")
}

// generateAndCommitMessage generates a message, saves it to a file, and commits it.
func generateAndCommitMessage() {
	message := generateMessage()
	if message == "" {
		fmt.Println("Failed to generate message.")
		return
	}

	// Create output directory if it doesn't exist
	err := os.MkdirAll(outputDir, 0755)
	if err != nil {
		fmt.Println("Error creating output directory:", err)
		return
	}

	filename := filepath.Join(outputDir, fmt.Sprintf("message_%s.%s", time.Now().Format("20060102T150405"),fileExtension))
	err = os.WriteFile(filename, []byte(message), 0644)
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
	_, err := c.AddFunc(cronSchedule, generateAndCommitMessage)
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

	// Create output directory if it doesn't exist
	err := os.MkdirAll(outputDir, 0755)
	if err != nil {
		fmt.Println("Test failed: unable to create output directory.")
		return
	}

	filename := filepath.Join(outputDir, fmt.Sprintf("test_message_%s.%s", time.Now().Format("20060102T150405"), fileExtension))
	err = os.WriteFile(filename, []byte(message), 0644)
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
