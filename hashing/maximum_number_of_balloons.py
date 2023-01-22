from collections import defaultdict

def maxNumberOfBalloons(text):
    balloon = {"b":0, "a":0, "l":0, "o":0, "n":0}
    for l in text:
        if l in balloon:
            balloon[l] += 1

    balloon["l"] = int(balloon["l"] / 2)
    balloon["o"] = int(balloon["o"] / 2)

    result = balloon["b"]
    for l, v in balloon.items():
        result = min(result, v)

    return result

test = [
    [
        "nlaebolko",
        1
    ],
    [
        "loonbalxballpoon",
        2
    ],
    [
        "leetcode",
        0
    ]
]

for i in test:
    result = maxNumberOfBalloons(i[0])

    if result != i[-1]:
        print("Error. ", result, i[-1])
    else:
        print("OK")

