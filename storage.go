package mimd

func newStorage() storage {
	return storage{
		values: make(map[string]*value),
	}
}

type storage struct {
	values map[string]*value
}

func (r *storage) MergeInto(dstId int8, outerStorage *storage) {
	for k, v := range r.values {
		if v.Deleted() {
			if dstId == 0 { // root storage
				delete(outerStorage.values, k)
			} else {
				outerStorage.values[k].Delete()
			}
		} else {
			outerStorage.values[k] = v
		}
	}
}

func newValue(val string, del bool) *value {
	return &value{
		val: val,
		del: del,
	}
}

type value struct {
	val string
	del bool
}

func (r *value) Value() string {
	return r.val
}

func (r *value) Deleted() bool {
	return r.del
}

func (r *value) Delete() {
	r.del = true
}
