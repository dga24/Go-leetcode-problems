package main

func main() {
	n := 5
	//bad := 4
	firstBadVersion(n)
}

func firstBadVersion(n int) int {
	low := 1
	high := n
	for low < high {
		var mid = low + (high-low)/2
		if isBadVersion(mid) {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return low
}

func isBadVersion(version int) bool