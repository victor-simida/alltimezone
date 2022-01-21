package alltimezone

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	t1 := time.Unix(1640077010, 0).In(time.UTC)
	t2 := t1.In(time.FixedZone("GMT+05:45",
		5*3600+45*60))

	t3 := time.Unix(1640077010, 0).In(time.FixedZone("GMT+05:45",
		5*3600+45*60))
	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println(t3)

	fmt.Println(CommonTimezone)

	for k := range CommonTimezone {
		temp := strings.TrimPrefix(string(k), "GMT")

		key := string(temp[0]) // + or -
		temp = string(temp[1:])
		var hour, mintue int
		_, err := fmt.Sscanf(temp, "%02d:%02d", &hour, &mintue)
		if err != nil {
			panic(err.Error())
		}

		if key == "-" {
			fmt.Printf("\"%s\", %s(3600 * %v + 60 * %v)\n", k, key, hour,
				mintue)
		} else {
			fmt.Printf("\"%s\",  3600 * %v + 60 * %v\n", k, hour, mintue)
		}

	}

}

func TestGmtZone(t *testing.T) {
	target := time.Unix(1640044800, 0) // 2021-12-21-Tuesday 00-00-00

	assert.Equal(t, target.In(GetLocation("US/Aleutian")).Format(
		"2006-01-02 15:04:05"), "2021-12-20 14:00:00")

	assert.Equal(t, target.In(GetLocation("America/Cancun")).Format(
		"2006-01-02 15:04:05"), "2021-12-20 19:00:00")

	assert.Equal(t, target.In(GetLocation("Asia/Kathmandu")).Format(
		"2006-01-02 15:04:05"), "2021-12-21 05:45:00")

	assert.Equal(t, target.In(GetLocation("Asia/Seoul")).Format(
		"2006-01-02 15:04:05"), "2021-12-21 09:00:00")

}
