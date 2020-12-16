package dblib

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
	"runtime"
	"strings"
	"sync/atomic"
)

type registeredQuery struct {
	file  string
	line  int
	query string
}

type registeredNamedQuery struct {
	file  string
	line  int
	query string
}

var disabledRegistering int32 = 0
var registeredQueries []registeredQuery
var registeredNamedQueries []registeredNamedQuery

// NewQuery registers a sqlx query
func NewQuery(q string) string {
	if atomic.LoadInt32(&disabledRegistering) != 0 {
		panic("Can NOT use NewQuery inside functions")
	}

	q = strings.TrimSpace(q)
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		panic("can't get caller info")
	}

	registeredQueries = append(registeredQueries, registeredQuery{
		file:  file,
		line:  line,
		query: q,
	})
	return q
}

// NewNamedQuery registers a named sqlx query
func NewNamedQuery(q string) string {
	if atomic.LoadInt32(&disabledRegistering) != 0 {
		panic("Can NOT use NewNamedQuery inside functions")
	}

	q = strings.TrimSpace(q)
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		panic("can't get caller info")
	}

	registeredNamedQueries = append(registeredNamedQueries, registeredNamedQuery{
		file:  file,
		line:  line,
		query: q,
	})
	return q
}

type checkResult struct {
	query registeredQuery
	err   error
}

type namedCheckResult struct {
	query registeredNamedQuery
	err   error
}

func checkNormalQueries(db *sqlx.DB, filter string, colorHighlight string, colorNone string) []checkResult {
	result := make([]checkResult, 0, len(registeredQueries))

	queries := registeredQueries
	var highlights []string
	if filter != "" {
		queries, highlights = fuzzyMatchNormal(registeredQueries, filter, colorHighlight, colorNone)
	}

	for i, q := range queries {
		stmt, err := db.Preparex(q.query)

		query := q
		if filter != "" {
			query.query = highlights[i]
		}

		result = append(result, checkResult{
			query: query,
			err:   err,
		})

		if err != nil {
			continue
		}
		_ = stmt.Close()
	}

	return result
}

func checkNamedQueries(db *sqlx.DB, filter string, colorHighlight, colorNone string) []namedCheckResult {
	result := make([]namedCheckResult, 0, len(registeredNamedQueries))

	queries := registeredNamedQueries
	var highlights []string
	if filter != "" {
		queries, highlights = fuzzyMatchNamed(registeredNamedQueries, filter, colorHighlight, colorNone)
	}

	for i, q := range queries {
		stmt, err := db.PrepareNamed(q.query)

		query := q
		if filter != "" {
			query.query = highlights[i]
		}

		result = append(result, namedCheckResult{
			query: query,
			err:   err,
		})

		if err != nil {
			continue
		}
		_ = stmt.Close()
	}

	return result
}

const (
	// ColorRed ...
	ColorRed = "\033[0;31m"
	// ColorOrange ...
	ColorOrange = "\033[0;33m"
	// ColorGreen ...
	ColorGreen = "\033[0;32m"
	// ColorNone ...
	ColorNone = "\033[0m"
)

var printFormat = strings.TrimSpace(`
=======================================================================
%s%s:%d%s
-----------------------------------------------------------------------
%s
`)

var printErrFormat = strings.TrimSpace(`
=======================================================================
%s%s:%d%s
-----------------------------------------------------------------------
%s
-----------------------------------------------------------------------
%s%v%s
`)

var endBar = strings.TrimSpace(`
=======================================================================
`)

// CheckOptions ...
type CheckOptions struct {
	Filter string
	// print when even not have error
	EnablePrint  bool
	DisableColor bool
}

// FinishRegisterQueries prevents NewQuery / NewNamedQuery from using inside functions
func FinishRegisterQueries() {
	atomic.StoreInt32(&disabledRegistering, 1)
}

// CheckQueries validates syntax of all registered queries
func CheckQueries(db *sqlx.DB, opts CheckOptions) {
	colorRed := ColorRed
	colorOrange := ColorOrange
	colorGreen := ColorGreen
	colorNone := ColorNone
	if opts.DisableColor {
		colorRed = ""
		colorOrange = ""
		colorGreen = ""
		colorNone = ""
	}

	normalResults := checkNormalQueries(db, opts.Filter, colorGreen, colorNone)
	namedResults := checkNamedQueries(db, opts.Filter, colorGreen, colorNone)

	exitValue := 0
	printEndBar := false
	for _, e := range normalResults {
		if e.err != nil {
			exitValue = -1
			fmt.Printf(printErrFormat,
				colorOrange, e.query.file, e.query.line, colorNone,
				e.query.query,
				colorRed, e.err, colorNone,
			)
			fmt.Println()

			printEndBar = true
		} else if opts.EnablePrint {
			fmt.Printf(printFormat,
				colorOrange, e.query.file, e.query.line, colorNone,
				e.query.query,
			)
			fmt.Println()

			printEndBar = true
		}
	}

	for _, e := range namedResults {
		if e.err != nil {
			exitValue = -1
			fmt.Printf(printErrFormat,
				colorOrange, e.query.file, e.query.line, colorNone,
				e.query.query,
				colorRed, e.err, colorNone,
			)
			fmt.Println()

			printEndBar = true
		} else if opts.EnablePrint {
			fmt.Printf(printFormat,
				colorOrange, e.query.file, e.query.line, colorNone,
				e.query.query,
			)
			fmt.Println()

			printEndBar = true
		}
	}

	if printEndBar {
		fmt.Println(endBar)
	}

	fmt.Println("Number of queries processed:", len(normalResults)+len(namedResults))

	if exitValue != 0 {
		os.Exit(exitValue)
	}
}
