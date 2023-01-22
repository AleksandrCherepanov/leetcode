def checkIfPangram(sentence):
    letters = set()
    for c in sentence:
        letters.add(c)

    return len(letters) == 26

test = [
    [
        "thequickbrownfoxjumpsoverthelazydog",
        True
    ],
    [
        "leetcode", 
        False
    ],
]

for i in test:
    result = checkIfPangram(i[0])

    if result != i[-1]:
        print("Error. ", result, i[-1])
    else:
        print("OK")

