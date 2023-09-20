var totalNQueens = function(n) {
    const check = (r, d, ad, c) => {
        if (r === n) {
            return 1;
        }

        let solutions = 0;

        for (let i = 0; i < n; i++) {
            const curr_d = r - i;
            const curr_ad = r + i;

            if (c.has(i) || d.has(curr_d) || ad.has(curr_ad)) {
                continue;
            }

            c.add(i);
            d.add(curr_d);
            ad.add(curr_ad);

            solutions += check(r + 1, d, ad, c);

            c.delete(i);
            d.delete(curr_d);
            ad.delete(curr_ad);
        }

        return solutions
    }

    return check(0, new Set(), new Set(), new Set());
};
