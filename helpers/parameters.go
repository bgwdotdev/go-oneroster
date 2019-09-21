package helpers

import ()

func SetOptions() (*mongo.FindOptions, error) {
	projection, err := getFields(url)
	options.
		Find().
		SetLimit(getLimit(url)).
		SetSkip(getOffset(url)).
		SetSort(bson.D{getSort(url), 1}).
		SetProjection(projection)
}

func getFilters() {}

func getFields(q url.Values, safeFields []string) ([]bson.DocElem, error) {
	v := q.Get("fields")
	d := bson.DocElem{"_id", 0}
	fields := []bson.DocElem{d}
	if v != "" {
		s := strings.Split(v, ",")
		for _, f := range s {
			val, err := validateField(f, safeFields)
			if err != nil {
				err.(*ErrorObject).CodeMinor = "invalid_selection_field"
				err.(*ErrorObject).Populate()
				return []bson.DocElem{d}, err
			}
			fields = append(fields, bson.DocElem{val, 1})
		}
	}
	return fields, nil
}

func getLimit(q url.Values) int {
	v := q.Get("limit")
	if v != "" {
		return v.ToInt()
	}
	return 100
}

func getOffset(q url.Values) int {
	v := q.Get("offset")
	if v != "" {
		return v.ToInt()
	}
	return 0
}

func getSort(q url.Values) string {
	v := q.Get("sort")
	if v != "" {
		return v
	}
	return "sourcedId"
}
