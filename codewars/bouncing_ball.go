package codewars

// Question: https://www.codewars.com/kata/5544c7a5cb454edb3c000047/train/go

func isValidExperiment(height, bounce, window float64) bool {
	return height > 0 && (bounce > 0 && bounce < 1) && window < height
}

func BouncingBall(h, bounce, window float64) int {

	if !isValidExperiment(h, bounce, window) {
		return -1
	}

	count := 1

	for isValidExperiment(h, bounce, window) {
		h = h * bounce

		if window < h {
			count += 2
		}
	}

	return count
}
