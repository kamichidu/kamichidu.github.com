---
layout: post
title: json-docletの仕様メモ
category: Development
tags: [Java]
---
{% include JB/setup %}

[json-doclet](https://github.com/kamichidu/java-json-doclet)はJavadocツール用のDoclet。
標準Docletと違い、HTMLではなくJSONを吐き出します。

JSONでドキュメントを吐き出すため、アプリケーションからJavadocを利用することがHTMLに比べて容易です。
以下、その仕様をまとめます。
*開発中のため、予告なく仕様は変更になる可能性があります。*

---

### 許容する引数

* -ofile (必須)

    出力するファイル名を指定する。

* -append

    -ofileで指定されたファイルに追記する。

* -pretty

    JSONを出力時、整形された状態で出力する。


### 出力されるJSONオブジェクト

#### <a name="root-object">ルートオブジェクト</a>

* classes - [クラスオブジェクト](#class-object)の配列

#### <a name="class-object">クラスオブジェクト</a>

* name - クラスの正規名

* interfaces - クラスが実装しているインタフェイスの正規名の配列

* superclass - 親クラスの正規名

* comment_text - クラスに付与されたJavadocコメント

* since - @sinceタグの情報

* see - @seeタグの情報を持つ配列

* constructors - [コンストラクタオブジェクト](#constructor-object)の配列

* fields - [フィールドオブジェクト](#field-object)の配列

* methods - [メソッドオブジェクト](#method-object)の配列

#### <a name="constructor-object">コンストラクタオブジェクト</a>

* name - コンストラクタメソッドの正規名

* comment_text - メソッドに付与されたJavadocコメント

* parameters - [メソッドパラメータオブジェクト](#method-parameter-object)の配列

* throws - [例外オブジェクト](#exception-object)の配列

#### <a name="exception-object">例外オブジェクト</a>

* name - 例外クラスの正規名

* comment_text - @throwsタグに付与されたJavadocコメント

#### <a name="method-parameter-object">メソッドパラメータオブジェクト</a>

* name - 変数名

* comment_text - @paramsタグに付与されたJavadocコメント

* type - 型の正規名

#### <a name="field-object">フィールドオブジェクト</a>

* name - フィールド名

* comment_text - フィールドに付与されたJavadocコメント

* type - 型の正規名

#### <a name="method-object">メソッドオブジェクト</a>

* name - メソッドの正規名

* comment_text - メソッドに付与されたJavadocコメント

* parameters - [メソッドパラメータオブジェクト](#method-parameter-object)の配列

* throws - [例外オブジェクト](#exception-object)の配列

* return_type - メソッドから返却される型の正規名
