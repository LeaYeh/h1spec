package h1spec

import (
	"fmt"
	"time"

	"github.com/LeaYeh/h1spec/log"
	"github.com/LeaYeh/h1spec/config"
)

func Run(c *config.Config) (bool, error) {
	total := 0
	success := true

	// spec := http1.Spec()

	start := time.Now()
	// for _, s := range specs {
	// 	s.Test(c)

	// 	if s.FailedCount > 0 {
	// 		success = false
	// 	}

	// 	total += s.FailedCount
	// 	total += s.SkippedCount
	// 	total += s.PassedCount
	// }
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
