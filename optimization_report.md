# Optimization report

## inuse_space

```text
File: server.exe
Build ID: D:\work\projects\go-metrics-collector\cmd\server\server.exe2026-10-05 19:16:12.4388163 +0300 MSK
Type: inuse_space
Time: 2026-10-05 19:12:54 MSK
Showing nodes accounting for 17867.67kB, 189.44% of 9431.91kB total
Dropped 2 nodes (cum <= 47.16kB)
      flat  flat%   sum%        cum   cum%
 9928.45kB 105.26% 105.26% 17858.82kB 189.34%  compress/flate.NewWriter (inline)
 6838.98kB 72.51% 177.77%  7930.37kB 84.08%  compress/flate.(*compressor).init
-2049.72kB 21.73% 156.04% -2049.72kB 21.73%  runtime.mallocgc
 1672.34kB 17.73% 173.77%  1672.34kB 17.73%  compress/flate.newDeflateFast (inline)
-1093.51kB 11.59% 162.18% -1093.51kB 11.59%  compress/flate.(*compressor).initDeflate (inline)
 1024.11kB 10.86% 173.04%  1024.11kB 10.86%  net/textproto.(*Reader).ReadLine (inline)
  525.43kB  5.57% 178.61%   525.43kB  5.57%  github.com/jackc/pgx/v5/internal/stmtcache.NewLRUCache (inline)
 -516.01kB  5.47% 173.14%  -516.01kB  5.47%  github.com/jackc/pgx/v5/internal/iobufpool.init.0.func1
  512.69kB  5.44% 178.57%   508.73kB  5.39%  github.com/jackc/pgx/v5/pgconn.connectOne
  512.56kB  5.43% 184.01%   512.56kB  5.43%  compress/flate.newHuffmanBitWriter (inline)
  512.44kB  5.43% 189.44%   512.44kB  5.43%  github.com/jackc/pgx/v5/pgproto3.(*DataRow).Decode
 -512.12kB  5.43% 184.01%  -512.12kB  5.43%  net/http.ListenAndServe (inline)
 -512.09kB  5.43% 178.58%  -512.09kB  5.43%  compress/flate.newHuffmanEncoder (inline)
  512.07kB  5.43% 184.01%   512.07kB  5.43%  net/url.parse
  512.05kB  5.43% 189.44%   512.05kB  5.43%  context.(*cancelCtx).Done
  512.05kB  5.43% 194.87%   512.05kB  5.43%  github.com/jackc/pgx/v5/pgproto3.(*AuthenticationSASL).Decode
 -512.05kB  5.43% 189.44%  -512.05kB  5.43%  net/textproto.readMIMEHeader
         0     0% 189.44%  -512.09kB  5.43%  compress/flate.generateFixedLiteralEncoding
         0     0% 189.44%  -512.09kB  5.43%  compress/flate.init
         0     0% 189.44% 17858.82kB 189.34%  compress/gzip.(*Writer).Write
         0     0% 189.44%   512.05kB  5.43%  database/sql.(*DB).Begin (inline)
         0     0% 189.44%   512.05kB  5.43%  database/sql.(*DB).BeginTx
         0     0% 189.44%   512.05kB  5.43%  database/sql.(*DB).BeginTx.func1
         0     0% 189.44%  1034.54kB 10.97%  database/sql.(*DB).QueryContext
         0     0% 189.44%  1034.54kB 10.97%  database/sql.(*DB).QueryContext.func1
         0     0% 189.44%   512.05kB  5.43%  database/sql.(*DB).begin
         0     0% 189.44%  1034.16kB 10.96%  database/sql.(*DB).conn
         0     0% 189.44%   512.05kB  5.43%  database/sql.(*DB).connectionOpener
         0     0% 189.44%  1034.54kB 10.97%  database/sql.(*DB).query
         0     0% 189.44%   512.44kB  5.43%  database/sql.(*DB).queryDC
         0     0% 189.44%   512.44kB  5.43%  database/sql.(*DB).queryDC.func1
         0     0% 189.44%  1546.60kB 16.40%  database/sql.(*DB).retry
         0     0% 189.44%   512.44kB  5.43%  database/sql.ctxDriverQuery
         0     0% 189.44%   512.44kB  5.43%  database/sql.withLock
         0     0% 189.44% -1996.09kB 21.16%  github.com/SlawaBE/go-metrics-collector/internal/gzip.(*compressWriter).Write
         0     0% 189.44%   512.05kB  5.43%  github.com/SlawaBE/go-metrics-collector/internal/handler.(*JSONUpdateMetricHandler).ServeHTTP
         0     0% 189.44% 20889.46kB 221.48%  github.com/SlawaBE/go-metrics-collector/internal/handler.(*ListMetricHandler).ServeHTTP
         0     0% 189.44% 19854.91kB 210.51%  github.com/SlawaBE/go-metrics-collector/internal/middleware.(*gzipResponseWriter).Write
         0     0% 189.44%  1546.60kB 16.40%  github.com/SlawaBE/go-metrics-collector/internal/retry.RetryWithBackoff
         0     0% 189.44% 19405.42kB 205.74%  github.com/SlawaBE/go-metrics-collector/internal/server.Run.GZip.func4
         0     0% 189.44% 19405.42kB 205.74%  github.com/SlawaBE/go-metrics-collector/internal/server.Run.RequestLogger.func5
         0     0% 189.44%  -512.12kB  5.43%  github.com/SlawaBE/go-metrics-collector/internal/server.Run.func1
         0     0% 189.44%  1034.54kB 10.97%  github.com/SlawaBE/go-metrics-collector/internal/service.(*MetricsService).List
         0     0% 189.44%   512.05kB  5.43%  github.com/SlawaBE/go-metrics-collector/internal/service.(*MetricsService).UpdateMetric
         0     0% 189.44%  1034.54kB 10.97%  github.com/SlawaBE/go-metrics-collector/internal/storage.(*DBStorage).GetValues
         0     0% 189.44%  1034.54kB 10.97%  github.com/SlawaBE/go-metrics-collector/internal/storage.(*DBStorage).GetValues.func1
         0     0% 189.44%   512.05kB  5.43%  github.com/SlawaBE/go-metrics-collector/internal/storage.(*DBStorage).UpdateMetric
         0     0% 189.44%   512.05kB  5.43%  github.com/SlawaBE/go-metrics-collector/internal/storage.(*DBStorage).UpdateMetric.func1
         0     0% 189.44% 21401.51kB 226.91%  github.com/go-chi/chi/v5.(*Mux).ServeHTTP
         0     0% 189.44% 21401.51kB 226.91%  github.com/go-chi/chi/v5.(*Mux).routeHTTP
         0     0% 189.44%   512.44kB  5.43%  github.com/jackc/pgx/v5.(*Conn).Query
         0     0% 189.44%  1034.16kB 10.96%  github.com/jackc/pgx/v5.ConnectConfig
         0     0% 189.44%  1034.16kB 10.96%  github.com/jackc/pgx/v5.connect
         0     0% 189.44%  -516.01kB  5.47%  github.com/jackc/pgx/v5/internal/iobufpool.Get
         0     0% 189.44%   512.44kB  5.43%  github.com/jackc/pgx/v5/pgconn.(*PgConn).ExecStatement
         0     0% 189.44%   512.44kB  5.43%  github.com/jackc/pgx/v5/pgconn.(*PgConn).execExtendedSuffix
         0     0% 189.44%  1024.49kB 10.86%  github.com/jackc/pgx/v5/pgconn.(*PgConn).peekMessage
         0     0% 189.44%  1024.49kB 10.86%  github.com/jackc/pgx/v5/pgconn.(*PgConn).receiveMessage
         0     0% 189.44%   512.44kB  5.43%  github.com/jackc/pgx/v5/pgconn.(*ResultReader).readUntilRowDescription
         0     0% 189.44%   512.44kB  5.43%  github.com/jackc/pgx/v5/pgconn.(*ResultReader).receiveMessage
         0     0% 189.44%   508.73kB  5.39%  github.com/jackc/pgx/v5/pgconn.ConnectConfig
         0     0% 189.44%   508.73kB  5.39%  github.com/jackc/pgx/v5/pgconn.connectPreferred
         0     0% 189.44%  1024.49kB 10.86%  github.com/jackc/pgx/v5/pgproto3.(*Frontend).Receive
         0     0% 189.44%  -516.01kB  5.47%  github.com/jackc/pgx/v5/pgproto3.NewFrontend
         0     0% 189.44%  -516.01kB  5.47%  github.com/jackc/pgx/v5/pgproto3.newChunkReader (inline)
         0     0% 189.44%   512.44kB  5.43%  github.com/jackc/pgx/v5/stdlib.(*Conn).QueryContext
         0     0% 189.44%  1034.16kB 10.96%  github.com/jackc/pgx/v5/stdlib.(*driverConnector).Connect
         0     0% 189.44% 19854.91kB 210.51%  io.WriteString
         0     0% 189.44%  1024.12kB 10.86%  net/http.(*conn).readRequest
         0     0% 189.44% 20429.54kB 216.60%  net/http.(*conn).serve
         0     0% 189.44% 19405.42kB 205.74%  net/http.HandlerFunc.ServeHTTP
         0     0% 189.44%  1024.12kB 10.86%  net/http.readRequest
         0     0% 189.44% 19405.42kB 205.74%  net/http.serverHandler.ServeHTTP
         0     0% 189.44%  -512.05kB  5.43%  net/textproto.(*Reader).ReadMIMEHeader (inline)
         0     0% 189.44%   512.07kB  5.43%  net/url.ParseRequestURI
         0     0% 189.44%     -513kB  5.44%  runtime.allocm
         0     0% 189.44%  -512.09kB  5.43%  runtime.doInit (inline)
         0     0% 189.44%  -512.09kB  5.43%  runtime.doInit1
         0     0% 189.44%     -513kB  5.44%  runtime.exitsyscallNoP
         0     0% 189.44%     -513kB  5.44%  runtime.findRunnable
         0     0% 189.44%  -512.25kB  5.43%  runtime.gcBgMarkWorker
         0     0% 189.44%     -513kB  5.44%  runtime.injectglist
         0     0% 189.44%     -513kB  5.44%  runtime.injectglist.func1
         0     0% 189.44%  -512.09kB  5.43%  runtime.main
         0     0% 189.44% -1024.47kB 10.86%  runtime.malg
         0     0% 189.44%    -1026kB 10.88%  runtime.mcall
         0     0% 189.44%      513kB  5.44%  runtime.mstart
         0     0% 189.44%      513kB  5.44%  runtime.mstart0
         0     0% 189.44%      513kB  5.44%  runtime.mstart1
         0     0% 189.44%     -513kB  5.44%  runtime.newm
         0     0% 189.44% -2049.72kB 21.73%  runtime.newobject
         0     0% 189.44% -1024.47kB 10.86%  runtime.newproc.func1
         0     0% 189.44% -1024.47kB 10.86%  runtime.newproc1
         0     0% 189.44%     -513kB  5.44%  runtime.park_m
         0     0% 189.44%     -513kB  5.44%  runtime.schedule
         0     0% 189.44%     -513kB  5.44%  runtime.startm
         0     0% 189.44% -1024.47kB 10.86%  runtime.systemstack
         0     0% 189.44%  -516.01kB  5.47%  sync.(*Pool).Get
```

