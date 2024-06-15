package array_list

type ArrayList struct {
	Size     int
	Capacity int
	Data     []any
}

func CreateArrayList() *ArrayList {
	list := &ArrayList{
		Size:     0,
		Capacity: 2,
		Data:     make([]any, 2),
	}
	return list
}

func (al *ArrayList) Append(data any) {
	if al.Size >= al.Capacity {
		// increase capacity by capacity * 2
		newData := make([]any, al.Capacity*2)
		for i, element := range al.Data {
			newData[i] = element
		}
		al.Data = newData
		al.Capacity *= 2
	}
	al.Size++
	al.Data[al.Size-1] = data
}
