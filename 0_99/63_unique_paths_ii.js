const uniquePathsWithObstacles = function(obstacleGrid) {
    const m = obstacleGrid.length;
    const n = obstacleGrid[0].length;

    const memo = [];
    for (let i = 0; i < m; i++) {
        memo.push(new Array(n).fill(-1));
    }

    const search = (row, col) => {
        if (row + col === 0 && obstacleGrid[row][col] !== 1) {
            return 1;
        }

        if (memo[row][col] > 0) {
            return memo[row][col];
        }

        let ans = 0;
        if (row > 0 && obstacleGrid[row][col] !== 1) {
            ans += search(row - 1, col);
        }

        if (col > 0 && obstacleGrid[row][col] !== 1) {
            ans += search(row, col - 1);
        }

        memo[row][col] = ans;
        return ans;
    }

    return search(m - 1, n - 1);
};
