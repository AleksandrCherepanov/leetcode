var subsets = function(nums) {
    const ans = [];

    const generate = (curr, i) => {
        if (i > nums.length) {
            return;
        }

        ans.push([...curr]);

        for (let j = i; j < nums.length; j++) {
            curr.push(nums[j]);
            generate(curr, j + 1)
            curr.pop();
        }
    }

    generate([], 0);

    return ans;
};
