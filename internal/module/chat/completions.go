// internal/module/chat/completions.go

package chat

import (
	"OpenTan/config"
	"OpenTan/internal/global"
	"OpenTan/internal/global/model"
	"OpenTan/internal/global/response"
	"OpenTan/utils"
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Completions(oreq model.OpenAICompletionsRequest) func(c *gin.Context) {
	//if oreq.Model == "" {
	//	oreq.Model = config.Get().UseModel
	//}
	isStream := true
	if oreq.Stream != nil {
		isStream = *oreq.Stream
	}
	content := ""
	for _, message := range oreq.Messages {
		switch contentValue := message.Content.(type) {
		case string:
			content += contentValue + "\n"
		case []interface{}:
			for _, item := range contentValue {
				if textMap, ok := item.(map[string]interface{}); ok {
					if text, ok := textMap["text"].(string); ok {
						content += text + "\n"
					}
				}
			}
		default:
			fmt.Printf("Unknown content type: %T\n", message.Content)
		}
	}
	convBody := model.NewConvBody{
		Content: content,
		Stream:  isStream,
		Conversation: model.ConvMetadata{
			Title: "temp",
			Model: config.Get().UseModel,
		},
	}

	return func(c *gin.Context) {
		req, err := utils.NewTanPostRequest(model.TAN_URL+"api/v2/messages", utils.Object2Body(convBody))
		if err != nil {
			response.NewServerError(500, "Internal Server Error")(c)
			return
		}
		utils.AddHeader(req, "accept", "*/*")
		utils.AddHeader(req, "authorization", config.Get().API_KEY)
		utils.AddHeader(req, "content-type", "application/json")
		utils.AddHeader(req, "origin", model.TAN_URL)
		utils.AddHeader(req, "priority", "u=1, i")
		utils.AddHeader(req, "referer", model.TAN_URL+"/chat")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			response.NewServerError(500, "Internal Server Error")(c)
			return
		}
		defer resp.Body.Close()

		convID := ""

		if isStream {
			c.Writer.Header().Set("Content-Type", "text/event-stream")
			c.Writer.Header().Set("Cache-Control", "no-cache")
			c.Writer.Header().Set("Connection", "keep-alive")

			reader := bufio.NewReader(resp.Body)
			firstChunkHandled := false
			for {
				lineBytes, _, err := reader.ReadLine()
				if err != nil {
					if err == io.EOF {
						rmFlag := global.RemoveConv(convID)
						if !rmFlag {
							fmt.Printf("Failed to remove conversation: %v\n", convID)
						}
						return
					}
					fmt.Printf("Stream error: %v\n", err)
					return
				}

				jsonStart := bytes.IndexByte(lineBytes, '{')
				if jsonStart != -1 {
					prefix := lineBytes[:jsonStart]
					lineData := lineBytes[jsonStart:]

					var raw map[string]interface{}
					if errRaw := json.Unmarshal(lineData, &raw); errRaw == nil {
						choices, ok := raw["choices"].([]interface{})
						if ok && len(choices) > 0 {
							if choiceMap, ok := choices[0].(map[string]interface{}); ok {
								if deltaMap, ok := choiceMap["delta"].(map[string]interface{}); ok {
									if _, hasReasoning := deltaMap["reasoning_content"]; !hasReasoning {
										// Ensure reasoning_content key exists for downstream clients
										deltaMap["reasoning_content"] = ""
										if patched, errMarshal := json.Marshal(raw); errMarshal == nil {
											lineData = patched
											lineBytes = append(prefix, patched...)
										}
									}
								}
							}
						}
					}

					var r model.StreamResponse
					if err = json.Unmarshal(lineData, &r); err != nil {
						if convID == "" {
							fmt.Printf("Error unmarshaling JSON: %v, body: %s\n", err, string(lineData))
							response.NewServerError(500, "Failed to unmarshal response")(c)
							return
						}
					} else {
						if convID == "" {
							convID = r.ConversationID
						}
						if !firstChunkHandled {
							firstChunkHandled = true
							if len(r.Choices) > 0 && r.Choices[0].Delta.Content == "" {
								continue
							}
						}
					}
				}

				// Directly write the line to the response writer.
				if _, err := c.Writer.Write(append(lineBytes, '\n')); err != nil {
					fmt.Printf("Failed to write to stream: %v\n", err)
					return
				}
				c.Writer.Flush()
			}
		} else {
			respBody, err := io.ReadAll(resp.Body)
			if err != nil {
				response.NewServerError(500, "Internal Server Error")(c)
				return
			}

			var data model.NonStreamResponse
			err = json.Unmarshal(respBody, &data)
			if err != nil {
				fmt.Printf("Error unmarshaling JSON: %v, body: %s\n", err, string(respBody))
				response.NewServerError(500, "Failed to unmarshal response")(c)
				return
			}
			c.JSON(http.StatusOK, data)
			convID = data.Data.ConversationID
			log.Printf("Trying to remove conversation by ConversationID: %s\n", convID)
			rmFlag := global.RemoveConv(convID)
			if !rmFlag {
				fmt.Printf("Failed to remove conversation: %v\n", rmFlag)
			}
		}
	}
}