## inuse_objects

```text
File: server.exe
Build ID: D:\work\projects\go-metrics-collector\cmd\server\server.exe2026-10-05 19:16:12.4388163 +0300 MSK
Type: inuse_objects
Time: 2026-10-05 19:12:54 MSK
Showing nodes accounting for 10601, 72.58% of 14606 total
Dropped 6 nodes (cum <= 73)
      flat  flat%   sum%        cum   cum%
      9363 64.10% 64.10%       9363 64.10%  net/textproto.(*Reader).ReadLine (inline)
      4681 32.05% 96.15%       4681 32.05%  context.(*cancelCtx).Done
      4681 32.05% 128.20%       4681 32.05%  github.com/jackc/pgx/v5/pgproto3.(*AuthenticationSASL).Decode
     -4681 32.05% 96.15%      -4681 32.05%  net/textproto.readMIMEHeader
      3641 24.93% 121.08%       3641 24.93%  net/url.parse
     -3465 23.72% 97.36%      -3465 23.72%  runtime.mallocgc
     -2979 20.40% 76.96%      -2979 20.40%  compress/flate.newHuffmanEncoder (inline)
     -2048 14.02% 62.94%      -2048 14.02%  net/http.ListenAndServe (inline)
       585  4.01% 66.95%        585  4.01%  github.com/jackc/pgx/v5/pgproto3.(*DataRow).Decode
       455  3.12% 70.06%        455  3.12%  compress/flate.newHuffmanBitWriter (inline)
       372  2.55% 72.61%       4989 34.16%  github.com/jackc/pgx/v5/pgconn.connectOne
       -64  0.44% 72.17%        -64  0.44%  github.com/jackc/pgx/v5/internal/iobufpool.init.0.func1
        45  0.31% 72.48%        506  3.46%  compress/flate.(*compressor).init
        15   0.1% 72.58%        521  3.57%  compress/flate.NewWriter (inline)
         0     0% 72.58%      -2979 20.40%  compress/flate.generateFixedLiteralEncoding
         0     0% 72.58%      -2979 20.40%  compress/flate.init
         0     0% 72.58%        521  3.57%  compress/gzip.(*Writer).Write
         0     0% 72.58%       4681 32.05%  database/sql.(*DB).Begin (inline)
         0     0% 72.58%       4681 32.05%  database/sql.(*DB).BeginTx
         0     0% 72.58%       4681 32.05%  database/sql.(*DB).BeginTx.func1
         0     0% 72.58%        912  6.24%  database/sql.(*DB).QueryContext
         0     0% 72.58%        912  6.24%  database/sql.(*DB).QueryContext.func1
         0     0% 72.58%       4681 32.05%  database/sql.(*DB).begin
         0     0% 72.58%       5008 34.29%  database/sql.(*DB).conn
         0     0% 72.58%       4681 32.05%  database/sql.(*DB).connectionOpener
         0     0% 72.58%        912  6.24%  database/sql.(*DB).query
         0     0% 72.58%        585  4.01%  database/sql.(*DB).queryDC
         0     0% 72.58%        585  4.01%  database/sql.(*DB).queryDC.func1
         0     0% 72.58%       5593 38.29%  database/sql.(*DB).retry
         0     0% 72.58%        585  4.01%  database/sql.ctxDriverQuery
         0     0% 72.58%        585  4.01%  database/sql.withLock
         0     0% 72.58%       4681 32.05%  github.com/SlawaBE/go-metrics-collector/internal/handler.(*JSONUpdateMetricHandler).ServeHTTP
         0     0% 72.58%       1449  9.92%  github.com/SlawaBE/go-metrics-collector/internal/handler.(*ListMetricHandler).ServeHTTP
         0     0% 72.58%        537  3.68%  github.com/SlawaBE/go-metrics-collector/internal/middleware.(*gzipResponseWriter).Write
         0     0% 72.58%       5593 38.29%  github.com/SlawaBE/go-metrics-collector/internal/retry.RetryWithBackoff
         0     0% 72.58%       6114 41.86%  github.com/SlawaBE/go-metrics-collector/internal/server.Run.GZip.func4
         0     0% 72.58%       6114 41.86%  github.com/SlawaBE/go-metrics-collector/internal/server.Run.RequestLogger.func5
         0     0% 72.58%      -2048 14.02%  github.com/SlawaBE/go-metrics-collector/internal/server.Run.func1
         0     0% 72.58%        912  6.24%  github.com/SlawaBE/go-metrics-collector/internal/service.(*MetricsService).List
         0     0% 72.58%       4681 32.05%  github.com/SlawaBE/go-metrics-collector/internal/service.(*MetricsService).UpdateMetric
         0     0% 72.58%        912  6.24%  github.com/SlawaBE/go-metrics-collector/internal/storage.(*DBStorage).GetValues
         0     0% 72.58%        912  6.24%  github.com/SlawaBE/go-metrics-collector/internal/storage.(*DBStorage).GetValues.func1
         0     0% 72.58%       4681 32.05%  github.com/SlawaBE/go-metrics-collector/internal/storage.(*DBStorage).UpdateMetric
         0     0% 72.58%       4681 32.05%  github.com/SlawaBE/go-metrics-collector/internal/storage.(*DBStorage).UpdateMetric.func1
         0     0% 72.58%       6130 41.97%  github.com/go-chi/chi/v5.(*Mux).ServeHTTP
         0     0% 72.58%       6130 41.97%  github.com/go-chi/chi/v5.(*Mux).routeHTTP
         0     0% 72.58%        585  4.01%  github.com/jackc/pgx/v5.(*Conn).Query
         0     0% 72.58%       5008 34.29%  github.com/jackc/pgx/v5.ConnectConfig
         0     0% 72.58%       5008 34.29%  github.com/jackc/pgx/v5.connect
         0     0% 72.58%        -64  0.44%  github.com/jackc/pgx/v5/internal/iobufpool.Get
         0     0% 72.58%        585  4.01%  github.com/jackc/pgx/v5/pgconn.(*PgConn).ExecStatement
         0     0% 72.58%        585  4.01%  github.com/jackc/pgx/v5/pgconn.(*PgConn).execExtendedSuffix
         0     0% 72.58%       5266 36.05%  github.com/jackc/pgx/v5/pgconn.(*PgConn).peekMessage
         0     0% 72.58%       5266 36.05%  github.com/jackc/pgx/v5/pgconn.(*PgConn).receiveMessage
         0     0% 72.58%        585  4.01%  github.com/jackc/pgx/v5/pgconn.(*ResultReader).readUntilRowDescription
         0     0% 72.58%        585  4.01%  github.com/jackc/pgx/v5/pgconn.(*ResultReader).receiveMessage
         0     0% 72.58%       4989 34.16%  github.com/jackc/pgx/v5/pgconn.ConnectConfig
         0     0% 72.58%       4989 34.16%  github.com/jackc/pgx/v5/pgconn.connectPreferred
         0     0% 72.58%       5266 36.05%  github.com/jackc/pgx/v5/pgproto3.(*Frontend).Receive
         0     0% 72.58%        -64  0.44%  github.com/jackc/pgx/v5/pgproto3.NewFrontend
         0     0% 72.58%        -64  0.44%  github.com/jackc/pgx/v5/pgproto3.newChunkReader (inline)
         0     0% 72.58%        585  4.01%  github.com/jackc/pgx/v5/stdlib.(*Conn).QueryContext
         0     0% 72.58%       5008 34.29%  github.com/jackc/pgx/v5/stdlib.(*driverConnector).Connect
         0     0% 72.58%        537  3.68%  io.WriteString
         0     0% 72.58%       8323 56.98%  net/http.(*conn).readRequest
         0     0% 72.58%      14437 98.84%  net/http.(*conn).serve
         0     0% 72.58%       6114 41.86%  net/http.HandlerFunc.ServeHTTP
         0     0% 72.58%       8323 56.98%  net/http.readRequest
         0     0% 72.58%       6114 41.86%  net/http.serverHandler.ServeHTTP
         0     0% 72.58%      -4681 32.05%  net/textproto.(*Reader).ReadMIMEHeader (inline)
         0     0% 72.58%       3641 24.93%  net/url.ParseRequestURI
         0     0% 72.58%       -256  1.75%  runtime.allocm
         0     0% 72.58%      -2979 20.40%  runtime.doInit (inline)
         0     0% 72.58%      -2979 20.40%  runtime.doInit1
         0     0% 72.58%       -256  1.75%  runtime.exitsyscallNoP
         0     0% 72.58%       -257  1.76%  runtime.findRunnable
         0     0% 72.58%      -1024  7.01%  runtime.gcBgMarkWorker
         0     0% 72.58%       -257  1.76%  runtime.injectglist
         0     0% 72.58%       -257  1.76%  runtime.injectglist.func1
         0     0% 72.58%      -2979 20.40%  runtime.main
         0     0% 72.58%      -2185 14.96%  runtime.malg
         0     0% 72.58%       -513  3.51%  runtime.mcall
         0     0% 72.58%        257  1.76%  runtime.mstart
         0     0% 72.58%        257  1.76%  runtime.mstart0
         0     0% 72.58%        257  1.76%  runtime.mstart1
         0     0% 72.58%       -256  1.75%  runtime.newm
         0     0% 72.58%      -3465 23.72%  runtime.newobject
         0     0% 72.58%      -2185 14.96%  runtime.newproc.func1
         0     0% 72.58%      -2185 14.96%  runtime.newproc1
         0     0% 72.58%       -257  1.76%  runtime.park_m
         0     0% 72.58%       -256  1.75%  runtime.schedule
         0     0% 72.58%       -256  1.75%  runtime.startm
         0     0% 72.58%      -2185 14.96%  runtime.systemstack
         0     0% 72.58%        -64  0.44%  sync.(*Pool).Get
```

