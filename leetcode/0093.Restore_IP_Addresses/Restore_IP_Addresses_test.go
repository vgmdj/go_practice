package Restore_IP_Addresses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestoreIP(t *testing.T) {
	ast := assert.New(t)

	ast.Equal([]string{"0.10.0.10", "0.100.1.0"}, restoreIpAddresses("010010"))
	ast.Equal([]string{"255.255.11.135", "255.255.111.35"}, restoreIpAddresses("25525511135"))

}
