package main

import(
	"fmt"
	"os"
)

func main() {
	sudoUser := os.Getenv("SUDO_USER")
	if sudoUser == "" {
		fmt.Println("It's a weird tree")
	} else {
		fmt.Println("     _              __")
    fmt.Println("    / `\\  (~._    ./  )")
    fmt.Println("    \\__/ __`-_\\__/ ./")
    fmt.Println("   _ \\ \\/  \\   \\ |_   __")
    fmt.Println(" (   )  \\__/ -^    \\ /  \\")
    fmt.Println("  \\_/ \"  \\  | o  o  |.. /  __")
    fmt.Println("       \\. --' ====  /  || /  \\ ")
    fmt.Println("         \\   .  .  |---__.\\__/")
    fmt.Println("         /  :     /   |   |")
    fmt.Println("         /   :   /     \\_/")
    fmt.Println("      --/ ::    (")
    fmt.Println("     (  |     (  (____")
    fmt.Println("   .--  .. ----**.____)")
    fmt.Println("   \\___/          ")
	}
}
