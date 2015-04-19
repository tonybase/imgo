/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50538
 Source Host           : localhost
 Source Database       : im

 Target Server Type    : MySQL
 Target Server Version : 50538
 File Encoding         : utf-8

 Date: 04/14/2015 14:19:17 PM
*/

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `im_conversation`
-- ----------------------------
DROP TABLE IF EXISTS `im_conversation`;
CREATE TABLE `im_conversation` (
  `id`          VARCHAR(255) NOT NULL
  COMMENT '会话的唯一标识',
  `creater`     VARCHAR(255) NOT NULL
  COMMENT '创建人',
  `create_date` DATE         NOT NULL
  COMMENT '创建日期',
  `receiver`    VARCHAR(255) NOT NULL
  COMMENT '接收人(可以是用户，群，讨论组)',
  `type`        CHAR(1)      NOT NULL DEFAULT '0'
  COMMENT '会话类型 0用户会话，1群会话，2讨论组会话',
  PRIMARY KEY (`id`)
)
  ENGINE = MyISAM
  DEFAULT CHARSET = utf8;

-- ----------------------------
--  Table structure for `im_crowd`
-- ----------------------------
DROP TABLE IF EXISTS `im_crowd`;
CREATE TABLE `im_crowd` (
  `id`          VARCHAR(255) NOT NULL
  COMMENT '群的唯一标识',
  `name`        VARCHAR(255) NOT NULL
  COMMENT '群名称',
  `creater`     VARCHAR(255) NOT NULL
  COMMENT '创建者',
  `create_date` DATE         NOT NULL
  COMMENT '创建日期',
  `user_num`    INT(11)      NOT NULL DEFAULT '100'
  COMMENT '群允许的用户数量',
  PRIMARY KEY (`id`)
)
  ENGINE = MyISAM
  DEFAULT CHARSET = utf8;

-- ----------------------------
--  Table structure for `im_discussion`
-- ----------------------------
DROP TABLE IF EXISTS `im_discussion`;
CREATE TABLE `im_discussion` (
  `id`          VARCHAR(255) NOT NULL
  COMMENT '讨论组的唯一标识',
  `name`        VARCHAR(255) NOT NULL
  COMMENT '讨论组名',
  `creater`     VARCHAR(255) NOT NULL
  COMMENT '创建者',
  `create_date` DATE         NOT NULL
  COMMENT '创建日期',
  PRIMARY KEY (`id`)
)
  ENGINE = MyISAM
  DEFAULT CHARSET = utf8;

-- ----------------------------
--  Table structure for `im_group`
-- ----------------------------
DROP TABLE IF EXISTS `im_group`;
CREATE TABLE `im_group` (
  `id`          VARCHAR(255) NOT NULL
  COMMENT '唯一标识',
  `name`        VARCHAR(255) NOT NULL
  COMMENT '分组名',
  `creater`     VARCHAR(255) DEFAULT NULL
  COMMENT '创建人',
  `create_date` DATE         DEFAULT NULL
  COMMENT '创建日期',
  PRIMARY KEY (`id`)
)
  ENGINE = MyISAM
  DEFAULT CHARSET = utf8;

-- ----------------------------
--  Table structure for `im_login`
-- ----------------------------
DROP TABLE IF EXISTS `im_login`;
CREATE TABLE `im_login` (
  `id`         VARCHAR(32) NOT NULL
  COMMENT '登录记录唯一标识',
  `user_id`    VARCHAR(32) NOT NULL
  COMMENT '用户ID',
  `token`      VARCHAR(32) NOT NULL
  COMMENT '用户token',
  `login_date` DATETIME    NOT NULL
  COMMENT '登录日期',
  `login_ip`   VARCHAR(20) NOT NULL
  COMMENT '用户登录IP',
  PRIMARY KEY (`id`)
)
  ENGINE = MyISAM
  DEFAULT CHARSET = utf8;

