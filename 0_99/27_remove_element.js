/**
 * @param {number[]} nums
 * @param {number} val
 * @return {number}
 */
var removeElement = function(nums, val) {
    let k = 0;
    let l = 0;
    let r = nums.length - 1;

    while (l <= r) {
        if (nums[l] !== val) {
            k++;
            l++;
        } else if (nums[r] === val) {
            r--;
        } else {
            const t = nums[r];
            nums[r] = nums[l];
            nums[l] = t;
            r--;
            l++;
            k++;
        }
    }

    return k;
};
