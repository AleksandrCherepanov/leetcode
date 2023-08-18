var maxNumberOfApples = function(weight) {
    weight.sort((a, b) => a - b);

    let limit = 5000;
    let ans = 0;

    for (let i = 0; i < weight.length; i++) {
        if (weight[i] <= limit) {
            ans++
            limit -= weight[i];
        }
    }

    return ans;
};
