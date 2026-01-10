import pytest
from group_anagrams import Solution

@pytest.mark.parametrize(
    "words, expected",
    [
        # Basic example
        (["eat","tea","tan","ate","nat","bat"], [["bat"],["nat","tan"],["ate","eat","tea"]]),

    ]
)

def test_group_anagrams(words, expected):
    sol = Solution()
    assert sorted(sol.group_anagrams(words)) == sorted(expected)
