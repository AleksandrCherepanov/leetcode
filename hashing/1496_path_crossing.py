from collections import defaultdict

def createKey(x, y):
    return x * 10000 + y

def isPathCrossing(path):
    p = list(path)
    prev = defaultdict(bool)
    prev[createKey(0,0)] = True
    
    x = y = 0
    for d in p:
        if d == "N":
            y += 1
        elif d == "S":
            y -= 1
        elif d == "E":
            x += 1
        else:
            x -= 1

        key = createKey(x, y)
        if key in prev:
            return True
        
        prev[key] = True

    return False

test = [
    [
        "NES",
        False
    ],
    [
        "NESWW",
        True
    ]
]

for i in test:
    result = isPathCrossing(i[0])

    if result != i[-1]:
        print("Error. ", result, i[-1])
    else:
        print("OK")

