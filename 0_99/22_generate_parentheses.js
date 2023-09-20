var generateParenthesis = function(n) {
    const ans = [];
    const generate = (curr, open, close) => {
        if (curr.length === n * 2) {
            ans.push(curr.join(''));
            return;
        }

        if (open < n) {
            curr.push('(');
            generate(curr, open + 1, close);
            curr.pop();
        }

        if (open > close) {
            curr.push(')');
            generate(curr, open, close + 1);
            curr.pop();
        }
    }

    generate(['('], 1, 0);

    return ans;
};
