// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: First Turn Winner
//
//  If the player wins on the first turn, then display
//  a special bonus message.
//
// RESTRICTION
//  The first picked random number by the computer should
//  match the player's guess.
//
// EXAMPLE SCENARIO
//  1. The player guesses 6
//  2. The computer picks a random number and it happens
//     to be 6
//  3. Terminate the game and print the bonus message
// ---------------------------------------------------------
func main() {
//	var (
//		maxTurn int
//		err     error
//		guess   int
//		ranLimit int
//		winMessage []string
//		loseMessage []string
//	)
//
//	winMessage = append(winMessage, "You won somehow","Lucky you","Too easy","Last Luck")
//	loseMessage = append(loseMessage, "You lose somehow","Too bad","Failure","Better Luck Next Time")
//	args := os.Args[1:]
//	query := args[2:]
//if len(args) < 2{
//	fmt.Printf("Input all 2 value\n")
//	return
//}
//maxTurn,err = strconv.Atoi(args[0])
//if err != nil{
//	fmt.Println("Max Turns not a number.")
//	return
//}
//guess,err = strconv.Atoi(args[1])
//if err != nil{
//	fmt.Println("Guess Not a number.")
//	return
//}
//
//if guess < 0{
//	fmt.Println("Please pick up a positive number")
//	return
//}
//
//if guess >= 10{
//	ranLimit = guess + 1
//}else{
//	ranLimit = 10
//}
//
//t := time.Now()
//rand.Seed(t.UnixNano())
//var result []string
//tempWin := winMessage[:]
//for _, q := range query {
//	for  i, w := range tempWin{
//		if strings.Contains(strings.ToLower(w),strings.ToLower(q)){
//			result = append(result, w)
//			tempWin[i-1] = tempWin[len(tempWin) - 1]
//			tempWin = tempWin[:len(tempWin) - 1]
//		}
//	}
//}
//
//if len(result) == 0 {
//	fmt.Println("Cant find word")
//}else{
//	fmt.Println("List Found Word :")
//	for _, s := range result {
//		fmt.Printf("-%2s\n",s)
//	}
//}
//
//for turn:= 0 ; turn < maxTurn; turn ++{
//	n := rand.Intn(ranLimit)
//	m := rand.Intn(ranLimit)
//	fmt.Println(n,m)
//	switch guess {
//	case n,m:
//		if turn == 0{
//			fmt.Println(winMessage[rand.Intn(4)])
//		}
//		fmt.Println(winMessage[rand.Intn(4)])
//		return
//	default:
//		continue
//	}
//}
//fmt.Println(loseMessage[rand.Intn(4)])
var list []int
list=append(list,2,7,11,15)
//fmt.Println(maxProfit(list))
//fmt.Println(maxProfit2(list))
//fmt.Println(getMinDistance3(list,11,2))
//	fmt.Println(twoSum(list,9))
//	res := addTwoNumbers(
//		&ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{}}}},
//		&ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: &ListNode{}}}},
//	)
//	fmt.Println(res.Val,res.Next.Val,res.Next.Next.Val)
//	fmt.Println(uniquePaths(3,7))
s := "abcabcbb"
fmt.Println(lengthOfLongestSubstring(s))
}


func maxProfit(prices []int) int {
	var (
		maxProfit = 0.0
		j = 0
	)
	for i,p := range prices{
		if p < prices[j] {
			j = i
		}else{
			maxProfit = math.Max(maxProfit,float64(p-prices[j]))
		}
	}
	return int(maxProfit)
}

func maxProfit2(prices []int) int {
	var (
		maxProfit int
		minBuy = prices[0]
		maxBuy = prices[0]
	)
	for i := 0; i < len(prices) - 1;{
		for i < len(prices) - 1 && prices[i] >= prices[i+1] {
			i++
		}
		minBuy = prices[i]
		for i < len(prices) - 1 && prices[i] <= prices[i+1] {
			i++
		}
		maxBuy = prices[i]
		maxProfit += maxBuy - minBuy
	}
	return maxProfit
}

func getMinDistance(nums []int, target int, start int) int {
	var(
		occ []int
		l = len(nums)
		i int
	)
	for i < l{
		if target == nums[i]{
			occ = append(occ,int(math.Abs(float64(i - start))))
		}
		i++
	}
	sort.Ints(occ)
	return occ[0]
}

func getMinDistance2(nums []int, target int, start int) int {
	var(
		l = len(nums)
		u bool
		d bool
		i = start
		y = start
	)
	for ;i < l;i++{
		if nums[i] == target{
			u = !u
			break
		}
	}

	for ;y > 0;y --{
		if nums[y] == target{
			d = !d
			break
		}
	}

	if u && d {
		return int(math.Min(math.Abs(float64(i-start)),math.Abs(float64(y-start))))
	}else if u{
		return int(math.Abs(float64(i-start)))
	}else {
		return int(math.Abs(float64(y-start)))
	}
	return 0
}

func getMinDistance3(nums []int, target int, start int) int {
	var(
		c = math.MaxInt16
		i = 0
		l = len(nums)
	)
	for ;i < l;i++{
		if nums[i] == target{
			if c > int(math.Abs(float64(i - start))){
				c = int(math.Abs(float64(i - start)))
			}
		}
	}
	return c
}

func twoSum(nums []int, target int) []int {
	var (
		ln []int
		i int
		y int
		l = len(nums)
	)
	for ;i < l - 1;i++{
		for y= i + 1; y < l; y++{
			if nums[i] + nums[y] == target{
				ln = append(ln, i , y)
				break
			}
		}
	}

	return ln
}


type ListNode struct {
     Val int
     Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1==nil || l2==nil{
		return nil
	}
	dummyNode := &ListNode{}
	carryOn := 0
	curNode := dummyNode

	for l1!=nil || l2!=nil || carryOn>0{
		i1 , i2 :=0, 0
		if l1!=nil{
			i1=l1.Val
		}
		if l2!=nil{
			i2=l2.Val
		}
		//create the values
		tmp := i1 + i2 + carryOn
		carryOn = tmp/10
		curNode.Next=&ListNode{
			Val:tmp%10,
		}
		//set the next
		curNode = curNode.Next
		if l1!=nil{
			l1=l1.Next
		}
		if l2!=nil {
			l2=l2.Next
		}
	}
	return dummyNode.Next
}

func uniquePaths(m int, n int) int {
	return recursivePath(m,n,0,0)
}

func recursivePath(m int,n int, i int, j int) int{
	if i >= m || j >= n {
		return 0
	}
	if i == m -1 && j == n -1{
		return 1
	}
	return recursivePath(m,n,i+1,j) + recursivePath(m,n,i,j+1)
}

func uniquePaths2(m int, n int) int {
	var (
		ans int
		i int
		j int
	)
	for i,j = m+n-2,1; i >= int(math.Max(float64(m), float64(n))); i-- {
		ans = (ans * i) / j
		j++
	}
	return ans
}

func lengthOfLongestSubstring(s string) int {
	var (
		maxLen int
		startIndex int
	)
	b := []byte(s)
	for i,v := range b{
		for strings.Contains(string([]byte(b[startIndex:i])),string([]byte{v})){
			startIndex++
		}
		if i+1 - startIndex > maxLen {
			maxLen = i + 1 - startIndex
		}
	}
	return maxLen
}

