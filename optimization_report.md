# Optimization report

## inuse_space

```text
File: server.exe
Build ID: D:\work\projects\go-metrics-collector\cmd\server\server.exe2026-10-07 10:38:59.6839136 +0300 MSK
Type: inuse_space
Time: 2026-10-05 19:12:54 MSK
Showing nodes accounting for 12652.60kB, 134.15% of 9431.91kB total
Dropped 13 nodes (cum <= 47.16kB)
      flat  flat%   sum%        cum   cum%
 6318.10kB 66.99% 66.99% 13668.29kB 144.92%  compress/flate.NewWriter (inline)
 5643.69kB 59.84% 126.82%  7350.19kB 77.93%  compress/flate.(*compressor).init
 2800.01kB 29.69% 156.51%  2800.01kB 29.69%  compress/flate.newDeflateFast (inline)
-1093.51kB 11.59% 144.92% -1093.51kB 11.59%  compress/flate.(*compressor).initDeflate (inline)
-1024.16kB 10.86% 134.06% -1024.16kB 10.86%  runtime.mallocgc
  525.43kB  5.57% 139.63%   525.43kB  5.57%  github.com/jackc/pgx/v5/internal/stmtcache.NewLRUCache (inline)
 -516.01kB  5.47% 134.16%  -516.01kB  5.47%  github.com/jackc/pgx/v5/internal/iobufpool.init.0.func1
    -514kB  5.45% 128.71%     -514kB  5.45%  bufio.NewReaderSize (inline)
  513.12kB  5.44% 134.15%   513.12kB  5.44%  compress/flate.(*huffmanEncoder).generate
 -512.12kB  5.43% 128.72%  -512.12kB  5.43%  net/http.ListenAndServe (inline)
 -512.09kB  5.43% 123.29%  -512.09kB  5.43%  compress/flate.newHuffmanEncoder (inline)
  512.08kB  5.43% 128.72%   512.08kB  5.43%  compress/gzip.NewWriterLevel
  512.05kB  5.43% 134.15%   512.05kB  5.43%  context.(*cancelCtx).Done
  512.05kB  5.43% 139.58%   512.05kB  5.43%  github.com/golang-migrate/migrate/v4.(*Migrate).lock.func2
 -512.05kB  5.43% 134.15%  -512.05kB  5.43%  net/textproto.readMIMEHeader
         0     0% 134.15%     -514kB  5.45%  bufio.NewReader (inline)
         0     0% 134.15%   513.12kB  5.44%  compress/flate.(*Writer).Close (inline)
         0     0% 134.15%   513.12kB  5.44%  compress/flate.(*compressor).close
         0     0% 134.15%   513.12kB  5.44%  compress/flate.(*compressor).encSpeed
         0     0% 134.15%   513.12kB  5.44%  compress/flate.(*huffmanBitWriter).indexTokens
         0     0% 134.15%   513.12kB  5.44%  compress/flate.(*huffmanBitWriter).writeBlockDynamic
         0     0% 134.15%  -512.09kB  5.43%  compress/flate.generateFixedLiteralEncoding
         0     0% 134.15%  -512.09kB  5.43%  compress/flate.init
         0     0% 134.15%   513.12kB  5.44%  compress/gzip.(*Writer).Close
         0     0% 134.15% 13668.29kB 144.92%  compress/gzip.(*Writer).Write
         0     0% 134.15%   512.05kB  5.43%  database/sql.(*DB).connectionOpener
         0     0% 134.15% -1996.09kB 21.16%  github.com/SlawaBE/go-metrics-collector/internal/gzip.(*compressWriter).Write
         0     0% 134.15% 16185.89kB 171.61%  github.com/SlawaBE/go-metrics-collector/internal/handler.(*ListMetricHandler).ServeHTTP
         0     0% 134.15%   513.12kB  5.44%  github.com/SlawaBE/go-metrics-collector/internal/middleware.(*gzipResponseWriter).Close
         0     0% 134.15% 15664.39kB 166.08%  github.com/SlawaBE/go-metrics-collector/internal/middleware.(*gzipResponseWriter).Write
         0     0% 134.15%   512.08kB  5.43%  github.com/SlawaBE/go-metrics-collector/internal/middleware.(*gzipResponseWriter).WriteHeader
         0     0% 134.15%   512.08kB  5.43%  github.com/SlawaBE/go-metrics-collector/internal/middleware.(*gzipResponseWriter).start
         0     0% 134.15%   512.08kB  5.43%  github.com/SlawaBE/go-metrics-collector/internal/middleware.init.func1
         0     0% 134.15% 14702.92kB 155.88%  github.com/SlawaBE/go-metrics-collector/internal/server.Run.GZip.func4
         0     0% 134.15% 14702.92kB 155.88%  github.com/SlawaBE/go-metrics-collector/internal/server.Run.RequestLogger.func5
         0     0% 134.15%  -512.12kB  5.43%  github.com/SlawaBE/go-metrics-collector/internal/server.Run.func1
         0     0% 134.15% 16185.89kB 171.61%  github.com/go-chi/chi/v5.(*Mux).ServeHTTP
         0     0% 134.15% 16185.89kB 171.61%  github.com/go-chi/chi/v5.(*Mux).routeHTTP
         0     0% 134.15%  -516.01kB  5.47%  github.com/jackc/pgx/v5/internal/iobufpool.Get
         0     0% 134.15%  -516.01kB  5.47%  github.com/jackc/pgx/v5/pgconn.ConnectConfig
         0     0% 134.15%  -516.01kB  5.47%  github.com/jackc/pgx/v5/pgconn.connectOne
         0     0% 134.15%  -516.01kB  5.47%  github.com/jackc/pgx/v5/pgconn.connectPreferred
         0     0% 134.15%  -516.01kB  5.47%  github.com/jackc/pgx/v5/pgproto3.NewFrontend
         0     0% 134.15%  -516.01kB  5.47%  github.com/jackc/pgx/v5/pgproto3.newChunkReader (inline)
         0     0% 134.15% 15664.39kB 166.08%  io.WriteString
         0     0% 134.15%  -512.05kB  5.43%  net/http.(*conn).readRequest
         0     0% 134.15% 13676.86kB 145.01%  net/http.(*conn).serve
         0     0% 134.15% 14702.92kB 155.88%  net/http.HandlerFunc.ServeHTTP
         0     0% 134.15%     -514kB  5.45%  net/http.newBufioReader
         0     0% 134.15%  -512.05kB  5.43%  net/http.readRequest
         0     0% 134.15% 14702.92kB 155.88%  net/http.serverHandler.ServeHTTP
         0     0% 134.15%  -512.05kB  5.43%  net/textproto.(*Reader).ReadMIMEHeader (inline)
         0     0% 134.15%   512.56kB  5.43%  runtime.allocm
         0     0% 134.15%  -512.09kB  5.43%  runtime.doInit (inline)
         0     0% 134.15%  -512.09kB  5.43%  runtime.doInit1
         0     0% 134.15%    -1026kB 10.88%  runtime.findRunnable
         0     0% 134.15%  -512.25kB  5.43%  runtime.gcBgMarkWorker
         0     0% 134.15%   512.56kB  5.43%  runtime.gcStart.func4
         0     0% 134.15%    -1026kB 10.88%  runtime.injectglist
         0     0% 134.15%    -1026kB 10.88%  runtime.injectglist.func1
         0     0% 134.15%   512.56kB  5.43%  runtime.mProfStackInit (inline)
         0     0% 134.15%  -512.09kB  5.43%  runtime.main
         0     0% 134.15%   512.56kB  5.43%  runtime.makeProfStackFP (inline)
         0     0% 134.15%   512.56kB  5.43%  runtime.makeslice
         0     0% 134.15% -1024.47kB 10.86%  runtime.malg
         0     0% 134.15%     -513kB  5.44%  runtime.mcall
         0     0% 134.15%   512.56kB  5.43%  runtime.mcommoninit
         0     0% 134.15%      513kB  5.44%  runtime.mstart
         0     0% 134.15%      513kB  5.44%  runtime.mstart0
         0     0% 134.15%      513kB  5.44%  runtime.mstart1
         0     0% 134.15%   512.56kB  5.43%  runtime.newm
         0     0% 134.15% -1536.72kB 16.29%  runtime.newobject
         0     0% 134.15% -1024.47kB 10.86%  runtime.newproc.func1
         0     0% 134.15% -1024.47kB 10.86%  runtime.newproc1
         0     0% 134.15%     -513kB  5.44%  runtime.park_m
         0     0% 134.15%     1026kB 10.88%  runtime.resetspinning
         0     0% 134.15%   512.56kB  5.43%  runtime.startTheWorldWithSema
         0     0% 134.15%  -511.91kB  5.43%  runtime.systemstack
         0     0% 134.15%     1026kB 10.88%  runtime.wakep
```

