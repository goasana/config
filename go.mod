module github.com/micro/go-config

require (
	cloud.google.com/go v0.37.4 // indirect
	github.com/BurntSushi/toml v0.3.1
	github.com/armon/circbuf v0.0.0-20190214190532-5111143e8da2 // indirect
	github.com/bitly/go-simplejson v0.5.0
	github.com/coreos/go-systemd v0.0.0-20190321100706-95778dfbb74e // indirect
	github.com/fsnotify/fsnotify v1.4.7
	github.com/ghodss/yaml v1.0.0
	github.com/golang/protobuf v1.3.1
	github.com/hashicorp/consul v1.4.5-0.20190327125850-f1fe406aa878
	github.com/hashicorp/consul/api v1.1.0
	github.com/hashicorp/hcl v1.0.0
	github.com/hashicorp/vault v1.1.0
	github.com/imdario/mergo v0.3.7
	github.com/mattn/go-colorable v0.1.1 // indirect
	github.com/micro/cli v0.1.0
	github.com/micro/go-micro v1.1.0
	github.com/pborman/uuid v1.2.0
	go.etcd.io/etcd v3.3.12+incompatible
	gocloud.dev v0.12.0
	golang.org/x/net v0.0.0-20190420063019-afa5a82059c6
	google.golang.org/grpc v1.20.1
	k8s.io/api v0.0.0-20190313235455-40a48860b5ab
	k8s.io/apimachinery v0.0.0-20190313205120-d7deff9243b1
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/utils v0.0.0-20190308190857-21c4ce38f2a7 // indirect
)

replace github.com/sourcegraph/go-diff => github.com/sourcegraph/go-diff v0.5.1

replace github.com/golang/lint v0.0.0-20190409202823-959b441ac422 => github.com/golang/lint v0.0.0-20190409202823-5614ed5bae6fb75893070bdc0996a68765fdd275

replace github.com/testcontainers/testcontainer-go => github.com/testcontainers/testcontainers-go v0.0.0-20181115231424-8e868ca12c0f
