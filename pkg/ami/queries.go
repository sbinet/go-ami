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

	jobs := make(chan int, c.nqueries)
	errch := make(chan error)
	runch := make([]chan int, len(periods))
	for i, period := range periods {
		amiargs := []string{"GetRunsForDataPeriod"}
		amiargs = append(amiargs,
			fmt.Sprintf("-projectName=data%02d%%", year),
			fmt.Sprintf("-period=%s", period),
		)
		runch[i] = make(chan int)

		go func(ich int) {
			jobs <- 1
			msg, err := c.Execute(amiargs...)
			if err != nil {
				//return nil, err
				errch <- err
			} else {

				for _, v := range msg.Result.Rows {
					m := v.Value()
					//fmt.Printf("--> %d (%d)\n", m["runNumber"], ich)
					//runset[m["runNumber"].(int)] = struct{}{}
					runch[ich] <- m["runNumber"].(int)
				}
				close(runch[ich])
			}
			<-jobs
		}(i)
	}
	done := 0
	for {
		//fmt.Printf("runch: %v, done: %v/%v\n", len(runch), done, len(periods))
		select {
		case err := <-errch:
			fmt.Printf("err: %v\n", err)
			return nil, err
		default:
			for _, ch := range runch {
				for run := range ch {
					runset[run] = struct{}{}
					//fmt.Printf("<<- %d\n", run)
				}
			}
			done += 1
		}
		if done == len(periods) {
			break
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
