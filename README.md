# DSA
Solving data structure and algorithm problems and adding notes I learned from attempting them

## Table of Contents
| id | Link To Problem  | Category | Difficulty | Source |
| -- | --  | -- | -- | -- |
| [0](#contains-duplicate) | [ContainsDuplicate](https://leetcode.com/problems/contains-duplicate/)  | Arrays | Easy | Blind 75 | 
| [1](#valid-anagram) | [ValidAnagram](https://leetcode.com/problems/valid-anagram/)  | Arrays | Easy | Blind 75 | 
| [2](#two-sum) | [TwoSum](https://leetcode.com/problems/two-sum/)  | Arrays | Easy | Blind 75 | 
| [3](#group-anagrams) | [GroupAnagrams](https://leetcode.com/problems/group-anagrams/)  | Arrays | Medium | Blind 75 | 


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