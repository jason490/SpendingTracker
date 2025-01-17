import { createSignal, For } from 'solid-js';

const SpendingForm = () => {
  const [cost, setCost] = createSignal('');
  const [reason, setReason] = createSignal('');
  const [selectedCard, setSelectedCard] = createSignal('');
  const [items, setItems] = createSignal([]);
  const [selectedItems, setSelectedItems] = createSignal(new Set());

  const cards = [
    { id: 'visa', name: 'Visa Card' },
    { id: 'mastercard', name: 'Mastercard' },
    { id: 'amex', name: 'American Express' },
    { id: 'discover', name: 'Discover Card' },
    { id: 'cash', name: 'Cash' }
  ];

  const handleSubmit = (e) => {
    e.preventDefault();
    if (cost() && reason() && selectedCard()) {
      setItems([...items(), { 
        id: Date.now(), 
        cost: cost(), 
        reason: reason(),
        card: selectedCard()
      }]);
      setCost('');
      setReason('');
      setSelectedCard('');
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
    setItems(items().filter(item => !selectedItems().has(item.id)));
    setSelectedItems(new Set());
  };

  return (
    <div class="container">
      <h2>Add/Remove Spending</h2>
      <form onSubmit={handleSubmit} class="mb-4">
        <div class="row g-3">
          <div class="col-md-3">
            <input
              type="number"
              class="form-control"
              placeholder="Cost"
              value={cost()}
              onInput={(e) => setCost(e.target.value)}
            />
          </div>
          <div class="col-md-4">
            <input
              type="text"
              class="form-control"
              placeholder="Reason"
              value={reason()}
              onInput={(e) => setReason(e.target.value)}
            />
          </div>
          <div class="col-md-3">
            <select 
              class="form-select"
              value={selectedCard()}
              onChange={(e) => setSelectedCard(e.target.value)}
            >
              <option value="">Select Card</option>
              {cards.map(card => (
                <option value={card.id}>{card.name}</option>
              ))}
            </select>
          </div>
          <div class="col-md-2">
            <button 
              type="submit" 
              class="btn btn-primary w-100"
              disabled={!cost() || !reason() || !selectedCard()}
            >
              Add
            </button>
          </div>
        </div>
      </form>

      {items().length > 0 && (
        <button
          class="btn btn-danger mb-3"
          onClick={removeSelected}
          disabled={selectedItems().size === 0}
        >
          Remove Selected
        </button>
      )}

      <div class="list-group">
        <For each={items()}>
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
    </div>
  );
};

export default SpendingForm;
