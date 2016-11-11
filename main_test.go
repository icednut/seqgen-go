package main

import(
    "io/ioutil"
    "net/http"
    "log"
    "testing"
    "strconv"
)

func callApi(apiUrl string) int {
    res, err := http.Get(apiUrl)
    if err != nil {
        log.Fatal(err)
    }
    seqNo, err := ioutil.ReadAll(res.Body)
    res.Body.Close()
    if err != nil {
        log.Fatal(err)
    }
    seq, _ := strconv.Atoi(string(seqNo))
    return seq
}

func removeDuplicates(elements []int) []int {
    encountered := map[int]bool{}
    result := []int{}

    for v := range elements {
        if encountered[elements[v]] == true {
        } else {
            encountered[elements[v]] = true
            result = append(result, elements[v])
        }
    }
    return result
}

func TestRunServer(t *testing.T) {
    go func() {
        runServer()
    }()

    arrayCount := 10
    seqNoArray := make([]int, 0)
    for i := 0; i < arrayCount; i++ {
        //go func() {
            seqNoArray = append(seqNoArray, callApi("http://localhost:8080"))
        //}()
    }
    uniqueSeqNoArray := removeDuplicates(seqNoArray)
    uniqueSeqNoArrayLen := len(uniqueSeqNoArray)
    seqNoArrayLen := len(seqNoArray)

    if uniqueSeqNoArrayLen != seqNoArrayLen || uniqueSeqNoArrayLen != arrayCount {
        t.Error("SeqNo must be unique. (uniqueSeqNoArray: ", uniqueSeqNoArray, ", seqNoArrayLen: ", seqNoArray, ")")
    }
}
