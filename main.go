package main

import "fmt"

func main() {
	tasks := Tasks{}
	tasks.add("Buy A")
	tasks.add("Buy B")
	fmt.Printf("%+v\n\n", tasks)
	tasks.delete(1)
	fmt.Printf("%+v", tasks)
	tasks.add("AB")
	tasks.add("BC")
}