-- ----------------------------
--  Records of `im_login`
-- ----------------------------
BEGIN;
INSERT INTO `im_login` VALUES
  ('bc795b095b75432f826b76732ab02bd5', '23232', 'e3c0c973-e7a1-49d4-a5c1-144c29e5', '0000-00-00 00:00:00',
   '[::1]:58139'),
  ('8f21bf0e8a294572bed41b488c7b1fd1', '23232', 'c59452ff-adb0-479a-b831-ccc1807e', '2015-04-14 00:00:00', '[::1]:58276'),
  ('4d2b9e0959a044b8b8d00949d8937318', '23232', '7b91f573-e686-4bf5-a68d-23628e2b', '2015-04-14 00:00:00', '[::1]:58417'),
  ('70f559dcea2a4fcaac7162eab6a60412', '23232', '61249db5-6497-4d24-a87f-ba22e552', '2015-04-14 00:00:00', '[::1]:58428'),
  ('c99c49c341e1485691472757b849698a', '23232', '00ede79c-5ed3-4160-b5f1-f02ba649', '2015-04-14 00:00:00', '127.0.0.1:58436'),
  ('87c61cd0863046bdafd3d3150a068279', '23232', 'd643d29c-5718-472f-9080-cebb15ed', '2015-04-14 00:00:00', '127.0.0.1'),
  ('b3035d82ba14404c8530a6fd29788b9b', '23232', '841c467c-b980-4170-b324-35c46a08', '2015-04-14 00:00:00', '<nil>'),
  ('bf3c8afc27814a4e9e18457734a3f487', '23232', '8ac7bd34-6dd1-4182-a419-8977df7c', '2015-04-14 00:00:00', '127.0.0.1'),
  ('932c618f7484491e9b6074e31ee6a853', '23232', 'a8f363a2-cf40-4b34-9f9c-7e6875a9', '2015-04-14 00:00:00', '127.0.0.1'),
  ('c95518f9c3dd4fa1b56921cd5571dfce', '23232', 'd2186000-1ecd-4ad2-a987-c92bf750', '2015-04-14 00:00:00', '127.0.0.1'),
  ('8c1292966b40416cae44812402cf7ae6', '23232', '6521a66f-df41-45d2-8bbc-0708c977', '2015-04-14 00:00:00', '127.0.0.1'),
  ('9ea82f820c2f410cb2dbf787d5f0610b', '23232', 'c765b539-a566-4f58-a3b9-1d01ce9a', '2015-04-14 12:26:34', '127.0.0.1'),
  ('4f708f2bf2ae4ae9955d4ad37afe926b', '23232', '4ccfbc13-db09-4860-9b2e-88c6f26c', '2015-04-14 12:48:49', '127.0.0.1'),
  ('d1f05d177a7d4b83968a060d95b563ae', '23232', '19bbead6-90b4-4e5a-b4ba-ec0b54d1', '2015-04-14 12:54:55', '127.0.0.1'),
  ('f02242a759bd43f7a5493f9c6bb9bb5e', '23232', 'db0c4ddd-cab5-4b6e-a367-5df4fe6e', '2015-04-14 13:16:50', '127.0.0.1'),
  ('9ac14d153be445ca9177bc2746fa01ff', '23232', '57acf516-088a-469d-ad20-eb992d73', '2015-04-14 13:17:32', '127.0.0.1'),
  ('c24396a628924fec80af9982d3f7a9ce', '23232', 'ee382063-d863-466d-a9ec-1baa33aa', '2015-04-14 13:23:57', '127.0.0.1'),
  ('596620d54f2141d1b351f787df5c72e9', '23232', 'c604666c-c9f2-48c4-b854-fa2d5b7f', '2015-04-14 13:24:02', '127.0.0.1'),
  ('4644f411c0b448f687e4fce66fa05da2', '23232', 'a48e4c85-c713-4989-9a25-125b4dbc', '2015-04-14 13:46:10', '127.0.0.1'),
  ('91efebc08fdb467098665a0f1ed61b1d', '23232', 'de99c743-9de8-495c-a644-bb7170ef', '2015-04-14 13:46:12', '127.0.0.1'),
  ('ae7a430b7c714ce685493b0cf0904259', '23232', '965acc6a-79aa-4760-ab83-5e44dc4d', '2015-04-14 14:00:41', '127.0.0.1');
