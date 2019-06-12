package Unique_Binary_Search_Trees

/*
当n=3是 num = [1, 2, 3]
此时根节点有三种情况：

1为根节点时，1的左子树是0个，右子树是2,3（两个）。所以1是根节点的种类树是: res[0] * res[2]
2为根节点时，2的左子树是1（1个），右子树是3（1个）。所以2是根节点的种类树是: res[1] * res[1]
3为根节点时，3的左子树是1,2（2个），右子树是0个。所以3是根节点的种类树是: res[2] * res[0]
所以 res[3] = res[0]*res[2] + res[1]*res[1] + res[2]*res[0]

*/

func numTrees(n int) int {
	if n < 3 {
		return n
	}

	var dp = make([]int, n+1)
	dp[0], dp[1], dp[2] = 1, 1, 2

	for i := 3; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i-j-1]
		}
	}

	return dp[n]

}
