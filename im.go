package mimd

func NewInMemoryDatabase() *DB {
	return &DB{
		storages: []storage{
			NewStorage(),
		},
	}
}

type DB struct {
	lastStorageId int8
	storages      []storage
}

func (r *DB) Get(k string) (string, bool) {
	for i := r.lastStorageId; i > -1; i-- {
		if v, ok := r.storages[i].values[k]; ok && !v.Deleted() {
			return v.Value(), ok
		}
	}

	return ``, false
}

func (r *DB) Set(k, v string) {
	r.storages[r.lastStorageId].values[k] = newValue(v, false)
}

func (r *DB) Delete(k string) {
	if _, ok := r.storages[r.lastStorageId].values[k]; ok {
		if r.lastStorageId != 0 {
			r.storages[r.lastStorageId].values[k].Delete()
		} else {
			delete(r.storages[r.lastStorageId].values, k)
		}
	} else {
		r.storages[r.lastStorageId].values[k] = newValue(``, true)
	}
}

func (r *DB) StartTransaction() {
	r.storages = append(r.storages, NewStorage())
	r.lastStorageId++
}

func (r *DB) Commit() {
	if r.lastStorageId != 0 {
		r.storages[r.lastStorageId].MergeInto(&r.storages[r.lastStorageId-1])
	}
}

func (r *DB) Rollback() {
	if r.lastStorageId != 0 {
		r.deleteLastStorage()
		r.lastStorageId--
	}
}

func (r *DB) deleteLastStorage() {
	r.storages = append(r.storages[:r.lastStorageId])
}