COMMIT;

-- ----------------------------
--  Table structure for `im_message`
-- ----------------------------
DROP TABLE IF EXISTS `im_message`;
CREATE TABLE `im_message` (
  `id`        VARCHAR(255) NOT NULL
  COMMENT '消息唯一标识',
  `sender`    VARCHAR(255) NOT NULL
  COMMENT '发送人(用户ID)',
  `contents`  VARCHAR(255) NOT NULL
  COMMENT '内容(支持富文本)',
  `date`      DATE         NOT NULL
  COMMENT '发送日期',
  `state`     CHAR(1)      NOT NULL DEFAULT '0'
  COMMENT '消息状态 0未发送，1送达，2已读，3取消，4删除',
  `direction` CHAR(1)      NOT NULL DEFAULT '0'
  COMMENT '方向 0发送，1接收',
  `type`      CHAR(1)      NOT NULL DEFAULT '0'
  COMMENT '消息类型 0聊天信息，1系统提示信息',
  `font`      VARCHAR(255)          DEFAULT NULL
  COMMENT '字体',
  `receiver`  VARCHAR(255) NOT NULL
  COMMENT '接收人(可以是用户，群，讨论组)',
  PRIMARY KEY (`id`)
)
  ENGINE = MyISAM
  DEFAULT CHARSET = utf8;

-- ----------------------------
--  Table structure for `im_relation_user_crowd`
-- ----------------------------
DROP TABLE IF EXISTS `im_relation_user_crowd`;
CREATE TABLE `im_relation_user_crowd` (
  `user_id`  VARCHAR(255) NOT NULL
  COMMENT '用户ID',
  `crowd_id` VARCHAR(255) NOT NULL
  COMMENT '群ID'
)
  ENGINE = MyISAM
  DEFAULT CHARSET = utf8;

-- ----------------------------
--  Table structure for `im_relation_user_discussion`
-- ----------------------------
DROP TABLE IF EXISTS `im_relation_user_discussion`;
CREATE TABLE `im_relation_user_discussion` (
  `user_id`       VARCHAR(255) NOT NULL
  COMMENT '用户ID',
  `discussion_id` VARCHAR(255) NOT NULL
  COMMENT '讨论组ID'
)
  ENGINE = MyISAM
  DEFAULT CHARSET = utf8;

-- ----------------------------
--  Table structure for `im_relation_user_group`
-- ----------------------------
DROP TABLE IF EXISTS `im_relation_user_group`;
CREATE TABLE `im_relation_user_group` (
  `user_id`  VARCHAR(255) DEFAULT NULL
  COMMENT '用户ID',
  `group_id` VARCHAR(255) DEFAULT NULL
  COMMENT '分组ID'
)
  ENGINE = MyISAM
  DEFAULT CHARSET = utf8;

-- ----------------------------
--  Table structure for `im_user`
-- ----------------------------
DROP TABLE IF EXISTS `im_user`;
CREATE TABLE `im_user` (
  `id`       VARCHAR(255) NOT NULL
  COMMENT '唯一标识',
  `account`  VARCHAR(255) NOT NULL
  COMMENT '账号',
  `password` VARCHAR(255) NOT NULL
  COMMENT '密码',
  `nick`     VARCHAR(255) NOT NULL
  COMMENT '用户昵称',
  `sign`     VARCHAR(255)          DEFAULT NULL
  COMMENT '个人前民',
  `avatar`   VARCHAR(255)          DEFAULT NULL
  COMMENT '头像',
  `status`   CHAR(1)      NOT NULL DEFAULT '0'
  COMMENT '状态 0离线，1在线，2离开，3请勿打扰，4忙碌，5Q我吧，6隐身',
  `remark`   VARCHAR(255)          DEFAULT NULL
  COMMENT '好友备注',
  PRIMARY KEY (`id`)
)
  ENGINE = MyISAM
  DEFAULT CHARSET = utf8;

-- ----------------------------
--  Records of `im_user`
-- ----------------------------
BEGIN;
INSERT INTO `im_user` VALUES ('23232', '44', '33', '22', '33', '33', '0', NULL);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
