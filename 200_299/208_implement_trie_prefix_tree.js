
const Trie = function() {
    this.end = false;
    this.children = new Map();
};

/**
 * @param {string} word
 * @return {void}
 */
Trie.prototype.insert = function(word) {
    let curr = this;
    for (let w of word) {
        if (!curr.children.has(w)) {
            curr.children.set(w, new Trie());
        }
        curr = curr.children.get(w);
    }
    curr.end = true;
};

/**
 * @param {string} word
 * @return {boolean}
 */
Trie.prototype.search = function(word) {
    let curr = this
    for (let w of word) {
        if (curr.children.has(w)) {
            curr = curr.children.get(w);
        } else {
            return false;
        }
    }

    return curr.end;
};

/**
 * @param {string} prefix
 * @return {boolean}
 */
Trie.prototype.startsWith = function(prefix) {
    let curr = this
    for (let w of prefix) {
        if (curr.children.has(w)) {
            curr = curr.children.get(w);
        } else {
            return false;
        }
    }

    return true;
};
