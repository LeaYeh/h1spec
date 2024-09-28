package config

import (
	"fmt"
	"strings"
	"time"
)

const (
	RunModeAll = iota
	RunModeGroup
	RunModeNone
)

type Config struct {
	AgentName	 string
	Host         string
	Port         int
	Path         string
	Timeout      time.Duration
	MaxHeaderLen int
	MaxBodyLen   int
	Strict       bool
	DryRun       bool
	Verbose      bool
	Sections     []string
	targetMap    map[string]bool
	Exec         string
	FromPort     int
}

// Addr returns the string concatenated with hostname and port number.
func (c *Config) Addr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

// RunMode returns a run mode of specified the section number.
// This is used to decide whether to run test cases.
func (c *Config) RunMode(section string) int {
	if c.targetMap == nil {
		c.buildTargetMap()
	}

	if len(c.targetMap) == 0 {
		return RunModeAll
	}

	comps := strings.Split(section, "/")
	compLen := len(comps)

	if compLen == 0 || compLen > 4 {
		return RunModeNone
	}

	keys := []string{comps[0]}

	for i := 1; i < compLen-1; i++ {
        nums := strings.Split(comps[i], ".")
        for j := range nums {
            key := fmt.Sprintf("%s/%s", strings.Join(comps[:i], "/"), strings.Join(nums[:j+1], "."))
            keys = append(keys, key)
        }
    }

	if compLen > 1 {
		keys = append(keys, section)
	}

	var result int
	for _, key := range keys {
		val, ok := c.targetMap[key]
		if ok {
			if val {
				return RunModeAll
			}
			result = RunModeGroup
		} else {
			result = RunModeNone
		}
	}

	return result
}

func (c *Config) buildTargetMap() {
	c.targetMap = map[string]bool{}

	for _, section := range c.Sections {
		comps := strings.Split(section, "/")
		compLen := len(comps)

		// Validate the format of the section string.
		if compLen == 0 || compLen > 4 {
			fmt.Printf("Invalid section: %s", section)
			continue
		}

		// Check the section string is root section or not.
		if compLen == 1 {
			c.targetMap[comps[0]] = true
			continue
		}

		_, ok := c.targetMap[comps[0]]
		if !ok {
			c.targetMap[comps[0]] = false
		}

		// The parent group of the test case associated with this section
		// must only run test cases included in the group.
		for i := 1; i < compLen-1; i++ {
            nums := strings.Split(comps[i], ".")
            for j := range nums {
                key := fmt.Sprintf("%s/%s", strings.Join(comps[:i], "/"), strings.Join(nums[:j+1], "."))
                c.targetMap[key] = false
            }
        }

		// The test case associated with this section string must be run.
		c.targetMap[section] = true
	}
}
