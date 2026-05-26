package pagination

// Manager handles the pagination.
type Manager struct {
	totalPages   int
	currentPage  int
	limit        uint
	offset       int
	currentTable string
}

// New returns a pointer to a Manager instance.
func New(limit uint, count int, currentTable string) (*Manager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NextPage increases the value currentPage.
func (m *Manager) NextPage() error { _ = "STUB: not implemented"; return nil }

// PreviousPage decreases the value of currentPage.
func (m *Manager) PreviousPage() error { _ = "STUB: not implemented"; return nil }

// Offset returns the limit.
func (m *Manager) Offset() int {
	_ = "STUB: not implemented"

	// Limit returns the limit.
	return 0
}

func (m *Manager) Limit() uint {
	_ = "STUB: not implemented"

	// TotalPages returns the total pages count.
	return 0
}

func (m *Manager) TotalPages() int { _ = "STUB: not implemented"; return 0 }

// CurrentPage returns the currentPage value.
func (m *Manager) CurrentPage() int { _ = "STUB: not implemented"; return 0 }

// SetCurrentTable sets the current table name.
func (m *Manager) SetCurrentTable(tableName string) { _ = "STUB: not implemented"; return }

// CurrentTable sets the current table name.
func (m *Manager) CurrentTable() string { _ = "STUB: not implemented"; return "" }

// setOffset calculates the offset based of the current page and the limit.
func (m *Manager) setOffset() { _ = "STUB: not implemented"; return }

// setTotalPages total pages = count / limit, if the limit is greater than 0.
func (m *Manager) setTotalPages(count int) error { _ = "STUB: not implemented"; return nil }
