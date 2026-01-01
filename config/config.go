package config

// UrlToCheck is the URL that will be checked when invoking the program.
// This variables should be set at compile time with linker flags, leaving it empty will produce a non-working binary.
var UrlToCheck string
