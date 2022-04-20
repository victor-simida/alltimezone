# alltimezone
A golang library to get all time.Locations in the world

```go
	target := time.Unix(1640044800, 0) // 2021-12-21-Tuesday 00-00-00

	assert.Equal(t, target.In(GetLocation("US/Aleutian")).Format(
		"2006-01-02 15:04:05"), "2021-12-20 14:00:00")

	assert.Equal(t, target.In(GetLocation("America/Cancun")).Format(
		"2006-01-02 15:04:05"), "2021-12-20 19:00:00")

	assert.Equal(t, target.In(GetLocation("Asia/Kathmandu")).Format(
		"2006-01-02 15:04:05"), "2021-12-21 05:45:00")

	assert.Equal(t, target.In(GetLocation("Asia/Seoul")).Format(
		"2006-01-02 15:04:05"), "2021-12-21 09:00:00")
```
