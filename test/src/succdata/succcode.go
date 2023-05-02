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
	go nestedFunc1()
}

func runGoroutine() {
	defer func() {
		recover()
	}()
}

func nestedFunc1() {
	// must have recover in parent caller
	nestedFunc2()
	recover()
}

func nestedFunc2() {}
