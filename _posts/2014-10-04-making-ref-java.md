---
layout: post
title: ref-javaというref-sourceを作っています
category: Development
tags: [Vim, Java]
---
{% include JB/setup %}

https://github.com/kamichidu/vim-ref-java

まだ触りだけしか作ってなくて、ほぼ動かないに等しい状態ですが、Java用のref-sourceを作成しています。

refというのは、[thinca](http://d.hatena.ne.jp/thinca/)さんが作成された、Vimからドキュメントを閲覧するためのVim Pluginです。
私は主にmanとperldocの閲覧に使っていました。
常々Javaのサポートが欲しかったのですが、先日[json-doclet](https://github.com/kamichidu/java-json-doclet)というJavadoc用のDocletを作成したため、refを使ってJavadocが見たい欲求を叶えようとしています。

設計中の仕組みを、備忘を兼ねてメモ。

---

1. [maven](http://maven.apache.org/)を使い、[Maven Central Repository](http://search.maven.org/)から依存するJarをダウンロード

    ここはPlugin側で自動的にダウンロードするより、-sources.jarをユーザが保持していることを前提としたほうが良くないか、と悩んでいます。
    要検討。

1. ダウンロードしたJarに対して、json-docletを使用してjson形式のJavadocファイルを作成

    json-docletで出力するファイルは、元のJarに対してディレクトリ構造を保持したままref-java用のキャッシュディレクトリに格納する。
    もしくは、ref本体のキャッシュ機能に任せてしまって、ref-javaは独自にキャッシュの管理は行わない。

    Jarの内容が変化することを考えて、キャッシュを行うのは前回キャッシュを行ってから変更のあったJarのみとする。

1. 作成されたJavadocファイルを、ref-javaから読み込み、ユーザからのクエリに従って表示

    Javadocファイルは複数になるため、読み込み終わったタイミングでJavadoc情報のマージを行う。
    描画時には、[wwwrenderer-vim](https://github.com/mattn/wwwrenderer-vim)を使用する。
    ただし、HTMLのより良い表示と速度のため、別の手段を考える必要があるかもしれない。

    ユーザからのクエリについては、使い勝手を考慮して独自の解釈を行う必要があるかも。

    このとき、クラスドキュメントの表示、メソッドやフィールドドキュメントの表示について、別途対応したい。
