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

	if packs[0] != 300 || packs[1] != 200 || packs[2] != 100 {
		t.Errorf("expected packs 300, 200 and 100 to be imported, got %v", packs)
	}
}

func TestPacksAreSorted(t *testing.T) {
	store.ImportPacks("100,200,300")
	packs := store.GetPacks()
	if len(packs) != 3 {
		t.Errorf("expected 3 packs, got %d", len(packs))
	}

	if packs[0] != 300 || packs[1] != 200 || packs[2] != 100 {
		t.Errorf("expected packs to be sorted, got %v", packs)
	}

	store.AddPack(150)
	packs = store.GetPacks()
	if len(packs) != 4 {
		t.Errorf("expected 4 packs, got %d", len(packs))
	}

	if packs[0] != 300 || packs[1] != 200 || packs[2] != 150 || packs[3] != 100 {
		t.Errorf("expected packs to be sorted, got %v", packs)
	}
}

func TestImportPacksDuplicatesNotAllowed(t *testing.T) {
	store.ImportPacks("100,200,300")
	packs := store.GetPacks()
	if len(packs) != 3 {
		t.Errorf("expected 3 packs, got %d", len(packs))
	}

	if packs[0] != 300 || packs[1] != 200 || packs[2] != 100 {
		t.Errorf("expected packs 100, 200 and 300 to be imported, got %v", packs)
	}

	store.AddPack(100)
	store.AddPack(100)

	packs = store.GetPacks()
	if len(packs) != 3 {
		t.Errorf("expected 3 packs, got %d", len(packs))
	}

	if packs[0] != 300 || packs[1] != 200 || packs[2] != 100 {
		t.Errorf("expected packs 100, 200 and 300 to be imported, got %v", packs)
	}
}

func TestImportPacksWithNans(t *testing.T) {
	store.ImportPacks("100,foo,300")
	packs := store.GetPacks()
	if len(packs) != 2 {
		t.Errorf("expected 2 packs, got %d", len(packs))
	}

	if packs[0] != 300 || packs[1] != 100 {
		t.Errorf("expected packs 300 and 100 to be imported, got %v", packs)
	}
}

func TestSolve(t *testing.T) {
	store.ImportPacks("100,200,300")
	solution := store.Solve(1000)
	if solution[100] != 1 || solution[200] != 0 || solution[300] != 3 {
		t.Errorf("expected solution to be {100: 1, 200: 0, 300: 3}, got %v", solution)
	}
}

func TestSolve2(t *testing.T) {
	store.ImportPacks("100,200,300")
	solution := store.Solve(1001)
	if solution[100] != 0 || solution[200] != 1 || solution[300] != 3 {
		t.Errorf("expected solution to be {100: 1, 200: 0, 300: 3}, got %v", solution)
	}
}

func TestSolve3(t *testing.T) {
	store.ImportPacks("250, 500, 1000, 2000, 5000")
	solution := store.Solve(1)
	if solution[250] != 1 || solution[500] != 0 || solution[1000] != 0 || solution[2000] != 0 || solution[5000] != 0 {
		t.Errorf("expected solution to be {250: 1, 500: 0, 1000: 0, 2000: 0, 5000: 0}, got %v", solution)
	}
}

func TestSolve4(t *testing.T) {
	store.ImportPacks("250, 500, 1000, 2000, 5000")
	solution := store.Solve(250)
	if solution[250] != 1 || solution[500] != 0 || solution[1000] != 0 || solution[2000] != 0 || solution[5000] != 0 {
		t.Errorf("expected solution to be {250: 1, 500: 0, 1000: 0, 2000: 0, 5000: 0}, got %v", solution)
	}
}

func TestSolve5(t *testing.T) {
	store.ImportPacks("250, 500, 1000, 2000, 5000")
	solution := store.Solve(251)
	if solution[250] != 0 || solution[500] != 1 || solution[1000] != 0 || solution[2000] != 0 || solution[5000] != 0 {
		t.Errorf("expected solution to be {250: 0, 500: 1, 1000: 0, 2000: 0, 5000: 0}, got %v", solution)
	}
}

func TestSolve6(t *testing.T) {
	store.ImportPacks("250, 500, 1000, 2000, 5000")
	solution := store.Solve(501)
	if solution[250] != 1 || solution[500] != 1 || solution[1000] != 0 || solution[2000] != 0 || solution[5000] != 0 {
		t.Errorf("expected solution to be {250: 1, 500: 1, 1000: 0, 2000: 0, 5000: 0}, got %v", solution)
	}
}

func TestSolve7(t *testing.T) {
	store.ImportPacks("250, 500, 1000, 2000, 5000")
	solution := store.Solve(12001)
	if solution[250] != 1 || solution[500] != 0 || solution[1000] != 0 || solution[2000] != 1 || solution[5000] != 2 {
		t.Errorf("expected solution to be {250: 1, 500: 0, 1000: 0, 2000: 1, 5000: 2}, got %v", solution)
	}
}

func TestSolve8(t *testing.T) {
	store.ImportPacks("23, 31, 53")
	solution := store.Solve(500000)
	if solution[23] != 2 || solution[31] != 7 || solution[53] != 9429 {
		t.Errorf("expected solution to be {23: 2, 31: 7, 53: 9429}, got %v", solution)
	}
}
