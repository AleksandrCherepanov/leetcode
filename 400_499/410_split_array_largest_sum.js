/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var splitArray = function(nums, k) {
    let l = Math.max(...nums);
    let r = nums.reduce((sum, el) => sum + el, 0);

    if (k === 0) {
        return r;
    }

    let check = (arr, s, k) => {
        let sum = 0;
        let count = 0;

        for (let a of arr) {
            if (sum + a <= s) {
                sum += a;
            } else {
                sum = a;
                count++;
            }
        }

        return count + 1 <= k;
    }

    while (l <= r) {
        let mid = Math.floor((l + r) / 2);
        if (check(nums, mid, k)) {
            r = mid - 1;
        } else {
            l = mid + 1;
        }
    }

    return r + 1;
};

console.log(splitArray([7,2,5,10,8], 2));
console.log(splitArray([1,2,3,4,5], 2));