# Spring Boot 2.7 vs Webflux 2.2

This project is part of the research comparing the Reactive implementation to the traditional Servlet one in Spring.

As part of the conclusion, Webflux would be more efficient than the Servlet one in low volume traffic, 
but could not tell in the high volume because it doesn't open the connection pool size configuration which was 1000.

On the other hand, Spring boot 3 with native image and Java virtual thread have been coming out so the battle would be more intensive. :)

## For 500 request/second
Servlet
| **Threads** | **Memory(MB)** | **Total execution time(ms)** | **Minimum(ms)** | **Maximum(ms)** | **Median(ms)** | **Average(ms)** | **Failed to connect** | **Failed to connect percentage(%)** | **With retry** | **Max. retry** | **Median retry** | **Average retry** |
| ----------- | -------------- | ---------------------------- | --------------- | --------------- | -------------- | --------------- | --------------------- | ----------------------------------- | -------------- | -------------- | ---------------- | ----------------- |
| 236         | 781            | 591.140                      | 15              | 489             | 299            | 241             | 0                     |                                     | 0              |                |                  |                   |
| 236         | 781            | 408.809                      | 20              | 278             | 145            | 141             | 0                     |                                     | 0              |                |                  |                   |
| 235         | 781            | 489.494                      | 11              | 403             | 173            | 150             | 0                     |                                     | 0              |                |                  |                   |
| 236         | 781            | 473.895                      | 42              | 344             | 178            | 174             | 0                     |                                     | 0              |                |                  |                   |
| 236         | 781            | 654.823                      | 7               | 456             | 191            | 183             | 0                     |                                     | 0              |                |                  |                   |

Webflux
| **Threads** | **Memory(MB)** | **Total execution time(ms)** | **Minimum(ms)** | **Maximum(ms)** | **Median(ms)** | **Average(ms)** | **Failed to connect** | **Failed to connect percentage(%)** | **With retry** | **Max. retry** | **Median retry** | **Average retry** |
| ----------- | -------------- | ---------------------------- | --------------- | --------------- | -------------- | --------------- | --------------------- | ----------------------------------- | -------------- | -------------- | ---------------- | ----------------- |
| 48          | 215            | 276                          | 4               | 33              | 20             | 19              | 0                     |                                     | 0              |                |                  |                   |
| 48          | 215            | 450.999                      | 8               | 333             | 151            | 172             | 1                     | 0.1                                 | 4              | 1              | 1                | 1                 |
| 48          | 215            | 648.049                      | 53              | 431             | 173            | 198             | 0                     |                                     | 0              |                |                  |                   |
| 48          | 215            | 308.028                      | 6               | 70              | 26             | 27              | 0                     |                                     | 0              |                |                  |                   |
| 48          | 215            | 507.115                      | 51              | 267             | 149            | 148             | 0                     |                                     | 0              |                |                  |                   |

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
| **Threads** | **Memory(MB)** | **Total execution time(ms)** | **Minimum(ms)** | **Maximum(ms)** | **Median(ms)** | **Average(ms)** | **Failed to connect** | **Failed to connect percentage(%)** | **With retry** | **Max. retry** | **Median retry** | **Average retry** |
| ----------- | -------------- | ---------------------------- | --------------- | --------------- | -------------- | --------------- | --------------------- | ----------------------------------- | -------------- | -------------- | ---------------- | ----------------- |
| 236         | 781            | 5791.966                     | 173             | 2925            | 2304           | 1956            | 0                     |                                     | 0              |                |                  |                   |
| 236         | 781            | 4105.346                     | 6               | 1282            | 759            | 718             | 0                     |                                     | 0              |                |                  |                   |
| 236         | 781            | 4807.511                     | 5               | 1985            | 646            | 785             | 0                     |                                     | 0              |                |                  |                   |
| 236         | 781            | 4130.371                     | 5               | 1463            | 908            | 808             | 0                     |                                     | 0              |                |                  |                   |
| 236         | 781            | 4233.456                     | 6               | 1508            | 774            | 761             | 0                     |                                     | 0              |                |                  |                   |

Webflux
| **Threads** | **Memory(MB)** | **Total execution time(ms)** | **Minimum(ms)** | **Maximum(ms)** | **Median(ms)** | **Average(ms)** | **Failed to connect** | **Failed to connect percentage(%)** | **With retry** | **Max. retry** | **Median retry** | **Average retry** |
| ----------- | -------------- | ---------------------------- | --------------- | --------------- | -------------- | --------------- | --------------------- | ----------------------------------- | -------------- | -------------- | ---------------- | ----------------- |
| 49          | 213            | 3994.040                     | 17              | 3607            | 264            | 348             | 704                   | 14.08                               | 1294           | 10             | 4                | 4                 |
| 49          | 213            | 4759.834                     | 10              | 3532            | 387            | 515             | 5                     | 0.1                                 | 85             | 2              | 1                | 1                 |
| 49          | 213            | 8291.384                     | 9               | 3773            | 335            | 511             | 481                   | 9.62                                | 2701           | 10             | 5                | 5                 |
| 49          | 213            | 6036.773                     | 34              | 2017            | 214            | 326             | 93                    | 1.86                                | 1545           | 10             | 3                | 3                 |
| 49          | 213            | 5702.884                     | 44              | 1390            | 268            | 334             | 729                   | 14.58                               | 1844           | 10             | 4                | 4                 |

The golang source file in this project is the request sender, and the spring applications for this test were the examples from microsoft azure site.
