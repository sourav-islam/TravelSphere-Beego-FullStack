# TravelSphere

A full-stack travel destination discovery and trip planner built with **Beego Framework (Go)**.  
Browse countries, explore attractions, manage a personal wishlist, and track your travel dashboard — all with SSR pages and AJAX partial updates.

---

## Project Overview

TravelSphere allows users to:

- Search and browse **250+ countries** with flags, capitals, population, currency, and languages
- View **destination detail pages** with live attractions via OpenTripMap
- Manage a personal **travel wishlist** (add, edit notes, change status, delete)
- View a **dashboard** with trip statistics (total saved, planned, visited)
- Use **AJAX partial updates** — no full page reloads for search, wishlist, or dashboard actions

---

## Tech Stack

| Layer | Technology |
|---|---|
| Backend framework | Beego v2 (Go) |
| Templating | Beego SSR templates (.tpl) |
| Routing | Beego router (SSR + JSON API) |
| Session management | Beego in-memory session |
| Country data | REST Countries API v3.1 |
| Attractions data | OpenTripMap API |
| Wishlist storage | In-memory store (no database) |
| Frontend | Vanilla JavaScript (Fetch API) |
| Styling | Plain CSS with CSS variables |
| Configuration | `conf/app.conf` |


---

## Project Structure
```
TravelSphere/
├── conf/
│   └── app.conf                  # Beego configuration + API keys
├── controllers/
│   ├── base.go                   # BaseController: layout, session, auth
│   ├── home.go                   # SSR: /, search suggestions
│   ├── country.go                # SSR: /countries, /countries/:slug
│   ├── wishlist.go               # SSR: /wishlist (protected)
│   ├── dashboard.go              # SSR: /dashboard (protected)
│   └── api/
│       ├── countries.go          # JSON: /api/countries, /api/countries/:slug
│       ├── wishlist.go           # JSON: /api/wishlist CRUD
│       └── dashboard.go          # JSON: /api/dashboard/summary
├── routers/
│   ├── router.go                 # SSR routes + logging/auth filters
│   └── api.go                    # /api/* routes
├── services/
│   ├── country_service.go        # Country fetch, search, cache
│   ├── attraction_service.go     # OpenTripMap geocode + attractions
│   ├── wishlist_service.go       # In-memory wishlist CRUD
│   └── dashboard_service.go      # Summary stats from wishlist
├── models/
│   ├── country.go                # Country, CountryDTO structs
│   ├── attraction.go             # AttractionDTO structs
│   └── wishlist.go               # WishlistItem, request/response structs
├── utils/
│   ├── api_client.go             # Reusable HTTP GET + JSON decode
│   ├── formatter.go              # Population, language, slug helpers
│   ├── validator.go              # Status, country name, slug validation
│   └── response.go               # Standard JSON response helpers
├── views/
│   ├── layout/
│   │   └── main.tpl              # Base HTML layout with navbar + footer
│   ├── home.tpl                  # Home page
│   ├── countries.tpl             # Country explorer
│   ├── destination.tpl           # Country detail + attractions
│   ├── wishlist.tpl              # Wishlist management
│   ├── dashboard.tpl             # Travel dashboard
│   └── 404.tpl                   # Not found page
├── static/
│   ├── css/
│   │   └── style.css             # All styles
│   └── js/
│       ├── home.js               # Search autocomplete
│       ├── countries.js          # AJAX search + filter
│       ├── destination.js        # Add to wishlist AJAX
│       ├── wishlist.js           # Save/delete row AJAX
│       └── dashboard.js          # Stats refresh AJAX
├── go.mod
├── go.sum

```
---

## Setup Instructions

### 1. Prerequisites

