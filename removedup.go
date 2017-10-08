package  main


// nums := {1,1,2}
import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for k := i + 1  ; k <  len(nums) ; k ++  {
		if nums[k] != nums[i] {
			i ++
			nums[i] = nums[k]
		}
	}
	return i+1
}
func main(){
	nums := []int{1,1,2,2,3,3,3,3}
	x := removeDuplicates(nums)
	fmt.Println(x)
}