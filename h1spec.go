package h1spec

import (
	"fmt"
	"time"

	"github.com/LeaYeh/h1spec/config"
	"github.com/LeaYeh/h1spec/report"
	"github.com/LeaYeh/h1spec/log"
	"github.com/LeaYeh/h1spec/spec"
	"github.com/LeaYeh/h1spec/spec/http1.1"
)

func Run(c *config.Config, rep report.Reporter) (bool, error) {
	total := 0
	success := true

	specs := []*spec.TestGroup{
		http11.Spec(),
	}

	start := time.Now()
	for _, s := range specs {
		s.Test(c, rep)

		if s.FailedCount > 0 {
			success = false
		}

		total += s.FailedCount
		total += s.SkippedCount
		total += s.PassedCount
	}
	end := time.Now()
	d := end.Sub(start)

	if c.DryRun {
		return true, nil
	}

	if total == 0 {
		log.SetIndentLevel(0)
		log.Println("No matched tests found.")
		return true, nil
	}

	if !success {
		log.SetIndentLevel(0)
	}

	log.SetIndentLevel(0)
	log.Println(fmt.Sprintf("Finished in %.4f seconds", d.Seconds()))

	return success, nil
}
