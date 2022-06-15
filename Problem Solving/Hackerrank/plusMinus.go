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
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) {
    n:=len(arr)
    var arrPositive []int32
    var arrNegative []int32
    var arrZero []int32
    
    for i:=0; i<n; i++ {
        if arr[i] > 0 {
            arrPositive = append(arrPositive, arr[i])
        } else if arr[i] < 0 {
            arrNegative = append(arrNegative, arr[i])
        } else if arr[i] == 0 {
            arrZero = append(arrZero, arr[i])
        }
    }
    
    resPlus  := float64(len(arrPositive)) / float64(n) 
    resMinus := float64(len(arrNegative)) / float64(n)
    resZero  := float64(len(arrZero)) / float64(n)
    
    fmt.Printf("%f\n",resPlus) 
    fmt.Printf("%f\n",resMinus) 
    fmt.Printf("%f\n",resZero) 
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var arr []int32

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    plusMinus(arr)
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
