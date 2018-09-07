package insert

import (
	"log"
	"encoding/json"
	"fmt"
)

type Test struct {
	Name   string `json:"name"`
	Id     int    `json:"id"`
	Second string `json:"second"`
}
type Cry struct {
	Test Test
}

func main() {

	/*	data, err := ioutil.ReadFile("./test.json")
		if err != nil {
			log.Println(err)
		}*/

	str := `{"name":"liliangbin","id":2,"second":"index"}`

	var test Test
	errr := json.Unmarshal([]byte(str), &test)
	if errr != nil {
		log.Println(errr)
	}

	fmt.Println(test)

}
