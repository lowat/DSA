# DSA
Solving data structure and algorithm problems and adding notes I learned from attempting them

## Table of Contents
| id | Link To Problem  | Category | Difficulty | Source |
| -- | --  | -- | -- | -- |
| [0](#contains-duplicate) | [ContainsDuplicate](https://leetcode.com/problems/contains-duplicate/)  | Arrays | Easy | Blind 75 | 
| [1](#valid-anagram) | [ValidAnagram](https://leetcode.com/problems/valid-anagram/)  | Arrays | Easy | Blind 75 | 
| [2](#two-sum) | [TwoSum](https://leetcode.com/problems/two-sum/)  | Arrays | Easy | Blind 75 | 
| [3](#group-anagrams) | [GroupAnagrams](https://leetcode.com/problems/group-anagrams/)  | Arrays | Medium | Blind 75 | 
| [4](#top-k-freq-elements) | [TopKFreqElements](https://leetcode.com/problems/top-k-frequent-elements/)  | Arrays | Medium | Blind 75 | 
| [5](#product-of-array-except-self) | [ProductOfArrayExceptSelf](https://leetcode.com/problems/product-of-array-except-self/)  | Arrays | Medium | Blind 75 | 
| [6](#encode-decode) | [EncodeDecode](https://leetcode.com/problems/encode-and-decode-strings/)  | Arrays | Medium | Blind 75 | 
| [7](#longest-consecutive-sequence) | [LongestConsecSeq](https://leetcode.com/problems/longest-consecutive-sequence/)   | Arrays | Medium | Blind 75 | 
| [8](#valid-palindrome) | [ValidPalindrome](https://leetcode.com/problems/valid-palindrome/)   | Two Pointers | Easy | Blind 75 | 

### Contains Duplicate
  - **Link To Problem**: [ContainsDuplicate](https://leetcode.com/problems/3sum/) 
  - **Category**: Arrays
  - **Difficulty**: Easy
  - **Source**: Blind 75
  - **Problem Notes**:
    - **Input**: List of integers
    - **Ouput**: Boolean representing if list contains duplicate numbers
    - **Clarifying Questions**:
  - **Solution Notes**:
    - **Steps**
      1. Maintain a hashset and add each element to hashset
      2. Return true if hashset already contains value
    - **Time Complexity**: O(n)
    - **Space Complexity**: O(n)
    - **Code**
      -  [Python](https://github.com/lowat/DSA/blob/main/Arrays/ContainsDuplicate/containsDuplicate.py)
      -  [Go](https://github.com/lowat/DSA/blob/main/Arrays/ContainsDuplicate/containsDuplicate.go)

### Valid Anagram
  - **Link To Problem**: [ValidAnagram](https://leetcode.com/problems/valid-anagram/) 
  - **Category**: Arrays
  - **Difficulty**: Easy
  - **Source**: Blind 75
  - **Problem Notes**:
    - **Input**: 2 strings
    - **Ouput**: Boolean representing if strings are anagrams of each other
    - **Clarifying Questions**:
  - **Solution Notes**:
    - **Steps**
      1. One approach is to make 2 arrays of 26 and iterate each index when you come across that char
      2. Other approach is to use a hashmap and build for each string, and then iterate through k,v and ensure they match. Also make sure lengths are same otherwise you might be missing a char that is on second hashmap
    - **Time Complexity**: O(n)
    - **Space Complexity**: O(1), bounded to length 26 if using char count array
    - **Code**
      -  [Python](https://github.com/lowat/DSA/blob/main/Arrays/ValidAnagram/validAnagram.py)
      -  [Go](https://github.com/lowat/DSA/blob/main/Arrays/ValidAnagram/validAnagram.go)

### Two Sum
  - **Link To Problem**: [TwoSum](https://leetcode.com/problems/two-sum/) 
  - **Category**: Arrays
  - **Difficulty**: Easy
  - **Source**: Blind 75
  - **Problem Notes**:
    - **Input**: Array of integers, Target Integer
    - **Ouput**: Boolean representing if any two numbers add up to target integer
    - **Clarifying Questions**:
  - **Solution Notes**:
    - **Steps**
      1. Maintain a dictionary of prev values as key and their index as the values
      2. Iterate through remaining values and check if target - current val matches something in the dict
      3. If yes then return index of other num and current index
      4. If no then add current val and index to dict
    - **Time Complexity**: O(n)
    - **Space Complexity**: O(n)
    - **Code**
      -  [Python](https://github.com/lowat/DSA/blob/main/Arrays/twoSum/twoSum.py)
      -  [Go](https://github.com/lowat/DSA/blob/main/Arrays/twoSum/twoSum.go)

### Group Anagrams
  - **Link To Problem**: [Group Anagrams](https://leetcode.com/problems/group-anagrams/) 
  - **Category**: Arrays
  - **Difficulty**: Medium
  - **Source**: Blind 75
  - **Problem Notes**:
    - **Input**: Array of Strings
    - **Ouput**: 2D Array where each index of result is an array of words that are anagrams of each other
    - **Clarifying Questions**:
  - **Solution Notes**:
    - **Steps**
      1. Brute force way is to sort each string 
      2. Better way is to take each string and determine the char count, since we are bounded to 26 lowercase letters this is more efficient than sorting
      3. Check if  tuple(char count array) maps to existing key and append string, else start a new list for that key
      4. Return dict values
    - **Time Complexity**: O(n * m), where n is number of words and m is longest word
    - **Space Complexity**: O(n), hash map would not contain more values then original number of words
    - **Code**
      -  [Python](https://github.com/lowat/DSA/blob/main/Arrays/GroupAnagrams/groupAnagrams.py)
      -  [Go](https://github.com/lowat/DSA/blob/main/Arrays/GroupAnagrams/groupAnagrams.go)


### Top K Freq Elements
  - **Link To Problem**: [TopKFreqElements](https://leetcode.com/problems/top-k-frequent-elements/)
  - **Category**: Arrays
  - **Difficulty**: Medium
  - **Source**: Blind 75
  - **Problem Notes**:
    - **Input**: Array of Intergers
    - **Ouput**: Array containing the K most freq elements
    - **Clarifying Questions**:
  - **Solution Notes**:
    - **Steps**
      1. Make hashmap of counts of all numbers in array
      2. Make an array of array of len(original array) and iterate through hash maps and add values to the inner array at that index. Index 5 of outer array would contain array of numbers that occurred 5 times for example.
      3. Build result while iterating backwards through freq array and adding values to res. Stop when arr length == k
    - **Time Complexity**: O(n)
    - **Space Complexity**: O(n)
    - **Code**
      -  [Python](https://github.com/lowat/DSA/blob/main/Arrays/TopKFreqElements/topKFreqElements.py)
      -  [Go](https://github.com/lowat/DSA/blob/main/Arrays/TopKFreqElements/topKFreqElements.go)

  
### Product of Array Except Self
  - **Link To Problem**: [ProductOfArrayExceptSelf](https://leetcode.com/problems/product-of-array-except-self/) 
  - **Category**: Arrays
  - **Difficulty**: Medium
  - **Source**: Blind 75
  - **Problem Notes**:
    - **Input**: Array of Integer
    - **Ouput**: Array of Integers where each index is the product of all other indexes except the current index
    - **Clarifying Questions**:
  - **Solution Notes**:
    - **Steps**
      1. Dynamic programming mindset
      2. Make res arr of [1]'s
      3. Iterate left to right and change each res value to the prevResValue * prevNum
      4. Iterate right to left and maintain a separate postfix agg variable 
      5. Res[i] *= postfix, postfix *= num[i]
    - **Time Complexity**: O(n)
    - **Space Complexity**: O(n)
    - **Code**
      -  [Python](https://github.com/lowat/DSA/blob/main/Arrays/ProdArrExceptSelf/prodArrExceptSelf.py)
      -  [Go](https://github.com/lowat/DSA/blob/main/Arrays/ProdArrExceptSelf/prodArrExceptSelf.go)

### Encode Decode
  - **Link To Problem**: [EncodeDecode](https://leetcode.com/problems/encode-and-decode-strings/) 
  - **Category**: Arrays
  - **Difficulty**: Medium
  - **Source**: Blind 75
  - **Problem Notes**:
    - **Input**: Array of Strings
    - **Ouput**: Array of Strings, Design an algorithm to encode a list of strings to a string. The encoded string is then sent over the network and is decoded back to the original list of strings.
    - **Clarifying Questions**:
  - **Solution Notes**:
    - **Steps**
      1. Len+#+str
      2. join res
      3. iterate until first #
      4. use pointers and slicing to grab len of str
      5. use slicing to grab str and append to result and then return"
    - **Time Complexity**: O(n)
    - **Space Complexity**: O(n)
    - **Code**
      -  [Python](https://github.com/lowat/DSA/blob/main/Arrays/EncodeDecode/encodeDecode.py)
      -  [Go](https://github.com/lowat/DSA/blob/main/Arrays/EncodeDecode/encodeDecode.go)

### Longest Consecutive Sequence
  - **Link To Problem**: [LongestConsecSeq](https://leetcode.com/problems/longest-consecutive-sequence/) 
  - **Category**: Arrays
  - **Difficulty**: Medium
  - **Source**: Blind 75
  - **Problem Notes**:
    - **Input**: Array of Ints
    - **Ouput**: Integer representing the longest sequence of consecutive nums
    - **Clarifying Questions**:
  - **Solution Notes**:
    - **Steps**
      1. Consider that if we make a set, any sequence beginning will not have n-1 in the set. otherwise it is part of an existing sequence
      2. If we do find a sequence start iterate while that num + the sequence length is in set
      3. When loop breaks just take max(maxsofar, currSeqLength)
    - **Time Complexity**: O(n)
    - **Space Complexity**: O(n)
    - **Code**
      -  [Python](https://github.com/lowat/DSA/blob/main/Arrays/LongestConsecSeq/LongestConsecSeq.py)
      -  [Go](https://github.com/lowat/DSA/blob/main/Arrays/LongestConsecSeq/LongestConsecSeq.go)

### Valid Palindrome
  - **Link To Problem**: [ValidPalindrome](https://leetcode.com/problems/valid-palindrome/)
  - **Category**: Two Pointers
  - **Difficulty**: Easy
  - **Source**: Blind 75
  - **Problem Notes**:
    - **Input**: String
    - **Ouput**: True/False depending on if string is valid palindrome. Need to remove non alpha numeric characters.
    - **Clarifying Questions**:
  - **Solution Notes**:
    - **Steps**
      1. Clean string
      2. l, r pointer converging to middle. If they dont equal then it is not palindrome
    - **Time Complexity**: O(n)
    - **Space Complexity**: O(1)
    - **Code**
      -  [Python](https://github.com/lowat/DSA/blob/main/TwoPointers/ValidPalindrome/ValidPalindrome.py)
      -  [Go](https://github.com/lowat/DSA/blob/main/TwoPointers/ValidPalindrome/ValidPalindrome.go)