var search = function(nums, target) {
    let l = 0;
    let r = nums.length - 1;

    while (l <= r) {
        const mid = parseInt('' + (l + r) / 2, 10);
        if (nums[mid] === target) {
            return mid
        }

        if (nums[mid] < target) {
            l = mid + 1;
        } else {
            r = mid - 1;
        }
    }

    return -1
};
