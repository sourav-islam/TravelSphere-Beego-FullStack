(function () {
  var searchInput  = document.getElementById('country-search');
  var regionFilter = document.getElementById('region-filter');
  var results      = document.getElementById('country-results');
  if (!results) return;

  var timer;

  // Pre-fill search if redirected from home with ?q=
  var params = new URLSearchParams(window.location.search);
  var preQ   = params.get('q');
  if (preQ && searchInput) {
    searchInput.value = preQ;
    fetchCountries();
  }

  if (searchInput) {
    searchInput.addEventListener('input', function () {
      clearTimeout(timer);
      timer = setTimeout(fetchCountries, 350);
    });
  }
  if (regionFilter) {
    regionFilter.addEventListener('change', fetchCountries);
  }

  function fetchCountries() {
    var search = searchInput ? searchInput.value.trim() : '';
    var region = regionFilter ? regionFilter.value : '';

    results.innerHTML = '<div class="state-msg"><div class="spinner"></div><br>Loading...</div>';

    fetch('/api/countries?search=' + encodeURIComponent(search) + '&region=' + encodeURIComponent(region))
      .then(function (r) { return r.json(); })
      .then(function (resp) {
        if (!resp.success || !resp.data || resp.data.length === 0) {
          results.innerHTML = '<div class="state-msg">No countries found.</div>';
          return;
        }
        results.innerHTML = resp.data.map(buildCard).join('');
      })
      .catch(function () {
        results.innerHTML = '<div class="state-msg">Error loading countries. Please try again.</div>';
      });
  }

  function buildCard(c) {
    return '<a href="/countries/' + esc(c.slug) + '" class="country-card">' +
      '<div class="flag-wrap"><img src="' + esc(c.flag_url) + '" alt="Flag of ' + esc(c.name) + '" loading="lazy" /></div>' +
      '<div class="card-body">' +
        '<h3>' + esc(c.name) + '</h3>' +
        '<p><strong>Capital:</strong> ' + esc(c.capital) + '</p>' +
        '<p><strong>Population:</strong> ' + esc(c.population) + '</p>' +
        '<p><strong>Currency:</strong> ' + esc(c.currencies) + '</p>' +
        '<p><strong>Languages:</strong> ' + esc(c.languages) + '</p>' +
      '</div></a>';
  }

  function esc(s) {
    return String(s||'').replace(/&/g,'&amp;').replace(/</g,'&lt;').replace(/>/g,'&gt;').replace(/"/g,'&quot;');
  }
})();