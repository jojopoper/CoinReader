# Go poloniex datas reader

*Description: This library for readout history trade datas and orders from https://www.bcex.ca

## Installation


```shell
go get -u github.com/jojopoper/CoinReader/bcex/...
```

## Usage


```go
cd exsample
go build main.go
main

>  BTC / ETH Open orders (Records length = 20)
>      ************ Buy ************                         ************ Sell ************
> Price         Amount          Total                   Price           Amount          Total
> 0.09202119    1.00000000      0.09202119              0.11785062      0.07820000      0.00921592
> 0.09201818    0.05730000      0.00527264              0.11785363      0.03000000      0.00353561
> 0.09201515    1.04660000      0.09630306              0.11785664      1.00000000      0.11785664
> 0.09200629    1.00000000      0.09200629              0.11786570      1.17990000      0.13906974
> 0.09200608    5.00000000      0.46003040              0.11788800      0.71260000      0.08400699
> 0.09112230    0.17000000      0.01549079              0.11890001      0.50000000      0.05945001
> 0.09112227    0.16770000      0.01528120              0.12190000      0.30000000      0.03657000
> 0.09088320    0.47800000      0.04344217              0.12800000      0.07760000      0.00993280
> 0.08503000    0.13000000      0.01105390              0.13490000      0.25700000      0.03466930
> 0.08500000    0.22600000      0.01921000              0.13490001      0.70000000      0.09443001
> 0.08120000    0.17500000      0.01421000              0.13500000      2.56470000      0.34623450
> 0.08100000    0.11730000      0.00950130              0.14899000      0.59620000      0.08882784
> 0.08075889    0.41000000      0.03311114              0.14999999      2.97440000      0.44615997
> 0.08075888    0.39200000      0.03165748              0.15600000      4.00000000      0.62400000
> 0.08073389    1.50000000      0.12110084              -               -               -
> 0.08073388    1.85000000      0.14935768              -               -               -
> 0.06000000    2.04040000      0.12242400              -               -               -
> 0.05840000    0.01500000      0.00087600              -               -               -
> 0.05330000    0.02780000      0.00148174              -               -               -
> 0.04000000    0.70000000      0.02800000              -               -               -


>  BTC / ETH Trade history datas (Records length = 20)
> DateTime              Type    Price           Amount          Total
> 2018-02-03 23:38:27   buy     0.11786788      0.01500000      0.00176802
> 2018-02-03 23:36:24   buy     0.10000000      0.01500000      0.00150000
> 2018-02-03 23:26:42   buy     0.11786579      0.04000000      0.00471463
> 2018-02-03 23:20:09   sell    0.10600000      0.02000000      0.00212000
> 2018-02-03 23:07:37   buy     0.11787498      0.33000000      0.03889874
> 2018-02-03 20:39:33   sell    0.10630301      0.48320000      0.05136561
> 2018-02-03 20:38:30   sell    0.10788000      0.02000000      0.00215760
> 2018-02-03 20:35:57   buy     0.10650000      0.00340000      0.00036210
> 2018-02-03 20:31:41   sell    0.10650000      0.02000000      0.00213000
> 2018-02-03 20:31:41   sell    0.10690000      0.02000000      0.00213800
> 2018-02-03 08:34:37   buy     0.10788880      0.03700000      0.00399189
> 2018-02-03 08:31:30   buy     0.10788879      0.06350000      0.00685094
> 2018-02-03 05:05:45   sell    0.10500000      0.76000000      0.07980000
> 2018-02-02 22:30:52   sell    0.10700000      0.91530000      0.09793710
> 2018-02-02 22:30:39   sell    0.10700001      0.06420000      0.00686940
> 2018-02-02 22:27:51   buy     0.10600000      0.02000000      0.00212000
> 2018-02-02 22:27:44   sell    0.10300301      0.02950000      0.00303859
> 2018-02-02 21:02:48   buy     0.10299999      0.03000000      0.00309000
> 2018-02-02 20:47:18   sell    0.09900000      0.05000000      0.00495000
> 2018-02-02 20:47:12   sell    0.09900301      1.00000000      0.09900301
```