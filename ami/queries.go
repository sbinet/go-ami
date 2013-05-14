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

	type response struct {
		run int
		err error
	}
	jobs := make(chan int, c.nqueries)
	runch := make([]chan response, len(periods))
	for i, period := range periods {
		amiargs := []string{"GetRunsForDataPeriod"}
		amiargs = append(amiargs,
			fmt.Sprintf("-projectName=data%02d%%", year),
			fmt.Sprintf("-period=%s", period),
		)
		runch[i] = make(chan response)

		go func(ich int) {
			jobs <- 1
			msg, err := c.Execute(amiargs...)
			if err != nil {
				//fmt.Printf("**error** %v\n", err)
				runch[ich] <- response{0, err}
			} else {

				for _, v := range msg.Result.Rows {
					m := v.Value()
					//fmt.Printf("--> %d (%d)\n", m["runNumber"], ich)
					//runset[m["runNumber"].(int)] = struct{}{}
					runch[ich] <- response{m["runNumber"].(int), nil}
				}
				close(runch[ich])
			}
			<-jobs
		}(i)
	}
	for _, ch := range runch {
		for r := range ch {
			if r.err != nil {
				//fmt.Printf("**error** %v\n", r.err)
				return nil, r.err
			}
			runset[r.run] = struct{}{}
			//fmt.Printf("<<- %d\n", r.run)
		}
	}

	runs := make([]int, 0, len(runset))
	for k, _ := range runset {
		runs = append(runs, k)
	}

	sort.Ints(runs)
	return runs, nil
}

// GetDatasets returns a list of ami.Dataset matching a pattern
func GetDatasets(c *Client, pattern, parent_type, order string, limit int, fields string, flatten bool, show_archived bool, opts map[string]interface{}) ([]string, error) {
	datasets := []string{}
	var err error

	if c.verbose {
		fmt.Printf("ami.GetDatasets: pattern=%q parent-type=%q order=%q limit=%d fields=%q flatten=%v show-archived=%v opts=%v\n", pattern, parent_type, order, limit, fields, flatten, show_archived, opts)
	}

	return datasets, err
}

// clean_dataset removes trailing slashes
func clean_dataset(dsname string) string {
	return strings.TrimRight(dsname, "/")
}

func search_query(c *Client, entity, cmd string, cmdargs []string, pattern, order string, limit int, fields string, flatten bool, mode, project_name, processing_step_name string, show_archived bool, opts map[string]interface{}) (interface{}, error) {

	table, ok := Tables[entity]
	if !ok {
		return nil, fmt.Errorf("ami: no such entity %q", entity)
	}
	pfield := table.Primary
	fmt.Printf("pfield: %v\n", pfield)
	return nil, nil
}

// expand_period_constraints takes a string and returns a suitable string for AMI db
// e.g.
//  period=B  -> period like B%
//  period=B2 -> period=B2
func expand_period_constraints(periods string) string {
	ps := strings.Split(periods, ",")
	sels := []string{}
	for _, p := range ps {
		switch len(p) {
		case 0:
		case 1:
			sels = append(sels, fmt.Sprintf(`period like "%s%%"`, p))
		default:
			sels = append(sels, fmt.Sprintf(`period="%s"`, p))
		}
	}
	return strings.Join(sels, " OR ")
}

// parse_fields returns the list of field names which are validated by the table
// fields must be a comma-separated list of field names.
func parse_fields(fields string, table *Table) ([]string, error) {
	lst := make([]string, 0)
	for _, field := range strings.Split(fields, ",") {
		if field != "" {
			lst = append(lst, field)
		}
	}
	query_fields := make([]string, 0, len(lst))
	for _, name := range lst {
		name, err := validate_field(name, table)
		if err != nil {
			return nil, err
		}
		query_fields = append(query_fields, name)
	}
	return query_fields, nil
}

func validate_field(field string, table *Table) (string, error) {
	if name, ok := table.Fields[field]; ok {
		return name, nil
	}

	name := strings.Replace(field, "-", "_", -1)
	if strings.Contains(name, ".") {
		foreign := strings.Split(name, ".")
		fname, ffield := foreign[0], foreign[1]
		fentity, ok := table.Foreign[fname]
		if !ok {
			return "", fmt.Errorf("ami: %q is not associated with %q", table.Name, fname)
		}
		found := false
		for _, v := range fentity.Fields {
			if v == ffield {
				found = true
				break
			}
		}
		if !found {
			ffield = fentity.Fields[ffield]
		}
		name = fmt.Sprintf("%s.%s", fname, ffield)
	} else {
		n, ok := table.Fields[name]
		if !ok {
			return "", fmt.Errorf("ami: Table.Field %q does not exist\n Valid fields are: %v", name, table.Fields)
		}
		name = n
	}
	return name, nil
}

// EOF
