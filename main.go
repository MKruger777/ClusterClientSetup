package main

import (
	"fmt"
	"runtime"
	"./src"
)

func main() {

	//the thing to notice here is that go does not require any break command in the case that was hit.
	//when the condition is met, the code will execute and it will exit the case

	//the default fall through - todo
	fmt.Print("Starting cluster CLI setup ...\n")
	fmt.Print("Host operating system is ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
		//do the Mac setup
		src.MacSetup()

	case "linux":
		fmt.Println("Linux.")
		//do the linux setup
		src.LinuxSetup()

	case "windows":
		fmt.Println("Windows.")
		//do the linux setup

	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
		fmt.Printf("My sincere appologies I cannot cope with this! Only Linux, Mac and Windows Operating systems are supported!")
	}
}

