--开启外键
PRAGMA FOREIGN_KEYS=ON;
CREATE TABLE smoe_comments (    "coid" INTEGER NOT NULL PRIMARY KEY,
                                "cid" int(10) NOT NULL default 0 ,
                                "created" int(10) NOT NULL default 0 ,
                                "author" varchar(150) NOT NULL default '',
                                "authorId" int(10) default 0 ,
                                "mail" varchar(150) NOT NULL default '',
                                "url" varchar(255) default NULL ,
                                "ip" varchar(64) NOT NULL default '',
                                "agent" varchar(511) NOT NULL default '',
                                "text" text ,
                                "status" varchar(16) default 'approved' ,
                                "parent" int(10) default 0 );

CREATE INDEX smoe_comments_cid ON smoe_comments ("cid");
CREATE INDEX smoe_comments_created ON smoe_comments ("created");

-- 用来保存文章和独立页面
CREATE TABLE smoe_contents (    "cid" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
                                "mid" INTEGER NOT NULL default 0 ,
                                "title" varchar(200) NOT NULL default '',
                                "slug" varchar(200) NOT NULL default '' ,
                                "created" int(10) NOT NULL default 0 ,
                                "text" text ,
                                "type" varchar(16) default 'post' ,
                                "status" varchar(16) default 'publish' ,
                                "allowComment" int(1) default 0 ,
                                "allowFeed" int(1) default 0 ,
                                "views" int(10) default 0 ,
                                "likes" int(10) default 0 ,
                                "coverList" varchar(32) default '',
                                "musicList" varchar(32) default '' ,
                                FOREIGN KEY("cid") REFERENCES "smoe_comments"("cid") ON DELETE CASCADE ON UPDATE CASCADE );

CREATE UNIQUE INDEX smoe_contents_slug ON smoe_contents ("slug");
CREATE INDEX smoe_contents_created ON smoe_contents ("created");

CREATE TABLE smoe_metas (    "mid" INTEGER NOT NULL PRIMARY KEY,
                             "name" varchar(150) default NULL ,
                             "slug" varchar(150) default NULL ,
                             "type" varchar(32) NOT NULL ,
                             "description" varchar(150) default NULL ,
                             "count" int(10) default 0 ,
                             "order" int(10) default 0 ,
                             "parent" int(10) default 0);

CREATE INDEX smoe_metas_slug ON smoe_metas ("slug");

CREATE TABLE smoe_options (    "name" varchar(32) NOT NULL ,
                               "user" int(10) NOT NULL default 0 ,
                               "value" text );

CREATE UNIQUE INDEX smoe_options_name_user ON smoe_options ("name", "user");

CREATE TABLE smoe_users (    "uid" INTEGER NOT NULL PRIMARY KEY,
                             "name" varchar(32) NOT NULL default '' ,
                             "password" varchar(64) NOT NULL default '' ,
                             "mail" varchar(150) NOT NULL default '' ,
                             "url" varchar(150) NOT NULL default '' ,
                             "screenName" varchar(32) NOT NULL default '' ,
                             "created" int(10) NOT NULL default 0 ,
                             "activated" int(10) NOT NULL default 0 ,
                             "logged" int(10) NOT NULL default 0 ,
                             "group" varchar(16) NOT NULL default 'visitor' ,
                             "authCode" varchar(64) NOT NULL default '');

CREATE UNIQUE INDEX smoe_users_name ON smoe_users ("name");
CREATE UNIQUE INDEX smoe_users_mail ON smoe_users ("mail");

CREATE TABLE "smoe_access_log" (
                                   "id"	INTEGER NOT NULL,
                                   "ua"	varchar(512) NOT NULL,
                                   "url"	varchar(255) NOT NULL,
                                   "path"	varchar(255) NOT NULL,
                                   "ip"	INTEGER NOT NULL,
                                   "referer"	varchar(255) NOT NULL,
                                   "time"	INTEGER NOT NULL,
                                   PRIMARY KEY("id" AUTOINCREMENT)