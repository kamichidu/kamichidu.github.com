+++
title = "VimConf2014で発表してきた"
Description = ""
Tags = ["Vim"]
Categories = ["Vim"]
date = "2014-11-02T00:00:00+09:00"

+++

昨日になりますが、[VimConf2014](http://vimconf.vim-jp.org/2014/)に参加し、発表をしてきました。
人生初の勉強会での発表だったので、非常に緊張しましたが、無事発表を終えることができました。
終わった後になって、やっぱり少しは発表練習するべきだったか、とか、スライドもう少し凝れば良かったとか、反省点は色々出てきました。
感想などをつらつらと。

... Now Loading

当初は10:00に会場へ行くつもりだったんですが、体温がなかなか上がらなかったため家を出られず、11:30に会場へ着きました。
着いて早々、derisさんとraaさんと共にラーメン凪へ。
黒マー油を頼みましたが、イカスミで真っ黒に染まったラーメンがとても美味しかったです。

... Now Loading

最初はkoronさんによる[Identity of the Vim](http://koron.github.io/vimconf-2014-koron/)でした。
かなり熱いお話で、非常に頷ける内容が多かったです。
Vimを始めてから色々な言語に触る機会が増え、同時に技術的に色々なものに触れる機会が増えた経験があったため、共感する部分が多かったです。
JavaだけいじりたいならEclipse使えっていう内容も、非常に同意。
無理にVim使う必要なんかない。
私の場合、Vimの操作が手に馴染んでしまい、Eclipseは便利だと思いつつも生産性はVimのほうが高いため、Vimでやってます。
でもやっぱりEclipseの機能が、Vimでも使いたいため、Java用のpluginを色々開発してたりする訳ですが。

続いてujihisaさんによる[PM2](https://docs.google.com/presentation/d/1u5A7F3Kd4XwJlIUQZAVmrwWfLcoLf9NURtqAEafi_oo/edit#slide=id.p)でした。
Civ5やろうっていうお話でした。
PM2に対する喜びの声に対して、私の登場率が高かったです。というか100%？

休憩を挟み、Lindaさんによる[f](https://speakerdeck.com/rhysd/vimconf-2014-f)でした。
f拡張プラグインについてのお話。
私はvim標準のfマッピングで満足しているので、あまり使おうとは思いませんが、何かのときには参考にしたいなーと。

で、私による[Hey, Java! Vim is coming.](https://docs.google.com/presentation/d/1zaPy82NJ6A3Iw1llKqU-lX88AJNt1EKy5O15nOp085c/edit#slide=id.p)でした。
手持ちのChromebookだと、プロジェクタを認識せずにLindaさんにPCをお借りしての発表になりました。
Lindaさんありがとうございました。
発表内容自体については、喋りたい内容の1割くらいは喋れたように思えます。
ですがもう少し色々と内容を練ればよかったなーと反省しています。
結局やっぱり、Java編集するならEclipseが便利で、InteliJやNetBeans等のIDEがやはり王道な訳です。
で、Vim使ってやりたいと思ったときにpluginを作成している訳ですが、pluginを作成するときにJavaならでは？の敵が待ち受けています。
じゃあどうするの、とか、他にも色々掘り進んだ話をしたいですので、次回のVimConfではJavaとVimに関する話をしたいと思いました。

cohamaさんによる[auto closing parenthesis](http://www.slideshare.net/cohama/auto-closing-parenthesis-vim-conf2014-41290298)では、cohamaさん作のleximaや他の括弧自動入力系pluginのお話でした。
個人的には、括弧自動入力については思考と指を乱されるため、好んでは使わないです。
ですがEclipseを使う同僚の話を聞くと、括弧自動入力は便利で欠かせないと言う人も多いです。
leximaは"."リピートが可能なので、vimらしいpluginではないかと思います。

その後、derisさんによる[怖くないマクロ入門](http://www.slideshare.net/deris0126/vimconf4)でした。
懇親会のお寿司を受け取りに出ていたりで、あまり聞けてないです。
マクロ漁船に乗りたくなってしまいそうなお話だった記憶です。

... Now Loading

休憩を挟んで後、thincaさんによる[Test for Vim script](https://gist.github.com/thinca/2cf4ae0df88a99423c9d)でした。
ライブコーディングが非常に便利でした。
最近はvim scriptのテストにはthemisを使わせてもらっています。

続いてShougoさんによる[Let's talk about neovim](http://www.slideshare.net/Shougo/lets-talk-about-neovim)。
私はあまりneovimを追いかけていないのですが、neovimの話題は楽しみました。
テキストエディタってなんぞや、という哲学的な問いがありました。

... Now Loading

休憩を挟み、LTの嵐。

supermomongaさんによる[かなりすごい発表（かなり）](http://www.slideshare.net/supermomonga/super-cool-presentation-at-vimconf2014)では、音を有効に使おうというお話でした。
ももんがさん相変らずおもろい人です。

pebble8888さんによる[XVim with MacVim and smartgrep](http://www.slideshare.net/pebble8888/using-xvim-with-macvim)では、XCodeでのVimEmuのお話、それとsmartgrepというコメントを無視してgrepしてくれる便利ツールのお話でした。
VimEmuは非常に便利だと思うのですが、触っているうちにやっぱりVimじゃないんだよねーと思ってしまう罪作りなやつだと思います。

haya14busaさんによる[/-improved](https://docs.google.com/presentation/d/1ie2VCSt9onXmoY3v_zxJdMjYJSbAelVR-QExdUQK-Tw/pub?start=false&loop=false&delayms=3000&slide=id.g4e7add63c_05)では"/"を有効に使おうというお話。
hayabusaさんはVim始めてまだ2年？なのにすごいと思いました。
勝てる気がしないです。

Kuniwakさんによる[vim script初心者に使ってもらいたい、転ばぬ先の杖「Vint」](https://speakerdeck.com/orgachem/zhuan-banuxian-falsezhang-vint)では、vim script用のlintツールの紹介でした。
Kuniwakさんは今回の開催にあたり、非常に尽力してくださっていてすごい人でした。
良さげなので今度使ってみようと思います、Vint。

raa0121さんによる[Jenkins + vimenvで最新のVimを使おう！](http://www.slideshare.net/raa0121/jenkinsvimenv-vim-vimconf2014)は、VimmerはpatchごとにVimをビルドして持っているらしいと聞きました。
驚愕しました。

... Now Loading

発表が終われば、懇親会になりました。
懇親会では何人かの方とお話をすることができました。
あと軽食が美味しかったです。
軽食らしく、ピザLサイズ2枚半、寿司1パック食べました。
THE 軽食ですね。
お腹いっぱいで幸せでした。

VimConf開催にあたって主催してくださったみなさま、本当にありがとうございました。
次回があれば、また発表をしたい気持ちです。
例えばJavaとVimについて。

... Now Loading

... Sorry, failed to initialize java virtual machine. Try it again.