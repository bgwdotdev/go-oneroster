

func getThing() { 
    table :=
    pubCols := 
    params := ParseURLParams(u, pubCols)
    rows := QueryDB(table, cols, params) 
    results := []map
    for rows.Next() {
        r :=FormatResults(rows)
        c := QueryNested
        p := QueryNested
        results = append(results, r)
    } 
    return render(results)
}

