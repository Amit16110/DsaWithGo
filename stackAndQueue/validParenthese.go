package stackandqueue

func validParenthes(s string) bool {
	stack := NewStack()

	for _, v := range s {
		if v == '(' || v == '{' || v == '[' {
			stack.Push(v)
		} else {
			if !stack.IsEmpty() {
				top := stack.list.Back()

				if v == ')' && top.Value == '(' ||
					v == '}' && top.Value == '{' ||
					v == ']' && top.Value == '[' {
					stack.list.Remove(top)
				} else {
					return false
				}

			} else {
				return false
			}
		}

	}

	if stack.list.Len() == 0 {
		return true
	} else {

		return false
	}
}
