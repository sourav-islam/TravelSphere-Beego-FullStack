<div class="page-header">
    <h1>Country Explorer</h1>
    <p>Browse every destination on first load. Search and filter update only the results below — no full page reload.</p>
</div>

<div class="filter-bar">
    <div class="filter-group">
        <label for="country-search">SEARCH</label>
        <input type="text" id="country-search" placeholder="Country or capital..." />
    </div>
    <div class="filter-group">
        <label for="region-filter">REGION</label>
        <select id="region-filter">
            <option value="">All regions</option>
            <option value="Africa">Africa</option>
            <option value="Americas">Americas</option>
            <option value="Asia">Asia</option>
            <option value="Europe">Europe</option>
            <option value="Oceania">Oceania</option>
        </select>
    </div>
</div>

<div id="country-results" class="country-grid-4">
    {{range .Countries}}
    <a href="/countries/{{.Slug}}" class="ccard">
        <div class="ccard-flag">
            <img src="{{.FlagURL}}" alt="Flag of {{.Name}}" loading="lazy" />
        </div>
        <div class="ccard-body">
            <h3 class="ccard-name">{{.Name}}</h3>
            <div class="ccard-row"><span class="ccard-lbl">Capital:</span> {{.Capital}}</div>
            <div class="ccard-row"><span class="ccard-lbl">Population:</span> {{.Population}}</div>
            <div class="ccard-row"><span class="ccard-lbl">Currency:</span> {{.Currencies}}</div>
            <div class="ccard-row"><span class="ccard-lbl">Languages:</span> {{.Languages}}</div>
        </div>
    </a>
    {{end}}
</div>

<script src="/static/js/countries.js"></script>