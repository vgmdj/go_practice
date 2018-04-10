package Subdomain_Visit_Count

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	var rst []string
	var visits = make(map[string]int)

	for _, v := range cpdomains {
		splitNum := strings.Split(v, " ")
		count, _ := strconv.Atoi(splitNum[0])
		subDomains := splitNum[1]

		visits[subDomains] += count
		for i, k := range subDomains {
			if k == '.' {
				visits[subDomains[i+1:]] += count
			}
		}

	}

	for k, v := range visits {
		visit := fmt.Sprintf("%s %s", strconv.Itoa(v), k)
		rst = append(rst, visit)
	}

	return rst

}
