
BENCHMARK_FUNC=BenchmarkMapCapacityFoo
BENCHMARK_FUNC_FILE=mapfoo_benchmark_test.go
BENCHMARK_FUNC_PPROF=$BENCHMARK_FUNC.md

CLEAN=$1
if [ -e $BENCHMARK_FUNC_PPROF ]; then
    if [[ $CLEAN == true ]]; then
        echo "rm $BENCHMARK_FUNC_PPROF"
        rm $BENCHMARK_FUNC_PPROF
    fi
fi

if [ ! -e $BENCHMARK_FUNC_PPROF ]; then
    touch $BENCHMARK_FUNC_PPROF
    chmod +wr $BENCHMARK_FUNC_PPROF
    echo '|count|capacity|run_duration/op|allocs_memory/op|allocs_times/op|' >> $BENCHMARK_FUNC_PPROF
    echo '|:-:|:-:|:-:|:-:|:-:|' >> $BENCHMARK_FUNC_PPROF
fi

TEST_COUNT=$2
if [ -z $TEST_COUNT ]; then
    TEST_COUNT=1
fi

PREVIOUS_COUNT=0
PREVIOUS_CAPACITY=0

for (( CURRENT_CAPACITY=1; CURRENT_CAPACITY <= $TEST_COUNT; CURRENT_CAPACITY++ )); do
    for (( CURRENT_COUNT=1; CURRENT_COUNT <= $CURRENT_CAPACITY; CURRENT_COUNT++ )); do
        # echo "sed $PREVIOUS_COUNT, $PREVIOUS_CAPACITY to $CURRENT_COUNT, $CURRENT_CAPACITY"
        sed -i "s/count: $PREVIOUS_COUNT, capacity: $PREVIOUS_CAPACITY/count: $CURRENT_COUNT, capacity: $CURRENT_CAPACITY/g" $BENCHMARK_FUNC_FILE

        go clean -testcache
        echo '|'$CURRENT_COUNT'|'$CURRENT_CAPACITY'|'`go test -benchmem -run=^$ -bench ^$BENCHMARK_FUNC$ . | grep $BENCHMARK_FUNC | awk -F ' ' '{print $3" "$4"|"$5" "$6"|"$7" "$8}'`'|' >> $BENCHMARK_FUNC_PPROF
        
        PREVIOUS_COUNT=$CURRENT_COUNT
        PREVIOUS_CAPACITY=$CURRENT_CAPACITY
    done
done