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

func (r *TestinyReporter) StartRun() error {
    runID, err := createTestRun(
        r.apiKey,                           // API key
        1,                                  // project ID (1 is the default project)
        fmt.Sprintf("h1spec auto run %s", time.Now().Format("2006-01-02 15:04:05")),
        "Auto test triggered by h1spec",    // description
    )
    if err != nil {
        fmt.Printf("Create test run failed: %v\n", err)
        return err
    }
    r.runID = runID
    fmt.Printf("Created test run, id = %d\n", runID)
    return nil
}

func (r *TestinyReporter) CaseDone(res Result) error {
    testCaseID, ok := r.titleToID[res.Title]
    if !ok {
        fmt.Printf("WARN: Test case not found in Testiny: %s\n", res.Title)
        return nil
    }
    err := addOrUpdateTestCaseToRun(r.apiKey, r.runID, testCaseID)
    if err != nil {
        fmt.Printf("Failed to add testcase %d to run %d: %v\n", testCaseID, r.runID, err)
        return err
    }
    status := "FAILED"
    if res.Passed {
        status = "PASSED"
    }
    err = setTestResult(r.apiKey, r.runID, testCaseID, status)
    if err != nil {
        fmt.Printf("Failed to set test result for case %s: %v\n", res.Title, err)
    }

    return nil
}

func (r *TestinyReporter) EndRun() error {
    fmt.Println("Test run completed.")
    return nil
}