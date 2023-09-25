const longestCommonSubsequence = function(text1, text2) {
    const memo = new Map();

    const dp = (i, j) => {
        if (i === text1.length || j === text2.length) {
            return 0;
        }

        if (memo.has('' + i + j)) {
            return memo.get('' + i + j);
        }

        if (text1[i] === text2[j]) {
            memo.set('' + i + j, 1 + dp(i + 1, j + 1));
        } else {
            memo.set('' + i + j, Math.max(dp(i + 1, j), dp(i, j + 1)));
        }

        return memo.get('' + i + j);
    }

    return dp(0, 0);
};
