var combine = function(n, k) {
    const ans = [];

    const genereate = (curr, i) => {
        if (curr.length === k) {
            ans.push([...curr]);
        }

        for (let j = i; j <= n; j++) {
            curr.push(j);
            genereate(curr, j + 1);
            curr.pop();
        }
    }

    genereate([], 1);

    return ans;
};
