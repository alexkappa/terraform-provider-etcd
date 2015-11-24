package main

import (
	etcd "github.com/coreos/etcd/client"
	"github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"endpoints": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"etcd_key": KeyResource(),
		},
		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (kv interface{}, err error) {
	var endpoints []string
	values := d.Get("endpoints").([]interface{})
	for _, value := range values {
		endpoints = append(endpoints, value.(string))
	}
	config := etcd.Config{
		Endpoints: endpoints,
	}
	client, err := etcd.New(config)
	if err != nil {
		return nil, err
	}
	return etcd.NewKeysAPI(client), nil
}
