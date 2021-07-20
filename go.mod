module main

go 1.16

require (
	agents/state v0.0.0-00010101000000-000000000000
	agents/worker v0.0.0-00010101000000-000000000000
	agents/world v0.0.0-00010101000000-000000000000
)

replace agents/world => ./entities/

replace agents/worker => ./workers/

replace agents/state => ./model/
