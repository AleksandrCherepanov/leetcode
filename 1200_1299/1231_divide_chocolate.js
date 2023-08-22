var maximizeSweetness = function(sweetness, k) {
    let l = 1;
    let r = sweetness.reduce((sum, el) => sum + el, 0);

    const check = (nums, s, k) => {
        let sum = 0;
        let count = 0;

        for (let n of nums) {
            sum += n;
            if (sum >= s) {
                count++;
                sum = 0
            }
        }

        return count >= k + 1;
    };

    while (l <= r) {
        const mid = Math.floor((l + r) / 2);
        if (check(sweetness, mid, k)) {
            l = mid + 1;
        } else {
            r = mid - 1;
        }
    }

    return r;
};

console.log(maximizeSweetness([1,2,3,4,5,6,7,8,9], 5));
console.log(maximizeSweetness([5,6,7,8,9,1,2,3,4], 8));
console.log(maximizeSweetness([1,2,2,1,2,2,1,2,2], 2));