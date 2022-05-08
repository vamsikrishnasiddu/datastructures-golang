package main

type Stack struct {
	items []byte
}

func (s *Stack) Push(data byte) {
	s.items = append(s.items, data)
}

func (s *Stack) Pop() {
	s.items = s.items[:len(s.items)-1]
}

func (s *Stack) Top() byte {

	return s.items[len(s.items)-1]

}

func (s *Stack) Empty() bool {
	if len(s.items) == 0 {
		return true
	}

	return false
}

func isValid(s string) bool {

	var st Stack

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			st.Push(s[i])
		} else {
			if st.Empty() {
				return false
			}

			ch := st.Top()

			if ch == '(' && s[i] == ')' || ch == '[' && s[i] == ']' || ch == '{' && s[i] == '}' {
				st.Pop()

			} else {
				return false
			}
		}
	}

	if st.Empty() {
		return true
	}

	return false
}

func main() {

}
