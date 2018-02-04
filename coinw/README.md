# Go poloniex datas reader

*Description: This library for readout history trade datas and orders from https://www.coinw.com

## Installation


```shell
go get -u github.com/jojopoper/CoinReader/coinw/...
```

## Usage


```go
cd exsample
go build main.go
main

>  CNY / ETH Open orders (Records length = 5)
>      ************ Buy ************                         ************ Sell ************
> Price         Amount          Total                   Price           Amount          Total
> 6210.00       0.08700000      540.27                  6230.00         12.44400000     77526.12
> 6200.10       1.10000000      6820.11                 6239.99         0.01300000      81.12
> 6200.00       0.00100000      6.20                    6240.00         6.36400000      39711.36
> 6199.90       4.03200000      24998.00                6250.00         0.01500000      93.75
> 6190.00       4.27000000      26431.30                6270.00         7.27200000      45595.44


>  CNY / ETH Trade history datas (Records length = 5)
> DateTime              Type    Price           Amount          Total
> 0000-01-01 15:24:05   buy     6211.06         44.69700000     277615.93
> 0000-01-01 15:24:02   sell    6210.00         0.17000000      1055.70
> 0000-01-01 15:24:01   sell    6210.12         54.54800000     338749.52
> 0000-01-01 15:23:29   buy     6210.57         11.15000000     69247.86
> 0000-01-01 15:23:26   sell    6210.00         0.31000000      1925.10
```