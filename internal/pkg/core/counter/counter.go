package counter

var id uint64 = 0

func GetId() uint64 {
	id++
	return id
}
