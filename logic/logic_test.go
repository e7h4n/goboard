package logic

func checkTestErr(err error) {
	if err != nil {
		panic(err)
	}
}
