package mimd

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
