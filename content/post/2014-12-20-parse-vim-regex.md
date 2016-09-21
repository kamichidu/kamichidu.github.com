+++
title = "Vimの正規表現を切り刻む"
Description = ""
Tags = ["Vim"]
Categories = ["Tips"]
date = "2014-12-20T00:00:00+09:00"

+++

この記事はVim Advent Calendar 2014の20日目の記事です。

19日目はthleapさんによる[マークアップや文章の編集をちょっと楽にする](http://chroma.hatenablog.com/entry/2014/12/19/224359)でした。

導入
---
みなさん、正規表現、使ってますか？
私は毎日使っています。
特にVimの正規表現、かなり変態で良さがありますよね。

ところでPerlには、[Regexp::Assemble](https://metacpan.org/pod/Regexp::Assemble)という非常に便利なモジュールがあります。
この子は何をする子かというと、自動で最適化された正規表現を生成してくれます。
正規表現のマッチングでは、うまくトライ木を作ることができれば、かなりマッチング速度が向上しますので、自動でそれをやってくれるモジュールと言えばわかりやすいでしょうか。
(いわゆるtrie optimization)

例えば、"public"、"protected"、"private"という3つの単語についてマッチする正規表現は、
純粋に`public|protected|private`と書けば良いのですが、これだと遅いので`(?:p(?:ublic|r(?:otected|ivate)))`と書いたほうが速いので、こう書きましょうよってことです。

ここから本題。
Regexp::Assembleのように、自動で正規表現を最適化するためには、まず正規表現を切り刻まなければなりません。
この記事では、正規表現を切り刻んで遊んでみることにします。


下準備
---
正規表現を切り刻む、Vim scriptコードを用意します。
以下を手元で実行し、`RegexpTokenize`コマンドを使えるようにしてください。

(簡単のため、一部の正規表現にしか対応させてません。)

```vim
function! s:tokenize(expr)
    let [chars, i]= [split(a:expr, '\zs'), 0]
    let tokens= []
    while i < len(chars)
        let ch= chars[i]

        if ch ==# '\'
            if !(i < len(chars))
                break
            endif

            let i+= 1
            let ch= chars[i]

            if ch ==# '+'
                let tokens+= ['\+']
            elseif ch ==# '=' || ch ==# '?'
                let tokens+= ['\' . ch]
            elseif ch ==# '('
                let tokens+= ['\(']
            elseif ch ==# ')'
                let tokens+= ['\)']
            elseif ch ==? 'm' || ch ==? 'v'
                let tokens+= ['\' . ch]
            elseif ch ==? 'c'
                let tokens+= ['\' . ch]
            elseif ch ==? 'i'
                let tokens+= ['\' . ch]
            elseif ch ==? 'k'
                let tokens+= ['\' . ch]
            elseif ch ==? 'f'
                let tokens+= ['\' . ch]
            elseif ch ==? 'p'
                let tokens+= ['\' . ch]
            elseif ch ==? 's'
                let tokens+= ['\' . ch]
            elseif ch ==? 'd'
                let tokens+= ['\' . ch]
            elseif ch ==? 'x'
                let tokens+= ['\' . ch]
            elseif ch ==? 'o'
                let tokens+= ['\' . ch]
            elseif ch ==? 'w'
                let tokens+= ['\' . ch]
            elseif ch ==? 'h'
                let tokens+= ['\' . ch]
            elseif ch ==? 'a'
                let tokens+= ['\' . ch]
            elseif ch ==? 'l'
                let tokens+= ['\' . ch]
            elseif ch ==? 'u'
                let tokens+= ['\' . ch]
            else
                throw printf("Unsupported regex: `%s'", a:expr)
            endif
        else
            let tokens+= [ch]
        endif

        let i+= 1
    endwhile
    return tokens
endfunction

command!
\   -nargs=1
\   RegexpTokenize
\   echo s:tokenize(<q-args>)
```


切り刻め！
---
以下のコマンドを実行してみましょう。
切り刻まれ、正規表現単位に分割された文字列が表示されます。

```vim
" => ['h', 'o', 'g', 'e', '\(', '\w', '\+', '\)', '\?']
RegexpTokenize hoge\(\w\+\)\?
" => ['p', 'u', 'b', 'l', 'i', 'c']
RegexpTokenize public
" => ['\d', '\+']
RegexpTokenize \d\+
" => ['\d', '\+', '\(', ',', '\d', '\+', '\)', '\=']
RegexpTokenize \d\+\(,\d\+\)\=
```


最後に
---
この記事で書いたようなことをするライブラリとして、[vim-regexp-assemble](https://github.com/kamichidu/vim-regexp-assemble)というものを作成しております。
切り刻みたいだけなら、Vital.Regexp.Lexerを使えばそれで終わってしまいますし、この記事に記載した一部の正規表現だけでなく、Vimの正規表現すべてに対して対応しています。

(help整備したら[vital本家](https://github.com/vim-jp/vital.vim)に入れてもらうかもしれませんが、ちょっと先の話になりそう。)

みなさんもっと正規表現で遊びましょう。
ということで、20日目の記事を終わりたいと思います。
明日はmitsuseさんです。
