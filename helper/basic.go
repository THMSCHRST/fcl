package helper

//The most basic helper functions for general use

func Check(err any) {
	if err != nil {
		panic(err)
	}
}
