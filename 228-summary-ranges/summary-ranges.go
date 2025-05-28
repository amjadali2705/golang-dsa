import "strconv"

func summaryRanges(nums []int) []string {
    ranges := []string{}

    i := 0
    for i < len(nums) {
        start := nums[i]
        j := i
        for j+1 < len(nums) && nums[j+1] == nums[j]+1 {
            j++
        }

        end := nums[j]

        if start == end {
            ranges = append(ranges, strconv.Itoa(start))
        } else {
            ranges = append(ranges, strconv.Itoa(start)+"->"+strconv.Itoa(end))
        }
        i = j+1
    }

    return ranges
}