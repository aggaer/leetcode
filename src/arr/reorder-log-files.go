package arr

import (
	"sort"
	"strconv"
	"strings"
)

func ReorderLogFiles(logs []string) []string {
	var req Log = logs
	sort.Sort(req)
	return req
}

func isNum(log string) bool {
	sig := strings.Split(log, " ")[1]
	_, err := strconv.Atoi(sig)
	return err == nil
}

type Log []string

func (l Log) Len() int {
	return len(l)
}

func (l Log) Less(i, j int) bool {
	strArrI := strings.Split(l[i], " ")
	strArrJ := strings.Split(l[j], " ")
	isNumI := isNum(l[i])
	isNumJ := isNum(l[j])
	if !isNumI && !isNumJ {
		cmp := strings.Compare(strArrI[1], strArrJ[1])
		if cmp != 0 {
			return cmp < 0
		}
		return strings.Compare(strArrI[0], strArrJ[0]) < 0
	} else if isNumI && !isNumJ || isNumI && isNumJ {
		return false
	} else {
		return true
	}
}

func (l Log) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}
