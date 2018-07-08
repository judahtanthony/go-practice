package practice

// Perm returns a slice of all possible permuations
// of the characters in an ascii encoded string.
func Perm(in string) []string {
	// allocate a buffer to work in
	buff := []byte(in)
	n := len(buff)
	if n == 0 {
		return []string{}
	} else if n == 1 {
		return []string{string(buff)}
	}

	store := []string{}

	var p func(n int)
	p = func(n int) {
		if n == 1 {
			store = append(store, string(buff))
		} else {
			for i := 0; i < n; i++ {
				p(n - 1)
				if n%2 == 1 {
					buff[n-1], buff[i] = buff[i], buff[n-1]
				} else {
					buff[n-1], buff[0] = buff[0], buff[n-1]
				}
			}
		}
	}

	p(n)

	return store
}

func PermChan(in string) <-chan string {
	// allocate a buffer to work in
	buff := []byte(in)
	n := len(buff)
	store := make(chan string)
	if n == 0 {
		close(store)
		return store
	}

	var p func(n int)
	p = func(n int) {
		if n == 1 {
			store <- string(buff)
		} else {
			for i := 0; i < n; i++ {
				p(n - 1)
				if n%2 == 1 {
					buff[n-1], buff[i] = buff[i], buff[n-1]
				} else {
					buff[n-1], buff[0] = buff[0], buff[n-1]
				}
			}
		}
	}

	go func() {
		p(n)
		close(store)
	}()

	return store
}
