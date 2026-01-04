from contains_duplicate import Solution

def test_example_1():
    s = Solution()
    assert s.contains_duplicate([1,2,3,1]) == True

def test_example_2():
    s = Solution()
    assert s.contains_duplicate([1,2,3,4]) == False

def test_example_3():
    s = Solution()
    assert s.contains_duplicate([1,1,1,3,3,4,3,2,4,2]) == True

