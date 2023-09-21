const coinChange = function(coins, amount) {
    let memo = new Map();

    const dp = (rest) => {
        if (rest < 0) {
            return -1;
        }

        if (rest === 0) {
            return 0;
        }

        if (memo.has(rest)) {
            return memo.get(rest);
        }

        let minCost = Infinity;
        for (let c of coins) {
            let cost = dp(rest - c);
            if (cost < 0) {
                continue;
            }

            minCost = Math.min(minCost, cost + 1);
        }

        memo.set(rest, minCost === Infinity ? - 1 : Math.min(minCost));
        return memo.get(rest);
    }

    return dp(amount);
};

console.log(coinChange([1,2,5], 11));
console.log(coinChange([2], 3));
console.log(coinChange([1], 0));
