def reverseWords(s):
    result = []
    start = 0
    i = 0
    for i in range(0, len(s)):
        if s[i] == " ":
            end = i - 1
            while end >= start:
                result.append(s[end])
                end -= 1

            result.append(s[i])
            start = i + 1
  
    end = i
    while end >= start:
        result.append(s[end])
        end -= 1

    return ''.join(result)

print(reverseWords("Let's take LeetCode contest"))
print(reverseWords("God Ding"))

