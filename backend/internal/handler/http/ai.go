package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"

	"qasr/backend/internal/utils"
)

type SlugSuggestRequest struct {
	URL string `json:"url"`
}

type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}

type OllamaResponse struct {
	Response string `json:"response"`
}

func (h *Handler) SuggestSlug(c *gin.Context) {
	var req SlugSuggestRequest
	if err := c.ShouldBindJSON(&req); err != nil || req.URL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	content := utils.FetchPageSummary(req.URL)

	fmt.Println("Fetched content:", content)
	// Avoid too long prompt input
	if len(content) > 800 {
		content = content[:800]
	}

	prompt := `
You are a backend service. 
Your job is to return only a single slug based on the following content. 

Rules:
- Return exactly one line
- No explanation
- No greeting
- No quotes
- No bullet points

Format example:
learn-english-online

Content:
` + content

	body, _ := json.Marshal(OllamaRequest{
		Model:  "llama2", // or llama3, mistral, etc.
		Prompt: prompt,
	})

	resp, err := http.Post("http://localhost:11434/api/generate", "application/json", bytes.NewReader(body))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ollama not reachable"})
		return
	}
	defer resp.Body.Close()

	var full string
	decoder := json.NewDecoder(resp.Body)

	for {
		var chunk OllamaResponse
		if err := decoder.Decode(&chunk); err == io.EOF {
			break
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse Ollama stream"})
			return
		}
		full += chunk.Response
	}

	fmt.Println("AI Response:", full)

	slug := utils.CleanSlug(full)

	c.JSON(http.StatusOK, gin.H{
		"slug": slug,
	})
}

func (h *Handler) DescribeSlug(c *gin.Context) {
	slug := c.Param("slug")
	link, err := h.Service.GetBySlug(slug)
	if err != nil || link == "" {
		c.JSON(404, gin.H{"error": "Slug not found"})
		return
	}

	prompt := "Describe this link in 1-2 sentences: " + link

	body, _ := json.Marshal(OllamaRequest{
		Model:  "llama3",
		Prompt: prompt,
	})

	resp, err := http.Post("http://localhost:11434/api/generate", "application/json", bytes.NewReader(body))
	if err != nil {
		c.JSON(500, gin.H{"error": "Ollama not reachable"})
		return
	}
	defer resp.Body.Close()

	raw, _ := io.ReadAll(resp.Body)
	var result OllamaResponse
	_ = json.Unmarshal(raw, &result)

	c.JSON(200, gin.H{
		"description": result.Response,
	})
}
