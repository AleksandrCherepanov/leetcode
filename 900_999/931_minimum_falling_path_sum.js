const minFallingPathSum = function(matrix) {
    for (let i = 1; i < matrix.length; i++) {
        for (let j = 0; j < matrix.length; j++) {
            let ans = matrix[i][j] + matrix[i - 1][j];
            if (j - 1 >= 0) {
                ans = Math.min(ans, matrix[i][j] + matrix[i - 1][j - 1]);
            }
            if (j + 1 < matrix.length) {
                ans = Math.min(ans, matrix[i][j] + matrix[i - 1][j + 1]);
            }
            matrix[i][j] = ans;
        }
    }

    let ans = Infinity;
    for (let i = 0; i < matrix.length; i++) {
        ans = Math.min(ans, matrix[matrix.length - 1][i]);
    }
    return ans;
};
