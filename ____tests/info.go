package main

import "fmt"

func main() {
	fmt.Println(
		"\033[36m",
`
      _____                    _____                    _____                _____                    _____          
     /\    \                  /\    \                  /\    \              /\    \                  /\    \         
    /::\    \                /::\    \                /::\    \            /::\    \                /::\    \        
    \:::\    \              /::::\    \              /::::\    \           \:::\    \              /::::\    \       
     \:::\    \            /::::::\    \            /::::::\    \           \:::\    \            /::::::\    \      
      \:::\    \          /:::/\:::\    \          /:::/\:::\    \           \:::\    \          /:::/\:::\    \     
       \:::\    \        /:::/__\:::\    \        /:::/__\:::\    \           \:::\    \        /:::/__\:::\    \    
       /::::\    \      /::::\   \:::\    \       \:::\   \:::\    \          /::::\    \       \:::\   \:::\    \   
      /::::::\    \    /::::::\   \:::\    \    ___\:::\   \:::\    \        /::::::\    \    ___\:::\   \:::\    \  
     /:::/\:::\    \  /:::/\:::\   \:::\    \  /\   \:::\   \:::\    \      /:::/\:::\    \  /\   \:::\   \:::\    \ 
    /:::/  \:::\____\/:::/__\:::\   \:::\____\/::\   \:::\   \:::\____\    /:::/  \:::\____\/::\   \:::\   \:::\____\
   /:::/    \::/    /\:::\   \:::\   \::/    /\:::\   \:::\   \::/    /   /:::/    \::/    /\:::\   \:::\   \::/    /
  /:::/    / \/____/  \:::\   \:::\   \/____/  \:::\   \:::\   \/____/   /:::/    / \/____/  \:::\   \:::\   \/____/ 
 /:::/    /            \:::\   \:::\    \       \:::\   \:::\    \      /:::/    /            \:::\   \:::\    \     
/:::/    /              \:::\   \:::\____\       \:::\   \:::\____\    /:::/    /              \:::\   \:::\____\    
\::/    /                \:::\   \::/    /        \:::\  /:::/    /    \::/    /                \:::\  /:::/    /    
 \/____/                  \:::\   \/____/          \:::\/:::/    /      \/____/                  \:::\/:::/    /     
                           \:::\    \               \::::::/    /                                 \::::::/    /      
                            \:::\____\               \::::/    /                                   \::::/    /       
                             \::/    /                \::/    /                                     \::/    /        
                              \/____/                  \/____/                                       \/____/         
                                                                                                                     
`,
		"\033[0m",
	)
	fmt.Println("\033[31mTEST ARE GONNA CLEAN DB, ARE U SURE?\033[0m")
}
