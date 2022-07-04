#!/bin/bash

BENCHMARK_FUNC=$1

if [[ -z $BENCHMARK_FUNC ]]; then
    echo "need benchmark function expr as \$1"
    exit
fi

BENCHMARK_FUNC_PPROF=${BENCHMARK_FUNC}.md

CLEAN=$2

if [ -e $BENCHMARK_FUNC_PPROF ]; then
    if [[ $CLEAN == true ]]; then
        echo "clean $BENCHMARK_FUNC_PPROF"
        rm $BENCHMARK_FUNC_PPROF
    fi
fi

if [ ! -e $BENCHMARK_FUNC_PPROF ]; then
    touch $BENCHMARK_FUNC_PPROF
    chmod +wr $BENCHMARK_FUNC_PPROF
    echo '|date|commit|run_duration/op|allocs_memory/op|allocs_times/op|' >> $BENCHMARK_FUNC_PPROF
    echo '|:-:|:-:|:-:|:-:|:-:|' >> $BENCHMARK_FUNC_PPROF
fi

go clean -testcache
echo '|'`date`'|'`git log --pretty=oneline -1`'|'`go test -benchmem -run=^$ -bench ^$1$ $3 | grep $1 | awk -F ' ' '{print $3" "$4"|"$5" "$6"|"$7" "$8}'`'|' >> $BENCHMARK_FUNC_PPROF
