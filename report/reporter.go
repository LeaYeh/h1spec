package report

type Reporter interface {
    StartRun() error
    CaseDone(res Result) error
    EndRun() error
}
