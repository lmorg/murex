package typemgmt_test

/*func TestScoping(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block: `set foobar=1
					out $foobar`,
			Stdout: "1\n",
		},
		{
			Block: `function scope-test {
						set foobar=1
						out $foobar
					}
					scope-test`,
			Stdout: "1\n",
		},
		{
			Block: `function scope-test {
						set foobar=2
					}
					set foobar=1
					scope-test
					out $foobar`,
			Stdout: "1\n",
		},
		{
			Block: `function scope-test {
						out $foobar
					}
					set foobar=1
					scope-test`,
			Stdout: "1\n",
		},
	}

	test.RunMurexTests(tests, t)
}*/
