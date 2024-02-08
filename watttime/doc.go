// The [Best Practices for API Usage](https://docs.watttime.org/#section/Best-Practices-for-API-Usage) read:
//
// "Because grid region boundaries are occasionally updated, it is important to re-query `/v3/region-from-loc` at least once a month
// to ensure devices are receiving the signal corresponding to their location.
// The `/v3/maps` endpoint provides a GeoJSON that can be used for offline geocoding.
// The GeoJSON file includes a `last_updated` field that changes whenever the grid regions change."
//
// This requires polling either `/v3/region-from-loc` or `/v3/maps`.
package watttime
