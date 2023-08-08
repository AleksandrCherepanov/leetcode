const nearestExit = function (maze, entrance) {
    const m = maze.length;
    const n = maze[0].length;

    const directions = [[1, 0], [0, 1], [-1, 0], [0, -1]];
    let seen = [];
    for (let i = 0; i < m; i++) {
        seen.push(new Array(n).fill(false));
    }

    const valid = (x, y) => {
        return x >= 0 && x < m && y >= 0 && y < n && maze[x][y] === ".";
    }

    let queue = [entrance];
    seen[entrance[0]][entrance[1]] = true;

    let steps = 0;

    while (queue.length) {
        const currLen = queue.length;
        for (let i = 0; i < currLen; i++) {
            const [x, y] = queue.shift();
            if (x === 0 || x === m - 1 || y === 0 || y === n - 1) {
                return steps;
            }

            for (const [dx,dy] in directions) {
                const newX = dx + x;
                const newY = dy + y;

                if (valid(newX, newY) && !seen[newX][newY]) {
                    seen[newX][newY] = true;
                    queue.push([newX, newY]);
                }
            }
        }
        steps++;
    }

    return -1;
};