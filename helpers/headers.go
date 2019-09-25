package helpers

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/url"
	"strconv"
)

// returns the relevant link headers for a given query
func GetLinkHeaders(totalCount int64, q url.Values) string {
	var link string
	offset, limit := parseOffsetLimit(q)
	no, nl := nextOffsetLimit(totalCount, offset, limit)
	if ok := testNextHeader(totalCount, no); ok {
		link += buildLinkHeader(q, "next", no, nl)
	}
	po, pl := prevOffsetLimit(totalCount, offset, limit)
	if ok := testPrevHeader(offset); ok {
		link += buildLinkHeader(q, "prev", po, pl)
	}
	return link
}

// gets the offset and limit values and converts to int
func parseOffsetLimit(q url.Values) (int64, int64) {
	sLimit := q.Get("limit")
	if sLimit == "" {
		sLimit = "100"
	}
	sOffset := q.Get("offset")
	if sOffset == "" {
		sOffset = "0"
	}
	limit, err := strconv.ParseInt(sLimit, 10, 64)
	if err != nil {
		log.Error(err)
	}
	offset, err := strconv.ParseInt(sOffset, 10, 64)
	if err != nil {
		log.Error(err)
	}
	return offset, limit
}

// calculates the offset and limit for the previous block of records
func prevOffsetLimit(totalCount, offset, limit int64) (int64, int64) {
	var prevOffset int64
	if offset > limit {
		prevOffset = offset - limit
	}
	prevLimit := limit
	if offset-limit <= 0 {
		prevLimit = offset
	}
	return prevOffset, prevLimit
}

// calculates the offset and limit for the next block of records
func nextOffsetLimit(totalCount, offset, limit int64) (int64, int64) {
	nextLimit := limit
	if totalCount < offset+limit {
		nextLimit = totalCount - offset
	}

	nextOffset := offset + nextLimit
	return nextOffset, nextLimit
}

// checks if previous header should exist
func testPrevHeader(offset int64) bool {
	if offset != 0 {
		return true
	}
	return false
}

// checks if next header should exist
func testNextHeader(totalCount, nextOffset int64) bool {
	if nextOffset != totalCount {
		return true
	}
	return false
}

// creates a link header string
func buildLinkHeader(q url.Values, ref string, limit, offset int64) string {
	return fmt.Sprintf(
		"<%v?limit=%v&offset=%v&%v>; ref=\"%v\",\n",
		"URLROOT",
		limit,
		offset,
		parseExistingParams(q),
		ref,
	)
}

// rebuilds original user query string params
// ignores limit/offset
func parseExistingParams(q url.Values) string {
	var s string
	fields := q.Get("fields")
	if fields != "" {
		s += "fields=" + fields + "&"
	}
	filter := q.Get("filter")
	if filter != "" {
		s += "filter=" + filter + "&"
	}
	sort := q.Get("sort")
	if sort != "" {
		s += "sort=" + sort + "&"
	}
	return s
}
