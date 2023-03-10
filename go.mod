module github.com/LastPossum/kamino_benchmarks

go 1.19

require (
	github.com/LastPossum/kamino v0.0.0-20230210145649-1f59216bc01d
	github.com/barkimedes/go-deepcopy v0.0.0-20220514131651-17c30cfc62df
	github.com/jinzhu/copier v0.3.5
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826
	github.com/vmihailenco/msgpack v4.0.4+incompatible
)

require (
	github.com/golang/protobuf v1.3.1 // indirect
	golang.org/x/net v0.0.0-20190603091049-60506f45cf65 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
)

replace github.com/LastPossum/kamino => ../kamino
