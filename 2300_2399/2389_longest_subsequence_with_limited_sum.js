var bs = function (nums, target) {
    let l = 0;
    let r = nums.length - 1;

    while (l <= r) {
        const mid = Math.floor((l+r)/2);
        if (nums[mid] === target) {
            return mid + 1;
        }

        if (nums[mid] > target) {
            r = mid - 1;
        } else {
            l = mid + 1;
        }
    }

    return l;
}

var answerQueries = function(nums, queries) {
    nums.sort((a,b) => a - b);
    const prefix = [nums[0]];
    for (let i = 1; i < nums.length; i++) {
        prefix.push(prefix[i-1] + nums[i]);
    }

    let answer = [];
    queries.forEach((q) => {
        answer.push(bs(prefix, q))
    });

    return answer;
};
