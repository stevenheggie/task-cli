module github.com/stevenheggie/task

go 1.19

require ( //TODO migrate to newer bolt library - https://github.com/etcd-io/bbolt
	github.com/boltdb/bolt v1.3.1 // direct
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/cobra v1.5.0 // direct
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/sys v0.0.0-20220919091848-fb04ddd9f9c8 // indirect

)