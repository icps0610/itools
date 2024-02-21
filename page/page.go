package page

import (
    "fmt"

    "iTools/script"
)

func GetPageTable(currentPage string, totalDataCount, pageCount int) Page {
    lastPage := totalDataCount / pageCount
    if totalDataCount%pageCount != 0 {
        lastPage++
    }

    currentPageI := script.To_i(currentPage)
    if currentPageI < 1 {
        currentPageI = 1
    }

    if currentPageI <= 1 {
        currentPage = "1"
    } else if currentPageI > lastPage {
        currentPageI = lastPage
        currentPage = fmt.Sprintf(`%v`, lastPage)
    }

    // 若超過範圍讓她變成-1 反白狀態
    peviousPage, nextPage := currentPageI-1, currentPageI+1

    if peviousPage == 0 {
        peviousPage = -1
    }
    if nextPage > lastPage {
        nextPage = -1
    }

    // 製作 currentTables
    var startPageI, endPageI int
    if currentPageI <= 3 {
        startPageI = 1
        endPageI = startPageI + 4
    } else if currentPageI > lastPage-3 {
        endPageI = lastPage
        startPageI = endPageI - 4
    } else {
        startPageI = currentPageI - 2
        endPageI = startPageI + 4
    }
    if lastPage < 5 {
        endPageI = lastPage
    }

    var currentTables []string
    for i := startPageI; i <= endPageI; i++ {
        currentTables = append(currentTables, fmt.Sprintf(`%v`, i))
    }

    // slice
    startIdx := (currentPageI - 1) * pageCount
    endIdx := startIdx + pageCount

    if endIdx > totalDataCount {
        endIdx = totalDataCount
    }
    if endIdx == 0 {
        startIdx = 0
    }

    return Page{startIdx, endIdx, peviousPage, currentPage, nextPage, currentTables, lastPage}
}
