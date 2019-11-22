# lmdb
Local Metadata Data Base (LMDB)



## What is this?

> lmdb is a local database that serves as a replacement for TheMovieDB/AniDB/TVDB/etc.
>
> It can be used for providing movie/tv metadata for the PlexMediaServer or any other Service that may request them over an API

## TODO

- [ ] Write API Server that exposes typical Endpoints for Metadata services
  - Show (Get, Search)
  - Movie (Get, Search)
  - Anime (Get, Search)
- [ ] Write MongoDB/Postgres Integration for Storing the Data
- [ ] Write Commandline Tools to export/import data from metadata.zip files
  - Write spec for the zip
- [ ] Write Commandline Tools to fetch bundle from remote sources