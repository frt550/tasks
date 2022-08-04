package counter

var id uint = 0

func GetId() uint {
	id++
	return id
}
