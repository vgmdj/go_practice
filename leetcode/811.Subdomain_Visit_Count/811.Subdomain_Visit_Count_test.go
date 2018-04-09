package Subdomain_Visit_Count

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubdomain(t *testing.T) {
	ast := assert.New(t)

	input := []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}

	checkArray(ast, subdomainVisits(input), []string{"901 mail.com",
		"50 yahoo.com", "900 google.mail.com", "5 wiki.org", "5 org",
		"1 intel.mail.com", "951 com"})

}

func checkArray(ast *assert.Assertions, src, dest []string) {
	ast.Equal(len(src), len(dest))

	for _, sv := range src {
		check := false

		for _, dv := range dest {
			if sv == dv {
				check = true
			}
		}

		ast.Equal(check, true)
	}

}
