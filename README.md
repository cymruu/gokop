# goKOP

goKOP is a wykop API wrapper

## Usage

```
package main
import (
	"fmt"

	"github.com/cymruu/gokop"
	"github.com/cymruu/gokop/v1"
	"github.com/cymruu/gokop/v1/models"
)

func main() {
	client := v1.CreateWykopV1API("apikey", "secret", "")
	entry := models.Entry{}
	err := client.MakeRequest("entries/index", entry, gokop.AddMethodParams([]string{"2018"}))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(entry)
}
```


