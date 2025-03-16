// internal/module/chat/completions.go

package chat

import (
    "OpenTan/config"
    "OpenTan/internal/global/model"
    "OpenTan/internal/global/response"
    "OpenTan/utils"
    "bufio"
    "encoding/json"
    "fmt"
    "github.com/gin-gonic/gin"
    "io"
    "net/http"
)

func Completions(oreq model.OpenAICompletionsRequest) func(c *gin.Context) {
    if oreq.Model == "" {
        oreq.Model = model.DefaultModel
    }
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
            Model: oreq.Model,
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

        if isStream {
            c.Writer.Header().Set("Content-Type", "text/event-stream")
            c.Writer.Header().Set("Cache-Control", "no-cache")
            c.Writer.Header().Set("Connection", "keep-alive")

            reader := bufio.NewReader(resp.Body)
            for {
                lineBytes, _, err := reader.ReadLine()
                if err != nil {
                    if err == io.EOF {
                        return
                    }
                    fmt.Printf("Stream error: %v\n", err)
                    return
                }

                // Directly write the line to the response writer.
                if _, err := c.Writer.Write(append(lineBytes, '\n')); err != nil {
                    fmt.Printf("Failed to write to stream: %v\n", err)
                    return
                }
                c.Writer.Flush()
            }
            //c.Stream(func(w io.Writer) bool {
            //    reader := bufio.NewReader(resp.Body)
            //    for {
            //        lineBytes, _, err := reader.ReadLine()
            //        line := string(lineBytes)
            //        if err != nil {
            //            if err == io.EOF {
            //                return false
            //            }
            //            fmt.Printf("Stream error: %v\n", err)
            //            return false
            //        }
            //        line = utils.TrimStreamLine(line)
            //        if line == "" {
            //            continue
            //        }
            //        if line == "[DONE]" {
            //            fmt.Println("Stream finished ([DONE])")
            //            return false
            //        }
            //        var data model.StreamResponse
            //        err = json.Unmarshal([]byte(line), &data)
            //        if err != nil {
            //            fmt.Printf("Error unmarshaling JSON: %v, line: %s\n", err, line)
            //            return false
            //        }
            //        fmt.Print("Response: ")
            //        for _, choice := range data.Choices {
            //            if _, err := w.Write([]byte(choice.Delta.Content)); err != nil {
            //                fmt.Printf("Failed to write to stream: %v\n", err)
            //                return false
            //            }
            //            fmt.Print(choice.Delta.Content)
            //        }
            //        fmt.Println()
            //        c.Writer.Flush()
            //        return true
            //    }
            //})
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
        }
    }
}
