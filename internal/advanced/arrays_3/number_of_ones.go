package arrays3

/*
Problem Description
Given an integer A, find and return the total number of digit 1 appearing in all non-negative integers less than or equal to A.

Problem Constraints
0 <= A <= 109

Input Format
The only argument given is the integer A.

Output Format
Return the total number of digit 1 appearing in all non-negative integers less than or equal to A.

Example Input
Input 1:
A = 10

Input 2:
A = 11

Example Output
Output 1:
2

Output 2:
4

Example Explanation
Explanation 1:
Digit 1 appears in 1 and 10 only and appears one time in each. So the answer is 2.

Explanation 2:
Digit 1 appears in 1(1 time) , 10(1 time) and 11(2 times) only. So the answer is 4.
*/

/*
Approach:
For a number N=3195 and divisor=10,
we will start with splitting the number into two parts:
quotient = N/divisor = (3195/10) = 319
digit = (N%divisor) / (divisor/10) = (3195%10) / (10/10) = 5

If we identify the pattern 01, 11, 21, 31......3181, there are 319 1s at unit place
If the given number had been 3190, this would have been our count for units place.
But since, N is 3195 and the digit 5>1, total ones at unit place would be 319+1=320

Now divisor=100
quotient = N/divisor = (3195/100) = 31
digit = (N%divisor) / (divisor/10) = (3195%100) / (100/10) = 9

If we identify the pattern 010-019, 110-119, 210-219.......3010-3019, each set has 10 1s and there are 31 such sets, 
so total 1s at tens place is 31*10=310
If the given number had been 3105, this would have been our count for tens place.
But since, N is 3195 and the digit 9>1, total ones at unit place would be 310+10=320

Now divisor=1000
quotient = N/divisor = (3195/1000) = 3
digit = (N%divisor) / (divisor/10) = (3195%1000) / (1000/10) = 1

If we identify the pattern 0100-0199, 1100-1199, 2100-2199, each set has 100 1s and there are 3 such sets, 
so total 1s at hundreds place is 3*100=300
If the given number had been 3095, this would have been our count for hundreds place.
Now the digit here is 1, so we cant add another 100 like we did in previous cases.
For this situtation, we will take the number which is right to the digit (95 in our case) and add 1 to it (95+1=96).
Proof: 3100-3195 has 96 1s at hundreds place.
Total ones at hundreds place is 300+96 = 396

Now divisor=10000
quotient = N/divisor = (3195/10000) = 0
digit = (N%divisor) / (divisor/10) = (3195%10000) / (10000/10) = 3

Since N/divisor has been 0, we would have to see if there are 1s at thousands place by looking at the digit.
The digit 3>1, total ones at thousands place would be 0+1000=1000

ANS=320+320+396+1000=2036
*/
func GetOneCount(A int) int {
	oneCount := 0 //count of ones so far
	divisor := 10
	quotient := A
	for quotient > 0 {
		subDivisor := divisor / 10
		quotient = A / divisor
		oneCount += quotient * subDivisor
		digit := (A % divisor) / subDivisor
		if digit > 1 {
			oneCount += subDivisor
		} else if digit == 1 {
			oneCount += (((A % divisor) % subDivisor) + 1)
		}
		divisor *= 10
	}
	return oneCount
}
