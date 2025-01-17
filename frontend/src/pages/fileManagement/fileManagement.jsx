import { createSignal, For } from 'solid-js';

const FileManagement = () => {
  const [importedItems, setImportedItems] = createSignal([]);
  const [selectedItems, setSelectedItems] = createSignal(new Set());

  const cards = [
    { id: 'visa', name: 'Visa Card' },
    { id: 'mastercard', name: 'Mastercard' },
    { id: 'amex', name: 'American Express' },
    { id: 'discover', name: 'Discover Card' },
    { id: 'cash', name: 'Cash' }
  ];

  const handleImport = (e) => {
    const file = e.target.files[0];
    if (file) {
      // Simulated CSV import with card information
      const mockData = [
        { id: 1, cost: "50.00", reason: "Groceries", card: "visa" },
        { id: 2, cost: "30.00", reason: "Gas", card: "mastercard" },
        { id: 3, cost: "25.00", reason: "Books", card: "amex" }
      ];
      setImportedItems(mockData);
    }
  };

  const toggleItem = (id) => {
    const newSelected = new Set(selectedItems());
    if (newSelected.has(id)) {
      newSelected.delete(id);
    } else {
      newSelected.add(id);
    }
    setSelectedItems(newSelected);
  };

  const removeSelected = () => {
    setImportedItems(importedItems().filter(item => !selectedItems().has(item.id)));
    setSelectedItems(new Set());
  };

  return (
    <div class="container">
      <h2>Import/Export Data</h2>
      <div class="row g-3 mb-4">
        <div class="col-md-6">
          <div class="d-grid">
            <button class="btn btn-primary">Export CSV</button>
          </div>
        </div>
        <div class="col-md-6">
          <div class="d-grid">
            <input
              type="file"
              accept=".csv"
              class="form-control"
              onChange={handleImport}
            />
          </div>
        </div>
      </div>

      {importedItems().length > 0 && (
        <>
          <button
            class="btn btn-danger mb-3"
            onClick={removeSelected}
            disabled={selectedItems().size === 0}
          >
            Remove Selected
          </button>

          <div class="list-group">
            <For each={importedItems()}>
              {(item) => (
                <div class="list-group-item">
                  <div class="d-flex align-items-center">
                    <input
                      type="checkbox"
                      class="form-check-input me-3"
                      checked={selectedItems().has(item.id)}
                      onChange={() => toggleItem(item.id)}
                    />
                    <div class="flex-grow-1">
                      <strong>${item.cost}</strong> - {item.reason}
                      <span class="badge bg-secondary ms-2">
                        {cards.find(card => card.id === item.card)?.name}
                      </span>
                    </div>
                  </div>
                </div>
              )}
            </For>
          </div>
        </>
      )}
    </div>
  );
};

export default FileManagement;