## inuse_objects

```text
File: server.exe
Build ID: D:\work\projects\go-metrics-collector\cmd\server\server.exe2026-10-07 10:38:59.6839136 +0300 MSK
Type: inuse_objects
Time: 2026-10-05 19:12:54 MSK
Showing nodes accounting for 294, 2.01% of 14606 total
      flat  flat%   sum%        cum   cum%
      4681 32.05% 32.05%       4681 32.05%  context.(*cancelCtx).Done
      4681 32.05% 64.10%       4681 32.05%  github.com/golang-migrate/migrate/v4.(*Migrate).lock.func2
     -4681 32.05% 32.05%      -4681 32.05%  net/textproto.readMIMEHeader
      3277 22.44% 54.48%       3277 22.44%  compress/gzip.NewWriterLevel
     -2979 20.40% 34.09%      -2979 20.40%  compress/flate.newHuffmanEncoder (inline)
     -2754 18.86% 15.23%      -2754 18.86%  runtime.mallocgc
     -2048 14.02%  1.21%      -2048 14.02%  net/http.ListenAndServe (inline)
       228  1.56%  2.77%        228  1.56%  compress/flate.(*huffmanEncoder).generate
      -128  0.88%  1.90%       -128  0.88%  bufio.NewReaderSize (inline)
       -64  0.44%  1.46%        -64  0.44%  github.com/jackc/pgx/v5/internal/iobufpool.init.0.func1
        34  0.23%  1.69%         52  0.36%  compress/flate.(*compressor).init
        33  0.23%  1.92%         33  0.23%  compress/flate.newDeflateFast (inline)
        19  0.13%  2.05%         19  0.13%  github.com/jackc/pgx/v5/internal/stmtcache.NewLRUCache (inline)
       -15   0.1%  1.94%        -15   0.1%  compress/flate.(*compressor).initDeflate (inline)
        10 0.068%  2.01%         62  0.42%  compress/flate.NewWriter (inline)
         0     0%  2.01%       -128  0.88%  bufio.NewReader (inline)
         0     0%  2.01%        228  1.56%  compress/flate.(*Writer).Close (inline)
         0     0%  2.01%        228  1.56%  compress/flate.(*compressor).close
         0     0%  2.01%        228  1.56%  compress/flate.(*compressor).encSpeed
         0     0%  2.01%        228  1.56%  compress/flate.(*huffmanBitWriter).indexTokens
         0     0%  2.01%        228  1.56%  compress/flate.(*huffmanBitWriter).writeBlockDynamic
         0     0%  2.01%      -2979 20.40%  compress/flate.generateFixedLiteralEncoding
         0     0%  2.01%      -2979 20.40%  compress/flate.init
         0     0%  2.01%        228  1.56%  compress/gzip.(*Writer).Close
         0     0%  2.01%         62  0.42%  compress/gzip.(*Writer).Write
         0     0%  2.01%        -45  0.31%  database/sql.(*DB).QueryContext
         0     0%  2.01%        -45  0.31%  database/sql.(*DB).QueryContext.func1
         0     0%  2.01%        -45  0.31%  database/sql.(*DB).conn
         0     0%  2.01%       4681 32.05%  database/sql.(*DB).connectionOpener
         0     0%  2.01%        -45  0.31%  database/sql.(*DB).query
         0     0%  2.01%        -45  0.31%  database/sql.(*DB).retry
         0     0%  2.01%        -16  0.11%  github.com/SlawaBE/go-metrics-collector/internal/gzip.(*compressWriter).Write
         0     0%  2.01%       3310 22.66%  github.com/SlawaBE/go-metrics-collector/internal/handler.(*ListMetricHandler).ServeHTTP
         0     0%  2.01%        228  1.56%  github.com/SlawaBE/go-metrics-collector/internal/middleware.(*gzipResponseWriter).Close
         0     0%  2.01%         78  0.53%  github.com/SlawaBE/go-metrics-collector/internal/middleware.(*gzipResponseWriter).Write
         0     0%  2.01%       3277 22.44%  github.com/SlawaBE/go-metrics-collector/internal/middleware.(*gzipResponseWriter).WriteHeader
         0     0%  2.01%       3277 22.44%  github.com/SlawaBE/go-metrics-collector/internal/middleware.(*gzipResponseWriter).start
         0     0%  2.01%       3277 22.44%  github.com/SlawaBE/go-metrics-collector/internal/middleware.init.func1
         0     0%  2.01%        -45  0.31%  github.com/SlawaBE/go-metrics-collector/internal/retry.RetryWithBackoff
         0     0%  2.01%       3522 24.11%  github.com/SlawaBE/go-metrics-collector/internal/server.Run.GZip.func4
         0     0%  2.01%       3522 24.11%  github.com/SlawaBE/go-metrics-collector/internal/server.Run.RequestLogger.func5
         0     0%  2.01%      -2048 14.02%  github.com/SlawaBE/go-metrics-collector/internal/server.Run.func1
         0     0%  2.01%        -45  0.31%  github.com/SlawaBE/go-metrics-collector/internal/service.(*MetricsService).List
         0     0%  2.01%        -45  0.31%  github.com/SlawaBE/go-metrics-collector/internal/storage.(*DBStorage).GetValues
         0     0%  2.01%        -45  0.31%  github.com/SlawaBE/go-metrics-collector/internal/storage.(*DBStorage).GetValues.func1
         0     0%  2.01%       3310 22.66%  github.com/go-chi/chi/v5.(*Mux).ServeHTTP
         0     0%  2.01%       3310 22.66%  github.com/go-chi/chi/v5.(*Mux).routeHTTP
         0     0%  2.01%        -45  0.31%  github.com/jackc/pgx/v5.ConnectConfig
         0     0%  2.01%        -45  0.31%  github.com/jackc/pgx/v5.connect
         0     0%  2.01%        -64  0.44%  github.com/jackc/pgx/v5/internal/iobufpool.Get
         0     0%  2.01%        -64  0.44%  github.com/jackc/pgx/v5/pgconn.ConnectConfig
         0     0%  2.01%        -64  0.44%  github.com/jackc/pgx/v5/pgconn.connectOne
         0     0%  2.01%        -64  0.44%  github.com/jackc/pgx/v5/pgconn.connectPreferred
         0     0%  2.01%        -64  0.44%  github.com/jackc/pgx/v5/pgproto3.NewFrontend
         0     0%  2.01%        -64  0.44%  github.com/jackc/pgx/v5/pgproto3.newChunkReader (inline)
         0     0%  2.01%        -45  0.31%  github.com/jackc/pgx/v5/stdlib.(*driverConnector).Connect
         0     0%  2.01%         78  0.53%  io.WriteString
         0     0%  2.01%      -4681 32.05%  net/http.(*conn).readRequest
         0     0%  2.01%      -1287  8.81%  net/http.(*conn).serve
         0     0%  2.01%       3522 24.11%  net/http.HandlerFunc.ServeHTTP
         0     0%  2.01%       -128  0.88%  net/http.newBufioReader
         0     0%  2.01%      -4681 32.05%  net/http.readRequest
         0     0%  2.01%       3522 24.11%  net/http.serverHandler.ServeHTTP
         0     0%  2.01%      -4681 32.05%  net/textproto.(*Reader).ReadMIMEHeader (inline)
         0     0%  2.01%        455  3.12%  runtime.allocm
         0     0%  2.01%      -2979 20.40%  runtime.doInit (inline)
         0     0%  2.01%      -2979 20.40%  runtime.doInit1
         0     0%  2.01%       -513  3.51%  runtime.findRunnable
         0     0%  2.01%      -1024  7.01%  runtime.gcBgMarkWorker
         0     0%  2.01%        455  3.12%  runtime.gcStart.func4
         0     0%  2.01%       -513  3.51%  runtime.injectglist
         0     0%  2.01%       -513  3.51%  runtime.injectglist.func1
         0     0%  2.01%        455  3.12%  runtime.mProfStackInit (inline)
         0     0%  2.01%      -2979 20.40%  runtime.main
         0     0%  2.01%        455  3.12%  runtime.makeProfStackFP (inline)
         0     0%  2.01%        455  3.12%  runtime.makeslice
         0     0%  2.01%      -2185 14.96%  runtime.malg
         0     0%  2.01%       -257  1.76%  runtime.mcall
         0     0%  2.01%        455  3.12%  runtime.mcommoninit
         0     0%  2.01%        257  1.76%  runtime.mstart
         0     0%  2.01%        257  1.76%  runtime.mstart0
         0     0%  2.01%        257  1.76%  runtime.mstart1
         0     0%  2.01%        455  3.12%  runtime.newm
         0     0%  2.01%      -3209 21.97%  runtime.newobject
         0     0%  2.01%      -2185 14.96%  runtime.newproc.func1
         0     0%  2.01%      -2185 14.96%  runtime.newproc1
         0     0%  2.01%       -257  1.76%  runtime.park_m
         0     0%  2.01%        513  3.51%  runtime.resetspinning
         0     0%  2.01%        455  3.12%  runtime.startTheWorldWithSema
         0     0%  2.01%      -1730 11.84%  runtime.systemstack
         0     0%  2.01%        513  3.51%  runtime.wakep
         0     0%  2.01%       3213 22.00%  sync.(*Pool).Get
```

