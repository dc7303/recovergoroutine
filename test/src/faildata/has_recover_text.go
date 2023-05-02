package faildata

func Run() {
	go func() {
		defer func() {
			// recover comment
		}()
	}()

	go func() {
		// recover variable
		var recover = 1
		foo(recover)
	}()

	func() {
		// not checked
	}()
}

func foo(_ int) {}
