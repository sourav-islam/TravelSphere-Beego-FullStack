function addToWishlist(countryName) {
  var btn      = document.getElementById('add-wishlist-btn');
  var feedback = document.getElementById('wishlist-feedback');
  if (btn) btn.disabled = true;
  feedback.textContent  = 'Adding...';
  feedback.style.color  = '#6c3fc5';

  fetch('/api/wishlist', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ country_name: countryName, status: 'Planned' })
  })
  .then(function (r) { return r.json(); })
  .then(function (resp) {
    if (resp.success) {
      feedback.textContent = '✓ Added to wishlist!';
      feedback.style.color = '#6c3fc5';
    } else {
      feedback.textContent = resp.error || 'Failed to add.';
      feedback.style.color = '#e8357a';
      if (btn) btn.disabled = false;
    }
  })
  .catch(function () {
    feedback.textContent = 'Connection error.';
    feedback.style.color = '#e8357a';
    if (btn) btn.disabled = false;
  });
}