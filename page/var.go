package page

type Page struct {
    StartIdx      int
    EndIdx        int
    Previous      int
    Current       string
    Next          int
    CurrentTables []string
    Last          int
}
