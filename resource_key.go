package main

import (
	etcd "github.com/coreos/etcd/client"
	"github.com/hashicorp/terraform/helper/schema"
	"golang.org/x/net/context"
)

func KeyResource() *schema.Resource {
	return &schema.Resource{
		Create: createKey,
		Read:   readKey,
		Update: createKey,
		Delete: deleteKey,
		Schema: map[string]*schema.Schema{
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func createKey(d *schema.ResourceData, m interface{}) error {
	key := d.Get("key").(string)
	value := d.Get("value").(string)

	kv := m.(etcd.KeysAPI)
	_, err := kv.Set(context.Background(), key, value, nil)
	if err != nil {
		return err
	}

	d.SetId(key)
	d.Set("key", key)
	d.Set("value", value)

	return nil
}

func readKey(d *schema.ResourceData, m interface{}) error {
	kv := m.(etcd.KeysAPI)
	_, err := kv.Get(context.Background(), d.Id(), nil)
	if err != nil {
		if cerr, ok := err.(etcd.Error); ok && cerr.Code == etcd.ErrorCodeKeyNotFound {
			d.SetId("")
			return nil
		}
		return err
	}

	return nil
}

func deleteKey(d *schema.ResourceData, m interface{}) error {
	kv := m.(etcd.KeysAPI)
	_, err := kv.Delete(context.Background(), d.Id(), nil)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}
