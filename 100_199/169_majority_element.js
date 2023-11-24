var majorityElement = function(nums) {
    const k = Math.floor(nums.length / 2);
    const m = new Map();

    for (let n of nums) {
        if (m.has(n)) {
            m.set(n, m.get(n) + 1);
        } else {
            m.set(n, 1);
        }

        if (m.get(n) > k) {
            return n;
        }
    }
};

// Boyer-Moore majority vote algorithm
var majorityElement2 = function(nums) {
    let candidate = 0;
    let counter = 0;

    for (let n of nums) {
        if (counter === 0) {
            candidate = n;
            counter++;
        }

        counter = n === candidate ? counter + 1 : counter - 1
    }

    return candidate;
};