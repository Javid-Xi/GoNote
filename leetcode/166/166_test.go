package _66

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [166]")
	examples := [][]string{
		{
			`1`, `2`,
			`"0.5"`,
		},
		{
			`2`, `3`,
			`"0.(6)"`,
		},
		{
			`3`, `1`,
			`"3"`,
		},

	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, fractionToDecimal, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}