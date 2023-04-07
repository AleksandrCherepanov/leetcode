from collections import defaultdict

def frequencySort(s):
    d = defaultdict(int)
    l = list(s)

    for c in l:
        d[c] += 1

    sortedDict = sorted(d.items(), key=lambda item: item[1], reverse=True)
    
    result = ''
    for t in sortedDict:
        result += t[0] * t[1]

    return result

test = [
    [
        "tree",
        "eert"
    ],
    [
        "cccaaa",
        "aaaccc"
    ],
    [
        "Aabb",
        "bbAa"
    ]
]

for i in test:
    result = frequencySort(i[0])

    if result != i[-1]:
        print("Error. ", result, i[-1])
    else:
        print("OK")

