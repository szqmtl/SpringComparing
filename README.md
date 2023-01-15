# Spring Boot 2.7 vs Webflux 2.2

This project is part of the research comparing the Reactive implementation to the traditional Servlet one in Spring.

As part of the conclusion, Webflux would be more efficient than the Servlet one in low volume traffic, 
but could not tell in the high volume because it doesn't open the connection pool size configuration which was 1000.

## For 1000 request/second
Servlet

Webflux

## For 5000 request/second
Servlet

Webflux
