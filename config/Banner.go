package config

import "fmt"

func ShowBanner() {
	banner := `_____/\\\\\\\\\\\\__/\\\\\\\\\\\\\\\\_
	___/\\\\//////////__\\///////\\\\/////__
	 __/\\\\______________________\\/\\\\_______
	  _\\/\\\\____/\\\\\\\\________\\/\\\\_______
	   _\\/\\\\___\\/////\\\\_______\\/\\\\_______
        _\\/\\\\_______\\/\\\\___\\/\\\\_______
		 _\\/\\\\_______\\/\\\\___\\/\\\\_______
		  _\\//\\\\\\\\\\\\/_______\\/\\\\_______
		   __\\////////////_________\\///________
   Author: kk
   Version: ` + ShowVersion() + `
   `

	fmt.Printf(banner)
}

func ShowVersion() string {
	return "0.0.1"
}