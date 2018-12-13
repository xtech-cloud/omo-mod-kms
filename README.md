# 安装

```
go get github.com/xtech-cloud/omo-mod-kms
```

# 使用

```go

package main

import (
    kms "github.com/xtech-cloud/omo-mod-kms"
)

func main() {
    //创建一个应用
    appkey,appsecret, pubkey, prikey, err := kms.CreateApp("omo")

    devicecode := "9AFE1346C68974C634960F7F4B876271"

    //生成一个无自定义数据的永久授权文件
    license0, _ := MakeLicense(appkey, appsecrect, devicecode, "", 0, pubkey, prikey)

    //生成一个含自定义数据的90天授权文件
    license90, _ := MakeLicense(appkey, appsecrect, devicecode, "{\"app\":\"omo\"}", 90, pubkey, prikey)
}
```

# 授权文件格式说明

```
key:
673b7d576ea8e6ae577f162274f084ec
code:
9AFE1346C68974C634960F7F4B876271
timestamp:
1544668747
expiry:
90
storage:
{"company":"omo"}
cer:
HbkuFbfAvYQaMeQlFEkFW9-aC3v_U-VSBkt5yjg0KXLCV1q1OIMuvzwIU5M_v00NVtaOHDOa5mIs0HDPvnsrXCaTovwQKVmXAns0IZqSw8meul295xyWiS-XqFBKJkY4aRz0M8m2GNvUVGy2J03lg5QM7G9G0_IfOirkK8bWw6m0nZ779U1Jw1oypNliRQlJ1DzKIGE6raoCXPosJ7S8EV-NQaG0Tc26M_GZR-ikT8JkqvMp3kpLDBm2gO3zycgyiIfHBhFlYb54XEkY-7onP1vsA88OzQF4g2aV_twQ_zpGwkuLjTUuxLl9LnKkUAYlhekH17Ihsc_Q7TVa0zd-PL_kgjhIDmJf0aPU-hyvEDDqdodPCJEHh0ZhvQMvWhGOYEF6lnIFwQP5H52_73zVWLd303uO6q1QQzDtoQZmFC7arP9mxyCo4_7SezeDqYYnNXNIOu5PcFf70GdKN0E3U1xnI34RkMr-iMqE1WZiYGmPVnHWb93NaA_D1PvIODLGoyoniq6ev2R3avLe5JY0_YRddMb0Q_xYyyi6iB-Dqi41v6qeU_DD6H_dx3T6styO1C6KBAD30OMTOEikUYsRMT_1IHzHon9QnixW12IUvyI=
sig:
pB4_s8FggA0M-9CxD7mYfQ8oQC-oLARjpZreWvgO5kJGskd-huAQxPMbArZdZ6xQ58DjwWtIeyAgrBdpTXYI-H9gSVJrW94cDPWV4ND-i4B0kCFsavqzWIbOAXGmhWralQZL3ozDp7et4QFYs327upuQf-reNNj21a8_1ZQcvZdF-hTfr3hY7YQ5D3QHDnBHpWMvxAYcvBwuDbwIxIkXN0wkfNBRHYdXwQnOi1LK548YvUT4CX4liqgpjWf1HmEVEgGfODvxzO9KG5SKMeOtEducyKPvmrqj4rzMp3fpZeWonfLm0TarWDyouNnO967XYnB_195UmVZ2EFZMa5kl3Q==
```

| 字段 | 说明 |
|:--|:--|
|key|AppID|
|code|授权码|
|timestamp|证书生成时间|
|expiry|有效期（天）,0表示永久有效|
|storage|数据存储，可存放自定义的数据|
|cer|证书|
|sig|证书的签名|

授权文件payload部分包含key、code、timestamp、expiry、storage,以明文显示便于其他应用对授权文件做直接解析。
将payload部分加密(AES)得到授权文件的证书部分，在文件末尾对证书进行签名(RSA)防止文件被篡改。
