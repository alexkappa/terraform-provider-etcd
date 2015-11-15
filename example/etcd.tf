provider "etcd" {
	endpoints = ["http://localhost:2379", "http://localhost:4001"]
}

resource "etcd_discovery" "token" {
    size = 3
}

output "etcd_discovery_url" {
	value = "${etcd_discovery.token.id}"
}

resource "etcd_key" "foo" {
	key = "foo"
	value = "bar"
}

output "etcd_key" {
	value = "${etcd_key.foo.id}"
}

output "etcd_value" {
	value = "${etcd_key.foo.value}"
}