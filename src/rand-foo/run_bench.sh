
for (( i=1; i <= $1; i++ )); do
    RES=`go clean -testcache && go test -v -timeout 30s -run ^TestGetRandSlice$ go-foo/rand-foo`
    if [[ $RES =~ 'PASS' ]]; then
        echo 'PASS'
        echo $RES
        exit
    fi
done