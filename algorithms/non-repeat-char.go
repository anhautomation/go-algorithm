package algorithms

func FirstNonRepeatedChar(s string) (byte) {
    charFrequency := make(map[byte]int)
    for i := 0; i < len(s); i++ {
        charFrequency[s[i]]++
    }
    for i := 0; i < len(s); i++ {
        if charFrequency[s[i]] == 1 {
            return s[i]
        }
    }
    return 0 
}