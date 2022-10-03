package prefixsum

func GetProductArrayAlternate(A []int) []int {
	productArray := make([]int, len(A))
	rightProduct := 1
	leftProdArray := getLeftExclusivePrefixProduct(A)
	for i := len(A) - 1; i >= 0; i-- {
		productArray[i] = leftProdArray[i] * rightProduct
		rightProduct = rightProduct * A[i]
	}
	return productArray
}

/*
for an input array [1,2,3,4,5]
below function will generate [1,1,2,6,24]
*/
func getLeftExclusivePrefixProduct(A []int) []int {
	leftProduct := 1
	var leftProdArray []int
	for i := range A {
		leftProdArray = append(leftProdArray, leftProduct)
		leftProduct = leftProduct * A[i]
	}
	return leftProdArray
}
