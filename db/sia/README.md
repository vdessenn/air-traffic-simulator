# AIXM source data

This directory is where the raw aeronautical database goes. It is **not versioned** —
the export is ~125 MB and is published by the French SIA, not by this project.

The simulation will not start without it: `cmd/api` resolves the newest export at
startup and aborts if none is found.

## Getting the export

1. Go to the SIA AIP publication site: <https://www.sia.aviation-civile.gouv.fr/>
2. Download the **AIXM 4.5** dataset for *France + Outre-Mer* (`export_xml_bd_sia`).
   Any recent AIRAC cycle works.
3. Unzip it here, keeping the archive's own directory name.

## Expected layout

```
db/sia/
└── export_xml_bd_sia_<YYYY-MM-DD>-v<NN>/
    └── AIXM4.5_all_FR_OM_<YYYY-MM-DD>.xml
```

Both names are globs, not fixed strings — see `FindLatestAIXMFile` in
`backend/internal/service/file_manager.go`. Several cycles can coexist; the most
recently modified XML wins.

## What happens next

Nothing else is manual. On startup the API converts this XML into `db/json/`
(`EnsureJSONData`, re-run when the JSON is older than 28 days) and downloads flight
routes into `db/routes/` (`EnsureRoutesData`, 7 days). Both directories are ignored
by git.
