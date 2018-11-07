package selfutil

func Swap(lhs *string, rhs *string) {
	*lhs, *rhs = *rhs, *lhs
}

