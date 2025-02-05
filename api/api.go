package api

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func GetAPIKey() string {
	// SECURITY WARNING: Never hardcode your API key!
	// Use environment variables or a secure configuration file.
	return "" // Replace with your actual key
}

func SaveOutputToMarkdown(output string) error {
	// Saves the output to a markdown file.
	file, err := os.Create("ai_output.md")
	if err != nil {
		return fmt.Errorf("Error creating markdown file: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(output)
	if err != nil {
		return fmt.Errorf("Error writing to markdown file: %v", err)
	}
	return nil
}

func AnalyzePorts(openPorts string, language string) {
	// Analyzes open ports using the Gemini API.
	ctx := context.Background()
	apiKey := GetAPIKey()

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-2.0-flash-exp") // or "gemini-pro"
	model.SetTemperature(1)
	model.SetTopK(40)
	model.SetTopP(0.95)
	model.SetMaxOutputTokens(8192)

	// Get current date
	currentTime := time.Now()
	reportDate := currentTime.Format("02 January 2006")

	systemPrompt := fmt.Sprintf(`You are a cybersecurity expert. Analyze the provided list of open network ports and create a detailed security assessment report.
Report Date: %s
It is very important to remember the information on previously analysed ports, if any.

Your analysis should include:
    - Identification of each open port and its standard service
    - Security implications and potential vulnerabilities
    - Risk level (Low/Medium/High) for each open port
    - Specific recommendations to secure each port
    - Overall system security status

Please be thorough and technical in your analysis while maintaining clarity for both technical and non-technical stakeholders. Please respond in this language: %s`, reportDate, language)

	// Add previous analysis information (if any) to the prompt.
	// This is a simplified example.  In a real application, you would
	// store and retrieve the previous analysis.
	previousAnalysis := "" // In a real application, load this from storage
	if previousAnalysis != "" {
		systemPrompt += "\n\nPrevious Analysis:\n" + previousAnalysis
	}

	prompt := fmt.Sprintf("%s\n\nOpen Ports:\n%s", systemPrompt, openPorts)

	// **RISKY: Lowering SafetySettings**
	// These settings lower the sensitivity thresholds for certain categories.
	// However, it's not possible to completely disable the filters.
	// Use with caution!

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Fatalf("Error generating content: %v", err)
	}

	// Combine the output
	output := ""
	if resp != nil && resp.Candidates != nil && len(resp.Candidates) > 0 {
		for _, part := range resp.Candidates[0].Content.Parts {
			if text, ok := part.(genai.Text); ok {
				output += string(text)
			}
		}
	} else if resp != nil && len(resp.PromptFeedback.SafetyRatings) > 0 {
		log.Printf("Content blocked due to safety filter. Feedback: %+v", resp.PromptFeedback)
		output = "Content blocked due to safety filter."
	} else {
		log.Println("Unexpected error or no response received.")
		output = "Unexpected error or no response received."
	}

	// Save the output to a markdown file
	err = SaveOutputToMarkdown(output)
	if err != nil {
		log.Fatalf("Error saving output to markdown: %v", err)
	}

	fmt.Println("Gemini Analysis saved to ai_output.md")
}
