package inmemdb

import "testing"

func TestInMemory_Set(t *testing.T) {
	db := NewInMemoryDatabase()
	db.Set("key1", "value1")
	db.StartTransaction()
	db.Set("key1", "value2")
	db.Commit()

	if v, _ := db.Get(`key1`); v != `value2` {
		t.Error(`Value mismatch`)
	}
}

func TestInMemory_RollBack(t *testing.T) {
	db := NewInMemoryDatabase()
	db.Set("key1", "value1")
	db.StartTransaction()
	if v, _ := db.Get(`key1`); v != `value1` {
		t.Error(`Value 1 mismatch`)
	}

	db.Set("key1", "value2")
	if v, _ := db.Get(`key1`); v != `value2` {
		t.Error(`Value 2 mismatch`)
	}

	db.Rollback()
	if v, _ := db.Get(`key1`); v != `value1` {
		t.Error(`Value 3 mismatch`)
	}
}

func TestInMemory_NT(t *testing.T) {
	db := NewInMemoryDatabase()
	db.Set("key1", "value1")
	db.StartTransaction()
	db.Set("key1", "value2")

	if v, _ := db.Get(`key1`); v != `value2` {
		t.Error(`Value 1 mismatch`)
	}

	db.StartTransaction()
	if v, _ := db.Get(`key1`); v != `value2` {
		t.Error(`Value 2 mismatch`)
	}

	db.Delete(`key1`)
	db.Commit()
	if _, ok := db.Get(`key1`); ok {
		t.Error(`Value 3 mismatch`)
	}

	db.Commit()
	if _, ok := db.Get(`key1`); ok {
		t.Error(`Value 4 mismatch`)
	}
}

func TestInMemory_NTRollBack(t *testing.T) {
	db := NewInMemoryDatabase()
	db.Set("key1", "value1")
	db.StartTransaction()
	db.Set("key1", "value2")
	if v, _ := db.Get(`key1`); v != `value2` {
		t.Error(`Value 1 mismatch`)
	}

	db.StartTransaction()
	if v, _ := db.Get(`key1`); v != `value2` {
		t.Error(`Value 2 mismatch`)
	}

	db.Delete(`key1`)
	db.Rollback()
	if v, _ := db.Get(`key1`); v != `value2` {
		t.Error(`Value 3 mismatch`)
	}
	db.Commit()
	if v, _ := db.Get(`key1`); v != `value2` {
		t.Error(`Value 4 mismatch`)
	}
}

func TestInMemory_ManyChanges(t *testing.T) {
	db := NewInMemoryDatabase()
	db.Set("key1", "value1")
	db.Set("key2", "value2")
	db.Set("key3", "value3")
	db.Delete("key3")

	db.StartTransaction()
	db.Set("key3", "new-value3")

	if v, _ := db.Get(`key3`); v != `new-value3` {
		t.Error(`Value 1 mismatch`)
	}

	db.Rollback()
	if _, ok := db.Get(`key3`); ok {
		t.Error(`Value 2 mismatch`)
	}

	db.StartTransaction()
	db.Set(`key2`, `new-value2`)
	db.Commit()
	if v, _ := db.Get(`key2`); v != `new-value2` {
		t.Error(`Value 2 mismatch`)
	}

	db.Commit()
	if v, _ := db.Get(`key2`); v != `new-value2` {
		t.Error(`Value 2 mismatch`)
	}
}
