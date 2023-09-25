const maxProfit = function(prices) {
    let hold = new Array(prices.length + 1).fill(-Infinity);
    let sold = new Array(prices.length + 1).fill(-Infinity);
    let empty = new Array(prices.length + 1).fill(0);

    for (let i = 1; i <= prices.length; i++) {
        sold[i] = hold[i - 1] + prices[i - 1];
        hold[i] = Math.max(hold[i - 1], empty[i - 1] - prices[i-1]);
        empty[i] = Math.max(empty[i - 1], sold[i - 1]);
    }

    return Math.max(sold[prices.length], empty[prices.length]);
};

