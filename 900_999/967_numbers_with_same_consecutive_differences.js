var numsSameConsecDiff = function(n, k) {
    const generate = (curr) => {
        if (curr.length === n) {
            ans.add(parseInt(curr.join(''), 10));
            return;
        }

        const lastDigit = parseInt(curr[curr.length - 1], 10);
        for (const d of [lastDigit + k, lastDigit - k]) {
            if (d <= 9 && d >= 0) {
                curr.push('' + d);
                generate(curr);
                curr.pop();
            }
        }
    }

    const ans = new Set();
    for (let i = 1; i <= 9; i++) {
        generate(['' + i]);
    }
    return Array.from(ans);
};
