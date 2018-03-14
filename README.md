# mitm proxy

## Introduction

本项目实现了一个 Man-in-the-middle 代理服务器。支持 HTTP 和 HTTPS 抓包。

It is a Man-in-the-middle proxy server, supports to capture HTTP/HTTPS request and response.

## How

对于 HTTP，直接修改头部，并转发请求。

而对于 HTTPS，则是在 CONNECT 请求所建立的 TCP 隧道中，又插入了一个 TLS 服务，并由该 TLS 服务先解析成明文，再进行 TLS 请求。该插入的 TLS 服务器需要为每个域名实时签发证书，故浏览器会提示证书错误。
