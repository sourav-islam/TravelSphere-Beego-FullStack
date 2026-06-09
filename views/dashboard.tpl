<div class="page-header">
    <h1>Travel Dashboard</h1>
    <p>Your saved trips at a glance. Stats refresh automatically when your wishlist changes.</p>
</div>

<div id="dashboard-stats" class="stats-grid">
    <div class="stat-card">
        <span class="stat-label">TOTAL SAVED</span>
        <span class="stat-value">{{.Summary.Total}}</span>
    </div>
    <div class="stat-card">
        <span class="stat-label">PLANNED</span>
        <span class="stat-value">{{.Summary.Planned}}</span>
    </div>
    <div class="stat-card">
        <span class="stat-label">VISITED</span>
        <span class="stat-value">{{.Summary.Visited}}</span>
    </div>
</div>

<div class="section-header">
    <h2>Saved destinations</h2>
</div>
<div class="saved-list">
    {{range .WishlistItems}}
    <div class="saved-item">
        <strong>{{.CountryName}}</strong>
        <span class="status-badge status-{{.Status}}">{{.Status}}</span>
        {{if .Note}}<span class="saved-note">&middot; {{.Note}}</span>{{end}}
    </div>
    {{end}}
    {{if not .WishlistItems}}
    <p class="state-msg">No saved destinations yet.</p>
    {{end}}
</div>

<script src="/static/js/dashboard.js"></script>