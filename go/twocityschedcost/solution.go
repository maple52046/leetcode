package twocityschedcost

func twoCitySchedCost(costs [][]int) int {
	N := len(costs) / 2
	dp := make([][]int, N+1)

	dp[0] = make([]int, N+1)
	for i := 1; i <= N; i++ {
		dp[i] = make([]int, N+1)
		dp[i][0] = dp[i-1][0] + costs[i-1][0]
		dp[0][i] = dp[0][i-1] + costs[i-1][1]
	}

	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			x := dp[i-1][j] + costs[i+j-1][0]
			y := dp[i][j-1] + costs[i+j-1][1]
			dp[i][j] = getMin(x, y)
		}
	}

	return dp[N][N]
}

func getMin(x int, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
