/**
 * @param {number[]} nums
 * @return {number}
 */
var removeDuplicates = function(nums) {
    let prev = nums[0];
    let l = 1;
    let most = 1;

    for (let i = 1; i < nums.length; i++) {
        if (nums[i] === prev) {
            if (most > 0) {
                most--;
                nums[l] = nums[i];
                l++;
            }
        } else {
            prev = nums[i];
            nums[l] = nums[i];
            l++;
            most = 1;
        }
    }

    return l;
};
