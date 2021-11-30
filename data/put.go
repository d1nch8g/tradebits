package data

// put key by some value to database (if value exists use Change()
// func instead)
func Put(key []byte, value []byte) {
	valueExists, _ := base.Has(key, nil)
	if valueExists {
		return
	}
	base.Put(key, value, nil)
	if reserveNodeConnected {
		obj := dataObject{
			key:   key,
			value: value,
		}
		DataQue = append(DataQue, obj)
	}
}
