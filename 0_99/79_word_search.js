var exist = function (board, word) {
    const directions = [[1, 0], [0, 1], [-1, 0], [0, -1]];
    const seen = new Set();
    const m = board.length;
    const n = board[0].length;

    const valid = (row, col) => {
        return 0 <= row && row < m && 0 <= col && col <= n;
    }

    const key = (row, col) => {
        return row.toString() + col.toString();
    }

    const check = (row, col, i) => {
        if (i === word.length) {
            return true;
        }

        for (const [dr, dc] of directions) {
            const nr = row + dr;
            const nc = col + dc;
            const k = key(nr, nc);

            if (valid(nr, nc) && !seen.has(k) && board[nr][nc] === word[i]) {
                seen.add(k);
                if (check(nr, nc, i + 1)) {
                    return true
                }
                seen.delete(k);
            }
        }

        return false;
    }

    for (let row = 0; row < m; row++) {
        for (let col = 0; col < n; col++) {
            if (board[row][col] === word[0]) {
                seen.add(key(row, col));
                if (check(row, col, 1)) {
                    return true;
                }
                seen.delete(key(row, col));
            }
        }
    }

    return false;
};
