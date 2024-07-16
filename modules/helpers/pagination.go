package helpers

import "strconv"

func Pagination(pageQuery, pageSizeQuery string) (int, int, error) {
	var page, pageSize int
	var err error

	if pageQuery == "" {
		page = 1
	} else {
		page, err = strconv.Atoi(pageQuery)
	}

	if pageSizeQuery == "" {
		pageSize = 10
	} else {
		pageSize, err = strconv.Atoi(pageSizeQuery)
	}

	if err != nil {
		return 0, 0, err
	}

	offset := pageSize * (page - 1)
	return pageSize, offset, nil

}
