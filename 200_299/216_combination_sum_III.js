var combinationSum3 = function(k, n) {
    const generate = (curr, sum, i) => {
        if (sum === n && curr.length === k) {
            ans.push([...curr]);
            return;
        }

        for (let j = i; j < 10; j++) {
            if (!curr.includes(j)) {
                curr.push(j);
                generate(curr, sum + j, j + 1);
                curr.pop()
            }
        }
    }

    const ans = [];
    generate([], 0, 1);
    return ans;

};
