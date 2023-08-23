var combinationSum = function(candidates, target) {
    const generate = (curr, i, sum) => {
        if (sum === target) {
            ans.push([...curr]);
            return;
        }

        for (let j = i; j < candidates.length; j++) {
            const c = candidates[j];
            if (sum + c <= target) {
                curr.push(c);
                generate(curr, j, sum + c);
                curr.pop();
            }
        }
    }

    const ans = [];
    generate([], 0, 0);

    return ans;
};

