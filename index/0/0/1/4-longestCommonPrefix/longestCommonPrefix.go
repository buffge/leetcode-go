package main

import "log"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	res := make([]byte, 0, 1)
	begin := strs[0]
	for i := 0; i < len(begin); i++ {
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || begin[i] != strs[j][i] {
				goto end
			}
		}
		res = append(res, begin[i])
	}
end:
	return string(res)
}
func longestCommonPrefixV2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	res := strs[0]
	for i := 1; i < len(strs); i++ {
		j := 0 // 有几个字符相同,这里要默认设置为0个相同
		for ; j < len(res) && j < len(strs[i]) && strs[i][j] == res[j]; j++ {
		}
		res = res[:j]
	}
	return res
}
func main() {
	log.Printf("%q", longestCommonPrefix([]string{"flower", "flow", "flight"}))
	log.Printf("%q", longestCommonPrefixV2([]string{"flower", "flow", "flight"}))
	log.Printf("%q", longestCommonPrefixV2([]string{"aa", "a"}))
}
