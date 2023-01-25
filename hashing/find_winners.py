from collections import defaultdict

def findWinners(matches):
    loosers = defaultdict(int)
    for m in matches:
        loosers[m[1]] += 1

    win = set()
    los = set()

    for k, v in loosers.items():
        if v == 1:
            los.add(k)
    
    for m in matches:
        if m[0] not in loosers:
            win.add(m[0])

    return [sorted(win), sorted(los)]

test = [
    [
        [[1,3],[2,3],[3,6],[5,6],[5,7],[4,5],[4,8],[4,9],[10,4],[10,9]],
        [[1,2,10],[4,5,7,8]]
    ],
    [
        [[2,3],[1,3],[5,4],[6,4]],
        [[1,2,5,6],[]]
    ]
]

for i in test:
    result = findWinners(i[0])

    if result != i[-1]:
        print("Error. ", result, i[-1])
    else:
        print("OK")

