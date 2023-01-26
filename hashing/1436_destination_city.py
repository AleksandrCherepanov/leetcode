from collections import defaultdict

def destCity(paths):
    fromCities = defaultdict(int)
    destCities = defaultdict(int)

    for p in paths:
        fromCities[p[0]] = 1
        destCities[p[1]] = 1
    
    for k, d in destCities.items():
        if k not in fromCities:
            return k

    return ""

test = [
    [
        [["London","New York"],["New York","Lima"],["Lima","Sao Paulo"]],
        "Sao Paulo" 
    ],
    [
        [["B","C"],["D","B"],["C","A"]],
        "A"
    ],
    [
        [["A","Z"]],
        "Z"
    ]
]

for i in test:
    result = destCity(i[0])

    if result != i[-1]:
        print("Error. ", result, i[-1])
    else:
        print("OK")

