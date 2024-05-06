package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://raw.githubusercontent.com/7-solutions/backend-challenge/main/files/hard.json"

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data from URL:", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {

		body, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return
		}
		// body := "[[59], [73, 41], [52, 40, 53], [26, 53, 6, 34]]"

		var arr [][]int
		_ = json.Unmarshal([]byte(body), &arr)

		fmt.Println(noRecur(arr))
		fmt.Println("no?")

		// for _, sublist := range arr {
		// 	for _, element := range sublist {
		// 		fmt.Print(element)
		// 	}
		// }
	} else {
		fmt.Println("Failed to fetch data from the URL. Status code:", response.StatusCode)
	}

}

// func recurs(arr [][]int, index int) int {
// 	if (len(arr)) == 1 {
// 		return arr[0][index]
// 	}
// 	fmt.Println("accessing")
// 	return max(recurs(arr[1:], index), recurs(arr[1:], index+1)) + arr[0][index]

// }
func noRecur(arr [][]int) int {
	n := len(arr)
	dp := make([]int, n)
	copy(dp, arr[n-1])
	for i := n - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			dp[j] = max(dp[j], dp[j+1]) + arr[i][j]
		}
	}
	return dp[0]
}