## alloc_space

```text
File: server.exe
Build ID: D:\work\projects\go-metrics-collector\cmd\server\server.exe2026-10-05 19:16:12.4388163 +0300 MSK
Type: alloc_space
Time: 2026-10-05 19:12:54 MSK
Showing nodes accounting for -2675.13MB, 81.45% of 3284.33MB total
Dropped 388 nodes (cum <= 16.42MB)
      flat  flat%   sum%        cum   cum%
-2253.82MB 68.62% 68.62% -2610.19MB 79.47%  compress/flate.NewWriter (inline)
 -512.09MB 15.59% 84.22%  -512.09MB 15.59%  compress/flate.(*compressor).initDeflate (inline)
  109.28MB  3.33% 80.89%  -356.36MB 10.85%  compress/flate.(*compressor).init
   62.46MB  1.90% 78.99%    62.46MB  1.90%  compress/flate.newDeflateFast (inline)
  -32.57MB  0.99% 79.98%   -32.57MB  0.99%  compress/flate.(*huffmanEncoder).generate
  -27.85MB  0.85% 80.83%   -27.85MB  0.85%  net/http.init.func16
  -14.03MB  0.43% 81.25%   -14.03MB  0.43%  sync.(*Pool).pinSlow
   -7.01MB  0.21% 81.47%   -16.01MB  0.49%  compress/flate.newHuffmanBitWriter (inline)
    4.02MB  0.12% 81.34%   198.87MB  6.06%  io.WriteString
   -2.51MB 0.076% 81.42%   183.78MB  5.60%  github.com/SlawaBE/go-metrics-collector/internal/handler.(*ListMetricHandler).ServeHTTP
      -1MB  0.03% 81.45% -2656.42MB 80.88%  github.com/SlawaBE/go-metrics-collector/internal/server.Run.RequestLogger.func5
         0     0% 81.45%   -32.57MB  0.99%  compress/flate.(*Writer).Close (inline)
         0     0% 81.45%   -32.57MB  0.99%  compress/flate.(*compressor).close
         0     0% 81.45%   -34.58MB  1.05%  compress/flate.(*compressor).deflate
         0     0% 81.45%   -34.58MB  1.05%  compress/flate.(*compressor).writeBlock
         0     0% 81.45%   -22.05MB  0.67%  compress/flate.(*huffmanBitWriter).indexTokens
         0     0% 81.45%   -34.58MB  1.05%  compress/flate.(*huffmanBitWriter).writeBlock
         0     0% 81.45%   -32.57MB  0.99%  compress/gzip.(*Writer).Close
         0     0% 81.45% -2610.19MB 79.47%  compress/gzip.(*Writer).Write
         0     0% 81.45%    96.21MB  2.93%  encoding/json.(*Encoder).Encode
         0     0% 81.45%   -34.58MB  1.05%  github.com/SlawaBE/go-metrics-collector/internal/gzip.(*compressWriter).Close
         0     0% 81.45% -2988.42MB 90.99%  github.com/SlawaBE/go-metrics-collector/internal/gzip.(*compressWriter).Write
         0     0% 81.45%    47.51MB  1.45%  github.com/SlawaBE/go-metrics-collector/internal/handler.(*JSONBatchUpdateMetricsHandler).ServeHTTP
         0     0% 81.45%    92.63MB  2.82%  github.com/SlawaBE/go-metrics-collector/internal/handler.(*JSONGetMetricHandler).ServeHTTP
         0     0% 81.45%    42.15MB  1.28%  github.com/SlawaBE/go-metrics-collector/internal/handler.(*JSONUpdateMetricHandler).ServeHTTP
         0     0% 81.45%   378.23MB 11.52%  github.com/SlawaBE/go-metrics-collector/internal/middleware.(*gzipResponseWriter).Write
         0     0% 81.45% -2654.42MB 80.82%  github.com/SlawaBE/go-metrics-collector/internal/server.Run.GZip.func4
         0     0% 81.45%   367.57MB 11.19%  github.com/go-chi/chi/v5.(*Mux).ServeHTTP
         0     0% 81.45%   366.07MB 11.15%  github.com/go-chi/chi/v5.(*Mux).routeHTTP
         0     0% 81.45%   -30.36MB  0.92%  net/http.(*Request).write
         0     0% 81.45% -2670.99MB 81.33%  net/http.(*conn).serve
         0     0% 81.45%   -30.36MB  0.92%  net/http.(*persistConn).writeLoop
         0     0% 81.45%   -29.36MB  0.89%  net/http.(*transferWriter).doBodyCopy
         0     0% 81.45%   -29.36MB  0.89%  net/http.(*transferWriter).writeBody
         0     0% 81.45% -2656.42MB 80.88%  net/http.HandlerFunc.ServeHTTP
         0     0% 81.45%   -29.36MB  0.89%  net/http.getCopyBuf (inline)
         0     0% 81.45% -2656.42MB 80.88%  net/http.serverHandler.ServeHTTP
         0     0% 81.45%   -45.42MB  1.38%  sync.(*Pool).Get
         0     0% 81.45%   -14.03MB  0.43%  sync.(*Pool).pin
```

