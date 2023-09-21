const maxProfit = function(k, prices) {
    const dp = (i, hold, remain) => {
        if (i === prices.length || remain === 0) {
            return 0;
        }

        if (memo[i][hold][remain] !== -1) {
            return memo[i][hold][remain];
        }

        let ans = dp(i + 1, hold, remain);
        if (!hold) {
            ans = Math.max(ans, -prices[i] + dp(i + 1, 1, remain));
        } else {
            ans = Math.max(ans, prices[i] + dp(i + 1, 0, remain - 1));
        }

        memo[i][hold][remain] = ans;
        return ans;
    }

    const memo = [];
    for (let i = 0; i < prices.length; i++) {
        memo.push([]);
        for (let j = 0; j < 2; j++) {
            memo[i].push(new Array(k + 1).fill(-1))
        }
    }

    return dp(0, 0, k);
};

console.log(maxProfit(2, [2,4,1]));
console.log(maxProfit(2, [3,2,6,5,0,3]));