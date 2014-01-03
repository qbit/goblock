goblock
=======

Go app to grab block lists from iblocklist.com

[![Build Status](https://drone.io/github.com/qbit/goblock/status.png)](https://drone.io/github.com/qbit/goblock/latest)

### usage ###

goblock expects a config file in CWD with the below format:

```
[global]
destination = /etc/pf/adblock.list
url = http://list.iblocklist.com/

[params]
pin = ****
id = potato
fileformat = cidr
archiveformat = gz

[list]
level1 = bt_level1
level2 = bt_level2
level3 = bt_level3
ads = bt_ads
spyware = bt_spyware
proxy = bt_proxy
spider = bt_spider
hijacked = bt_hijacked
```

It will download all the files in the ``list`` section using the ``params``
specified.

Later it will combine all the files as it downloads them into ``destination``
