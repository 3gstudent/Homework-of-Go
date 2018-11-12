@echo off
if not exist %1 (
	SET OUTNAME=%1
	SET CGO_ENABLED=0
	SET GOOS=android
	SET GOARCH=arm
	go build -o %OUTNAME%_%GOOS%_%GOARCH%

	SET GOOS=darwin
	SET GOARCH=386
	go build -o %OUTNAME%_%GOOS%_%GOARCH%

	SET GOOS=darwin
	SET GOARCH=amd64
	go build -o %OUTNAME%_%GOOS%_%GOARCH%

	SET GOOS=darwin
	SET GOARCH=arm
	go build -o %OUTNAME%_%GOOS%_%GOARCH%

	SET GOOS=darwin
	SET GOARCH=arm64
	go build -o %OUTNAME%_%GOOS%_%GOARCH%

	SET GOOS=dragonfly
	SET GOARCH=amd64
	go build -o %OUTNAME%_%GOOS%_%GOARCH%

	SET GOOS=freebsd
	SET GOARCH=386
	go build -o %OUTNAME%_%GOOS%_%GOARCH%

	SET GOOS=freebsd
	SET GOARCH=amd64
	go build -o %OUTNAME%_%GOOS%_%GOARCH%

	SET GOOS=freebsd
	SET GOARCH=arm
	go build -o %OUTNAME%_%GOOS%_%GOARCH%

	SET GOOS=linux
	SET GOARCH=386
	go build -o %OUTNAME%_%GOOS%_%GOARCH%

	SET GOOS=linux
	SET GOARCH=amd64
	go build -o %OUTNAME%_%GOOS%_%GOARCH%

	SET GOOS=linux
	SET GOARCH=arm
	go build -o %OUTNAME%_%GOOS%_%GOARCH%

	SET GOOS=linux
	SET GOARCH=arm64
	go build -o %OUTNAME%_%GOOS%_%GOARCH%

	SET GOOS=linux
	SET GOARCH=ppc64
	go build -o %OUTNAME%_%GOOS%_%GOARCH%

	SET GOOS=linux
	SET GOARCH=ppc64le
	go build -o %OUTNAME%_%GOOS%_%GOARCH%

	SET GOOS=linux
	SET GOARCH=mips
	go build -o %OUTNAME%_%GOOS%_%GOARCH%

	SET GOOS=linux
	SET GOARCH=mipsle
	go build -o %OUTNAME%_%GOOS%_%GOARCH%

	SET GOOS=linux
	SET GOARCH=mips64
	go build -o %OUTNAME%_%GOOS%_%GOARCH%

	SET GOOS=linux
	SET GOARCH=mips64le
	go build -o %OUTNAME%_%GOOS%_%GOARCH%

	SET GOOS=linux
	SET GOARCH=s390x
	go build -o %OUTNAME%_%GOOS%_%GOARCH%

	SET GOOS=netbsd
	SET GOARCH=386
	go build -o %OUTNAME%_%GOOS%_%GOARCH%

	SET GOOS=netbsd
	SET GOARCH=amd64
	go build -o %OUTNAME%_%GOOS%_%GOARCH%

	SET GOOS=netbsd
	SET GOARCH=arm
	go build -o %OUTNAME%_%GOOS%_%GOARCH%

	SET GOOS=openbsd
	SET GOARCH=386
	go build -o %OUTNAME%_%GOOS%_%GOARCH%

	SET GOOS=openbsd
	SET GOARCH=amd64
	go build -o %OUTNAME%_%GOOS%_%GOARCH%

	SET GOOS=openbsd
	SET GOARCH=arm
	go build -o %OUTNAME%_%GOOS%_%GOARCH%

	SET GOOS=plan9
	SET GOARCH=386
	go build -o %OUTNAME%_%GOOS%_%GOARCH%

	SET GOOS=plan9
	SET GOARCH=amd64
	go build -o %OUTNAME%_%GOOS%_%GOARCH%

	SET GOOS=solaris
	SET GOARCH=amd64
	go build -o %OUTNAME%_%GOOS%_%GOARCH%

	SET GOOS=windows
	SET GOARCH=386
	go build -o %OUTNAME%_%GOOS%_%GOARCH%

	SET GOOS=windows
	SET GOARCH=amd64
	go build -o %OUTNAME%_%GOOS%_%GOARCH%
)


