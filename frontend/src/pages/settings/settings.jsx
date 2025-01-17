import { createSignal, Show, For } from 'solid-js';

const Settings = () => {
    const [showResetConfirm, setShowResetConfirm] = createSignal(false);
    const [exportFormat, setExportFormat] = createSignal('csv');
    const [currency, setCurrency] = createSignal('USD');
    const [cards, setCards] = createSignal([
        { id: 1, name: "Visa" },
        { id: 2, name: "Mastercard" }
    ]);
    const [newCardName, setNewCardName] = createSignal('');

    const currencies = [
        { code: 'USD', symbol: '$' },
        { code: 'EUR', symbol: '€' },
        { code: 'GBP', symbol: '£' },
        { code: 'JPY', symbol: '¥' }
    ];

    const handleAddCard = (e) => {
        e.preventDefault();
        if (newCardName().trim()) {
            setCards([...cards(), {
                id: cards().length + 1,
                name: newCardName().trim()
            }]);
            setNewCardName('');
        }
    };

    const handleRemoveCard = (id) => {
        setCards(cards().filter(card => card.id !== id));
    };

    const handleReset = () => {
        if (showResetConfirm()) {
            // Implement your reset logic here
            setShowResetConfirm(false);
        } else {
            setShowResetConfirm(true);
        }
    };

    return (
        <div class="container">
            <h2 class="mb-4">Settings</h2>

            {/* Theme Toggle */}
            <div class="card mb-4">
                <div class="card-header">
                    <h5 class="mb-0">Appearance</h5>
                </div>
                <div class="card-body">
                    <div class="mb-3">
                        <label class="form-label">Theme</label>
                        <div class="form-check form-switch">
                            <input
                                class="form-check-input"
                                type="checkbox"
                                role="switch"
                                id="themeToggle"
                            // Add your onChange handler here
                            />
                            <label class="form-check-label" for="themeToggle">
                                Dark Mode
                            </label>
                        </div>
                    </div>
                </div>
            </div>

            {/* Card Management */}
            <div class="card mb-4">
                <div class="card-header">
                    <h5 class="mb-0">Payment Methods</h5>
                </div>
                <div class="card-body">
                    <form class="mb-3 row g-3" onSubmit={handleAddCard}>
                        <div class="col-auto">
                            <input
                                type="text"
                                class="form-control"
                                placeholder="Card Name"
                                value={newCardName()}
                                onInput={(e) => setNewCardName(e.target.value)}
                            />
                        </div>
                        <div class="col-auto">
                            <button
                                type="submit"
                                class="btn btn-primary"
                                disabled={!newCardName().trim()}
                            >
                                Add Card
                            </button>
                        </div>
                    </form>
                    <div class="table-responsive">
                        <table class="table">
                            <thead>
                                <tr>
                                    <th>Card Name</th>
                                    <th>Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                <For each={cards()}>
                                    {(card) => (
                                        <tr>
                                            <td>{card.name}</td>
                                            <td>
                                                <button
                                                    class="btn btn-sm btn-outline-danger"
                                                    onClick={() => handleRemoveCard(card.id)}
                                                >
                                                    Remove
                                                </button>
                                            </td>
                                        </tr>
                                    )}
                                </For>
                            </tbody>
                        </table>
                    </div>
                </div>
            </div>

            {/* Data Export Settings */}
            <div class="card mb-4">
                <div class="card-header">
                    <h5 class="mb-0">Export Settings</h5>
                </div>
                <div class="card-body">
                    <div class="mb-3">
                        <label class="form-label">Default Export Format</label>
                        <select
                            class="form-select"
                            value={exportFormat()}
                            onChange={(e) => setExportFormat(e.target.value)}
                        >
                            <option value="csv">CSV</option>
                            <option value="json">JSON</option>
                            <option value="xlsx">Excel</option>
                        </select>
                    </div>
                </div>
            </div>

            {/* Currency Settings */}
            <div class="card mb-4">
                <div class="card-header">
                    <h5 class="mb-0">Currency Settings</h5>
                </div>
                <div class="card-body">
                    <div class="mb-3">
                        <label class="form-label">Display Currency</label>
                        <select
                            class="form-select"
                            value={currency()}
                            onChange={(e) => setCurrency(e.target.value)}
                        >
                            {currencies.map(curr => (
                                <option value={curr.code}>
                                    {curr.code} ({curr.symbol})
                                </option>
                            ))}
                        </select>
                    </div>
                </div>
            </div>

            {/* Reset Application */}
            <div class="card mb-4">
                <div class="card-header">
                    <h5 class="mb-0">Reset Application</h5>
                </div>
                <div class="card-body">
                    <Show
                        when={!showResetConfirm()}
                        fallback={
                            <div>
                                <p class="text-danger">Are you sure? This will delete all your data!</p>
                                <button class="btn btn-danger me-2" onClick={handleReset}>
                                    Yes, Delete
                                </button>
                                <button
                                    class="btn btn-secondary"
                                    onClick={() => setShowResetConfirm(false)}
                                >
                                    Cancel
                                </button>
                            </div>
                        }
                    >
                        <button class="btn btn-outline-danger" onClick={handleReset}>
                            Reset All Data
                        </button>
                    </Show>
                </div>
            </div>
        </div>
    );
};

export default Settings;

