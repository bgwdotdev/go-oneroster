package handlers

import (
	"fmt"
)

func TestNestedProperty() {
	// setup
	c1 := map[string]interface{}{
		"sourcedId": "1",
		"type":      "s",
	}
	c2 := map[string]interface{}{
		"sourcedId": "2",
		"type":      "s",
	}
	o := map[string]interface{}{
		"Children": []map[string]interface{}{
			c1,
			c2,
		},
	}
	sourcedId := "123"

	// execute
	fmt.Println(o)
	stmt := fmt.Sprintf("SELECT sourcedId, type FROM orgs WHERE parentSourcedId=%v", sourcedId)

	children := o["children"]
	for rows.Next() {
		child := FormatResults(rows)
		children = append(children, child)
	}
	org["children"] = children
	JSON.marshal(org)

	// validate
}
