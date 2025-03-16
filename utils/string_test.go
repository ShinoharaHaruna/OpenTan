// utils/string_test.go

package utils

import (
    "fmt"
    "testing"
)

func TestTrimStreamLine(t *testing.T) {
    type args struct {
        line string
    }
    tests := []struct {
        name string
        args args
        want string
    }{
        {
            name: "TestTrimStreamLine",
            args: args{
                line: "data: {\"id\":\"XXX\",\"parent_id\":\"XXX\",\"conversation_id\":\"XXX\",\"question_id\":\"XXX\",\"model\":\"gemini-2.0-flash\",\"choices\":[{\"delta\":{\"content\":\"XXX\",\"role\":\"assistant\"},\"logprobs\":null,\"finish_reason\":null,\"index\":0}],\"created_time\":\"YYYY-MM-DDTHH:MM:SS.MSSZ\"}\n",
            },
            want: "{\"id\":\"XXX\",\"parent_id\":\"XXX\",\"conversation_id\":\"XXX\",\"question_id\":\"XXX\",\"model\":\"gemini-2.0-flash\",\"choices\":[{\"delta\":{\"content\":\"XXX\",\"role\":\"assistant\"},\"logprobs\":null,\"finish_reason\":null,\"index\":0}],\"created_time\":\"YYYY-MM-DDTHH:MM:SS.MSSZ\"}",
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := TrimStreamLine(tt.args.line)
            fmt.Println(got)
            if got != tt.want {
                t.Errorf("TrimStreamLine() = %v, want %v", got, tt.want)
            }
        })
    }
}
