package consul

import (
	"encoding/json"
	"strings"

	"github.com/hashicorp/consul/api"
)

func makeMap(kv api.KVPairs) map[string]interface{} {
	data := make(map[string]interface{})

	for _, v := range kv {
		// split on prefix
		keys := strings.Split(v.Key, "/")

		var vals interface{}
		json.Unmarshal(v.Value, &vals)

		// set data for first iteration
		kvals := data

		// iterate the keys and make maps
		for i, k := range keys {
			kval, ok := kvals[k].(map[string]interface{})
			if !ok {
				// create next map
				kval = make(map[string]interface{})
				// set it
				kvals[k] = kval
			}

			// last key: write vals
			if l := len(keys) - 1; i == l {
				kvals[k] = vals
				break
			}

			// set kvals for next iterator
			kvals = kval
		}

	}

	return data
}
