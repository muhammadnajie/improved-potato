def isPalindromePermutation(phrase):
    temp_dict = {}

    for char in phrase:
        if char in temp_dict:
            del temp_dict[char]
        else:
            temp_dict[char] = True
    
    return len(temp_dict) <= 1

print(isPalindromePermutation('gigi'))
print(isPalindromePermutation('kkassrurua'))
print(isPalindromePermutation('pena'))