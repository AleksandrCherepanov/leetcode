var minCostClimbingStairs = function(cost) {
    const dp = (i) => {
        if (i <= 1) {
            return 0;
        }

        if (memo.has(i)) {
            return memo.get(i);
        }

        memo.set(i, Math.min(dp(i - 1) + cost[i - 1], dp(i-2) + cost[i - 2]));
        return memo.get(i);
    }

    const memo = new Map();
    return dp(cost.length);
};

console.log(minCostClimbingStairs([10,15,20]));
console.log(minCostClimbingStairs([1,100,1,1,1,100,1,1,100,1]));