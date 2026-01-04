import pytest
from valid_anagram import Solution

@pytest.mark.parametrize(
    "s, t, expected",
    [
        # LeetCode examples
        ("anagram", "nagaram", True),
        ("rat", "car", False),

        # Edge cases
        ("", "", True),
        ("a", "a", True),
        ("a", "b", False),

        # Length mismatch
        ("a", "aa", False),
        ("abc", "ab", False),

        # Repeated characters
        ("aabb", "bbaa", True),
        ("aabb", "abbb", False),

        # Case sensitivity
        ("a", "A", False),

        # Classic anagram
        ("listen", "silent", True),
    ],
)

def test_is_anagram(s, t, expected):
    sol = Solution()
    assert sol.is_valid_anagram(s, t) == expected
