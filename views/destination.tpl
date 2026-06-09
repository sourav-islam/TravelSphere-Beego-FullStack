<div class="dest-hero-card">
    <div class="dest-flag">
        <img src="{{.Country.FlagURL}}" alt="Flag of {{.Country.Name}}" />
    </div>
    <div class="dest-info">
        <span class="region-badge">{{.Country.Region}}</span>
        <h1>{{.Country.Name}}</h1>
        <p class="official">Republic of {{.Country.Name}}</p>
        <div class="dest-meta">
            <div class="dest-meta-item">
                <label>CAPITAL</label>
                <span>{{.Country.Capital}}</span>
            </div>
            <div class="dest-meta-item">
                <label>POPULATION</label>
                <span>{{.Country.Population}}</span>
            </div>
            <div class="dest-meta-item">
                <label>REGION</label>
                <span>{{.Country.Region}} - {{.Country.Subregion}}</span>
            </div>
            <div class="dest-meta-item">
                <label>CURRENCY</label>
                <span>{{.Country.Currencies}}</span>
            </div>
            <div class="dest-meta-item">
                <label>LANGUAGES</label>
                <span>{{.Country.Languages}}</span>
            </div>
        </div>
    </div>
</div>

{{if .IsLoggedIn}}
<div class="wishlist-action-bar">
    <button id="add-wishlist-btn" class="btn-add-wishlist" onclick="addToWishlist('{{.Country.Name}}')">
        + Add to Wishlist
    </button>
    <div id="wishlist-feedback"></div>
</div>
{{end}}

<div class="dest-bottom">
    <div class="dest-panel">
        <h2>Travel weather</h2>
        <p class="weather-note">
            Weather data is optional. I don't have time to implement this.
        </p>
    </div>

    <div class="dest-panel">
        <h2>Attractions &amp; landmarks</h2>
        {{if .Attractions}}
        <div class="attraction-list">
            {{range .Attractions}}
            <div class="attraction-item">
                <span class="attr-name">{{.Name}}</span>
                <span class="attr-kinds">{{.Kinds}}</span>
            </div>
            {{end}}
        </div>
        {{else}}
        <p class="state-msg">No attraction data available for this destination.</p>
        {{end}}
    </div>
</div>

<script src="/static/js/destination.js"></script>