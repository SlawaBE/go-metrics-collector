// Package main implements a multichecker — a configurable static analyzer
// for the go-metrics-collector project.
//
// # How to run
//
// Build the binary from the cmd/staticlint directory:
//
//	go build -o staticlint ./cmd/staticlint
//
// Then run it, passing the packages to analyze:
//
//	./staticlint ./...
//	./staticlint ./cmd/agent ./internal/handler
//
// The analyzer behaves like the standard go vet tool: it performs type and
// AST analysis of Go source and prints diagnostics to the standard output. Use
// the -json flag to output reports in JSON format for external tools:
//
//	./staticlint -json ./...
//
// The list of analyzers can be obtained through the common -help flag:
//
//	./staticlint -help
//
// # Analyzer composition
//
// The multichecker combines several groups of analyzers:
//
//  1. The standard analyzers from the golang.org/x/tools/go/analysis/passes
//     package used by the go vet tool. They check correctness, common errors
//     and warnings (see the list below).
//
//  2. All SA-class analyzers from staticcheck.io (staticcheck.Analyzers).
//     The SA class covers checks for possible errors: incorrect conditions,
//     unreachable code, errors working with pointers, channels and so on.
//
//  3. Analyzers from other staticcheck.io classes. The build includes the S
//     (simple, code simplification), ST (stylecheck, style issues) and QF
//     (quickfix, automatically fixable findings) classes.
//
//  4. Third-party public analyzers:
//     errcheck (github.com/kisielk/errcheck) — finds unchecked errors;
//     whitespace (github.com/ultraware/whitespace) — finds extra whitespace
//     lines at the start and end of blocks.
//
//  5. The custom noexit analyzer, which forbids a direct os.Exit call in the
//     main function of the main package.
//
// # Standard go vet analyzers
//
// The build includes the following standard analyzers from the
// golang.org/x/tools/go/analysis/passes package:
//
//   - appends — checks that append is not assigned to the source slice itself
//     when the ranges overlap, which could corrupt data;
//   - asmdecl — checks that assembly function declarations match their Go
//     declarations;
//   - assign — detects erroneous assignments (including self-assignment);
//   - atomic — checks the correctness of the sync/atomic package usage;
//   - atomicalign — checks the alignment of atomic operation arguments;
//   - bools — finds erroneous conditions with logical operators (for example,
//     a == b || a == b);
//   - buildtag — checks the correctness of //go:build and // +`build` comments;
//   - cgocall — detects violations when using cgo from goroutines;
//   - composite — checks the correctness of composite type literals (including
//     the absence of disallowed keys);
//   - copylock — forbids copying values that contain sync.Mutex and similar
//     primitives;
//   - deepequalerrors — finds comparisons via reflect.DeepEqual involving
//     values of the error type;
//   - defers — checks common mistakes when using defer inside loops;
//   - directive — checks the correctness of special directives in comments
//     (//go:xxx);
//   - errorsas — checks that the second argument of errors.As is a pointer to
//     a type implementing error;
//   - fieldalignment — finds redundant padding bytes in structs;
//   - findcall — finds calls to the given function (via the -name flag);
//   - framepointer — checks the correctness of the build for the amd64
//     architecture;
//   - httpresponse — finds the use of an http response result before checking
//     the error;
//   - ifaceassert — finds redundant interface assertions when the type is
//     already known;
//   - loopclosure — detects capture of loop variables in closures executed in
//     new goroutines;
//   - lostcancel — finds lost cancel() calls when using context;
//   - nilfunc — finds function comparisons with nil via ==;
//   - printf — checks the correctness of format strings and arguments in
//     Printf-like functions;
//   - reflectvaluecompare — finds reflect.Value comparisons via ==;
//   - shadow — finds shadowing variable declarations;
//   - shift — checks the correctness of shifts, including shifts by a
//     negative or too large amount;
//   - sigchanyzer — finds passing an unbuffered channel to signal.Notify,
//     which can cause signal loss;
//   - slog — checks the correctness of pkg/log/slog usage (key-value pairing);
//   - sortslice — finds sort.Slice calls with slices when the slice and the
//     function are the same and this could lead to problems;
//   - sqlrowserr — checks that rows.Err() is checked after rows.Scan();
//   - stdmethods — checks that signatures match standard methods (Error,
//     String and so on);
//   - stdversion — finds the use of symbols unavailable in the declared Go
//     version;
//   - stringintconv — finds conversions between strings and integers;
//   - structtag — checks the correctness of struct field tags;
//   - testinggoroutine — finds Fatal/Error method calls from goroutines during
//     tests;
//   - tests — checks the correctness of tests (signatures, function names);
//   - timeformat — checks that correct 2006/01/02 templates are used in
//     time.Parse/Format;
//   - unmarshal — finds json.Unmarshal calls with wrong (non-pointer) types;
//   - unreachable — finds unreachable code;
//   - unsafeptr — checks the correctness of converting integers to pointers;
//   - unusedresult — finds calls to functions whose result should be checked
//     or used;
//   - unusedwrite — finds writes to slices and maps that are never read.
package main

