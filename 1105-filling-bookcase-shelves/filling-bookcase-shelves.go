package _105_filling_bookcase_shelves

import "fmt"

/*
https://leetcode.com/problems/filling-bookcase-shelves/

1105. Filling Bookcase Shelves
Medium

145

7

Favorite

Share
We have a sequence of books: the i-th book has thickness books[i][0] and height books[i][1].

We want to place these books in order onto bookcase shelves that have total width shelf_width.

We choose some of the books to place on this shelf (such that the sum of their thickness is <= shelf_width), then build another level of shelf of the bookcase so that the total height of the bookcase has increased by the maximum height of the books we just put down.  We repeat this process until there are no more books to place.

Note again that at each step of the above process, the order of the books we place is the same order as the given sequence of books.  For example, if we have an ordered list of 5 books, we might place the first and second book onto the first shelf, the third book on the second shelf, and the fourth and fifth book on the last shelf.

Return the minimum possible height that the total bookshelf can be after placing shelves in this manner.



Example 1:


Input: books = [[1,1],[2,3],[2,3],[1,1],[1,1],[1,1],[1,2]], shelf_width = 4
Output: 6
Explanation:
The sum of the heights of the 3 shelves are 1 + 3 + 2 = 6.
Notice that book number 2 does not have to be on the first shelf.


Constraints:

1 <= books.length <= 1000
1 <= books[i][0] <= shelf_width <= 1000
1 <= books[i][1] <= 1000

*/

func minHeightShelves(books [][]int, shelf_width int) int {
	width := 0
	height := 0

	sumHeight := 0
	for index, value := range books {
		if index == 0 {
			width = value[0]
			height = value[1]
			fmt.Printf("\n %d", value)
			continue
		}
		nextWidth := width + value[0]
		if nextWidth > shelf_width {
			if value[0] > shelf_width {
				break
			}
			sumHeight += height
			fmt.Printf("\n sumheight:%d height:%d", sumHeight, height)
			width = value[0]
			height = value[1]
		} else {
			if height < value[1] {
				nextMax := nextMaxHeight(books, index, shelf_width)
				fmt.Printf("\nindex:%d nextMax:%d currentHeight:%d", index, nextMax, height)
				if height < nextMax {
					sumHeight += height
					fmt.Printf("\n sumheight:%d height:%d", sumHeight, height)
					width = value[0]
				} else {
					width += value[0]
				}
				height = value[1]
			} else {
				width += value[0]
			}
		}
		fmt.Printf("\n %d", value)
	}
	sumHeight += height
	fmt.Printf("\n sumheight:%d height:%d\n", sumHeight, height)
	return sumHeight
}

//find max height of the books after index within index's bounds,
func nextMaxHeight(books [][]int, index int, width int) int {
	w := books[index][0]
	height := 0
	length := len(books)
	begin := index + 1
	for {
		index++
		if index < length {
			w += books[index][0]
			if w > width {
				break
			}
			h := books[index][1]
			if h > height && index >= begin {
				height = h
			}
		} else {
			break
		}
	}
	return height
}
