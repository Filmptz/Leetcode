package solutions



import (
    "strconv"
	"fmt"
)
func numDecodings(s string) int {
    var dp = make([]int,len(s)+1)
    dp[0] = 1
    if string(s[0]) != "0" {
        dp[1] = 1
    }
    
    for i:=2; i<=len(s); i++ {
        oneDigit, err := strconv.Atoi(string(s[i-1])); 
        if err != nil {
            fmt.Println(err)
        }
        twoDigits, err := strconv.Atoi(string(s[i-2:i]));
         if err != nil {
            fmt.Println(err)
        }

        if(oneDigit >=1){
            dp[i] = dp[i]+dp[i-1]
        }
        
        if(twoDigits >=10 && twoDigits <= 26){
            dp[i] = dp[i]+dp[i-2]
        }
    }

    return dp[len(s)]
}