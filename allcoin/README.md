# Go poloniex datas reader

*Description: This library for readout history trade datas and orders from https://www.allcoin.com

## Installation


```shell
go get -u github.com/jojopoper/CoinReader/allcoin/...
```

## Usage


```go
cd exsample
go build main.go
main

>  BTC / ETH Open orders (Records length = 20)
>      ************ Buy ************                         ************ Sell ************
> Price         Amount          Total                   Price           Amount          Total
> 0.10100930    0.17320000      0.01749481              0.10588790      0.78900000      0.08354555
> 0.10099930    2.24500000      0.22674343              0.10589800      4.47850000      0.47426419
> 0.10098920    0.01280000      0.00129266              0.10598990      0.05130000      0.00543728
> 0.10096910    1.17390000      0.11852763              0.10600000      0.01000000      0.00106000
> 0.09611000    1.71120000      0.16446343              0.11600000      2.93380000      0.34032080
> 0.09550000    0.10000000      0.00955000              0.11610160      1.89700000      0.22024474
> 0.09520000    0.10000000      0.00952000              0.11899900      0.38280000      0.04555282
> 0.09000000    0.08590000      0.00773100              0.12200000      0.09000000      0.01098000
> 0.08980010    2.66480000      0.23929931              0.12610160      1.89700000      0.23921474
> 0.08901030    0.02070000      0.00184251              0.15000000      0.38080000      0.05712000
> 0.08400000    0.09000000      0.00756000              0.19995800      0.04540000      0.00907809
> 0.08000100    0.78690000      0.06295279              0.19996800      2.64010000      0.52793552
> 0.08000000    2.50620000      0.20049600              0.20000000      1.66930000      0.33386000
> 0.07800000    1.50900000      0.11770200              0.21000000      0.01160000      0.00243600
> 0.05840000    0.15000000      0.00876000              -               -               -
> 0.04400000    0.14000000      0.00616000              -               -               -
> 0.04000000    0.11490000      0.00459600              -               -               -
> 0.03400000    0.10000000      0.00340000              -               -               -
> 0.03380000    0.03900000      0.00131820              -               -               -
> 0.03313040    0.01080000      0.00035781              -               -               -


>  BTC / ETH Trade history datas (Records length = 20)
> DateTime              Type    Price           Amount          Total
> 2018-02-05 15:14:35   buy     0.10218880      0.06080000      0.00621308
> 2018-02-05 15:14:35   buy     0.10218870      0.02520000      0.00257516
> 2018-02-05 15:09:52   buy     0.10218870      0.02340000      0.00239122
> 2018-02-05 14:48:55   buy     0.10218870      0.02680000      0.00273866
> 2018-02-05 14:25:44   buy     0.10100920      0.11300000      0.01141404
> 2018-02-05 14:25:43   sell    0.10100920      0.08650000      0.00873730
> 2018-02-05 14:06:57   sell    0.10098920      0.32820000      0.03314466
> 2018-02-05 14:06:34   sell    0.10097920      0.54430000      0.05496298
> 2018-02-05 14:06:34   sell    0.10098920      0.08870000      0.00895774
> 2018-02-05 14:05:58   sell    0.10098920      0.19130000      0.01931923
> 2018-02-05 14:05:34   sell    0.10097920      0.24440000      0.02467932
> 2018-02-05 14:05:34   sell    0.10098930      0.07570000      0.00764489
> 2018-02-05 14:05:34   sell    0.10099930      0.08680000      0.00876674
> 2018-02-05 14:04:54   sell    0.10099930      0.47690000      0.04816657
> 2018-02-05 14:04:18   sell    0.10098920      0.04920000      0.00496867
> 2018-02-05 13:57:04   sell    0.10097910      0.32300000      0.03261625
> 2018-02-05 13:51:42   buy     0.10140150      0.66640000      0.06757396
> 2018-02-05 13:51:15   sell    0.10140150      0.41810000      0.04239597
> 2018-02-05 13:48:31   sell    0.10139140      0.07810000      0.00791867
> 2018-02-05 13:47:48   sell    0.10139140      0.00990000      0.00100377
```