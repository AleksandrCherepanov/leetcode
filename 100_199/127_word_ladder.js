var ladderLength = function(beginWord, endWord, wordList) {
    const wordListSet = new Set(wordList);
    if (!wordListSet.has(endWord)) {
        return 0;
    }

    const alphabet = Array.from(Array(26)).map((e, i) => i + 97).map((x) => String.fromCharCode(x));

    let queue = [beginWord];
    let seen = new Map([[beginWord, true]]);
    let steps = 1;

    while (queue.length > 0) {
        const currLen = queue.length;
        for (let i = 0; i < currLen; i++) {
            const word = queue.shift();
            if (word === endWord) {
                return steps;
            }
            for (let j = 0; j < word.length; j++) {
                for (let k = 0; k < alphabet.length; k++) {
                    const newWord = word.slice(0, j) + alphabet[k] + word.slice(j+1, word.length);
                    if (!seen.has(newWord) && wordListSet.has(newWord)) {
                        queue.push(newWord);
                        seen.set(newWord, true);
                    }
                }
            }
        }
        steps++;
    }

    return 0;
};
