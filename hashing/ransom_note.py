from collections import defaultdict

def canConstruct(ransomNote, magazine):
    first = dict() 
    for i in ransomNote:
        if i in first:
            first[i] += 1
        else:
            first[i] = 1
    
    for i in magazine:
        if i in first:
            first[i] -= 1
            if first[i] == 0:
                del first[i]
    return len(first) == 0 

test = [
    [
        "a",
        "b",
        False
    ],
    [
        "aa",
        "ab",
        False
    ],
    [
        "aa",
        "aab",
        True
    ]
]

for i in test:
    result = canConstruct(i[0], i[1])

    if result != i[-1]:
        print("Error. ", result, i[-1])
    else:
        print("OK")

