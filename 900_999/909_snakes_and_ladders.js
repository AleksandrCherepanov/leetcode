function snakesAndLadders(board: number[][]): number {
    const n = board.length;
    let seen = new Array(n * n + 1).fill(-1);
    let cells = new Array(n*n+1).fill([]);
    let columns = [];
    let label = 1;
    for (let i = 0; i < n; i++) {
        columns.push(i);
    }

    for(let i = n - 1; i >= 0; i--) {
        for (const col of columns) {
            cells[label++] = [i, col];
        }
        columns.reverse()
    }

    let queue = [1]
    seen[1] = 0;

    while (queue.length > 0) {
        let curr = queue.shift();
        for (let next = curr + 1; next <= Math.min(curr + 6, n * n); next++) {
            let [row, col] = cells[next];
            let dest = board[row][col] !== -1 ? board[row][col] : next;
            if (seen[dest] === -1) {
                seen[dest] = seen[curr] + 1;
                queue.push(dest);
            }
        }
    }

    return seen[n*n];
};
