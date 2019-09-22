package helpers

import ()

// returns the database query parameters based on user url request
// e.g. ?limit=1&fields=id
func GetOptions() (*mongo.FindOptions, []error) {
	var errP []error
	projection, err := getFields(url)
	if err != nil {
		errP = append(errP, err)
	}
	sort, err := getSort(url)
	if err != nil {
		errP = append(errP, err)
	}
	o := options.
		Find().
		SetLimit(getLimit(url)).
		SetSkip(getOffset(url)).
		SetSort(bson.D{sort, 1}).
		SetProjection(bson.D{projection})
	return o, ep
}

// returns the filtering query based on user url request
// e.g. ?filter=id>='1'
func GetFilters(q url.Values, safeFields []string) (bson.D, error) {
	v := q.Get("filter")
	if v != "" {
		lo := parseFilterLo(v)
		var filter []bson.DocElem
		queries := splitFilterQuery(v)
		for _, f := range queries {
			ff, err := parseFilterField(f, safeFields)
			if err != nil {
				return bson.D{}, err
			}
			fp := parseFilterPredicate(f)
			fv := parseFilterValue(f)
			filter = append(filter, bson.DocElem{ff, bson.D{fp, fv}})
		}
		return bson.D{{lo, filter}}, nil
	}
	return bson.D{}, nil
}

// returns a bson of field filtering for mongodb from url
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

// returns the user requested field to sort by
// validated against a field whitelist
func getSort(q url.Values, safeFields []string) (string, error) {
	v := q.Get("sort")
	d := "sourcedId"
	if v != "" {
		f, err := validateField(v, safeCol)
		if err != nil {
			err.(*ErrorObject).CodeMinor = "invalid_sort_field"
			err.(*ErrorObject).Populate()
			return d, err
		}
		return f, nil
	}
	return d
}

// returns the max doc count requested by user
func getLimit(q url.Values) int {
	v := q.Get("limit")
	if v != "" {
		return v.ToInt()
	}
	return 100
}

// returns the doc skip requested by user
func getOffset(q url.Values) int {
	v := q.Get("offset")
	if v != "" {
		return v.ToInt()
	}
	return 0
}
