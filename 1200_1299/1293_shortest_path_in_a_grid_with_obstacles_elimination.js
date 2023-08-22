let shortestPath = function (grid, k) {
    let m = grid.length
    let n = grid[0].length

    let valid = (x, y) => {
        return x >= 0 && y >= 0 && x < m && y < n;
    }

    let queue = [[0, 0, k]];
    let directions = [[1, 0], [0, 1], [-1, 0], [0, -1]];
    let steps = 0;

    let seen = [];
    for (let i = 0; i < m; i++) {
        seen.push(new Array(n).fill(-1))
    }

    while (queue.length > 0) {
        const currLen = queue.length;
        for (let j = 0; j < currLen; j++) {
            let [x, y, remain] = queue.shift();
            if (x === m - 1 && y === n - 1) {
                return steps
            }

            if (grid[x][y] === 1) {
                if (remain === 0) {
                    continue
                } else {
                    remain--;
                }
            }

            if (seen[x][y] >= remain) {
                continue
            }

            seen[x][y] = remain;
            for (const [dx, dy] of directions) {
                const newX = x + dx;
                const newY = y + dy;

                if (valid(newX, newY)) {
                    queue.push([newX, newY, remain])
                }
            }
        }
        console.log(seen)
        steps++
    }

    return -1
};
