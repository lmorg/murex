package crash

/*
	We want tests to panic!

	So this weird source file disabled the crash handler for tests
*/

func init() {
	disable_handler = true
}
