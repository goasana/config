package etcd

import (
	"encoding/json"
	"strings"

	"github.com/coreos/etcd/mvcc/mvccpb"
)

func makeMap(kv []*mvccpb.KeyValue, stripPrefix string) map[string]interface{} {
	data := make(map[string]interface{})

	for _, v := range kv {
		// remove prefix if non empty, and ensure leading / is removed as well
		vkey := strings.TrimPrefix(strings.TrimPrefix(string(v.Key), stripPrefix), "/")
		// split on prefix
		keys := strings.Split(vkey, "/")

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
