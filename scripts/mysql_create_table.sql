drop table if exists  s_user_miner;
CREATE TABLE s_user_miner
(
  app_id varchar(255) ,
  client_ip varchar(255),
  client_time varchar(255),
  server_time varchar(255),
  sdk_version varchar(255),
  resolution varchar(255),
  os varchar(255),
  event_id varchar(255),
  page varchar(255),
  arg1 varchar(255),
  arg2 varchar(255),
  arg3 varchar(255),
  args longtext
);