def reversePrefix(word, ch):
    ss = list(word)
   
    index = -1
    for i, c in enumerate(ss):
        if c == ch:
            index = i
            break

    if index < 0:
        return word

    j = 0
    i = index
    while j < i:
        t = ss[i]
        ss[i] = ss[j]
        ss[j] = t
        i -= 1
        j += 1

    return ''.join(ss)


test = [
    [
        "abcdefd",
        "d",
        "dcbaefd"
    ],
    [
        "xyxzxe",
        "z",
        "zxyxxe"
    ],
    [
        "abcd",
        "z",
        "abcd"
    ],
]

for i in test:
    result = reversePrefix(i[0], i[1])

    if result != i[2]:
        print("Error. ", result, i[2])
    else:
        print("OK")


