var lengthOfLongestSubstringFast = function (s) {
  let longest = 0;
  let curr = 0;
  let map = new Map();
  if (s.length < 2) {
    return s.length;
  }
  for (let i = 0; i < s.length; i++) {
    if (map.get(s[i]) == null) {
      curr += 1;
    } else {
      curr = Math.min(i - map.get(s[i]), curr + 1);
    }
    longest = Math.max(longest, curr);
    map.set(s[i], i);
  }
  return longest;
};

console.log(lengthOfLongestSubstring("  "));
