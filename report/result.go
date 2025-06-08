package report

// Result represents a single test case result in the report.
type Result struct {
    Title    string
    Passed   bool
    Duration int64
    Message  string
}