import (
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
	"golang.org/x/tools/go/analysis/passes/appends"
	"golang.org/x/tools/go/analysis/passes/asmdecl"
	"golang.org/x/tools/go/analysis/passes/assign"
	"golang.org/x/tools/go/analysis/passes/atomic"
	"golang.org/x/tools/go/analysis/passes/atomicalign"
	"golang.org/x/tools/go/analysis/passes/bools"
	"golang.org/x/tools/go/analysis/passes/buildtag"
	"golang.org/x/tools/go/analysis/passes/cgocall"
	"golang.org/x/tools/go/analysis/passes/composite"
	"golang.org/x/tools/go/analysis/passes/copylock"
	"golang.org/x/tools/go/analysis/passes/deepequalerrors"
	"golang.org/x/tools/go/analysis/passes/defers"
	"golang.org/x/tools/go/analysis/passes/directive"
	"golang.org/x/tools/go/analysis/passes/errorsas"
	"golang.org/x/tools/go/analysis/passes/fieldalignment"
	"golang.org/x/tools/go/analysis/passes/findcall"
	"golang.org/x/tools/go/analysis/passes/framepointer"
	"golang.org/x/tools/go/analysis/passes/httpresponse"
	"golang.org/x/tools/go/analysis/passes/ifaceassert"
	"golang.org/x/tools/go/analysis/passes/loopclosure"
	"golang.org/x/tools/go/analysis/passes/lostcancel"
	"golang.org/x/tools/go/analysis/passes/nilfunc"
	"golang.org/x/tools/go/analysis/passes/printf"
	"golang.org/x/tools/go/analysis/passes/reflectvaluecompare"
	"golang.org/x/tools/go/analysis/passes/shadow"
	"golang.org/x/tools/go/analysis/passes/shift"
	"golang.org/x/tools/go/analysis/passes/sigchanyzer"
	"golang.org/x/tools/go/analysis/passes/slog"
	"golang.org/x/tools/go/analysis/passes/sortslice"
	"golang.org/x/tools/go/analysis/passes/sqlrowserr"
	"golang.org/x/tools/go/analysis/passes/stdmethods"
	"golang.org/x/tools/go/analysis/passes/stdversion"
	"golang.org/x/tools/go/analysis/passes/stringintconv"
	"golang.org/x/tools/go/analysis/passes/structtag"
	"golang.org/x/tools/go/analysis/passes/testinggoroutine"
	"golang.org/x/tools/go/analysis/passes/tests"
	"golang.org/x/tools/go/analysis/passes/timeformat"
	"golang.org/x/tools/go/analysis/passes/unmarshal"
	"golang.org/x/tools/go/analysis/passes/unreachable"
	"golang.org/x/tools/go/analysis/passes/unsafeptr"
	"golang.org/x/tools/go/analysis/passes/unusedresult"
	"golang.org/x/tools/go/analysis/passes/unusedwrite"

	"github.com/SlawaBE/go-metrics-collector/cmd/staticlint/noexit"
	"github.com/kisielk/errcheck/errcheck"
	"github.com/ultraware/whitespace"
	"honnef.co/go/tools/analysis/lint"
	"honnef.co/go/tools/quickfix/qf1001"
	"honnef.co/go/tools/simple/s1000"
	"honnef.co/go/tools/staticcheck"
	"honnef.co/go/tools/stylecheck/st1000"
)

func main() {
	multichecker.Main(analyzers()...)
}

func analyzers() []*analysis.Analyzer {
	analyzersList := []*analysis.Analyzer{
		// Standard go vet analyzers.
		appends.Analyzer,
		asmdecl.Analyzer,
		assign.Analyzer,
		atomic.Analyzer,
		atomicalign.Analyzer,
		bools.Analyzer,
		buildtag.Analyzer,
		cgocall.Analyzer,
		composite.Analyzer,
		copylock.Analyzer,
		deepequalerrors.Analyzer,
		defers.Analyzer,
		directive.Analyzer,
		errorsas.Analyzer,
		fieldalignment.Analyzer,
		findcall.Analyzer,
		framepointer.Analyzer,
		httpresponse.Analyzer,
		ifaceassert.Analyzer,
		loopclosure.Analyzer,
		lostcancel.Analyzer,
		nilfunc.Analyzer,
		printf.Analyzer,
		reflectvaluecompare.Analyzer,
		shadow.Analyzer,
		shift.Analyzer,
		sigchanyzer.Analyzer,
		slog.Analyzer,
		sortslice.Analyzer,
		sqlrowserr.Analyzer,
		stdmethods.Analyzer,
		stdversion.Analyzer,
		stringintconv.Analyzer,
		structtag.Analyzer,
		testinggoroutine.Analyzer,
		tests.Analyzer,
		timeformat.Analyzer,
		unmarshal.Analyzer,
		unreachable.Analyzer,
		unsafeptr.Analyzer,
		unusedresult.Analyzer,
		unusedwrite.Analyzer,

		// Third-party public analyzers.
		errcheck.Analyzer,
		whitespace.NewAnalyzer(&whitespace.Settings{}),

		// Custom analyzer.
		noexit.Analyzer,
	}

	// All SA-class analyzers (staticcheck).
	analyzersList = append(analyzersList, toGoAnalyzers(staticcheck.Analyzers)...)
	// One of S-class analyzer (simple).
	analyzersList = append(analyzersList, toGoAnalyzer(s1000.SCAnalyzer))
	// One of ST-class analyzer (stylecheck).
	analyzersList = append(analyzersList, toGoAnalyzer(st1000.SCAnalyzer))
	// One of QF-class analyzer (quickfix).
	analyzersList = append(analyzersList, toGoAnalyzer(qf1001.SCAnalyzer))

	return analyzersList
}

func toGoAnalyzers(list []*lint.Analyzer) []*analysis.Analyzer {
	result := make([]*analysis.Analyzer, 0, len(list))
	for _, a := range list {
		result = append(result, toGoAnalyzer(a))
	}
	return result
}

func toGoAnalyzer(a *lint.Analyzer) *analysis.Analyzer {
	lint.InitializeAnalyzer(a)
	return a.Analyzer
}
