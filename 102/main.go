package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	s := `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus ut ex suscipit, faucibus nibh a, faucibus dui. Nunc vitae eleifend enim, sed pharetra nisi. Quisque et felis sagittis, viverra nisl ultrices, vehicula tellus. Donec quis rutrum leo, ut ullamcorper eros. Cras porttitor gravida dolor, sed ultricies nulla vehicula sed. Morbi vitae est ullamcorper, lobortis purus a, euismod diam. Curabitur vehicula luctus felis, ut porta odio volutpat id. Sed eu dolor vel est vulputate vulputate laoreet non nunc. Nullam nec neque lorem. Proin eget felis vitae ipsum efficitur eleifend sagittis quis lacus.`

	encodeStd := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

	s64 := base64.NewEncoding(encodeStd).EncodeToString([]byte(s))

	fmt.Println(len(s))
	fmt.Println(len(s64))
	fmt.Println(s)
	fmt.Println(s64)
}
