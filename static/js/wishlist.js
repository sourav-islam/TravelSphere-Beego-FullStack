function saveWishlistRow(btn) {
    const row = btn.closest('tr');
    const id = row.dataset.id;
    const note = row.querySelector('.note-input').value;
    const status = row.querySelector('.status-select').value;

    fetch(`/api/wishlist/${id}`, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ note, status })
    })
    .then(r => r.json())
    .then(resp => {
        if (resp.success) {
            btn.textContent = 'Saved!';
            setTimeout(() => btn.textContent = 'Save', 1500);
            refreshDashboardStats();
        } else {
            alert(resp.error || 'Save failed');
        }
    });
}

function deleteWishlistRow(btn) {
    const row = btn.closest('tr');
    const id = row.dataset.id;
    if (!confirm('Remove this destination?')) return;

    fetch(`/api/wishlist/${id}`, { method: 'DELETE' })
    .then(r => r.json())
    .then(resp => {
        if (resp.success) {
            row.remove();
            refreshDashboardStats();
        } else {
            alert(resp.error || 'Delete failed');
        }
    });
}

function refreshDashboardStats() {
    const statsEl = document.getElementById('dashboard-stats');
    if (!statsEl) return;
    fetch('/api/dashboard/summary')
        .then(r => r.json())
        .then(resp => {
            if (resp.success) {
                const d = resp.data;
                statsEl.innerHTML = `
                    <div class="stat-card"><span class="stat-label">TOTAL SAVED</span><span class="stat-value">${d.total}</span></div>
                    <div class="stat-card"><span class="stat-label">PLANNED</span><span class="stat-value">${d.planned}</span></div>
                    <div class="stat-card"><span class="stat-label">VISITED</span><span class="stat-value">${d.visited}</span></div>`;
            }
        });
}