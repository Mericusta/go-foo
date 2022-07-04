
for (( i=1; i <= $2; i++ )); do
    RES=`go clean -testcache && go test -v -timeout 30s -run ^$1$ game_server/pkg/game_testing/game_server`
    if [[ $RES =~ 'FAIL' ]]; then
        echo "$i FAIL"
    else
        echo "$i PASS"
        echo $RES
        exit
    fi
done