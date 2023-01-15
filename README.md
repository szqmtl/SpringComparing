# Spring Boot 2.7 vs Webflux 2.2

This project is part of the research comparing the Reactive implementation to the traditional Servlet one in Spring.

As part of the conclusion, Webflux would be more efficient than the Servlet one in low volume traffic, 
but could not tell in the high volume because it doesn't open the connection pool size configuration which was 1000.

## For 1000 request/second
Servlet
| **Threads** | **Memory(MB)** | **Total execution time(ms)** | **Minimum(ms)** | **Maximum(ms)** | **Median(ms)** | **Average(ms)** | **Failed to connect** | **Failed to connect percentage(%)** | **With retry** | **Max. retry** | **Median retry** | **Average retry** |
| ----------- | -------------- | ---------------------------- | --------------- | --------------- | -------------- | --------------- | --------------------- | ----------------------------------- | -------------- | -------------- | ---------------- | ----------------- |
| 236         | 781            | 701.320                      | 6               | 254             | 116            | 110             | 0                     |                                     | 0              |                |                  |                   |
| 236         | 781            | 968.922                      | 27              | 494             | 308            | 267             | 0                     |                                     | 0              |                |                  |                   |
| 235         | 780            | 660.142                      | 4               | 296             | 119            | 108             | 0                     |                                     | 0              |                |                  |                   |
| 236         | 781            | 967.630                      | 5               | 482             | 221            | 221             | 0                     |                                     | 0              |                |                  |                   |
| 236         | 781            | 669.355                      | 4               | 226             | 69             | 75              | 0                     |                                     | 0              |                |                  |                   |

Webflux
| **Threads** | **Memory(MB)** | **Total execution time(ms)** | **Minimum(ms)** | **Maximum(ms)** | **Median(ms)** | **Average(ms)** | **Failed to connect** | **Failed to connect percentage(%)** | **With retry** | **Max. retry** | **Median retry** | **Average retry** |
| ----------- | -------------- | ---------------------------- | --------------- | --------------- | -------------- | --------------- | --------------------- | ----------------------------------- | -------------- | -------------- | ---------------- | ----------------- |
| 48          | 211            | 985.598                      | 10              | 622             | 227            | 232             | 0                     |                                     | 0              |                |                  |                   |
| 48          | 211            | 867.847                      | 60              | 701             | 191            | 235             | 1                     | 0.1                                 | 4              | 1              | 1                | 1                 |
| 48          | 212            | 715.904                      | 5               | 355             | 155            | 149             | 0                     |                                     | 0              |                |                  |                   |
| 48          | 211            | 712.836                      | 14              | 227             | 119            | 119             | 0                     |                                     | 0              |                |                  |                   |
| 48          | 211            | 726.003                      | 11              | 317             | 108            | 111             | 0                     |                                     | 0              |                |                  |                   |

## For 5000 request/second
Servlet

Webflux
