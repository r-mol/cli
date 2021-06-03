// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package terminal

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/temporalio/shared-go/collection"
	"github.com/urfave/cli/v2"
)

// Paginate creates an interactive CLI mode to control the printing of items
func Paginate(c *cli.Context, paginationFn collection.PaginationFn, fields []string) error {
	more := c.Bool(FlagMore)
	isJSONView := c.Bool(FlagJSON)
	pageSize := c.Int(FlagPageSize)
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	iter := collection.NewPagingIterator(paginationFn)

	var pageItems []interface{}
	for iter.HasNext() {
		item, err := iter.Next()
		if err != nil {
			return err
		}

		pageItems = append(pageItems, item)
		if len(pageItems) == pageSize || !iter.HasNext() {
			if isJSONView {
				printJSON(pageItems)
			} else {
				printTable(pageItems, fields)
			}

			if !more || !showNextPage() {
				break
			}
			pageItems = pageItems[:0]
		}
	}

	return nil
}

func showNextPage() bool {
	fmt.Printf("Press %s to show next page, press %s to quit: ",
		color.GreenString("Enter"), color.RedString("any other key then Enter"))
	var input string
	_, _ = fmt.Scanln(&input)
	return strings.Trim(input, " ") == ""
}