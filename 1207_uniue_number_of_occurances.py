from collections import defaultdict

def uniqueOccurrences(arr):
    freq = defaultdict(int)
    uniq = set()

    for a in arr:
        freq[a] += 1
    
    for f in freq.values():
        uniq.add(f)

    return len(freq) == len(uniq)

test = [
    [
        [1,2,2,1,1,3],
        True
    ],
    [
        [1,2],
        False
    ],
    [
        [-3,0,1,-3,1,1,1,-3,10,0],
        True
    ]
]

for i in test:
    result = uniqueOccurrences(i[0])

    if result != i[-1]:
        print("Error. ", result, i[-1])
    else:
        print("OK")

