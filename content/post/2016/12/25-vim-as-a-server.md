+++
categories = ["Article"]
tags = ["Vim"]
date = "2016-12-25T20:03:08+09:00"
title = "Vimをサーバとして稼動させる"

+++

さて、みなさん大好きテキストエディタのVimですが、実はサーバとして稼動させることができることをご存じでしょうか。\\
[参考](http://mattn.kaoriya.net/software/vim/20120127204609.htm)

Vimにも比較的最近、ソケット通信を扱うためのChannel機能というものが入りました。
そこで、2016年最後のネタとして、VimのChannel機能を使い、Vim Scriptでecho serverを実装してみようと思います。

まず、Vim本体に以下のパッチを当ててください。
このパッチは、私が開発に利用しているArch Linuxでしか動作確認を取っておらず、面倒だったのでWinSock対応はしておりません。
また、既存のChannel機能に無理矢理合わせるために、開いたソケットを閉じるなどの処理を一切省いてしまいました。
片手落ち感は否めませんが、ここで重要なのは『Vim Scriptでecho serverを実装したい』という願望を叶えることと割り切りました。

[パッチはこちら](https://gist.github.com/anonymous/558d625b8be4d8cf7a913381c851fec6)\\
※gistaの設定をミスって、anonymousとして上がってしまいました....

パッチを当てたVimをmakeすると、src/vimが出来あがります。
では早速、Vim Scriptを書いていきましょう。

```vim
let ch= ch_listen('0.0.0.0:12345', {
\   'mode': 'raw',
\   'waittime': -1,
\})
while 1
    let st= ch_status(ch)
    echomsg st
    if st ==# 'fail' || st ==# 'closed'
        break
    endif

    let msg= ch_readraw(ch)
    echomsg 'Received: ' . msg
    if !empty(msg)
        call ch_sendraw(ch, msg)
    endif
endwhile
```

さて、このVim Scriptを `echo-server.vim` として保存し、 `./src/vim -N -u NONE -U NONE --noplugin -S echo-server.vim` を実行しましょう。
おめでとうございます、 `netstat -tanlp` すると、12345ポートでVimがListenしている勇姿が見えるでしょうか。

動作確認に入ります。
`telnet 127.0.0.1 12345` を実行し、何か打ってみてください。
Vimがechoを返してくるはずです。

![動作の様子](/images/2016-12-25-demo.gif)

以上、ネタでした。
