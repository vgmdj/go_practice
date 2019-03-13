package _383_Ransom_Note

func canConstruct(ransomNote string, magazine string) bool {
	count := [26]int{}
	for i := range magazine {
		count[magazine[i]-'a']++
	}

	for i := range ransomNote {
		count[ransomNote[i]-'a']--
		if count[ransomNote[i]-'a'] < 0 {
			return false
		}
	}

	return true

}
