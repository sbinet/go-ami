package ami

import (
	"fmt"
	"sort"
	"strings"
)

// GetPeriods returns all periods at a specified detail level in the given year
func GetPeriods(c *Client, year, level int) ([]RunPeriod, error) {

	if year > 2000 {
		year %= 1000
	}

	cmd := []string{
		"ListDataPeriods",
		"-createdSince=2009-01-01 00:00:00",
		fmt.Sprintf("-projectName=data%02d%%", year),
	}

	switch level {
	case 1, 2, 3:
		cmd = append(cmd, fmt.Sprintf("-periodLevel=%d", level))
	default:
		return nil, fmt.Errorf("ami.GetPeriods: level must be 1,2 or 3 (got: %v)", level)
	}

	msg, err := c.Execute(cmd...)
	if err != nil {
		return nil, err
	}

	periods := make([]RunPeriod, 0, len(msg.Result.Rows))
	for _, row := range msg.Result.Rows {
		m := row.Value()
		periods = append(periods,
			RunPeriod{
				Project:     m["projectName"].(string),
				Year:        year,
				Name:        m["period"].(string),
				Level:       level,
				Status:      m["status"].(string),
				Description: m["description"].(string),
			})
	}

	return periods, nil
}

// GetRuns returns all runs contained in the given periods in the specified year
func GetRuns(c *Client, str_periods string, year int) ([]int, error) {

	runset := make(map[int]struct{})
	var periods []string
	{
		if str_periods != "" {
			periods = strings.Split(str_periods, ",")
			//fmt.Printf("=============> %v %v %q\n", periods, len(periods), str_periods)
		} else {
			//fmt.Printf("=============\n")
			// get all periods
			level := 1
			runperiods, err := GetPeriods(c, year, level)
			if err != nil {
				return nil, err
			}

			for _, rp := range runperiods {
				periods = append(periods, rp.Name)
			}
		}
	}

	for _, period := range periods {
		amiargs := []string{"GetRunsForDataPeriod"}
		amiargs = append(amiargs,
			fmt.Sprintf("-projectName=data%02d%%", year),
			fmt.Sprintf("-period=%s", period),
		)

		msg, err := c.Execute(amiargs...)
		if err != nil {
			return nil, err
		}

		for _, v := range msg.Result.Rows {
			m := v.Value()
			//fmt.Printf("%d\n", m["runNumber"])
			runset[m["runNumber"].(int)] = struct{}{}
		}
	}

	runs := make([]int, 0, len(runset))
	for k, _ := range runset {
		runs = append(runs, k)
	}

	sort.Ints(runs)
	return runs, nil
}

// EOF
