---
layout: post
title: watchdogsを使ってjavaのシンタックスチェックを非同期に行う
category: Tips
tags: [Vim Java]
---
{% include JB/setup %}

[vim-watchdogs](https://github.com/osyo-manga/vim-watchdogs)とは、[vim-quickrun](https://github.com/thinca/vim-quickrun)を使って非同期にシンタックスチェックを行う便利pluginです。
詳しくはhelpか、[おしょーさんの記事](http://d.hatena.ne.jp/osyo-manga/20120924/1348473304)を参照してください。

javaのコンパイルを行うためにはclasspathの設定やら何やら、色々と煩雑な手順を踏む必要があります。
mavenなんかを使えばこの辺楽になりますが、いちいちmaven compileなんてやってたら日が暮れてしまいます。
そこでwatchdogsですよ。

以下の設定を$MYVIMRCに記述して、javaの編集中に:WatchdogsRunすれば、いい感じにqfixwindowが開くと思います。

```vim
let g:quickrun_config['java/watchdogs_checker']= {'type': 'watchdogs_checker/javac'}
let g:quickrun_config['watchdogs_checker/javac']= {
\   'command': '$JAVA_HOME/bin/javac',
\   'cmdopt': join([
\       '-Xlint:all',
\       '-d $TEMP',
\       '-sourcepath "%{javaclasspath#source_path()}"',
\       '-classpath "%{javaclasspath#classpath()}"',
\       '-deprecation',
\   ]),
\   'exec': '%c %o %S',
\   'errorformat': '%A%f:%l: %m,%-Z%p^,%+C%.%#,%-G%.%#',
\}
```

errorformatについては、:help errorformat-javacから引っぱってきました。
ちょっと不満点あるんで、そのうちどうにかするかも。

ミソは[vim-javaclasspath](https://github.com/kamichidu/vim-javaclasspath)を使ってclasspathの設定やsourcepathの設定を簡略化してるところ。
vim-javaclasspathが対応している限り、特に自分で何かしなくてもwatchdogsで非同期チェックが可能になります。

参考までに、私のqfixwindow関係の設定は、[thincaさんの記事](http://d.hatena.ne.jp/thinca/20130708/1373210009)を参考に入れています。
