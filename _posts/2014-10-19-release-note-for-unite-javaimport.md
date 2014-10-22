---
layout: post
title: unite-javaimportがメジャーバージョンアップしました
category: Announcement
tags: [Vim Java]
---
{% include JB/setup %}

今までバージョン0.2.xだった[unite-javaimport](https://github.com/kamichidu/vim-unite-javaimport)が、この度0.3.xになりました。
内部構造が大幅に変更されているため、0.2.xと比べて

1. メモリ使用量の削減

    0.2.xでは500M程使ってしまっていましたが、大きく改善されています。

1. 速度の向上

    0.2.xでは内部的にローカルサーバを立ち上げて、javaimportが起動されるごとに計算を行っていました。
    0.3.xでは主にキャッシュ周りのファイル構成を見直し、一度起動した後は主にキャッシュから候補を取得するため、速度が向上しました。

以上の点が改善されています。

また、新規機能の追加として、

1. static import機能の追加

    以下のunite-sourceが追加されています。

    1. javaimport/field

    1. javaimport/method

    また同時に、javaimport/classのアクションにexpandが追加されています。
    javaimport/classで選択したクラスをjavaimport/fieldおよびjavaimport/methodに渡して起動することができます。

1. [ctrlp](https://github.com/ctrlpvim/ctrlp.vim)をサポートしました

    以下のコマンドが追加されています。
    ただしhelpはまだないので、実験的なものとなります。

    1. CtrlPJavaImportClass

    1. CtrlPJavaImportField

    1. CtrlPJavaImportMethod

以上の機能が追加されています。

また今後の開発についてですが、unite-sourceを主として開発するのではなく、メインのUIを別に作成することを考えています。
uniteのサポートは今後とも継続するつもりですが、あくまでjavaimportのサポートする1つのインタフェイスという位置付けになります。
今回ctrlpのサポートを追加したのも、この変更に対する伏線の意図があります。
なお、本格的な変更は次回のメジャーバージョンアップで行う予定です。

不具合や要望があれば、[githubのissues](https://github.com/kamichidu/vim-unite-javaimport/issues)までお願いいたします。
