package benchmark

import (
    "bytes"
    "strings"
)

func ConcatenateBuffer(first string, second string) string {
    var buffer bytes.Buffer
    buffer.WriteString(first)
    buffer.WriteString(second)
    return buffer.String()
}

func ConcatenateJoin(first string, second string) string {
    return strings.Join([]string{first, second}, "")
}

func DoSomeThing() {
    a := make(map[string]int, 10000)
    a["1"] = 1;
}

