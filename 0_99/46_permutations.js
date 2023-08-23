var permute = function(nums) {
    let ans = [];

    const generate = (curr) => {
        if (curr.length === nums.length) {
            ans.push([...curr]);
            return;
        }

        for (let n of nums) {
            if (!curr.includes(n)) {
                curr.push(n);
                generate(curr);
                curr.pop();
            }
        }
    }

    generate([]);

    return ans;
};
