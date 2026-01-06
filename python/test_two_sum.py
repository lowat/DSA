import pytest
from two_sum import Solution

@pytest.mark.parametrize(
    "nums, target, expected",
    [
        # Basic example
        ([2, 7, 11, 15], 9, [0, 1]),

        # Numbers out of order
        ([3, 2, 4], 6, [1, 2]),

        # Duplicate values allowed (but different indices)
        ([3, 3], 6, [0, 1]),

        # Negative numbers
        ([-1, -2, -3, -4, -5], -8, [2, 4]),

        # Mixed positive and negative
        ([-3, 4, 3, 90], 0, [0, 2]),

        # Zero involved
        ([0, 4, 3, 0], 0, [0, 3]),

        # Larger array
        ([1, 5, 7, 2, 8, 11], 9, [2, 3]),
    ]
)

def test_is_anagram(nums, target, expected):
    sol = Solution()
    assert sorted(sol.two_sum(nums, target)) == sorted(expected)
