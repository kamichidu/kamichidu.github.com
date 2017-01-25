+++
title = "JSON Web Token（JWT）のClaimについて"
tags = [
    "JWT"
]
categories = [
    "Article"
]
date = "2017-01-24T19:19:50+09:00"

+++

REST APIなどを開発すると、避けては通れないものに認証があります。
最近はOAuth2 Providerを実装することが多いのですが、発行するアクセストークンにJSON Web Token（JWT）を利用しています。
JWTは名前の通り、[RFC7519](https://tools.ietf.org/html/rfc7519)によって定義されたスキーマを持つJSONです。
JSONの各キーとして、RFCで定義されている標準的なキーと値のペア（Claim）を取ることにより、標準的な取り扱いが可能になります。

本記事では、JWTのClaimについて、OAuth2 Providerでのアクセストークンを発行する立場から、備忘録的にまとめたいと思います。

なお、JWTには[JOSE Header](https://tools.ietf.org/html/rfc7519#section-5)のような、Claimとは別の仕様も含まれていますが、本記事ではClaimのみを対象とします。


アクセストークンとしてJWTを利用することの利点
------------------------------------------------------------------------------------------------------------------------
JWTは単なるJSONのため、アクセストークンとして用いることにより様々な情報を含めることができます。
それにより、認証サーバとアプリケーションサーバを分割しているようなアーキテクチャであれば、認証サーバの負荷を軽減することに寄与します。
[こちらの記事がよくまとまっており、わかりやすいです。](http://d.hatena.ne.jp/ritou/20140927/1411811648)

また、本記事の範囲からは離れますが、JWTと関連する[JSON Web Encryption（JWE）](https://tools.ietf.org/html/rfc7516)や
[JSON Web Signature（JWS）](https://tools.ietf.org/html/rfc7515)などと併用することにより、
規格化された方法のみで一般的なREST APIの認証で必要となる要件を満たすことができます。
規格化されているため、技術者間の共通言語としても利用できますし、サービスごとに独自仕様を策定するという無駄な手間を省くことができます。


Claimの種類
------------------------------------------------------------------------------------------------------------------------
JWTには、[3種類のClaimが定義されています。](https://tools.ietf.org/html/rfc7519#section-4)
ここでは、それぞれのClaimについて簡単に説明します。

### Registered Claim Names

IANAの"JSON Web Token Claims"に登録された、一般的な用途で利用されることを想定されたClaimです。
基本的にこれらのClaimの利用は任意ですが、JWTを発行する際には可能な限り含めておいたほうが良いと思います。

* "iss" (Issuer) (Optional)

    JWTの発行者を意味します。
    値として文字列かURIを取ることができ、アプリケーション名やドメイン名が入ると思います。
    値は大文字/小文字を区別して扱われます。

* "sub" (Subject) (Optional)

    JWTの用途を意味します。
    値として文字列かURIを取ることができます。
    値は大文字/小文字が区別され、同じIssuer内でユニーク、または全世界でユニークである必要があります。

* "aud" (Audience) (Optional)

    JWTの想定利用者を意味します。
    値として文字列かURI、またはそれらの配列を取ることができます。
    値は大文字/小文字が区別されます。

    発行する側はJWTの発行要求をしてきた相手を識別する文字列やURIを入れ込み、
    発行された側はAudience Claimが存在する場合は自分向けに発行されたJWTなのかどうかを検証することに用います。

* "exp" (Expiration Time) (Optional)

    JWTが失効する日時を意味します。
    値としては、数値表現された日付となります。（例: "12345"）

    失効する日時を表すUNIX時刻などを入れるのが良いと思います。

* "nbf" (Not Before) (Optional)

    JWTが有効になる日時を意味します。
    取り得る値は"exp"と同様です。

* "iat" (Issued At) (Optional)

    JWTが発行された日時を意味します。
    取り得る値は"exp"や"nbf"と同様です。

* "jti" (JWT ID) (Optional)

    JWTのユニーク性を担保するID値を意味します。
    "jti"が異なるJWTは、全く別のJWTとして扱われます。
    値は大文字/小文字が区別されます。

    "jti"の存在は、同じJWTを使い回すことを抑制することを目的にしています。
    JWTの発行ごとに、UUIDでも入れておくのが良いのでしょうか。


### Public Claim Names

[IANA "JSON Web Token Claims"](https://www.iana.org/assignments/jwt/jwt.xhtml)に登録されているClaimを言います。
Registered Claimは基本的にIANAに登録されており、Registered ClaimはPublic Claimのサブセットになっています。
Public ClaimはIANAに登録されたClaimを言い、主に汎用的な用途のClaimを衝突する可能性を（仕様的な意味で）排除する目的で利用されます。

一般的なサービスを開発するとかいう場合には、恐らくこの種類のClaimを使うことはないでしょう。


### Private Claim Names

このClaimは、JWTのIssuerとAudienceの間で取り決めされた仕様に応じて何でも定義することができ、
Private ClaimはRegistered ClaimやPublic Claimで予約された以外の名前を使うことができます。
JWTをアクセストークンとして利用する場合は、大抵Private Claimにガシガシ値を突っ込んでいくことになります。

基本的にはどのような値も入れることはできますが、Claim名の衝突を避けるため、プレフィクスや企業のドメイン名をつけるなど工夫をしたほうが良いと思います。


アクセストークンの構築例
------------------------------------------------------------------------------------------------------------------------
ここでは今まで説明した内容を踏まえて、実際のアクセストークンがどのようになるかを説明します。

### 想定ケース

想定するサービスは、ログインユーザの情報として以下を持つとします。

1. ユーザID
1. ユーザ名
1. メールアドレス
1. 年齢
1. 性別

また、サービスは当然独自のドメイン（example.com）を持っているものとし、
認証系と機能系はそれぞれサブドメイン（idp.example.comとapi.example.com）で構成されているとします。

### アクセストークン設計

アクセストークンとしてJWTを発行するのは認証系であり、発行されたJWTを利用するのは機能系となります。
場合によっては、認証系にはユーザ情報を検索する等の機能が含まれている場合もありますので、それも加味します。
つまりJWTの発行者はidp.example.comとなり、想定利用者はapi.example.comおよびidp.example.comとなります。

発行するJWTの用途としては、アクセストークンです。
セキュリティを考えて、発行から1時間で失効するのが良いでしょう。
また、発行したJWTが有効になる時刻は、各サーバで時刻同期が仮にずれていた場合を考えて発行時刻の5秒前としましょう。

以上をまとめると、Registered ClaimのみでJWTを構成すると以下のようになります。

```
{
  "jti": "92f46647-90a2-4174-bca9-27d7f69a8fb7",
  "iss": "https://idp.example.com/",
  "sub": "AccessToken",
  "aud": [
    "https://api.example.com/",
    "https://idp.example.com/"
  ],
  "exp": 1485320878,
  "nbf": 1485317273,
  "iat": 1485317278
}
```

続いてサービスに固有のClaimを考えます。
Private Claimとして、ログインユーザの情報をすべて含めてしまいましょう。
名前が衝突するのを防止するため、idp.example.comをプレフィクスとして構成することにし、
ユーザ情報の各属性を、そのまま英語に直したものをサフィックスとしてしまいましょう。

例として以下のようなユーザ情報を考えます。
（golangのstructで示します。）

```
user := &struct{
	Id    string
	Name  string
	Email string
	Age   int
	Sex   string
}{
	Id:    "72f5fbcc-75bb-4393-b2dd-76e74ad0fd87",
	Name:  "pico jiro",
	Email: "picojiro@example.com",
	Age:   30,
	Sex:   "male",
}
```

このユーザ情報からPrivate Claimを構成すると以下のようになります。

```
{
  "https://idp.example.com/claim-types/user-id": "72f5fbcc-75bb-4393-b2dd-76e74ad0fd87",
  "https://idp.example.com/claim-types/user-name": "pico jiro",
  "https://idp.example.com/claim-types/user-email": "picojiro@example.com",
  "https://idp.example.com/claim-types/user-age": 30,
  "https://idp.example.com/claim-types/user-sex": "male"
}
```

Registered ClaimとPrivate Claimを合わせて、最終的に以下のようなJWTが得られました。

```
{
  "jti": "37852e99-ab54-460c-92b2-18231d3ba823",
  "iss": "https://idp.example.com/",
  "sub": "AccessToken",
  "aud": [
    "https://api.example.com/",
    "https://idp.example.com/"
  ],
  "exp": 1485322113,
  "nbf": 1485318508,
  "iat": 1485318513,
  "https://idp.example.com/claim-types/user-id": "72f5fbcc-75bb-4393-b2dd-76e74ad0fd87",
  "https://idp.example.com/claim-types/user-name": "pico jiro",
  "https://idp.example.com/claim-types/user-email": "picojiro@example.com",
  "https://idp.example.com/claim-types/user-age": 30,
  "https://idp.example.com/claim-types/user-sex": "male"
}
```

まとめ
------------------------------------------------------------------------------------------------------------------------
JWTのClaimについて、REST API等のアクセストークンに適用するという観点から仕様をまとめました。
またアクセストークンを構築する例を通して、実際にアクセストークンとしてJWTを適用する方法について説明しました。

JWTはJWEやJWSと組み合わせて、セキュアなアクセストークンの仕組みを標準的な方法で構築することができるため、非常に有用だと思います。

今回アクセストークンの例を出力するために書いたgolangのコードを貼っておきますので、参考程度にどうぞ。
なおJWTの確認には、[jwt.io](https://jwt.io)が便利です。

```
package main

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/satori/go.uuid"
)

func main() {
	user := &struct {
		Id    string
		Name  string
		Email string
		Age   int
		Sex   string
	}{
		Id:    "72f5fbcc-75bb-4393-b2dd-76e74ad0fd87",
		Name:  "pico jiro",
		Email: "picojiro@example.com",
		Age:   30,
		Sex:   "male",
	}
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodNone, jwt.MapClaims{
		"iss": "https://idp.example.com/",
		"sub": "AccessToken",
		"aud": []string{"https://api.example.com/", "https://idp.example.com/"},
		"exp": now.Add(1 * time.Hour).Unix(),
		"nbf": now.Add(-5 * time.Second).Unix(),
		"iat": now.Unix(),
		"jti": uuid.NewV4().String(),
		"https://idp.example.com/claim-types/user-id":    user.Id,
		"https://idp.example.com/claim-types/user-name":  user.Name,
		"https://idp.example.com/claim-types/user-email": user.Email,
		"https://idp.example.com/claim-types/user-age":   user.Age,
		"https://idp.example.com/claim-types/user-sex":   user.Sex,
	})
	s, _ := token.SigningString()
	println(s)
}
```
