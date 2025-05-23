package store_test

import (
	"packr/internal/store"
	"testing"
)

func TestAddAndRemovePack(t *testing.T) {
	store.AddPack(100)
	store.AddPack(200)

	packs := store.GetPacks()
	if len(packs) != 2 {
		t.Errorf("expected 2 packs, got %d", len(packs))
	}

	store.RemovePack(100)
	packs = store.GetPacks()
	if len(packs) != 1 {
		t.Errorf("expected 1 pack, got %d", len(packs))
	}

	if packs[0] != 200 {
		t.Errorf("expected pack 200 to remain, got %d", packs[0])
	}
}

func TestImportPacks(t *testing.T) {
	store.ImportPacks("100,200,300")
	packs := store.GetPacks()
	if len(packs) != 3 {
		t.Errorf("expected 3 packs, got %d", len(packs))
	}

	if packs[0] != 100 || packs[1] != 200 || packs[2] != 300 {
		t.Errorf("expected packs 100, 200 and 300 to be imported, got %v", packs)
	}
}

func TestPacksAreSorted(t *testing.T) {
	store.ImportPacks("100,200,300")
	packs := store.GetPacks()
	if len(packs) != 3 {
		t.Errorf("expected 3 packs, got %d", len(packs))
	}

	if packs[0] != 100 || packs[1] != 200 || packs[2] != 300 {
		t.Errorf("expected packs to be sorted, got %v", packs)
	}

	store.AddPack(150)
	packs = store.GetPacks()
	if len(packs) != 4 {
		t.Errorf("expected 4 packs, got %d", len(packs))
	}

	if packs[0] != 100 || packs[1] != 150 || packs[2] != 200 || packs[3] != 300 {
		t.Errorf("expected packs to be sorted, got %v", packs)
	}
}

func TestImportPacksDuplicatesAllowed(t *testing.T) {
	store.ImportPacks("100,200,300")
	packs := store.GetPacks()
	if len(packs) != 3 {
		t.Errorf("expected 3 packs, got %d", len(packs))
	}

	if packs[0] != 100 || packs[1] != 200 || packs[2] != 300 {
		t.Errorf("expected packs 100, 200 and 300 to be imported, got %v", packs)
	}

	store.AddPack(100)
	store.AddPack(100)

	packs = store.GetPacks()
	if len(packs) != 5 {
		t.Errorf("expected 4 packs, got %d", len(packs))
	}

	if packs[0] != 100 || packs[1] != 100 || packs[2] != 100 || packs[3] != 200 || packs[4] != 300 {
		t.Errorf("expected packs 100, 100, 100, 200 and 300 to be imported, got %v", packs)
	}
}

func TestImportPacksWithNans(t *testing.T) {
	store.ImportPacks("100,foo,300")
	packs := store.GetPacks()
	if len(packs) != 2 {
		t.Errorf("expected 2 packs, got %d", len(packs))
	}

	if packs[0] != 100 || packs[1] != 300 {
		t.Errorf("expected packs 100 and 300 to be imported, got %v", packs)
	}
}
