package succdata

func anonymousFuncCall() {
	go func() {
		defer func() {
			if r := recover(); r != nil {
			}
		}()
	}()

	go func() {
		defer func() {
			recover()
		}()
	}()

	go func() {
		recover()
	}()
}

func funcCall() {
	go runGoroutine()
}

func runGoroutine() {
	defer func() {
		recover()
	}()
}
