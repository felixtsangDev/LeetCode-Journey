var lengthOfLongestSubstring = function (s) {
  let str = "";
  let maxCount = 0;

  Array.from(s).forEach((c, i) => {
    if (str.indexOf(c) !== -1) {
      str = str.slice(str.indexOf(c) + 1);
    }
    str += c;

    if (str.length > maxCount) {
      maxCount = str.length;
    }
  });

  return maxCount;
};

console.log(
  `lengthOfLongestSubstring: ${lengthOfLongestSubstring("abcabcbb")}`
);
