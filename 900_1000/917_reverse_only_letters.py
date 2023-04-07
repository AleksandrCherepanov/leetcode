def reverseOnlyLetters(s):
    result = list(s)

    i = 0
    j = len(s) - 1
    while i < j:
        l = (ord(s[i]) >= 65 and ord(s[i]) <= 90) or (ord(s[i]) >= 97 and ord(s[i]) <= 122)
        r = (ord(s[j]) >= 65 and ord(s[j]) <= 90) or (ord(s[j]) >= 97 and ord(s[j]) <= 122)
        
        if l and r:
            tmp = result[i]
            result[i] = result[j]
            result[j] = tmp

            j -= 1
            i += 1
            continue

        if l:
            j -= 1
        elif r:
            i += 1
        else:
            j -= 1
            i += 1

    return ''.join(result)

data = ["ab-cd", "a-bC-dEf-ghIj", "Test1ng-Leet=code-Q!", "7_28]"]
exp = ["dc-ba", "j-Ih-gfE-dCba", "Qedo1ct-eeLg=ntse-T!", "7_28]"]

for i in range(0, len(data)):
    result = reverseOnlyLetters(data[i])
    if exp[i] != result:
        print("Error!", exp[i], result)
    else:
        print("OK")

