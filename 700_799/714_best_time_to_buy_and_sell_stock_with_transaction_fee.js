var maxProfit = function(prices, fee) {
    const dp = (i, holding) => {
        if (i === prices.length) {
            return 0;
        }

        if (memo[i][holding] !== -1) {
            return memo[i][holding];
        }

        let ans = dp(i + 1, holding);
        if (holding) {
            ans = Math.max(ans, prices[i] + dp(i + 1, 0) - fee);
        } else {
            ans = Math.max(ans, -prices[i] + dp(i + 1, 1));
        }

        memo[i][holding] = ans;
        return ans;
    }

    const memo = [];
    for (let i = 0; i < prices.length; i++) {
        memo.push([-1, -1]);
    }

    return dp(0, 0);
};

console.log(maxProfit([1,3,2,8,4,9], 2));
console.log(maxProfit([1,3,7,5,10,3], 3));