## alloc_space

```text
File: server.exe
Build ID: D:\work\projects\go-metrics-collector\cmd\server\server.exe2026-10-07 10:38:59.6839136 +0300 MSK
Type: alloc_space
Time: 2026-10-05 19:12:54 MSK
Showing nodes accounting for -2857.16MB, 86.99% of 3284.33MB total
Dropped 400 nodes (cum <= 16.42MB)
      flat  flat%   sum%        cum   cum%
-2361.36MB 71.90% 71.90% -2810.28MB 85.57%  compress/flate.NewWriter (inline)
 -512.09MB 15.59% 87.49%  -512.09MB 15.59%  compress/flate.(*compressor).initDeflate (inline)
   52.35MB  1.59% 85.90%  -448.93MB 13.67%  compress/flate.(*compressor).init
  -33.07MB  1.01% 86.90%   -33.07MB  1.01%  compress/flate.(*huffmanEncoder).generate
   28.33MB  0.86% 86.04%    28.33MB  0.86%  compress/flate.newDeflateFast (inline)
  -26.82MB  0.82% 86.86%   -26.82MB  0.82%  net/http.init.func16
   -8.01MB  0.24% 87.10%   -17.51MB  0.53%  compress/flate.newHuffmanBitWriter (inline)
    5.02MB  0.15% 86.95%    68.72MB  2.09%  io.WriteString
   -2.51MB 0.076% 87.02%    52.19MB  1.59%  github.com/SlawaBE/go-metrics-collector/internal/handler.(*ListMetricHandler).ServeHTTP
    0.50MB 0.015% 87.01%   -14.74MB  0.45%  github.com/jackc/pgx/v5.connect
   -0.50MB 0.015% 87.02%   -17.24MB  0.52%  github.com/jackc/pgx/v5/stdlib.(*driverConnector).Connect
    0.50MB 0.015% 87.01%   -16.74MB  0.51%  database/sql.(*DB).conn
    0.50MB 0.015% 86.99%    23.50MB  0.72%  github.com/SlawaBE/go-metrics-collector/internal/handler.(*JSONUpdateMetricHandler).ServeHTTP
         0     0% 86.99%   -33.07MB  1.01%  compress/flate.(*Writer).Close (inline)
         0     0% 86.99%   -33.07MB  1.01%  compress/flate.(*compressor).close
         0     0% 86.99%   -34.58MB  1.05%  compress/flate.(*compressor).deflate
         0     0% 86.99%   -34.58MB  1.05%  compress/flate.(*compressor).writeBlock
         0     0% 86.99%   -22.05MB  0.67%  compress/flate.(*huffmanBitWriter).indexTokens
         0     0% 86.99%   -34.58MB  1.05%  compress/flate.(*huffmanBitWriter).writeBlock
         0     0% 86.99%   -33.07MB  1.01%  compress/gzip.(*Writer).Close
         0     0% 86.99% -2810.28MB 85.57%  compress/gzip.(*Writer).Write
         0     0% 86.99%   -19.78MB   0.6%  database/sql.(*DB).QueryContext
         0     0% 86.99%   -19.78MB   0.6%  database/sql.(*DB).QueryContext.func1
         0     0% 86.99%   -19.78MB   0.6%  database/sql.(*DB).query
         0     0% 86.99%   -16.24MB  0.49%  database/sql.(*DB).retry
         0     0% 86.99%    39.90MB  1.21%  encoding/json.(*Encoder).Encode
         0     0% 86.99%   -34.58MB  1.05%  github.com/SlawaBE/go-metrics-collector/internal/gzip.(*compressWriter).Close
         0     0% 86.99% -2988.42MB 90.99%  github.com/SlawaBE/go-metrics-collector/internal/gzip.(*compressWriter).Write
         0     0% 86.99%    41.58MB  1.27%  github.com/SlawaBE/go-metrics-collector/internal/handler.(*JSONBatchUpdateMetricsHandler).ServeHTTP
         0     0% 86.99%    31.16MB  0.95%  github.com/SlawaBE/go-metrics-collector/internal/handler.(*JSONGetMetricHandler).ServeHTTP
         0     0% 86.99%   178.13MB  5.42%  github.com/SlawaBE/go-metrics-collector/internal/middleware.(*gzipResponseWriter).Write
         0     0% 86.99%   -16.20MB  0.49%  github.com/SlawaBE/go-metrics-collector/internal/retry.RetryWithBackoff
         0     0% 86.99% -2873.06MB 87.48%  github.com/SlawaBE/go-metrics-collector/internal/server.Run.GZip.func4
         0     0% 86.99% -2874.57MB 87.52%  github.com/SlawaBE/go-metrics-collector/internal/server.Run.RequestLogger.func5
         0     0% 86.99%   149.93MB  4.56%  github.com/go-chi/chi/v5.(*Mux).ServeHTTP
         0     0% 86.99%   148.42MB  4.52%  github.com/go-chi/chi/v5.(*Mux).routeHTTP
         0     0% 86.99%   -14.74MB  0.45%  github.com/jackc/pgx/v5.ConnectConfig
         0     0% 86.99%   -29.33MB  0.89%  net/http.(*Request).write
         0     0% 86.99% -2889.15MB 87.97%  net/http.(*conn).serve
         0     0% 86.99%   -29.33MB  0.89%  net/http.(*persistConn).writeLoop
         0     0% 86.99%   -28.32MB  0.86%  net/http.(*transferWriter).doBodyCopy
         0     0% 86.99%   -28.32MB  0.86%  net/http.(*transferWriter).writeBody
         0     0% 86.99% -2874.57MB 87.52%  net/http.HandlerFunc.ServeHTTP
         0     0% 86.99%   -28.32MB  0.86%  net/http.getCopyBuf (inline)
         0     0% 86.99% -2874.57MB 87.52%  net/http.serverHandler.ServeHTTP
         0     0% 86.99%   -46.91MB  1.43%  sync.(*Pool).Get
```

