var hammingDistance = function(x, y) {
    let xor = x ^ y;
    let d = 0;
    while (xor !== 0) {
        if (xor % 2 === 1) {
            d++;
        }
        xor = xor >> 1;
    }

    return d;
};

//Brian Kernighan's Algorithm
var hammingDistance2 = function(x, y) {
    let xor = x ^ y;
    let d = 0;
    while (xor !== 0) {
        xor = xor & (xor - 1);
        d++;
    }

    return d;
};