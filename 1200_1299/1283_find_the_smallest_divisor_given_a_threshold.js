var smallestDivisor = function(nums, threshold) {
    const check = (nums, d) => {
        let s = 0;
        for (let n of nums) {
            s += Math.ceil(n / d);
        }

        return s <= threshold;
    }

    let l = 1;
    let r = Math.max(...nums);

    while (l <= r) {
        const mid = Math.floor((l + r) / 2);
        if (check(nums, mid)) {
            r = mid - 1;
        } else {
            l = mid + 1;
        }
    }

    return l;
};
