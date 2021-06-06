def isAdj(all):
    for i in range(len(all)-1):
        if all[i] == all[i+1]:
            return True
    return False


def splitEachDigit(num):
    sp = [int(d) for d in str(num)]
    # print(sp)
    return sp


def findDepth(num):
    if isAdj(splitEachDigit(num)):
        return 0

    Q = []
    first = (num, 0)
    Q.append(first)
    while(1):
        curr = Q.pop(0)
        print(curr)
        currNum = curr[0]
        depth = curr[1]
        splitted = splitEachDigit(currNum)
        for multiplier in splitted:
            all = splitEachDigit(currNum) + splitEachDigit(currNum * multiplier)
            if isAdj(all):
                return depth + 1
            Q.append((currNum * multiplier, depth + 1))


if __name__ == '__main__':
    num = int(input())
    depth = findDepth(num)
    print(depth)