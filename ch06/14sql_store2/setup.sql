/*
DROP TABLE hoge CASCADE: 削除するテーブルに依存するオブジェクトを自動的に削除する
下記では comments テーブルが posts テーブルに依存している
*/
drop table posts cascade;
drop table comments;

create table posts (
  id      serial primary key,
  content text,
  author  varchar(255)
);

create table comments (
  id      serial primary key,
  content text,
  author  varchar(255),
  post_id integer references posts(id)
);