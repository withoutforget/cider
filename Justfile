[doc("Run Go app outside Docker")]
@run:
    go run ./cmd/app/main.go
[doc("Show lines of code")]
@lines:
    cloc -vcs=git .