## alloc_objects

```text
File: server.exe
Build ID: D:\work\projects\go-metrics-collector\cmd\server\server.exe2026-10-07 10:38:59.6839136 +0300 MSK
Type: alloc_objects
Time: 2026-10-05 19:12:54 MSK
Showing nodes accounting for -355962, 18.91% of 1882544 total
Dropped 103 nodes (cum <= 9412)
      flat  flat%   sum%        cum   cum%
     98305  5.22%  5.22%      98305  5.22%  github.com/jackc/pgx/v5/pgtype.scanPlanString.Scan
    -98305  5.22%     0%     -98305  5.22%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next.func6
    -77088  4.09%  4.09%    -181384  9.64%  github.com/SlawaBE/go-metrics-collector/internal/service.(*MetricsService).List
    -65537  3.48%  7.58%     -65537  3.48%  github.com/jackc/pgx/v5/pgproto3.(*AuthenticationSASL).Decode
    -65537  3.48% 11.06%      32768  1.74%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next.func14
    -65536  3.48% 14.54%     -65536  3.48%  encoding/json.(*decodeState).literalStore
     43691  2.32% 12.22%     -21845  1.16%  encoding/json.(*decodeState).object
     43690  2.32%  9.90%      43690  2.32%  net.JoinHostPort (inline)
    -37632  2.00% 11.90%     -37632  2.00%  compress/flate.newHuffmanEncoder (inline)
     34257  1.82% 10.08%      34257  1.82%  net/textproto.MIMEHeader.Add (inline)
    -32962  1.75% 11.83%     -26259  1.39%  net/http.(*persistConn).readLoop
     32769  1.74% 10.09%      32769  1.74%  net/textproto.MIMEHeader.Set (inline)
     32769  1.74%  8.35%      32769  1.74%  syscall.(*RawSockaddrAny).Sockaddr
     32768  1.74%  6.61%      37450  1.99%  database/sql.(*Tx).grabConn
    -32768  1.74%  8.35%     -32768  1.74%  encoding/json.(*scanner).pushParseState
    -32768  1.74% 10.09%     -32768  1.74%  internal/strconv.FormatInt
    -32768  1.74% 11.83%     -32768  1.74%  net.copyIP (inline)
     32768  1.74% 10.09%      32768  1.74%  net/http.(*connReader).startBackgroundRead
    -32767  1.74% 11.83%     -32767  1.74%  syscall.UTF16FromString
     30630  1.63% 10.20%     -60606  3.22%  github.com/SlawaBE/go-metrics-collector/internal/storage.(*DBStorage).GetValues.func1
    -26214  1.39% 11.59%     -17536  0.93%  github.com/jackc/pgx/v5/pgconn.(*PgConn).Prepare
    -22942  1.22% 12.81%     -22942  1.22%  io.init.func1
     21846  1.16% 11.65%      32769  1.74%  net/textproto.(*Reader).ReadLine (inline)
    -21846  1.16% 12.81%     -21846  1.16%  net/textproto.NewReader (inline)
    -21845  1.16% 13.97%     -34133  1.81%  crypto/internal/fips140/hmac.New[go.shape.interface { BlockSize int; Reset; Size int; Sum []uint8; Write  }]
    -21845  1.16% 15.13%     -21845  1.16%  encoding/base64.(*Encoding).DecodeString
    -21845  1.16% 16.29%     -21845  1.16%  internal/strconv.FormatFloat (inline)
     19662  1.04% 15.25%       4897  0.26%  context.withCancel (inline)
     18257  0.97% 14.28%       9120  0.48%  net.(*sysDialer).dialParallel
    -17028   0.9% 15.18%     -51681  2.75%  compress/flate.newHuffmanBitWriter (inline)
     16384  0.87% 14.31%      16384  0.87%  crypto/internal/fips140/sha256.(*Digest).Sum
    -16384  0.87% 15.18%     -16384  0.87%  fmt.(*buffer).write (inline)
     16384  0.87% 14.31%       5461  0.29%  github.com/jackc/pgx/v5/pgconn/ctxwatch.(*ContextWatcher).Watch
    -16384  0.87% 15.18%     -16384  0.87%  github.com/jackc/pgx/v5/pgproto3.(*ParameterStatus).Decode
     16384  0.87% 14.31%      16384  0.87%  github.com/jackc/pgx/v5/pgproto3.(*StartupMessage).Encode
    -16384  0.87% 15.18%     -81921  4.35%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
    -15474  0.82% 16.00%     -15474  0.82%  maps.Copy[go.shape.map[string]string,go.shape.map[string]string,go.shape.string,go.shape.string] (inline)
    -15051   0.8% 16.80%     -15051   0.8%  compress/flate.(*huffmanEncoder).generate
    -13892  0.74% 17.54%     -13892  0.74%  sync.(*Pool).pinSlow
    -13653  0.73% 18.27%     -13198   0.7%  github.com/jackc/pgx/v5/pgtype.NewMap
    -12605  0.67% 18.94%     -21967  1.17%  context.(*cancelCtx).propagateCancel
    -12482  0.66% 19.60%      -3119  0.17%  time.NewTimer
    -12288  0.65% 20.25%     -12288  0.65%  crypto/internal/fips140/sha256.New (inline)
    -10923  0.58% 20.83%     -10923  0.58%  context.WithValue
     10923  0.58% 20.25%      10923  0.58%  internal/reflectlite.Swapper
     10923  0.58% 19.67%      10923  0.58%  net.(*Resolver).internetAddrList.func1
     10923  0.58% 19.09%      10923  0.58%  net.sockaddrToTCP
    -10923  0.58% 19.67%     -10923  0.58%  net/url.UserPassword (inline)
     10764  0.57% 19.10%      10764  0.57%  strings.(*Builder).WriteString (inline)
    -10013  0.53% 19.63%     -47822  2.54%  github.com/jackc/pgx/v5/pgconn.parseEnvSettings
     -9363   0.5% 20.13%      -9363   0.5%  context.AfterFunc
      9363   0.5% 19.63%       9363   0.5%  time.newTimer
      9362   0.5% 19.14%       2160  0.11%  context.WithDeadlineCause
     -8738  0.46% 19.60%      -5097  0.27%  net/http.NewRequestWithContext
      8402  0.45% 19.15%       8402  0.45%  database/sql.(*DB).addDepLocked (inline)
      8192  0.44% 18.72%      14746  0.78%  database/sql.(*Rows).initContextClose
      8192  0.44% 18.28%     -73729  3.92%  database/sql.(*Rows).nextLocked
     -8192  0.44% 18.72%     -38380  2.04%  database/sql.(*driverConn).prepareLocked
      8192  0.44% 18.28%       8192  0.44%  encoding/json.Marshal
     -8192  0.44% 18.72%     -16612  0.88%  encoding/json.newEncodeState
      8192  0.44% 18.28%     -43202  2.29%  github.com/SlawaBE/go-metrics-collector/internal/handler.(*JSONUpdateMetricHandler).ServeHTTP
      8192  0.44% 17.85%     -40313  2.14%  github.com/SlawaBE/go-metrics-collector/internal/storage.(*DBStorage).GetMetric
     -8192  0.44% 18.28%      -8192  0.44%  github.com/jackc/pgx/v5/internal/stmtcache.StatementName
     -8192  0.44% 18.72%      -8192  0.44%  github.com/jackc/pgx/v5/pgconn/internal/bgreader.New (inline)
      8192  0.44% 18.28%       8192  0.44%  internal/poll.(*FD).pin
      8192  0.44% 17.85%       8192  0.44%  net.ipToSockaddr
      8192  0.44% 17.41%       8192  0.44%  net/http.readTransfer
     -7726  0.41% 17.82%      -7726  0.41%  compress/flate.(*compressor).initDeflate (inline)
      6827  0.36% 17.46%       6827  0.36%  sync.(*poolChain).pushHead
     -6554  0.35% 17.81%     -24090  1.28%  github.com/jackc/pgx/v5.(*Conn).Prepare
      6242  0.33% 17.48%      -3121  0.17%  net.(*Resolver).lookupIP
     -6170  0.33% 17.80%      -6170  0.33%  net/textproto.readMIMEHeader
     -5846  0.31% 18.12%      -5846  0.31%  bufio.NewReaderSize (inline)
     -5042  0.27% 18.38%     -37842  2.01%  net.(*Resolver).lookupIP.func1
     -5042  0.27% 18.65%     -37809  2.01%  syscall.Getenv
      4917  0.26% 18.39%       4917  0.26%  net/http.send.func1 (inline)
     -4916  0.26% 18.65%      12703  0.67%  net/http.readRequest
      4916  0.26% 18.39%       4916  0.26%  net/http.setupRewindBody (inline)
      4915  0.26% 18.13%       4915  0.26%  encoding/json.NewDecoder (inline)
     -4724  0.25% 18.38%      -4724  0.25%  github.com/jackc/pgx/v5/internal/stmtcache.NewLRUCache (inline)
      4681  0.25% 18.13%       4453  0.24%  fmt.Sprintf
      4681  0.25% 17.88%       4681  0.25%  github.com/golang-migrate/migrate/v4.(*Migrate).lock.func2
      4681  0.25% 17.63%      -2958  0.16%  net/http.(*Transport).dialConn
     -4681  0.25% 17.88%       1873 0.099%  net/http.(*Transport).getConn
     -4681  0.25% 18.13%      15536  0.83%  net/http.(*persistConn).roundTrip
      4370  0.23% 17.90%      23627  1.26%  net/http.(*conn).readRequest
      4096  0.22% 17.68%    -149119  7.92%  database/sql.(*DB).conn
     -4096  0.22% 17.90%      -4096  0.22%  github.com/jackc/pgx/v5/pgproto3.(*SASLResponse).Encode
     -3734   0.2% 18.10%     -62538  3.32%  compress/flate.NewWriter (inline)
      3641  0.19% 17.90%     -34739  1.85%  database/sql.(*DB).prepareDC
     -3641  0.19% 18.10%      -3641  0.19%  github.com/jackc/pgx/v5.(*Conn).getRows
     -3641  0.19% 18.29%      -3641  0.19%  net/http.cloneURL (inline)
      3641  0.19% 18.10%      -7282  0.39%  net/url.parse
      3278  0.17% 17.92%       3278  0.17%  net/http.(*Request).WithContext (inline)
     -3277  0.17% 18.10%      11677  0.62%  database/sql.(*DB).queryDC
     -3277  0.17% 18.27%      -3277  0.17%  net/http.newTransferWriter
     -2979  0.16% 18.43%      -2979  0.16%  internal/bytealg.MakeNoZero
     -2754  0.15% 18.58%      -2754  0.15%  runtime.mallocgc
      2732  0.15% 18.43%     -79872  4.24%  github.com/jackc/pgx/v5/pgconn.(*PgConn).receiveMessage
      2731  0.15% 18.29%      15572  0.83%  net.(*Dialer).DialContext
     -2731  0.15% 18.43%      -2731  0.15%  os.newFile
     -2185  0.12% 18.55%      -2413  0.13%  go.uber.org/zap/internal/stacktrace.Capture
     -2185  0.12% 18.66%      -2185  0.12%  internal/poll.(*FD).Accept
     -2184  0.12% 18.78%      -2184  0.12%  github.com/jackc/pgx/v5/pgconn.configTLS
     -2048  0.11% 18.89%      30720  1.63%  encoding/json.(*Decoder).refill
     -2048  0.11% 19.00%      -2048  0.11%  net/http.ListenAndServe (inline)
      1820 0.097% 18.90%      12743  0.68%  github.com/jackc/pgx/v5/pgconn.defaultSettings
      1820 0.097% 18.80%       1820 0.097%  github.com/jackc/pgx/v5/pgconn.newScramClient
     -1092 0.058% 18.86%    -163192  8.67%  github.com/jackc/pgx/v5/stdlib.(*driverConnector).Connect
     -1076 0.057% 18.92%    -120955  6.43%  github.com/jackc/pgx/v5/pgconn.connectOne
      -642 0.034% 18.95%    -179216  9.52%  github.com/SlawaBE/go-metrics-collector/internal/handler.(*ListMetricHandler).ServeHTTP
       512 0.027% 18.92%     -94675  5.03%  github.com/jackc/pgx/v5.connect
       307 0.016% 18.91%     -58804  3.12%  compress/flate.(*compressor).init
         0     0% 18.91%       8192  0.44%  bufio.(*Reader).Peek
         0     0% 18.91%      10923  0.58%  bufio.(*Reader).ReadLine
         0     0% 18.91%      10923  0.58%  bufio.(*Reader).ReadSlice
         0     0% 18.91%      19115  1.02%  bufio.(*Reader).fill
         0     0% 18.91%     -24538  1.30%  bufio.(*Writer).Flush
         0     0% 18.91%      -2979  0.16%  bytes.Join
         0     0% 18.91%     -15051   0.8%  compress/flate.(*Writer).Close (inline)
         0     0% 18.91%     -15051   0.8%  compress/flate.(*compressor).close
         0     0% 18.91%     -15735  0.84%  compress/flate.(*compressor).deflate
         0     0% 18.91%     -15735  0.84%  compress/flate.(*compressor).writeBlock
         0     0% 18.91%     -10034  0.53%  compress/flate.(*huffmanBitWriter).indexTokens
         0     0% 18.91%     -15735  0.84%  compress/flate.(*huffmanBitWriter).writeBlock
         0     0% 18.91%      -2979  0.16%  compress/flate.generateFixedLiteralEncoding
         0     0% 18.91%      -2979  0.16%  compress/flate.init
         0     0% 18.91%     -15051   0.8%  compress/gzip.(*Writer).Close
         0     0% 18.91%     -62538  3.32%  compress/gzip.(*Writer).Write
         0     0% 18.91%      -3277  0.17%  compress/gzip.NewWriter (inline)
         0     0% 18.91%      14620  0.78%  context.WithCancel
         0     0% 18.91%      -9723  0.52%  context.WithCancelCause
         0     0% 18.91%       2160  0.11%  context.WithDeadline (inline)
         0     0% 18.91%       4681  0.25%  context.WithTimeout
         0     0% 18.91%     -30037  1.60%  crypto/hmac.New
         0     0% 18.91%      -8192  0.44%  crypto/hmac.New.UnwrapNew[go.shape.interface { BlockSize int; Reset; Size int; Sum []uint8; Write  }].func1
         0     0% 18.91%      16384  0.87%  crypto/internal/fips140/hmac.(*HMAC).Sum
         0     0% 18.91%      -4096  0.22%  crypto/internal/fips140/pbkdf2.Key[go.shape.interface { BlockSize int; Reset; Size int; Sum []uint8; Write  }]
         0     0% 18.91%      -4096  0.22%  crypto/pbkdf2.Key[go.shape.interface { BlockSize int; Reset; Size int; Sum []uint8; Write  }]
         0     0% 18.91%      -4096  0.22%  crypto/pbkdf2.Key[go.shape.interface { BlockSize int; Reset; Size int; Sum []uint8; Write  }].UnwrapNew[go.shape.interface { BlockSize int; Reset; Size int; Sum []uint8; Write  }].func1
         0     0% 18.91%     -12288  0.65%  crypto/sha256.New
         0     0% 18.91%     -80515  4.28%  database/sql.(*DB).Begin (inline)
         0     0% 18.91%     -80515  4.28%  database/sql.(*DB).BeginTx
         0     0% 18.91%     -80515  4.28%  database/sql.(*DB).BeginTx.func1
         0     0% 18.91%     -57610  3.06%  database/sql.(*DB).QueryContext
         0     0% 18.91%     -57610  3.06%  database/sql.(*DB).QueryContext.func1
         0     0% 18.91%     -31911  1.70%  database/sql.(*DB).QueryRowContext (inline)
         0     0% 18.91%     -80515  4.28%  database/sql.(*DB).begin
         0     0% 18.91%       4681  0.25%  database/sql.(*DB).connectionOpener
         0     0% 18.91%     -38380  2.04%  database/sql.(*DB).prepareDC.func2
         0     0% 18.91%      -8402  0.45%  database/sql.(*DB).putConn
         0     0% 18.91%     -57610  3.06%  database/sql.(*DB).query
         0     0% 18.91%     -88290  4.69%  database/sql.(*DB).retry
         0     0% 18.91%     -16594  0.88%  database/sql.(*Row).Scan
         0     0% 18.91%      -8402  0.45%  database/sql.(*Rows).Close
         0     0% 18.91%     -73729  3.92%  database/sql.(*Rows).Next
         0     0% 18.91%     -73729  3.92%  database/sql.(*Rows).Next.func1
         0     0% 18.91%      -8402  0.45%  database/sql.(*Rows).close
         0     0% 18.91%      -2521  0.13%  database/sql.(*Stmt).Close
         0     0% 18.91%      49152  2.61%  database/sql.(*Stmt).ExecContext
         0     0% 18.91%      49152  2.61%  database/sql.(*Stmt).ExecContext.func1
         0     0% 18.91%      32768  1.74%  database/sql.(*Stmt).connStmt
         0     0% 18.91%      -2749  0.15%  database/sql.(*Tx).Commit
         0     0% 18.91%     -30057  1.60%  database/sql.(*Tx).PrepareContext
         0     0% 18.91%      -2521  0.13%  database/sql.(*Tx).closePrepared
         0     0% 18.91%      -8402  0.45%  database/sql.(*driverConn).Close
         0     0% 18.91%      -8402  0.45%  database/sql.(*driverConn).finalClose
         0     0% 18.91%      -8402  0.45%  database/sql.(*driverConn).finalClose.func2
         0     0% 18.91%      -8402  0.45%  database/sql.(*driverConn).releaseConn
         0     0% 18.91%      -2521  0.13%  database/sql.(*driverStmt).Close
         0     0% 18.91%     -30188  1.60%  database/sql.ctxDriverPrepare
         0     0% 18.91%      16384  0.87%  database/sql.ctxDriverStmtExec
         0     0% 18.91%      16384  0.87%  database/sql.resultFromStatement
         0     0% 18.91%    -120531  6.40%  database/sql.withLock
         0     0% 18.91%     -23893  1.27%  encoding/json.(*Decoder).Decode
         0     0% 18.91%      -2048  0.11%  encoding/json.(*Decoder).readValue
         0     0% 18.91%     -15969  0.85%  encoding/json.(*Encoder).Encode
         0     0% 18.91%     -21845  1.16%  encoding/json.(*decodeState).unmarshal
         0     0% 18.91%     -21845  1.16%  encoding/json.(*decodeState).value
         0     0% 18.91%     -32768  1.74%  encoding/json.stateBeginValue
         0     0% 18.91%     -16384  0.87%  fmt.(*fmt).fmtBs
         0     0% 18.91%     -16384  0.87%  fmt.(*fmt).pad
         0     0% 18.91%     -16384  0.87%  fmt.(*pp).doPrintf
         0     0% 18.91%     -16384  0.87%  fmt.(*pp).fmtBytes
         0     0% 18.91%     -16384  0.87%  fmt.(*pp).printArg
         0     0% 18.91%     -16840  0.89%  fmt.Appendf
         0     0% 18.91%     -15735  0.84%  github.com/SlawaBE/go-metrics-collector/internal/gzip.(*compressWriter).Close
         0     0% 18.91%     -63746  3.39%  github.com/SlawaBE/go-metrics-collector/internal/gzip.(*compressWriter).Write
         0     0% 18.91%      -2979  0.16%  github.com/SlawaBE/go-metrics-collector/internal/gzip.(*compressWriter).WriteHeader
         0     0% 18.91%      -3277  0.17%  github.com/SlawaBE/go-metrics-collector/internal/gzip.NewCompressWriter
         0     0% 18.91%     -48473  2.57%  github.com/SlawaBE/go-metrics-collector/internal/handler.(*JSONBatchUpdateMetricsHandler).ServeHTTP
         0     0% 18.91%     -75398  4.01%  github.com/SlawaBE/go-metrics-collector/internal/handler.(*JSONGetMetricHandler).ServeHTTP
         0     0% 18.91%       7973  0.42%  github.com/SlawaBE/go-metrics-collector/internal/middleware.(*gzipResponseWriter).WriteHeader
         0     0% 18.91%       3505  0.19%  github.com/SlawaBE/go-metrics-collector/internal/middleware.(*gzipResponseWriter).start
         0     0% 18.91%       3277  0.17%  github.com/SlawaBE/go-metrics-collector/internal/middleware.init.func1
         0     0% 18.91%    -173280  9.20%  github.com/SlawaBE/go-metrics-collector/internal/retry.RetryWithBackoff
         0     0% 18.91%    -438531 23.29%  github.com/SlawaBE/go-metrics-collector/internal/server.Run.GZip.func4
         0     0% 18.91%    -439215 23.33%  github.com/SlawaBE/go-metrics-collector/internal/server.Run.RequestLogger.func5
         0     0% 18.91%      -2048  0.11%  github.com/SlawaBE/go-metrics-collector/internal/server.Run.func1
         0     0% 18.91%     -26370  1.40%  github.com/SlawaBE/go-metrics-collector/internal/service.(*AuditService).SendMetric
         0     0% 18.91%     -57124  3.03%  github.com/SlawaBE/go-metrics-collector/internal/service.(*AuditService).SendMetrics
         0     0% 18.91%     -83494  4.44%  github.com/SlawaBE/go-metrics-collector/internal/service.(*AuditService).notify
         0     0% 18.91%     126178  6.70%  github.com/SlawaBE/go-metrics-collector/internal/service.(*AuditService).notifySubscribers
         0     0% 18.91%     126178  6.70%  github.com/SlawaBE/go-metrics-collector/internal/service.(*AuditService).run
         0     0% 18.91%     -11151  0.59%  github.com/SlawaBE/go-metrics-collector/internal/service.(*FileAuditSubscriber).Notify
         0     0% 18.91%      57613  3.06%  github.com/SlawaBE/go-metrics-collector/internal/service.(*HTTPAuditSubscriber).Notify
         0     0% 18.91%     -40313  2.14%  github.com/SlawaBE/go-metrics-collector/internal/service.(*MetricsService).GetMetric
         0     0% 18.91%     -69540  3.69%  github.com/SlawaBE/go-metrics-collector/internal/service.(*MetricsService).UpdateMetric
         0     0% 18.91%       5371  0.29%  github.com/SlawaBE/go-metrics-collector/internal/service.(*MetricsService).UpdateMetrics
         0     0% 18.91%     -48505  2.58%  github.com/SlawaBE/go-metrics-collector/internal/storage.(*DBStorage).GetMetric.func1
         0     0% 18.91%     -60606  3.22%  github.com/SlawaBE/go-metrics-collector/internal/storage.(*DBStorage).GetValues
         0     0% 18.91%       5371  0.29%  github.com/SlawaBE/go-metrics-collector/internal/storage.(*DBStorage).UpdateAll
         0     0% 18.91%       5371  0.29%  github.com/SlawaBE/go-metrics-collector/internal/storage.(*DBStorage).UpdateAll.func1
         0     0% 18.91%     -69540  3.69%  github.com/SlawaBE/go-metrics-collector/internal/storage.(*DBStorage).UpdateMetric
         0     0% 18.91%     -69540  3.69%  github.com/SlawaBE/go-metrics-collector/internal/storage.(*DBStorage).UpdateMetric.func1
         0     0% 18.91%     -32768  1.74%  github.com/SlawaBE/go-metrics-collector/internal/utils.ConvertCounter (inline)
         0     0% 18.91%     -21845  1.16%  github.com/SlawaBE/go-metrics-collector/internal/utils.ConvertGauge
         0     0% 18.91%    -353478 18.78%  github.com/go-chi/chi/v5.(*Mux).ServeHTTP
         0     0% 18.91%    -346289 18.39%  github.com/go-chi/chi/v5.(*Mux).routeHTTP
         0     0% 18.91%      -8402  0.45%  github.com/jackc/pgx/v5.(*Conn).Close
         0     0% 18.91%      -7202  0.38%  github.com/jackc/pgx/v5.(*Conn).Deallocate
         0     0% 18.91%      16156  0.86%  github.com/jackc/pgx/v5.(*Conn).Exec
         0     0% 18.91%      16156  0.86%  github.com/jackc/pgx/v5.(*Conn).exec
         0     0% 18.91%      16384  0.87%  github.com/jackc/pgx/v5.(*Conn).execPrepared
         0     0% 18.91%      -2094  0.11%  github.com/jackc/pgx/v5.(*Conn).getStatementDescription
         0     0% 18.91%     -94799  5.04%  github.com/jackc/pgx/v5.ConnectConfig
         0     0% 18.91%     -67301  3.58%  github.com/jackc/pgx/v5.ParseConfig (inline)
         0     0% 18.91%     -67301  3.58%  github.com/jackc/pgx/v5.ParseConfigWithOptions
         0     0% 18.91%       4096  0.22%  github.com/jackc/pgx/v5/pgconn.(*MultiResultReader).Close (inline)
         0     0% 18.91%       4096  0.22%  github.com/jackc/pgx/v5/pgconn.(*MultiResultReader).receiveMessage
         0     0% 18.91%      -8402  0.45%  github.com/jackc/pgx/v5/pgconn.(*PgConn).Close
         0     0% 18.91%      -7202  0.38%  github.com/jackc/pgx/v5/pgconn.(*PgConn).Deallocate
         0     0% 18.91%      -2521  0.13%  github.com/jackc/pgx/v5/pgconn.(*PgConn).Exec
         0     0% 18.91%      23587  1.25%  github.com/jackc/pgx/v5/pgconn.(*PgConn).ExecStatement
         0     0% 18.91%      23586  1.25%  github.com/jackc/pgx/v5/pgconn.(*PgConn).execExtendedPrefix
         0     0% 18.91%     -82832  4.40%  github.com/jackc/pgx/v5/pgconn.(*PgConn).peekMessage
         0     0% 18.91%      -4096  0.22%  github.com/jackc/pgx/v5/pgconn.(*PgConn).rxSASLFinal
         0     0% 18.91%     -61332  3.26%  github.com/jackc/pgx/v5/pgconn.(*PgConn).scramAuth
         0     0% 18.91%     -20599  1.09%  github.com/jackc/pgx/v5/pgconn.(*scramClient).clientFinalMessage
         0     0% 18.91%     -16612  0.88%  github.com/jackc/pgx/v5/pgconn.(*scramClient).clientFirstMessage
         0     0% 18.91%       4096  0.22%  github.com/jackc/pgx/v5/pgconn.(*scramClient).recvServerFinalMessage
         0     0% 18.91%     -21845  1.16%  github.com/jackc/pgx/v5/pgconn.(*scramClient).recvServerFirstMessage
         0     0% 18.91%     -77265  4.10%  github.com/jackc/pgx/v5/pgconn.ConnectConfig
         0     0% 18.91%      43690  2.32%  github.com/jackc/pgx/v5/pgconn.NetworkAddress
         0     0% 18.91%     -67301  3.58%  github.com/jackc/pgx/v5/pgconn.ParseConfigWithOptions
         0     0% 18.91%      43690  2.32%  github.com/jackc/pgx/v5/pgconn.buildConnectOneConfigs
         0     0% 18.91%     -17749  0.94%  github.com/jackc/pgx/v5/pgconn.computeClientProof
         0     0% 18.91%     -13653  0.73%  github.com/jackc/pgx/v5/pgconn.computeHMAC
         0     0% 18.91%       4096  0.22%  github.com/jackc/pgx/v5/pgconn.computeServerSignature
         0     0% 18.91%    -120955  6.43%  github.com/jackc/pgx/v5/pgconn.connectPreferred
         0     0% 18.91%     -15474  0.82%  github.com/jackc/pgx/v5/pgconn.mergeSettings (inline)
         0     0% 18.91%     -14564  0.77%  github.com/jackc/pgx/v5/pgconn.parseURLSettings
         0     0% 18.91%     -82832  4.40%  github.com/jackc/pgx/v5/pgproto3.(*Frontend).Receive
         0     0% 18.91%      12288  0.65%  github.com/jackc/pgx/v5/pgproto3.(*Frontend).Send
         0     0% 18.91%      -8402  0.45%  github.com/jackc/pgx/v5/stdlib.(*Conn).Close
         0     0% 18.91%      16384  0.87%  github.com/jackc/pgx/v5/stdlib.(*Conn).ExecContext
         0     0% 18.91%     -30188  1.60%  github.com/jackc/pgx/v5/stdlib.(*Conn).PrepareContext
         0     0% 18.91%      -2521  0.13%  github.com/jackc/pgx/v5/stdlib.(*Stmt).Close
         0     0% 18.91%      16384  0.87%  github.com/jackc/pgx/v5/stdlib.(*Stmt).ExecContext
         0     0% 18.91%      -4462  0.24%  go.uber.org/zap.(*Logger).Info
         0     0% 18.91%      -2641  0.14%  go.uber.org/zap.(*Logger).check
         0     0% 18.91%      -1821 0.097%  go.uber.org/zap/zapcore.(*CheckedEntry).Write
         0     0% 18.91%      -1821 0.097%  go.uber.org/zap/zapcore.(*ioCore).Write
         0     0% 18.91%      -4096  0.22%  internal/poll.(*FD).ConnectEx
         0     0% 18.91%      18203  0.97%  internal/poll.(*FD).Read
         0     0% 18.91%       5915  0.31%  internal/poll.(*FD).execIO
         0     0% 18.91%      -3121  0.17%  internal/singleflight.(*Group).doCall
         0     0% 18.91%     -24082  1.28%  io.Copy (inline)
         0     0% 18.91%     -24082  1.28%  io.CopyN
         0     0% 18.91%     -24082  1.28%  io.copyBuffer
         0     0% 18.91%     -24082  1.28%  io.discard.ReadFrom
         0     0% 18.91%      -7202  0.38%  net.(*Dialer).dialCtx
         0     0% 18.91%      10923  0.58%  net.(*Resolver).internetAddrList
         0     0% 18.91%     -37842  2.01%  net.(*Resolver).lookupIP.func2
         0     0% 18.91%      -3121  0.17%  net.(*Resolver).lookupIPAddr.func1
         0     0% 18.91%      10923  0.58%  net.(*Resolver).resolveAddrList
         0     0% 18.91%       8192  0.44%  net.(*TCPAddr).sockaddr
         0     0% 18.91%      -2185  0.12%  net.(*TCPListener).Accept
         0     0% 18.91%      -2185  0.12%  net.(*TCPListener).accept
         0     0% 18.91%      18203  0.97%  net.(*conn).Read
         0     0% 18.91%      18203  0.97%  net.(*netFD).Read
         0     0% 18.91%      -2185  0.12%  net.(*netFD).accept
         0     0% 18.91%      -2536  0.13%  net.(*netFD).connect
         0     0% 18.91%      49348  2.62%  net.(*netFD).dial
         0     0% 18.91%      31598  1.68%  net.(*sysDialer).dialParallel.func1
         0     0% 18.91%      49348  2.62%  net.(*sysDialer).dialSerial
         0     0% 18.91%      49348  2.62%  net.(*sysDialer).dialSingle
         0     0% 18.91%      49348  2.62%  net.(*sysDialer).dialTCP
         0     0% 18.91%      49348  2.62%  net.(*sysDialer).doDialTCP (inline)
         0     0% 18.91%      49348  2.62%  net.(*sysDialer).doDialTCPProto
         0     0% 18.91%      10923  0.58%  net.filterAddrList
         0     0% 18.91%      -3121  0.17%  net.init.func1
         0     0% 18.91%      49348  2.62%  net.internetSocket
         0     0% 18.91%      49348  2.62%  net.socket
         0     0% 18.91%      23240  1.23%  net/http.(*Client).Do (inline)
         0     0% 18.91%      23240  1.23%  net/http.(*Client).do
         0     0% 18.91%      23240  1.23%  net/http.(*Client).send
         0     0% 18.91%      -5047  0.27%  net/http.(*Request).write
         0     0% 18.91%      -2185  0.12%  net/http.(*Server).ListenAndServe
         0     0% 18.91%      -2185  0.12%  net/http.(*Server).Serve
         0     0% 18.91%      17283  0.92%  net/http.(*Transport).RoundTrip
         0     0% 18.91%      -2178  0.12%  net/http.(*Transport).dial
         0     0% 18.91%      -2958  0.16%  net/http.(*Transport).dialConnFor
         0     0% 18.91%      17283  0.92%  net/http.(*Transport).roundTrip
         0     0% 18.91%      -2958  0.16%  net/http.(*Transport).startDialConnForLocked.func1
         0     0% 18.91%      32768  1.74%  net/http.(*body).Read
         0     0% 18.91%      32768  1.74%  net/http.(*body).readLocked
         0     0% 18.91%     -24538  1.30%  net/http.(*chunkWriter).Write
         0     0% 18.91%     -24538  1.30%  net/http.(*chunkWriter).writeHeader
         0     0% 18.91%    -440511 23.40%  net/http.(*conn).serve
         0     0% 18.91%      10923  0.58%  net/http.(*connReader).Read
         0     0% 18.91%       8192  0.44%  net/http.(*persistConn).Read
         0     0% 18.91%      -5047  0.27%  net/http.(*persistConn).writeLoop
         0     0% 18.91%     -24538  1.30%  net/http.(*response).finishRequest
         0     0% 18.91%    -439215 23.33%  net/http.HandlerFunc.ServeHTTP
         0     0% 18.91%      34257  1.82%  net/http.Header.Add (inline)
         0     0% 18.91%      32769  1.74%  net/http.Header.Set (inline)
         0     0% 18.91%      14564  0.77%  net/http.NewRequest (inline)
         0     0% 18.91%     -22302  1.18%  net/http.newTextprotoReader
         0     0% 18.91%      26881  1.43%  net/http.send
         0     0% 18.91%    -439215 23.33%  net/http.serverHandler.ServeHTTP
         0     0% 18.91%       4681  0.25%  net/http.setRequestCancel
         0     0% 18.91%      -6170  0.33%  net/textproto.(*Reader).ReadMIMEHeader (inline)
         0     0% 18.91%      10923  0.58%  net/textproto.(*Reader).readLineSlice
         0     0% 18.91%     -10923  0.58%  net/url.Parse
         0     0% 18.91%       3641  0.19%  net/url.ParseRequestURI
         0     0% 18.91%     -10923  0.58%  net/url.parseAuthority
         0     0% 18.91%     -37809  2.01%  os.Getenv
         0     0% 18.91%      -2731  0.15%  os.OpenFile
         0     0% 18.91%      -2731  0.15%  os.openFileNolog
         0     0% 18.91%      10923  0.58%  path/filepath.Join (inline)
         0     0% 18.91%      10923  0.58%  path/filepath.join
         0     0% 18.91%      -2979  0.16%  runtime.doInit (inline)
         0     0% 18.91%      -2979  0.16%  runtime.doInit1
         0     0% 18.91%      -4481  0.24%  runtime.main
         0     0% 18.91%      -2185  0.12%  runtime.malg
         0     0% 18.91%      -3209  0.17%  runtime.newobject
         0     0% 18.91%      -2185  0.12%  runtime.newproc.func1
         0     0% 18.91%      -2185  0.12%  runtime.newproc1
         0     0% 18.91%      10923  0.58%  sort.Slice
         0     0% 18.91%     -21845  1.16%  strconv.FormatFloat (inline)
         0     0% 18.91%     -32768  1.74%  strconv.FormatInt (inline)
         0     0% 18.91%     -33894  1.80%  sync.(*Pool).Get
         0     0% 18.91%       6371  0.34%  sync.(*Pool).Put
         0     0% 18.91%     -13892  0.74%  sync.(*Pool).pin
         0     0% 18.91%      32769  1.74%  syscall.Getsockname
         0     0% 18.91%     -32767  1.74%  syscall.UTF16PtrFromString (inline)
```

