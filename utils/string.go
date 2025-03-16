// utils/string.go

package utils

import "strings"

// TrimStreamLine 去除行首的 "data: " 和行尾的换行符
func TrimStreamLine(line string) string {
    trimmed := line
    prefix := "data: "
    if len(line) > len(prefix) && line[:len(prefix)] == prefix {
        trimmed = line[len(prefix):]
    }

    // 去除末尾的换行符
    trimmed = strings.TrimSuffix(trimmed, "\n")
    return trimmed
}
