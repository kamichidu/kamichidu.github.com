---
layout: page
title: なるはやで いい感じの 動作確認
tagline: Supporting tagline
---
{% include JB/setup %}

<ul class="posts">
  {% for post in site.posts %}
    <li>
        <p><span>{{post.date | date_to_string}}</span> &raquo; <a href="{{BASE_PATH}}{{post.url}}">{{post.title}}</a></p>
        <p>{{post.content}}</p>
    </li>
  {% endfor %}
</ul>
