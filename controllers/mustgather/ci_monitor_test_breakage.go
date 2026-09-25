package mustgather

// This file is intentionally broken as part of sandbox testing for
// openshift-eng/oape-ai-e2e#67 (CI-Monitor). It is isolated in its own
// file so it does not modify any existing production code, and is not
// referenced from anywhere else. Safe to delete.
//
// undefinedHelperFunction does not exist anywhere in this package or its
// imports, so this triggers a genuine, non-trivial compile error that
// CI-Monitor should classify as build-failure and route to "investigate"
// rather than attempt an automatic fix.
var ciMonitorTestBreakageResult = undefinedHelperFunction()
