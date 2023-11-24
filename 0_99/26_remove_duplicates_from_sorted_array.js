/**
 * @param {number[]} nums
 * @return {number}
 */
var removeDuplicates = function(nums) {
    let prev = nums[0];
    let l = 1;
    let r = 1;

    for (let i = r; i < nums.length; i++) {
        if (nums[i] !== prev) {
            prev = nums[i];
            nums[l] = nums[i];
            l++;
        }
    }

    return l;
};
