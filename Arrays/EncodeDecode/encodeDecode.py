def encode(strings):
    res = []
    for string in strings:
        res.append(str(len(string)))
        res.append("#")
        res.append(string)
    return "".join(res)

def decode(string):
    res, left, i = [], 0, 0
    
    while i < len(string):
        c = string[i]
        if c != "#":
            i += 1
            continue
        length = int(string[left:i])
        res.append(string[i+1:i+length+1])
        i += 1 + length
        left = i
    return res

test = ["lint","code","love","you"]
res = encode(test)
print(res)
res2 = decode(res)
print(res2)