package toy

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerCreate,
		Read:   resourceServerRead,
		Update: resourceServerUpdate,
		Delete: resourceServerDelete,

		Schema: map[string]*schema.Schema{
			"address": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceServerCreate(d *schema.ResourceData, m interface{}) error {
	address := d.Get("address").(string)
	d.SetId(address)
	writeAddressToFile(address)
	return resourceServerRead(d, m)
}

func resourceServerRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("read from upstream API")
	address := readAddressFromFile()
	d.Set("address", address)
	return nil
}

func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
	d.Partial(true)

	if d.HasChange("address") {
		address := d.Get("address").(string)
		writeAddressToFile(address)
		d.SetPartial("address")
	}
	d.Partial(false)

	return resourceServerRead(d, m)
}

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	deleteAddressFile()
	return nil
}

const ADDRESS_FILE = "./metadata.txt"

func readAddressFromFile() string {
	log.Println("start read address from file")

	data, err := ioutil.ReadFile(ADDRESS_FILE)
	if err != nil {
		log.Printf("readAddressFromFile err: %+v\n", err)
		return ""
	}
	return string(data)
}

func writeAddressToFile(address string) error {
	log.Println("start read address file")

	pwd, _ := os.Getwd()
	executable, _ := os.Executable()
	log.Printf("Pwd: %s\n", pwd)
	log.Printf("Executable: %s\n", executable)

	err := ioutil.WriteFile(ADDRESS_FILE, []byte(address), 0644)
	if err != nil {
		log.Printf("writeAddressToFile err: %+v\n", err)
	}
	return err
}

func deleteAddressFile() error {
	log.Println("start delete address file")
	return os.Remove(ADDRESS_FILE)
}
