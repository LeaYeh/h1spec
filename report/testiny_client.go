package report

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "bytes"
)

type TestinyCase struct {
    ID    int    `json:"id"`
    Title string `json:"title"`
}

func getTestCases(apiKey string) ([]TestinyCase, error) {
    url := "https://app.testiny.io/api/v1/testcase"
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }
    req.Header.Add("Accept", "application/json")
    req.Header.Add("X-Api-Key", apiKey)

    res, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer res.Body.Close()

    if res.StatusCode != 200 {
        b, _ := io.ReadAll(res.Body)
        return nil, fmt.Errorf("Testiny API error: %s %s", res.Status, string(b))
    }

    var parsed struct {
        Data []TestinyCase `json:"data"`
    }
    if err := json.NewDecoder(res.Body).Decode(&parsed); err != nil {
        return nil, err
    }
    return parsed.Data, nil
}

func BuildTitleToIDMap(cases []TestinyCase) map[string]int {
    m := make(map[string]int)
    for _, tc := range cases {
        m[tc.Title] = tc.ID
    }
    return m
}

func createTestRun(apiKey string, projectID int, title, description string) (int, error) {
    payload := map[string]interface{}{
        "title":       title,
        "project_id":  projectID,
        "description": description,
    }
    body, _ := json.Marshal(payload)

    req, err := http.NewRequest("POST", "https://app.testiny.io/api/v1/testrun", bytes.NewReader(body))
    if err != nil {
        return 0, err
    }
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("X-Api-Key", apiKey)

    res, err := http.DefaultClient.Do(req)
    if err != nil {
        return 0, err
    }
    defer res.Body.Close()

    if res.StatusCode < 200 || res.StatusCode > 299 {
        msg, _ := io.ReadAll(res.Body)
        return 0, fmt.Errorf("Testiny API error: %s", string(msg))
    }

    var resp struct {
        ID int `json:"id"`
    }
    if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
        return 0, err
    }
    return resp.ID, nil
}

func addOrUpdateTestCaseToRun(apiKey string, testRunID, testCaseID int) error {
    body := []map[string]interface{}{
        {
            "ids": map[string]int{
                "testcase_id": testCaseID,
                "testrun_id":  testRunID,
            },
            "mapped": map[string]string{
                "result_status": "NOTRUN",
                "assigned_to": "ANY",
            },
        },
    }
    jsonBody, _ := json.Marshal(body)
    url := "https://app.testiny.io/api/v1/testrun/mapping/bulk/testcase:testrun?op=add_or_update"
    req, err := http.NewRequest("POST", url, bytes.NewReader(jsonBody))
    if err != nil { return err }
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("X-Api-Key", apiKey)
    res, err := http.DefaultClient.Do(req)
    if err != nil { return err }
    defer res.Body.Close()
    if res.StatusCode >= 300 {
        b, _ := io.ReadAll(res.Body)
        return fmt.Errorf("Testiny: %s %s", res.Status, string(b))
    }
    return nil
}

func setTestResult(apiKey string, testRunID int, testCaseID int, status string) error {
    body := []map[string]interface{}{
        {
            "ids": map[string]int{
                "testcase_id": testCaseID,
                "testrun_id":  testRunID,
            },
            "mapped": map[string]string{
                "result_status": status, // "PASSED", "FAILED", ...
            },
        },
    }
    fmt.Println("Setting result for case ID:", testCaseID, "in run ID:", testRunID, "with status:", status)
    jsonBody, _ := json.Marshal(body)

    url := "https://app.testiny.io/api/v1/testrun/mapping/bulk/testcase:testrun?op=update"
    req, err := http.NewRequest("POST", url, bytes.NewReader(jsonBody))
    if err != nil {
        return err
    }
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("X-Api-Key", apiKey)
    res, err := http.DefaultClient.Do(req)
    if err != nil {
        return err
    }
    defer res.Body.Close()
    b, _ := io.ReadAll(res.Body)
    if res.StatusCode >= 300 {
        return fmt.Errorf("Testiny: %s %s", res.Status, string(b))
    }
    return nil
}
