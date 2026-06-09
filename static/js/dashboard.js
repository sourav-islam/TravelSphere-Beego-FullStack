function refreshStats() {
  fetch('/api/dashboard/summary')
    .then(function (r) { return r.json(); })
    .then(function (resp) {
      if (!resp.success) return;
      var d  = resp.data;
      var el = document.getElementById('dashboard-stats');
      if (!el) return;
      el.innerHTML =
        '<div class="stat-card"><span class="stat-label">TOTAL SAVED</span><span class="stat-value">' + d.total   + '</span></div>' +
        '<div class="stat-card"><span class="stat-label">PLANNED</span><span class="stat-value">'     + d.planned + '</span></div>' +
        '<div class="stat-card"><span class="stat-label">VISITED</span><span class="stat-value">'     + d.visited + '</span></div>';
    });
}

// Auto-refresh every 30 seconds
setInterval(refreshStats, 30000);