<section class="hero">
    <h1>Discover your next destination</h1>
    <p>Search countries, explore attractions, and curate your personal travel wishlist.</p>
    <div class="search-box">
        <label>WHERE TO NEXT?</label>
        <input type="text" id="home-search" placeholder="Country or capital..." autocomplete="off" />
        <div id="search-suggestions" class="suggestions-dropdown"></div>
    </div>
</section>

<section class="featured-section">
    <h2>Featured destinations</h2>
    <div class="country-grid">
        {{range .FeaturedCountries}}
        <a href="/countries/{{.Slug}}" class="country-card">
            <div class="flag-img-wrap">
                <img src="{{.FlagURL}}" alt="{{.FlagAlt}}" loading="lazy" />
            </div>
            <div class="card-info">
                <h3>{{.Name}}</h3>
                <p>{{.Capital}} &middot; {{.Region}}</p>
            </div>
        </a>
        {{end}}
    </div>
</section>

<section class="attractions-section">
    <h2>Popular attractions</h2>
    <div class="attraction-list">
        {{range .PopularAttractions}}
        <div class="attraction-item">
            <span class="attr-name">{{.Name}}</span>
            <span class="attr-kinds">{{.Kinds}}</span>
        </div>
        {{end}}
    </div>
</section>

<script src="/static/js/home.js"></script>