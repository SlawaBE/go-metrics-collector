#!/bin/bash

FILTERS="inuse_space inuse_objects alloc_space alloc_objects"
REPORT_FILE="optimization_report.md"

echo -e "# Optimization report\n" > ${REPORT_FILE}

for sample_index in ${FILTERS}
do
  {
    echo '## '${sample_index}
    echo
    echo '```text'
    go tool pprof -top -sample_index=${sample_index} -diff_base=profiles/base.pprof profiles/result.pprof
    echo '```'
    echo
  } >> ${REPORT_FILE}
done
