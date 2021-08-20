package solutions

func numDecodings(s string) int {
    dp := make([]int,len(s)+1) //Memory
    dp[0] = 1
    dp[1] = decodeOneDigit(string(s[0])) 
    
    for i:=2; i<=len(s); i++{
        current := string(s[i-1])
        prev := string(s[i-2])
        dp[i] = dp[i] + (dp[i-1]*decodeOneDigit(current))
        dp[i] = dp[i] + (dp[i-2]*decodeTwoDigits(prev,current))
        
        dp[i]= dp[i]%1000000007
    }
    return dp[len(s)]
}

func decodeOneDigit(digit string) int{
    if digit == "*" {
        return 9
    }else if digit == "0"{
        return 0
    }
    return 1
}

func decodeTwoDigits(first,second string) int{
    switch first {
        case "*" :{
            if second == "*" {
                return 15
            }else if second >= "0" && second <= "6" {
                return 2
            }else {
                return 1
            }
        }
        case "1" :{
            if second == "*" {
                return 9
            }else {
                return 1
            }
        }
        case "2" :{
            if second == "*" {
                return 6
            }else if second >= "0" && second <= "6" {
                return 1
            }else {
                return 0
            }
        }
    }
    return 0
}