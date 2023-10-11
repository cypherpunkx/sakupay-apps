package common

func CountTotalPage(total, perPage int) int {
	if (total % perPage) > 0 {
		return (total / perPage) + 1
	}

	return total / perPage
}
