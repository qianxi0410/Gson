package tokenizer

type rs struct {
	val rune
	next *rs
}

type reader struct {
	// a head of list with no-value
	head 	*rs
	cur 	*rs
}

func (r *reader) hasMore() bool {
	return r.cur != nil
}

// read peek value and go next
func (r *reader) read() rune {
	res := r.cur.val
	r.cur = r.cur.next
	return res
}

// just return peek value
func (r *reader) peek() rune {
	return r.cur.val
}

