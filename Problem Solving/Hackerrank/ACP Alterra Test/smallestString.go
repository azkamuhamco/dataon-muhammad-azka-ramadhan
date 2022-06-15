package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)



/*
 * Complete the 'smallestString' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING_ARRAY substrings as parameter.
 */

func smallestString(substrings []string) string {
    var hasil string
    
    for i:=0; i<len(substrings); i++ {
        for j:=i+1; j<len(substrings); j++ {
            if substrings[i] + substrings[j] > substrings[j] + substrings[i] {
                s := substrings[i]
                substrings[i] = substrings[j]
                substrings[j] = s
            }
        }
        hasil += substrings[i]
    }
    
    return hasil
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    substringsCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    var substrings []string

    for i := 0; i < int(substringsCount); i++ {
        substringsItem := readLine(reader)
        substrings = append(substrings, substringsItem)
    }

    result := smallestString(substrings)

    fmt.Fprintf(writer, "%s\n", result)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