## alloc_objects

```text
File: server.exe
Build ID: D:\work\projects\go-metrics-collector\cmd\server\server.exe2026-10-05 19:16:12.4388163 +0300 MSK
Type: alloc_objects
Time: 2026-10-05 19:12:54 MSK
Showing nodes accounting for 67512, 3.59% of 1882544 total
Dropped 38 nodes (cum <= 9412)
      flat  flat%   sum%        cum   cum%
    -93624  4.97%  4.97%     -93624  4.97%  github.com/jackc/pgx/v5/pgproto3.(*AuthenticationSASL).Decode
     91185  4.84%  0.13%     147923  7.86%  github.com/SlawaBE/go-metrics-collector/internal/storage.(*DBStorage).GetValues.func1
     65537  3.48%  3.35%      76459  4.06%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next.func14
     65537  3.48%  6.83%      65537  3.48%  reflect.unsafe_New
    -65536  3.48%  3.35%     -65536  3.48%  encoding/json.(*decodeState).literalStore
     42982  2.28%  5.63%      42982  2.28%  net/textproto.readMIMEHeader
     38229  2.03%  7.67%      38229  2.03%  github.com/jackc/pgx/v5/pgproto3.(*ParameterStatus).Decode
     37449  1.99%  9.65%      57134  3.03%  github.com/jackc/pgx/v5/pgconn/ctxwatch.(*ContextWatcher).Watch
    -37177  1.97%  7.68%     -37177  1.97%  compress/flate.newHuffmanEncoder (inline)
     35811  1.90%  9.58%     -12673  0.67%  net/http.(*Transport).dialConn
     33280  1.77% 11.35%     -80903  4.30%  github.com/jackc/pgx/v5.connect
     32768  1.74% 13.09%      21845  1.16%  database/sql.(*Rows).close
    -32768  1.74% 11.35%     -32768  1.74%  encoding/json.(*scanner).pushParseState
     32768  1.74% 13.09%      22755  1.21%  github.com/jackc/pgx/v5/pgconn.(*scramClient).clientFinalMessage
    -32768  1.74% 11.35%     -32768  1.74%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next.func6
     32768  1.74% 13.09%      32768  1.74%  internal/strconv.FormatFloat (inline)
     32768  1.74% 14.83%      32768  1.74%  internal/syscall/windows.errnoErr (inline)
     32768  1.74% 16.57%      32768  1.74%  net.IP.String
     32768  1.74% 18.31%      32768  1.74%  net.JoinHostPort (inline)
    -32768  1.74% 16.57%     -32768  1.74%  net.addrList.partition (inline)
     31278  1.66% 18.23%      31278  1.66%  net/textproto.MIMEHeader.Set (inline)
    -29491  1.57% 16.67%     -30175  1.60%  github.com/jackc/pgx/v5/pgconn.(*PgConn).Prepare
    -27821  1.48% 15.19%     152870  8.12%  github.com/SlawaBE/go-metrics-collector/internal/service.(*MetricsService).List
    -22748  1.21% 13.98%     -22748  1.21%  io.init.func1
    -21846  1.16% 12.82%     -21846  1.16%  net/textproto.NewReader (inline)
    -21845  1.16% 11.66%     -21845  1.16%  encoding/base64.(*Encoding).DecodeString
    -21845  1.16% 10.50%      -3309  0.18%  net/http.(*persistConn).readLoop
    -19114  1.02%  9.48%     -23210  1.23%  crypto/internal/fips140/hmac.New[go.shape.interface { BlockSize int; Reset; Size int; Sum []uint8; Write  }]
     18724  0.99% 10.48%       7801  0.41%  context.AfterFunc
     16384  0.87% 11.35%      21064  1.12%  database/sql.resultFromStatement
    -16384  0.87% 10.48%     -16384  0.87%  fmt.(*buffer).write (inline)
    -16384  0.87%  9.61%     -16384  0.87%  maps.Copy[go.shape.map[string]string,go.shape.map[string]string,go.shape.string,go.shape.string] (inline)
     16384  0.87% 10.48%      14894  0.79%  net/http.(*Client).makeHeadersCopier
     16383  0.87% 11.35%      16383  0.87%  syscall.(*RawSockaddrAny).Sockaddr
    -15480  0.82% 10.53%     -49678  2.64%  compress/flate.newHuffmanBitWriter (inline)
    -14823  0.79%  9.74%     -14823  0.79%  compress/flate.(*huffmanEncoder).generate
    -14348  0.76%  8.98%     -14348  0.76%  sync.(*Pool).pinSlow
    -13653  0.73%  8.25%     -13653  0.73%  github.com/jackc/pgx/v5/pgtype.NewMap
    -13444  0.71%  7.54%     -13443  0.71%  context.(*cancelCtx).propagateCancel
    -13106   0.7%  6.84%     -13106   0.7%  net/http.NewRequestWithContext
    -12483  0.66%  6.18%      -3120  0.17%  time.NewTimer
    -12482  0.66%  5.52%     -12482  0.66%  net/textproto.(*Reader).ReadLine (inline)
     10923  0.58%  6.10%      10923  0.58%  github.com/SlawaBE/go-metrics-collector/internal/middleware.newGzipResponseWriter
     10923  0.58%  6.68%      79874  4.24%  github.com/SlawaBE/go-metrics-collector/internal/service.(*AuditService).notify
    -10923  0.58%  6.10%     -10923  0.58%  net.sockaddrToTCP
     10923  0.58%  6.68%      10923  0.58%  net/netip.ParseAddr
    -10923  0.58%  6.10%     -10923  0.58%  net/url.UserPassword (inline)
     10922  0.58%  6.68%      10922  0.58%  github.com/jackc/pgx/v5/pgtype.scanPlanString.Scan
    -10468  0.56%  6.12%      -7947  0.42%  github.com/jackc/pgx/v5/pgconn.parseEnvSettings
     -9831  0.52%  5.60%      -9831  0.52%  database/sql.driverArgsConnLocked
      9364   0.5%  6.10%       9364   0.5%  net.newFD (inline)
      9363   0.5%  6.59%       9363   0.5%  time.newTimer
      9362   0.5%  7.09%      14044  0.75%  context.WithDeadlineCause
     -9362   0.5%  6.59%     -32410  1.72%  net.(*sysDialer).dialParallel
      8192  0.44%  7.03%       8192  0.44%  database/sql.(*Rows).initContextClose
     -8192  0.44%  6.59%     -45221  2.40%  database/sql.(*driverConn).prepareLocked
      8192  0.44%  7.03%       8192  0.44%  encoding/json.Marshal
     -8192  0.44%  6.59%     -16612  0.88%  encoding/json.newEncodeState
     -8192  0.44%  6.16%      -8192  0.44%  github.com/jackc/pgx/v5/internal/stmtcache.StatementName
      8192  0.44%  6.59%       8192  0.44%  github.com/jackc/pgx/v5/pgconn.buildConnectOneConfigs
     -8192  0.44%  6.16%      -8192  0.44%  github.com/jackc/pgx/v5/pgconn/internal/bgreader.New (inline)
     -8192  0.44%  5.72%      -8192  0.44%  internal/poll.(*FD).pin
     -7726  0.41%  5.31%      -7726  0.41%  compress/flate.(*compressor).initDeflate (inline)
      7282  0.39%  5.70%       7282  0.39%  github.com/jackc/pgx/v5/pgconn.makeDefaultDialer (inline)
     -6746  0.36%  5.34%      -6746  0.36%  bufio.NewReaderSize (inline)
      6554  0.35%  5.69%       -648 0.034%  context.withCancel (inline)
     -6554  0.35%  5.34%      -6554  0.35%  github.com/jackc/pgx/v5/pgconn.(*ResultReader).Read
     -6554  0.35%  4.99%     -13901  0.74%  github.com/jackc/pgx/v5/stdlib.(*Conn).QueryContext
     -6453  0.34%  4.65%      -6453  0.34%  strings.(*Builder).WriteString (inline)
      5958  0.32%  4.97%     -28960  1.54%  github.com/jackc/pgx/v5/pgconn.ParseConfigWithOptions
      5042  0.27%  5.23%       5042  0.27%  syscall.Getenv
      4917  0.26%  5.49%       4917  0.26%  net/http.(*Request).WithContext (inline)
      4915  0.26%  5.76%       4915  0.26%  encoding/json.NewDecoder (inline)
      4681  0.25%  6.00%    -108485  5.76%  github.com/jackc/pgx/v5/pgconn.connectOne
     -4681  0.25%  5.76%     -14044  0.75%  net.(*Resolver).lookupIP
     -4370  0.23%  5.52%     -14345  0.76%  net/http.(*conn).readRequest
     -4096  0.22%  5.31%      -4096  0.22%  crypto/internal/fips140/sha256.New (inline)
      4096  0.22%  5.52%      10650  0.57%  database/sql.(*DB).beginDC
      4096  0.22%  5.74%     -73484  3.90%  database/sql.(*DB).conn
     -4033  0.21%  5.53%     -34208  1.82%  github.com/jackc/pgx/v5.(*Conn).Prepare
      3641  0.19%  5.72%     -41580  2.21%  database/sql.(*DB).prepareDC
     -3641  0.19%  5.53%      -3641  0.19%  github.com/jackc/pgx/v5.(*Conn).getRows
      3641  0.19%  5.72%      18536  0.98%  net/http.ReadResponse
     -3641  0.19%  5.53%      -3641  0.19%  net/http.cloneURL (inline)
     -3563  0.19%  5.34%     -59635  3.17%  compress/flate.NewWriter (inline)
     -3465  0.18%  5.15%      -3465  0.18%  runtime.mallocgc
     -3278  0.17%  4.98%      69081  3.67%  github.com/SlawaBE/go-metrics-collector/internal/server.Run.RequestLogger.func5
     -3278  0.17%  4.81%      -9975  0.53%  net/http.readRequest
     -3277  0.17%  4.63%      -3277  0.17%  compress/gzip.NewWriterLevel
     -3277  0.17%  4.46%      -8986  0.48%  database/sql.(*DB).queryDC
     -3277  0.17%  4.28%      -3277  0.17%  net/http.newTransferWriter
     -2980  0.16%  4.13%      -2980  0.16%  net/http.Header.Clone (inline)
     -2731  0.15%  3.98%     -52304  2.78%  net.(*Dialer).DialContext
     -2521  0.13%  3.85%      -2521  0.13%  database/sql.(*DB).addDepLocked (inline)
     -2521  0.13%  3.71%      -2553  0.14%  net.(*Resolver).lookupIP.func1
      2426  0.13%  3.84%       2426  0.13%  bytes.growSlice
     -2185  0.12%  3.72%      -2185  0.12%  go.uber.org/zap/internal/stacktrace.Capture
     -2185  0.12%  3.61%      -2185  0.12%  internal/poll.(*FD).Accept
     -2184  0.12%  3.49%      -2184  0.12%  github.com/jackc/pgx/v5/pgconn.configTLS
     -2048  0.11%  3.38%      -2048  0.11%  net/http.ListenAndServe (inline)
      1820 0.097%  3.48%      -6582  0.35%  github.com/jackc/pgx/v5/pgconn.defaultSettings
      1820 0.097%  3.58%      -9103  0.48%  github.com/jackc/pgx/v5/pgconn.parseURLSettings
     -1756 0.093%  3.48%      -1756 0.093%  github.com/jackc/pgx/v5/pgproto3.(*DataRow).Decode
      1638 0.087%  3.57%       1638 0.087%  net/http.setupRewindBody (inline)
     -1489 0.079%  3.49%      -1489 0.079%  github.com/jackc/pgx/v5/pgconn.(*Config).Copy
      1489 0.079%  3.57%       1489 0.079%  net/textproto.MIMEHeader.Add (inline)
     -1365 0.073%  3.50%      -1365 0.073%  sync.(*poolChain).pushHead
      1277 0.068%  3.57%       1277 0.068%  net/http.(*Transport).getConn
     -1260 0.067%  3.50%      -1260 0.067%  github.com/jackc/pgx/v5/pgtype.(*Map).planEncodeDepth
      1260 0.067%  3.57%       1260 0.067%  unicode/utf16.Encode
      1028 0.055%  3.62%       2800  0.15%  io.WriteString
      -891 0.047%  3.57%       -891 0.047%  net/http.init.func16
       707 0.038%  3.61%     -56072  2.98%  compress/flate.(*compressor).init
      -642 0.034%  3.58%     157828  8.38%  github.com/SlawaBE/go-metrics-collector/internal/handler.(*ListMetricHandler).ServeHTTP
       625 0.033%  3.61%        625 0.033%  compress/flate.newDeflateFast (inline)
      -455 0.024%  3.59%       -881 0.047%  github.com/jackc/pgx/v5/pgproto3.NewFrontend
         0     0%  3.59%     -24344  1.29%  bufio.(*Writer).Flush
         0     0%  3.59%      -1670 0.089%  bytes.(*Buffer).Write
         0     0%  3.59%       4096  0.22%  bytes.(*Buffer).WriteByte
         0     0%  3.59%       2426  0.13%  bytes.(*Buffer).grow
         0     0%  3.59%     -14823  0.79%  compress/flate.(*Writer).Close (inline)
         0     0%  3.59%     -14823  0.79%  compress/flate.(*compressor).close
         0     0%  3.59%     -15735  0.84%  compress/flate.(*compressor).deflate
         0     0%  3.59%        912 0.048%  compress/flate.(*compressor).encSpeed
         0     0%  3.59%     -15735  0.84%  compress/flate.(*compressor).writeBlock
         0     0%  3.59%     -10034  0.53%  compress/flate.(*huffmanBitWriter).indexTokens
         0     0%  3.59%     -15735  0.84%  compress/flate.(*huffmanBitWriter).writeBlock
         0     0%  3.59%        456 0.024%  compress/flate.(*huffmanBitWriter).writeBlockDynamic
         0     0%  3.59%        456 0.024%  compress/flate.(*huffmanBitWriter).writeBlockHuff
         0     0%  3.59%      -2979  0.16%  compress/flate.generateFixedLiteralEncoding
         0     0%  3.59%      -2979  0.16%  compress/flate.init
         0     0%  3.59%     -14823  0.79%  compress/gzip.(*Writer).Close
         0     0%  3.59%     -59635  3.17%  compress/gzip.(*Writer).Write
         0     0%  3.59%      -3277  0.17%  compress/gzip.NewWriter (inline)
         0     0%  3.59%       6554  0.35%  context.WithCancel
         0     0%  3.59%      -7202  0.38%  context.WithCancelCause
         0     0%  3.59%      14044  0.75%  context.WithDeadline (inline)
         0     0%  3.59%       9362   0.5%  context.WithTimeout
         0     0%  3.59%     -19114  1.02%  crypto/hmac.New
         0     0%  3.59%      -4096  0.22%  crypto/internal/fips140/pbkdf2.Key[go.shape.interface { BlockSize int; Reset; Size int; Sum []uint8; Write  }]
         0     0%  3.59%      -4096  0.22%  crypto/pbkdf2.Key[go.shape.interface { BlockSize int; Reset; Size int; Sum []uint8; Write  }]
         0     0%  3.59%      -4096  0.22%  crypto/pbkdf2.Key[go.shape.interface { BlockSize int; Reset; Size int; Sum []uint8; Write  }].UnwrapNew[go.shape.interface { BlockSize int; Reset; Size int; Sum []uint8; Write  }].func1
         0     0%  3.59%      -4096  0.22%  crypto/sha256.New
         0     0%  3.59%    -106583  5.66%  database/sql.(*DB).Begin (inline)
         0     0%  3.59%    -106583  5.66%  database/sql.(*DB).BeginTx
         0     0%  3.59%    -106583  5.66%  database/sql.(*DB).BeginTx.func1
         0     0%  3.59%      34763  1.85%  database/sql.(*DB).QueryContext
         0     0%  3.59%      34763  1.85%  database/sql.(*DB).QueryContext.func1
         0     0%  3.59%      21717  1.15%  database/sql.(*DB).QueryRowContext (inline)
         0     0%  3.59%    -106583  5.66%  database/sql.(*DB).begin
         0     0%  3.59%       4681  0.25%  database/sql.(*DB).connectionOpener
         0     0%  3.59%     -45221  2.40%  database/sql.(*DB).prepareDC.func2
         0     0%  3.59%      -1561 0.083%  database/sql.(*DB).putConn
         0     0%  3.59%      34763  1.85%  database/sql.(*DB).query
         0     0%  3.59%     -13901  0.74%  database/sql.(*DB).queryDC.func1
         0     0%  3.59%     -50756  2.70%  database/sql.(*DB).retry
         0     0%  3.59%      87381  4.64%  database/sql.(*Row).Scan
         0     0%  3.59%      21845  1.16%  database/sql.(*Rows).Close
         0     0%  3.59%      43691  2.32%  database/sql.(*Rows).Next
         0     0%  3.59%      43691  2.32%  database/sql.(*Rows).Next.func1
         0     0%  3.59%      65537  3.48%  database/sql.(*Rows).Scan
         0     0%  3.59%      43691  2.32%  database/sql.(*Rows).nextLocked
         0     0%  3.59%      65537  3.48%  database/sql.(*Rows).scanLocked
         0     0%  3.59%      16745  0.89%  database/sql.(*Stmt).Close
         0     0%  3.59%      21064  1.12%  database/sql.(*Stmt).ExecContext
         0     0%  3.59%      21064  1.12%  database/sql.(*Stmt).ExecContext.func1
         0     0%  3.59%      25879  1.37%  database/sql.(*Tx).Commit
         0     0%  3.59%     -46261  2.46%  database/sql.(*Tx).PrepareContext
         0     0%  3.59%       9362   0.5%  database/sql.(*Tx).close (inline)
         0     0%  3.59%      16745  0.89%  database/sql.(*Tx).closePrepared
         0     0%  3.59%      -4681  0.25%  database/sql.(*Tx).grabConn
         0     0%  3.59%      -1561 0.083%  database/sql.(*driverConn).Close
         0     0%  3.59%      -1561 0.083%  database/sql.(*driverConn).finalClose
         0     0%  3.59%      -1561 0.083%  database/sql.(*driverConn).finalClose.func2
         0     0%  3.59%      -1561 0.083%  database/sql.(*driverConn).releaseConn
         0     0%  3.59%      30247  1.61%  database/sql.(*driverConn).resetSession
         0     0%  3.59%      16745  0.89%  database/sql.(*driverStmt).Close
         0     0%  3.59%      65537  3.48%  database/sql.convertAssignRows
         0     0%  3.59%     -37029  1.97%  database/sql.ctxDriverPrepare
         0     0%  3.59%     -13901  0.74%  database/sql.ctxDriverQuery
         0     0%  3.59%      14511  0.77%  database/sql.ctxDriverStmtExec
         0     0%  3.59%     -17220  0.91%  database/sql.withLock
         0     0%  3.59%     -98303  5.22%  encoding/json.(*Decoder).Decode
         0     0%  3.59%     -32767  1.74%  encoding/json.(*Decoder).readValue
         0     0%  3.59%     -10983  0.58%  encoding/json.(*Encoder).Encode
         0     0%  3.59%     -65536  3.48%  encoding/json.(*decodeState).object
         0     0%  3.59%     -65536  3.48%  encoding/json.(*decodeState).unmarshal
         0     0%  3.59%     -65536  3.48%  encoding/json.(*decodeState).value
         0     0%  3.59%     -32768  1.74%  encoding/json.stateBeginValue
         0     0%  3.59%     -16384  0.87%  fmt.(*fmt).fmtBs
         0     0%  3.59%     -16384  0.87%  fmt.(*fmt).pad
         0     0%  3.59%     -16384  0.87%  fmt.(*pp).doPrintf
         0     0%  3.59%     -16384  0.87%  fmt.(*pp).fmtBytes
         0     0%  3.59%     -16384  0.87%  fmt.(*pp).printArg
         0     0%  3.59%     -16840  0.89%  fmt.Appendf
         0     0%  3.59%       -684 0.036%  fmt.newPrinter
         0     0%  3.59%     -15735  0.84%  github.com/SlawaBE/go-metrics-collector/internal/gzip.(*compressWriter).Close
         0     0%  3.59%     -63746  3.39%  github.com/SlawaBE/go-metrics-collector/internal/gzip.(*compressWriter).Write
         0     0%  3.59%      -2979  0.16%  github.com/SlawaBE/go-metrics-collector/internal/gzip.(*compressWriter).WriteHeader
         0     0%  3.59%      -3277  0.17%  github.com/SlawaBE/go-metrics-collector/internal/gzip.NewCompressWriter
         0     0%  3.59%     -29077  1.54%  github.com/SlawaBE/go-metrics-collector/internal/handler.(*JSONBatchUpdateMetricsHandler).ServeHTTP
         0     0%  3.59%      70136  3.73%  github.com/SlawaBE/go-metrics-collector/internal/handler.(*JSONGetMetricHandler).ServeHTTP
         0     0%  3.59%     -54251  2.88%  github.com/SlawaBE/go-metrics-collector/internal/handler.(*JSONUpdateMetricHandler).ServeHTTP
         0     0%  3.59%        912 0.048%  github.com/SlawaBE/go-metrics-collector/internal/middleware.(*gzipResponseWriter).Close
         0     0%  3.59%       4111  0.22%  github.com/SlawaBE/go-metrics-collector/internal/middleware.(*gzipResponseWriter).Write
         0     0%  3.59%       1489 0.079%  github.com/SlawaBE/go-metrics-collector/internal/middleware.(*gzipResponseWriter).WriteHeader
         0     0%  3.59%      -1670 0.089%  github.com/SlawaBE/go-metrics-collector/internal/middleware.(*responseRecorder).Write
         0     0%  3.59%      -1490 0.079%  github.com/SlawaBE/go-metrics-collector/internal/middleware.(*responseWriter).WriteHeader
         0     0%  3.59%     151120  8.03%  github.com/SlawaBE/go-metrics-collector/internal/retry.RetryWithBackoff
         0     0%  3.59%       2497  0.13%  github.com/SlawaBE/go-metrics-collector/internal/server.Run
         0     0%  3.59%      71783  3.81%  github.com/SlawaBE/go-metrics-collector/internal/server.Run.GZip.func4
         0     0%  3.59%      -2048  0.11%  github.com/SlawaBE/go-metrics-collector/internal/server.Run.func1
         0     0%  3.59%      52430  2.79%  github.com/SlawaBE/go-metrics-collector/internal/service.(*AuditService).SendMetric
         0     0%  3.59%      27444  1.46%  github.com/SlawaBE/go-metrics-collector/internal/service.(*AuditService).SendMetrics
         0     0%  3.59%      -8420  0.45%  github.com/SlawaBE/go-metrics-collector/internal/service.(*FileAuditSubscriber).Notify
         0     0%  3.59%      75460  4.01%  github.com/SlawaBE/go-metrics-collector/internal/service.(*HTTPAuditSubscriber).Notify
         0     0%  3.59%     109098  5.80%  github.com/SlawaBE/go-metrics-collector/internal/service.(*MetricsService).GetMetric
         0     0%  3.59%     -77183  4.10%  github.com/SlawaBE/go-metrics-collector/internal/service.(*MetricsService).UpdateMetric
         0     0%  3.59%     -28718  1.53%  github.com/SlawaBE/go-metrics-collector/internal/service.(*MetricsService).UpdateMetrics
         0     0%  3.59%     109098  5.80%  github.com/SlawaBE/go-metrics-collector/internal/storage.(*DBStorage).GetMetric
         0     0%  3.59%     109098  5.80%  github.com/SlawaBE/go-metrics-collector/internal/storage.(*DBStorage).GetMetric.func1
         0     0%  3.59%     147923  7.86%  github.com/SlawaBE/go-metrics-collector/internal/storage.(*DBStorage).GetValues
         0     0%  3.59%     -28718  1.53%  github.com/SlawaBE/go-metrics-collector/internal/storage.(*DBStorage).UpdateAll
         0     0%  3.59%     -28718  1.53%  github.com/SlawaBE/go-metrics-collector/internal/storage.(*DBStorage).UpdateAll.func1
         0     0%  3.59%     -77183  4.10%  github.com/SlawaBE/go-metrics-collector/internal/storage.(*DBStorage).UpdateMetric
         0     0%  3.59%     -77183  4.10%  github.com/SlawaBE/go-metrics-collector/internal/storage.(*DBStorage).UpdateMetric.func1
         0     0%  3.59%      32768  1.74%  github.com/SlawaBE/go-metrics-collector/internal/utils.ConvertGauge
         0     0%  3.59%     145685  7.74%  github.com/go-chi/chi/v5.(*Mux).ServeHTTP
         0     0%  3.59%     144636  7.68%  github.com/go-chi/chi/v5.(*Mux).routeHTTP
         0     0%  3.59%      -6242  0.33%  github.com/jackc/pgx/v5.(*Conn).Close
         0     0%  3.59%      12064  0.64%  github.com/jackc/pgx/v5.(*Conn).Deallocate
         0     0%  3.59%      14283  0.76%  github.com/jackc/pgx/v5.(*Conn).Exec
         0     0%  3.59%      -7347  0.39%  github.com/jackc/pgx/v5.(*Conn).Query
         0     0%  3.59%      14283  0.76%  github.com/jackc/pgx/v5.(*Conn).exec
         0     0%  3.59%      14511  0.77%  github.com/jackc/pgx/v5.(*Conn).execPrepared
         0     0%  3.59%      -5371  0.29%  github.com/jackc/pgx/v5.(*Conn).getStatementDescription
         0     0%  3.59%      -1260 0.067%  github.com/jackc/pgx/v5.(*ExtendedQueryBuilder).Build
         0     0%  3.59%      -1260 0.067%  github.com/jackc/pgx/v5.(*ExtendedQueryBuilder).appendParam
         0     0%  3.59%      -1260 0.067%  github.com/jackc/pgx/v5.(*ExtendedQueryBuilder).encodeExtendedParamValue
         0     0%  3.59%     -81027  4.30%  github.com/jackc/pgx/v5.ConnectConfig
         0     0%  3.59%     -28960  1.54%  github.com/jackc/pgx/v5.ParseConfig (inline)
         0     0%  3.59%     -28960  1.54%  github.com/jackc/pgx/v5.ParseConfigWithOptions
         0     0%  3.59%       -426 0.023%  github.com/jackc/pgx/v5/internal/iobufpool.Get
         0     0%  3.59%      -6242  0.33%  github.com/jackc/pgx/v5/pgconn.(*PgConn).Close
         0     0%  3.59%      12064  0.64%  github.com/jackc/pgx/v5/pgconn.(*PgConn).Deallocate
         0     0%  3.59%      30247  1.61%  github.com/jackc/pgx/v5/pgconn.(*PgConn).Exec
         0     0%  3.59%      23990  1.27%  github.com/jackc/pgx/v5/pgconn.(*PgConn).ExecStatement
         0     0%  3.59%      30247  1.61%  github.com/jackc/pgx/v5/pgconn.(*PgConn).Ping
         0     0%  3.59%      25746  1.37%  github.com/jackc/pgx/v5/pgconn.(*PgConn).execExtendedPrefix
         0     0%  3.59%      -1756 0.093%  github.com/jackc/pgx/v5/pgconn.(*PgConn).execExtendedSuffix
         0     0%  3.59%     -62159  3.30%  github.com/jackc/pgx/v5/pgconn.(*PgConn).peekMessage
         0     0%  3.59%     -61931  3.29%  github.com/jackc/pgx/v5/pgconn.(*PgConn).receiveMessage
         0     0%  3.59%      -4096  0.22%  github.com/jackc/pgx/v5/pgconn.(*PgConn).rxSASLFinal
         0     0%  3.59%     -33451  1.78%  github.com/jackc/pgx/v5/pgconn.(*PgConn).scramAuth
         0     0%  3.59%      -1756 0.093%  github.com/jackc/pgx/v5/pgconn.(*ResultReader).readUntilRowDescription
         0     0%  3.59%      -1756 0.093%  github.com/jackc/pgx/v5/pgconn.(*ResultReader).receiveMessage
         0     0%  3.59%     -16612  0.88%  github.com/jackc/pgx/v5/pgconn.(*scramClient).clientFirstMessage
         0     0%  3.59%     -13653  0.73%  github.com/jackc/pgx/v5/pgconn.(*scramClient).recvServerFinalMessage
         0     0%  3.59%     -21845  1.16%  github.com/jackc/pgx/v5/pgconn.(*scramClient).recvServerFirstMessage
         0     0%  3.59%    -100293  5.33%  github.com/jackc/pgx/v5/pgconn.ConnectConfig
         0     0%  3.59%      -5461  0.29%  github.com/jackc/pgx/v5/pgconn.computeClientProof
         0     0%  3.59%     -19114  1.02%  github.com/jackc/pgx/v5/pgconn.computeHMAC
         0     0%  3.59%     -13653  0.73%  github.com/jackc/pgx/v5/pgconn.computeServerSignature
         0     0%  3.59%    -108485  5.76%  github.com/jackc/pgx/v5/pgconn.connectPreferred
         0     0%  3.59%     -16384  0.87%  github.com/jackc/pgx/v5/pgconn.mergeSettings (inline)
         0     0%  3.59%      -5008  0.27%  github.com/jackc/pgx/v5/pgconn/internal/bgreader.(*BGReader).Read
         0     0%  3.59%     -62159  3.30%  github.com/jackc/pgx/v5/pgproto3.(*Frontend).Receive
         0     0%  3.59%      -5008  0.27%  github.com/jackc/pgx/v5/pgproto3.(*chunkReader).Next
         0     0%  3.59%       -426 0.023%  github.com/jackc/pgx/v5/pgproto3.newChunkReader (inline)
         0     0%  3.59%      -1260 0.067%  github.com/jackc/pgx/v5/pgtype.(*Map).Encode
         0     0%  3.59%      -1260 0.067%  github.com/jackc/pgx/v5/pgtype.(*Map).PlanEncode (inline)
         0     0%  3.59%      -1561 0.083%  github.com/jackc/pgx/v5/stdlib.(*Conn).Close
         0     0%  3.59%      14511  0.77%  github.com/jackc/pgx/v5/stdlib.(*Conn).ExecContext
         0     0%  3.59%     -37029  1.97%  github.com/jackc/pgx/v5/stdlib.(*Conn).PrepareContext
         0     0%  3.59%      30247  1.61%  github.com/jackc/pgx/v5/stdlib.(*Conn).ResetSession
         0     0%  3.59%      43691  2.32%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
         0     0%  3.59%      16745  0.89%  github.com/jackc/pgx/v5/stdlib.(*Stmt).Close
         0     0%  3.59%      14511  0.77%  github.com/jackc/pgx/v5/stdlib.(*Stmt).ExecContext
         0     0%  3.59%    -109987  5.84%  github.com/jackc/pgx/v5/stdlib.(*driverConnector).Connect
         0     0%  3.59%       2487  0.13%  go.uber.org/zap.(*Logger).Info
         0     0%  3.59%      -2413  0.13%  go.uber.org/zap.(*Logger).check
         0     0%  3.59%       -684 0.036%  go.uber.org/zap/internal/pool.(*Pool[go.shape.*uint8]).Get (inline)
         0     0%  3.59%       4900  0.26%  go.uber.org/zap/zapcore.(*CheckedEntry).Write
         0     0%  3.59%       4900  0.26%  go.uber.org/zap/zapcore.(*ioCore).Write
         0     0%  3.59%       4096  0.22%  go.uber.org/zap/zapcore.(*jsonEncoder).AddReflected
         0     0%  3.59%       3640  0.19%  go.uber.org/zap/zapcore.(*jsonEncoder).EncodeEntry
         0     0%  3.59%       -456 0.024%  go.uber.org/zap/zapcore.(*jsonEncoder).clone
         0     0%  3.59%       4096  0.22%  go.uber.org/zap/zapcore.(*jsonEncoder).encodeReflected
         0     0%  3.59%       1260 0.067%  go.uber.org/zap/zapcore.(*lockedWriteSyncer).Write
         0     0%  3.59%       4096  0.22%  go.uber.org/zap/zapcore.Field.AddTo
         0     0%  3.59%       4096  0.22%  go.uber.org/zap/zapcore.addFields (inline)
         0     0%  3.59%      -4324  0.23%  internal/poll.(*FD).ConnectEx
         0     0%  3.59%      19568  1.04%  internal/poll.(*FD).Read
         0     0%  3.59%       1260 0.067%  internal/poll.(*FD).Write
         0     0%  3.59%      23436  1.24%  internal/poll.(*FD).execIO
         0     0%  3.59%       1260 0.067%  internal/poll.(*FD).writeConsole
         0     0%  3.59%     -14044  0.75%  internal/singleflight.(*Group).doCall
         0     0%  3.59%      32768  1.74%  internal/syscall/windows.WSAGetOverlappedResult
         0     0%  3.59%     -23888  1.27%  io.Copy (inline)
         0     0%  3.59%     -23888  1.27%  io.CopyN
         0     0%  3.59%      -5008  0.27%  io.ReadAtLeast
         0     0%  3.59%     -23888  1.27%  io.copyBuffer
         0     0%  3.59%     -23888  1.27%  io.discard.ReadFrom
         0     0%  3.59%       2497  0.13%  main.main
         0     0%  3.59%       4682  0.25%  net.(*Dialer).dialCtx
         0     0%  3.59%      10923  0.58%  net.(*Resolver).internetAddrList
         0     0%  3.59%      -2553  0.14%  net.(*Resolver).lookupIP.func2
         0     0%  3.59%      10923  0.58%  net.(*Resolver).lookupIPAddr
         0     0%  3.59%     -14044  0.75%  net.(*Resolver).lookupIPAddr.func1
         0     0%  3.59%      10923  0.58%  net.(*Resolver).resolveAddrList
         0     0%  3.59%      32768  1.74%  net.(*TCPAddr).String
         0     0%  3.59%       2497  0.13%  net.(*TCPListener).Accept
         0     0%  3.59%       2497  0.13%  net.(*TCPListener).accept
         0     0%  3.59%      19568  1.04%  net.(*conn).Read
         0     0%  3.59%      19568  1.04%  net.(*netFD).Read
         0     0%  3.59%       2497  0.13%  net.(*netFD).accept
         0     0%  3.59%     -11527  0.61%  net.(*netFD).connect
         0     0%  3.59%      -6067  0.32%  net.(*netFD).dial
         0     0%  3.59%       9181  0.49%  net.(*sysDialer).dialParallel.func1
         0     0%  3.59%      -1385 0.074%  net.(*sysDialer).dialSerial
         0     0%  3.59%      -1385 0.074%  net.(*sysDialer).dialSingle
         0     0%  3.59%      -1385 0.074%  net.(*sysDialer).dialTCP
         0     0%  3.59%      -1385 0.074%  net.(*sysDialer).doDialTCP (inline)
         0     0%  3.59%      -1385 0.074%  net.(*sysDialer).doDialTCPProto
         0     0%  3.59%     -14044  0.75%  net.init.func1
         0     0%  3.59%      -1385 0.074%  net.internetSocket
         0     0%  3.59%      32768  1.74%  net.ipEmptyString (inline)
         0     0%  3.59%      -1385 0.074%  net.socket
         0     0%  3.59%      80375  4.27%  net/http.(*Client).Do (inline)
         0     0%  3.59%      80375  4.27%  net/http.(*Client).do
         0     0%  3.59%      65481  3.48%  net/http.(*Client).send
         0     0%  3.59%      -5080  0.27%  net/http.(*Request).write
         0     0%  3.59%       2497  0.13%  net/http.(*Server).ListenAndServe
         0     0%  3.59%       2497  0.13%  net/http.(*Server).Serve
         0     0%  3.59%      69122  3.67%  net/http.(*Transport).RoundTrip
         0     0%  3.59%      32768  1.74%  net/http.(*Transport).connectMethodForRequest
         0     0%  3.59%     -41738  2.22%  net/http.(*Transport).dial
         0     0%  3.59%     -12673  0.67%  net/http.(*Transport).dialConnFor
         0     0%  3.59%      69122  3.67%  net/http.(*Transport).roundTrip
         0     0%  3.59%     -12673  0.67%  net/http.(*Transport).startDialConnForLocked.func1
         0     0%  3.59%     -24344  1.29%  net/http.(*chunkWriter).Write
         0     0%  3.59%     -24344  1.29%  net/http.(*chunkWriter).writeHeader
         0     0%  3.59%      73826  3.92%  net/http.(*conn).serve
         0     0%  3.59%      24576  1.31%  net/http.(*connReader).backgroundRead
         0     0%  3.59%      18536  0.98%  net/http.(*persistConn).readResponse
         0     0%  3.59%      40641  2.16%  net/http.(*persistConn).roundTrip
         0     0%  3.59%      -5080  0.27%  net/http.(*persistConn).writeLoop
         0     0%  3.59%      -1490 0.079%  net/http.(*response).WriteHeader
         0     0%  3.59%     -13421  0.71%  net/http.(*response).finishRequest
         0     0%  3.59%      -1575 0.084%  net/http.(*transferWriter).doBodyCopy
         0     0%  3.59%      -1575 0.084%  net/http.(*transferWriter).writeBody
         0     0%  3.59%      69081  3.67%  net/http.HandlerFunc.ServeHTTP
         0     0%  3.59%       1489 0.079%  net/http.Header.Add (inline)
         0     0%  3.59%      31278  1.66%  net/http.Header.Set (inline)
         0     0%  3.59%       -456 0.024%  net/http.Header.WriteSubset (inline)
         0     0%  3.59%       -684 0.036%  net/http.Header.sortedKeyValues
         0     0%  3.59%       -684 0.036%  net/http.Header.writeSubset
         0     0%  3.59%      32768  1.74%  net/http.canonicalAddr
         0     0%  3.59%      -1490 0.079%  net/http.cloneOrMakeHeader
         0     0%  3.59%      -1575 0.084%  net/http.getCopyBuf (inline)
         0     0%  3.59%     -22302  1.18%  net/http.newTextprotoReader
         0     0%  3.59%      10923  0.58%  net/http.putBufioWriter
         0     0%  3.59%      69122  3.67%  net/http.send
         0     0%  3.59%      69081  3.67%  net/http.serverHandler.ServeHTTP
         0     0%  3.59%      42982  2.28%  net/textproto.(*Reader).ReadMIMEHeader (inline)
         0     0%  3.59%     -10923  0.58%  net/url.Parse
         0     0%  3.59%     -10923  0.58%  net/url.parse
         0     0%  3.59%     -10923  0.58%  net/url.parseAuthority
         0     0%  3.59%       1260 0.067%  os.(*File).Write
         0     0%  3.59%       1260 0.067%  os.(*File).write (inline)
         0     0%  3.59%       5042  0.27%  os.Getenv
         0     0%  3.59%     -10923  0.58%  path/filepath.Join (inline)
         0     0%  3.59%     -10923  0.58%  path/filepath.join
         0     0%  3.59%      65537  3.48%  reflect.New
         0     0%  3.59%      -2979  0.16%  runtime.doInit (inline)
         0     0%  3.59%      -2979  0.16%  runtime.doInit1
         0     0%  3.59%      -1024 0.054%  runtime.gcBgMarkWorker
         0     0%  3.59%       -482 0.026%  runtime.main
         0     0%  3.59%      -2185  0.12%  runtime.malg
         0     0%  3.59%       -513 0.027%  runtime.mcall
         0     0%  3.59%      -3465  0.18%  runtime.newobject
         0     0%  3.59%      -2185  0.12%  runtime.newproc.func1
         0     0%  3.59%      -2185  0.12%  runtime.newproc1
         0     0%  3.59%      -2185  0.12%  runtime.systemstack
         0     0%  3.59%      32768  1.74%  strconv.FormatFloat (inline)
         0     0%  3.59%     -36817  1.96%  sync.(*Pool).Get
         0     0%  3.59%      -2277  0.12%  sync.(*Pool).Put
         0     0%  3.59%     -14348  0.76%  sync.(*Pool).pin
         0     0%  3.59%       8192  0.44%  syscall.Getpeername
         0     0%  3.59%       8191  0.44%  syscall.Getsockname
```

