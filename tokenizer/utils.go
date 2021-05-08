package tokenizer

func isUseless(r rune) bool {
	if r == ' ' || r == '\t' || r == '\n' {
		return true
	}
	return false
}


func isNotEndForNull(r rune) bool {
	if r != '}' && r != ']' && r != ',' {
		return true
	}
	return false
}

func isNotEndForBool(r rune) bool {
	if r != '}' && r != ']' && r != ',' {
		return true
	}
	return false
}

func isNotEndForNumber(r rune) bool {
	if r != '}' && r != ']' && r != ',' {
		return true
	}
	return false
}

// make string to reader
func convStr2Reader(str string) *reader {
	// a head of list with no-value
	head := &rs{}
	// a help-pointer
	p := head
	// the size of string's rune array
	for _, v := range str {
		if isUseless(v) {
			continue
		}

		t := &rs{
			val:  v,
		}
		p.next = t
		p = p.next
	}
	// help gc
	p = nil
	return &reader{
		head: head,
		cur: head.next,
	}
}


func getTokenList() *TokenList {
	tokens := &Token{}

	return &TokenList{
		tokens: tokens,
		cur: tokens,
	}
}