const searchInput = document.getElementById('home-search');
const suggestionsBox = document.getElementById('search-suggestions');
let debounceTimer;

if (searchInput) {
    searchInput.addEventListener('input', function () {
        clearTimeout(debounceTimer);
        const q = this.value.trim();
        if (q.length < 2) {
            suggestionsBox.innerHTML = '';
            return;
        }
        debounceTimer = setTimeout(() => {
            fetch(`/search?q=${encodeURIComponent(q)}`)
                .then(r => r.json())
                .then(countries => {
                    suggestionsBox.innerHTML = '';
                    countries.forEach(c => {
                        const item = document.createElement('div');
                        item.className = 'suggestion-item';
                        item.textContent = `${c.name} — ${c.capital}`;
                        item.addEventListener('click', () => {
                            window.location.href = `/countries/${c.slug}`;
                        });
                        suggestionsBox.appendChild(item);
                    });
                });
        }, 300);
    });

    document.addEventListener('click', (e) => {
        if (!searchInput.contains(e.target)) {
            suggestionsBox.innerHTML = '';
        }
    });
}