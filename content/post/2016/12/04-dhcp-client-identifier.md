+++
categories = ["Article"]
tags = ["Network", "DHCP"]
date = "2016-12-04T19:20:31+09:00"
title = "ルータの設定で固定IPを設定していたらハマった話"

+++

我が家では、家庭内でのみ利用するためのサービスを稼動させているサーバがあります。
基本的にmDNSによりホスト名とIPを紐づけているんですが、出先からアクセスしたいときのためにsshdも稼動させている関係でIPを固定したいサーバがあります。

基本的にルータの設定でできるならそちらに倒したかったため、DHCP固定割り当ての機能を利用しようとしました。
NICのMAC Addressを調べて、該当MAC Addressへ割り当てるIPを固定し、ルータ再起動。
sshdサーバ上で落ちてくるIPを見ると、なぜか固定IP設定したものが落ちてきません。

色々調べているうちに、DHCPサーバではClient Identifierというものをベースとして、IPを管理しているということがわかりました。
無意識に、MAC Addressベースで管理していると思っていたら、Client Identifierとして一般的なのがMAC Addressというだけの話っぽい。
初めて知った...。

とまれ、sshdサーバ上で稼動しているdhcpクライアントの設定を確認してみると、丁度まさにClient IdentifierとしてGUIDを使う設定になってました。
Client Identifierがわかればなんでも良かったんですが、まぁMAC Addressを使うように設定して `systemctl restart dhcpcd` すると、ちゃんと設定した固定IPが落ちてくるようになりました。

ああ、ネットワーク系にもうちょい強くなりたい。
