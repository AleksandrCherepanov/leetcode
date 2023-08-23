var letterCombinations = function (digits) {
    if (digits.length === 0) {
        return [];
    }

    const buttons = {
        '0': [],
        '1': [],
        '2': ['a', 'b', 'c'],
        '3': ['d', 'e', 'f'],
        '4': ['g', 'h', 'i'],
        '5': ['j', 'k', 'l'],
        '6': ['m', 'n', 'o'],
        '7': ['p', 'q', 'r', 's'],
        '8': ['t', 'u', 'v'],
        '9': ['w', 'x', 'y', 'z']
    };

    const ans = [];

    const generate = (curr, i) => {
        if (curr.length === digits.length) {
            ans.push(curr.join(''))
            return;
        }

        for (let l of buttons[digits[i]]) {
            curr.push(l);
            generate(curr, i + 1);
            curr.pop();
        }
    }

    generate([], 0);

    return ans;
};
