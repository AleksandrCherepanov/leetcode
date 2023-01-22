def largestAltitude(gain):
    largest = max(0, gain[0])

    for i in range(1, len(gain)):
        gain[i] += gain[i - 1]
        largest = max(largest, gain[i])

    return largest

test = [
    [
        [-5,1,5,0,-7],
        1
    ],
    [
        [-4,-3,-2,-1,4,3,2],
        0
    ],
    [
        [52,-91,72],
        52
    ]
]

for i in test:
    result = largestAltitude(i[0])

    if result != i[-1]:
        print("Error. ", result, i[-1])
    else:
        print("OK")

