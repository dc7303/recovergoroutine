package succdata

func whenLit() {
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
		rec := func() {
			defer func() {
				recover()
			}()
		}

		defer rec()
	}()
}

func whenIdent() {
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
	defer func() {
		recover()
	}()
}

func nestedFunc2() {}
