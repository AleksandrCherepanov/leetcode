/**
 * @param {number[]} nums1
 * @param {number} m
 * @param {number[]} nums2
 * @param {number} n
 * @return {void} Do not return anything, modify nums1 in-place instead.
 */
var merge = function(nums1, m, nums2, n) {
    let l = m - 1;
    let r = n - 1;
    let idx = m + n - 1;

    while(l >= 0 && r >= 0) {
        if (nums1[l] > nums2[r]) {
            nums1[idx] = nums1[l];
            l--;
        } else {
            nums1[idx] = nums2[r];
            r--;
        }
        idx--;
    }

    for(let i = r; i >= 0; i--) {
        nums1[i] = nums2[i];
    }
};
