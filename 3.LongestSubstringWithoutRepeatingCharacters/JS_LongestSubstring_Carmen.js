var lengthOfLongestSubstring = function (s) {
	let str = "";
	let maxCount = 0;

	Array.from(s).forEach((c, i) => {
		let tempIndex = str.indexOf(c);
		if (tempIndex !== -1) {
			str = str.slice(tempIndex + 1);
		}
		str += c;

		if (str.length > maxCount) {
			maxCount = str.length;
		}
	});

	return maxCount;
};

console.log(`lengthOfLongestSubstring: ${lengthOfLongestSubstring("abcabcbb")}`);
