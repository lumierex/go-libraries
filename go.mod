module go-libraries

go 1.14

require sfpxm v0.0.0
require sfpxm-rpc v0.0.0


replace (
	sfpxm-rpc => ./sfpxm-rpc
	sfpxm => ./sfpxm
)


