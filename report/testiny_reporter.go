package report

import (
    "fmt"
    "time"
    // "net/http"
    // "encoding/json"
)

type TestinyReporter struct {
    apiKey     string
    runID      int
    titleToID  map[string]int
}

func NewTestinyReporter(apiKey string) (*TestinyReporter, error) {
    cases, err := getTestCases(apiKey)
    if err != nil {
        return nil, err
    }
    fmt.Printf("Found %d test cases in Testiny\n", len(cases))
    for _, c := range cases {
        fmt.Printf("Case: %s, ID: %d\n", c.Title, c.ID)
    }
    mapping := BuildTitleToIDMap(cases)
    return &TestinyReporter{
        apiKey:    apiKey,
        titleToID: mapping,
    }, nil
}
