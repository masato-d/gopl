package training5_19

func noReturn() (r string) {
	defer func() {
		p := recover()
		r = p.(string)
	}()
	panic("not zero value")
}