- Go 1.21 or higher
- Internet access (for REST Countries API and OpenTripMap API)
- OpenTripMap API key — free at [https://opentripmap.io](https://opentripmap.io)

### 2. Clone the repository

```bash
git clone https://github.com/sourav-islam/TravelSphere-Beego-FullStack
cd TravelSphere
```

### 3. Install dependencies

```bash
go mod tidy
```

## ⚙️ Quick Configuration (`conf/app.conf`)
Beego reads this automatically. No `.env` needed.

```ini
appname = TravelSphere
httpport = 8081
runmode = dev

sessionon = true
sessionprovider = memory
copyrequestbody = true

restcountries_api_BASE_URL = https://restcountries.com/v3.1/all
OPENTRIPMAP_BASE_URL = https://api.opentripmap.com/0.1/en/places/radius
opentripmap_api_key = your_key_here

### 5. Run the application

```bash
go run main.go
```

### 6. Open in browser
http://localhost:8081

---
### Why in-memory

- No database setup required for the assessment
- Zero external dependencies
- Thread-safe with `sync.RWMutex`
- Fits the assessment requirement: *"wishlist data accessed through service-layer API calls"*

### Limitation

Data resets when the server restarts. This is expected and documented per assessment requirements.  
To add persistence, swap `wishlist_service.go`'s store with a file-based JSON adapter — the service interface stays identical.

### Wishlist entity

```go
type WishlistItem struct {
    ID          string    // uuid v4, auto-generated
    CountryName string    // required
    Note        string    // optional, editable
    Status      string    // "Planned" or "Visited"
    CreatedAt   time.Time // auto-set on creation
    UserID      string    // from session
}
```

---

## AJAX Behavior

All AJAX interactions update **only the targeted container** — the browser never reloads the full page.

| Page | Trigger | Container updated | API called |
|---|---|---|---|
| `/` | Typing in search box | `#search-suggestions` | `GET /search?q=` |
| `/countries` | Search input / region filter | `#country-results` | `GET /api/countries` |
| `/countries/:slug` | Click "Add to Wishlist" | `#wishlist-feedback` | `POST /api/wishlist` |
| `/wishlist` | Click Save on a row | Row in `#wishlist-rows` | `PUT /api/wishlist/:id` |
| `/wishlist` | Click Delete on a row | Row removed from DOM | `DELETE /api/wishlist/:id` |
| `/dashboard` | Auto every 30 seconds | `#dashboard-stats` | `GET /api/dashboard/summary` |

---

## External APIs

### REST Countries API
- **URL:** `https://restcountries.com/v3.1/all`
- **Used for:** country name, flag, capital, population, region, languages, currencies
- **Auth:** None required
- **Caching:** Results cached in memory on first load — subsequent requests served from cache

### OpenTripMap API
- **URL:** `https://api.opentripmap.com/0.1/en/places/`
- **Used for:** geocoding capital cities, fetching nearby attractions
- **Auth:** API key in `conf/app.conf` → `opentripmap_api_key`
- **Flow:**
  1. `geoname` endpoint resolves capital name → lat/lon
  2. `radius` endpoint fetches attractions within 50km

---


## Architecture Notes

### MVC separation

- **Controllers** handle HTTP — parse request, call service, set template data or return JSON
- **Services** contain all business logic — no HTTP knowledge, no direct storage access
- **Models** are pure data structs — no methods, no logic
- **Utils** are stateless helpers — formatting, validation, HTTP client, response builders

### Filters / Middleware

Two filters registered in `routers/router.go`:

1. **Logging filter** — logs `[REQ] METHOD /path` before and `[RES] METHOD /path — duration` after every request
2. **Auth filter** — applied to `/wishlist` and `/dashboard` — redirects to `/` if no session

### Prepare() usage

`BaseController.Prepare()` runs before every SSR controller action and:
- Sets `c.Layout = "layout/main.tpl"`
- Reads session username → sets `c.IsLoggedIn`, `c.CurrentUser`
- Passes `CurrentUser`, `IsLoggedIn`, `ActiveMenu` to all templates

### Country slug resolution

Country names from REST Countries API are converted to slugs at cache-load time:
- Lowercased
- Spaces replaced with hyphens
- Apostrophes and commas removed

The same function runs on every search, ensuring consistent matching.

