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

 Date: 05/27/2015 10:40:11 AM
*/

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `im_buddy_request`
-- ----------------------------
DROP TABLE IF EXISTS `im_buddy_request`;
CREATE TABLE `im_buddy_request` (
  `id` varchar(255) NOT NULL COMMENT 'ID',
  `sender` varchar(255) NOT NULL COMMENT '发送者',
  `sender_category_id` varchar(255) NOT NULL COMMENT '发送者好友分类ID',
  `receiver` varchar(255) NOT NULL COMMENT '接收者',
  `receiver_category_id` varchar(255) DEFAULT NULL,
  `send_at` datetime NOT NULL COMMENT '发送请求日期',
  `status` char(1) NOT NULL DEFAULT '0' COMMENT '状态 0未读、1同意、2拒绝',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `im_buddy_request`
-- ----------------------------
BEGIN;
INSERT INTO `im_buddy_request` VALUES ('66708492-6201-4c03-a7ef-6e7bbc6f589c', '22', '44', '11', null, '2015-05-26 10:59:38', '2'), ('d44369b7-27c1-4de4-bcfc-44385094f7e1', '22', '44', '11', '33', '2015-05-26 11:00:01', '1');
COMMIT;

-- ----------------------------
--  Table structure for `im_category`
-- ----------------------------
DROP TABLE IF EXISTS `im_category`;
CREATE TABLE `im_category` (
  `id` varchar(255) NOT NULL COMMENT '唯一标识',
  `name` varchar(255) NOT NULL COMMENT '分类名',
  `creator` varchar(255) DEFAULT NULL COMMENT '创建人 user_id',
  `create_at` datetime DEFAULT NULL COMMENT '创建日期',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `im_category`
-- ----------------------------
BEGIN;
INSERT INTO `im_category` VALUES ('33', '我的好友', '11', '2015-05-04 21:55:31'), ('44', '我的好友', '22', '2015-05-04 21:57:38');
COMMIT;

-- ----------------------------
--  Table structure for `im_conn`
-- ----------------------------
DROP TABLE IF EXISTS `im_conn`;
CREATE TABLE `im_conn` (
  `id` varchar(255) NOT NULL COMMENT '连接唯一标识',
  `user_id` varchar(255) NOT NULL COMMENT '用户ID',
  `token` varchar(255) NOT NULL COMMENT '连接TOKEN',
  `create_at` datetime NOT NULL COMMENT '创建日期',
  `update_at` datetime NOT NULL COMMENT '更新日期',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

-- ----------------------------
--  Records of `im_conn`
-- ----------------------------
BEGIN;
INSERT INTO `im_conn` VALUES ('d7e8aab3-2546-418a-a66b-4f4d7ac7dd6d', '11', '9b8c7bea-369d-4d8a-8145-748ac54748fa', '2015-05-27 10:19:08', '2015-05-27 10:19:08');
COMMIT;

-- ----------------------------
--  Table structure for `im_login`
-- ----------------------------
DROP TABLE IF EXISTS `im_login`;
CREATE TABLE `im_login` (
  `id` varchar(255) NOT NULL COMMENT '登录记录唯一标识',
  `user_id` varchar(255) NOT NULL COMMENT '用户ID',
  `token` varchar(255) NOT NULL COMMENT '用户token',
  `login_at` datetime NOT NULL COMMENT '登录日期',
  `login_ip` varchar(32) NOT NULL COMMENT '用户登录IP',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `im_login`
-- ----------------------------
BEGIN;
INSERT INTO `im_login` VALUES ('3b1901d1-5788-46e7-97f3-0b796cdf0e6d', '11', 'a09932a6-0654-4cfe-8199-d877cd01ae5e', '2015-05-26 11:03:54', '127.0.0.1'), ('0e5cc251-b881-47ce-9da3-d746f56d9597', '11', 'd0d6a544-bc54-4b5f-a824-b744800b70de', '2015-05-26 11:08:21', '127.0.0.1'), ('d933314c-89f9-46ef-b175-503e1649a6c0', '11', 'c84c4962-fe94-451f-9926-57cd80393f54', '2015-05-26 11:11:36', '127.0.0.1'), ('2f231cd5-5687-4e1f-9683-75c3da3c3209', '11', '45c1c7cf-a4c1-4825-8459-9ce5d1869462', '2015-05-26 11:14:10', '127.0.0.1'), ('9a18e073-6e26-4808-b081-0ddc1245f505', '11', '889a96ff-9f8c-420c-9389-cc3ed5cb56e6', '2015-05-26 11:19:46', '127.0.0.1'), ('bda8ae16-9b71-436b-8071-cbb1881b15e0', '11', '2c2bc2e8-9d71-41e6-a374-d552544327de', '2015-05-26 11:20:26', '127.0.0.1'), ('73a71b3f-64bc-45e3-abd1-cca4ca13fa6f', '11', '534fba51-5687-4546-9b04-97f8c7936426', '2015-05-26 11:28:33', '127.0.0.1'), ('51bda13c-fcd9-4c88-981e-13262d5b9915', '11', '90b7a455-8cc3-407b-adde-081b7534fa83', '2015-05-26 11:31:27', '127.0.0.1'), ('079f1ad1-a655-44ce-b0d4-516f7042f91a', '11', '22ee84c3-3978-4ea8-b199-13f448763c5c', '2015-05-27 08:56:16', '127.0.0.1'), ('0fcd0b70-da2a-493c-8f0e-d4252d2a0166', '22', '952a2ace-e434-4236-b04e-0381add5773f', '2015-05-27 08:56:55', '127.0.0.1'), ('6cd10533-7364-4be7-8ca8-f2318c19b16d', '11', 'd05afbb0-48dc-4e49-95cd-b09e07f72159', '2015-05-27 08:59:15', '127.0.0.1'), ('bfd8d7ec-6b31-418e-9022-1f669504de62', '11', '6316e29a-eed0-4a9c-b4ef-8abae5421056', '2015-05-27 09:00:01', '127.0.0.1'), ('fcc4a579-a164-4c0b-8c8c-e8981a4be173', '11', '9b8c7bea-369d-4d8a-8145-748ac54748fa', '2015-05-27 10:19:08', '127.0.0.1'), ('cbaa5779-7338-4532-9f80-ce6324e5f2f9', '11', '95f82794-5b25-48e2-8af9-ae81024c7b44', '2015-05-27 10:24:05', '127.0.0.1'), ('383ed174-f6e3-4f7f-98cb-9435dfd5e6d0', '11', 'b172a467-03dd-4d3f-902a-647e60ed558d', '2015-05-27 10:27:08', '127.0.0.1'), ('abbe3dde-0f26-4ca5-8eae-c503f3aac1b0', '11', 'ee4155aa-b38f-4d1a-bfdb-3294f0b1df1b', '2015-05-27 10:36:05', '127.0.0.1');
COMMIT;

-- ----------------------------
--  Table structure for `im_message`
-- ----------------------------
DROP TABLE IF EXISTS `im_message`;
CREATE TABLE `im_message` (
  `id` varchar(255) NOT NULL COMMENT '消息唯一标识',
  `sender` varchar(255) NOT NULL COMMENT '发送人(用户ID)',
  `contents` varchar(255) NOT NULL COMMENT '内容(支持富文本)',
  `send_at` datetime NOT NULL COMMENT '发送日期',
  `state` char(1) NOT NULL DEFAULT '0' COMMENT '消息状态 0未发送，1送达，2已读，3取消，4删除',
  `direction` char(1) NOT NULL DEFAULT '0' COMMENT '方向 0发送，1接收',
  `type` char(1) NOT NULL DEFAULT '0' COMMENT '消息类型 0聊天信息，1系统提示信息',
  `font` varchar(255) DEFAULT NULL COMMENT '字体',
  `receiver` varchar(255) NOT NULL COMMENT '接收人(可以是用户，群，讨论组)',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `im_relation_user_category`
-- ----------------------------
DROP TABLE IF EXISTS `im_relation_user_category`;
CREATE TABLE `im_relation_user_category` (
  `user_id` varchar(255) NOT NULL COMMENT '用户ID',
  `category_id` varchar(255) NOT NULL COMMENT '分类ID',
  `create_at` datetime NOT NULL COMMENT '建立好友关系时间'
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `im_relation_user_category`
-- ----------------------------
BEGIN;
INSERT INTO `im_relation_user_category` VALUES ('22', '33', '2015-05-04 21:55:44'), ('11', '44', '0000-00-00 00:00:00');
COMMIT;

-- ----------------------------
--  Table structure for `im_relation_user_room`
-- ----------------------------
DROP TABLE IF EXISTS `im_relation_user_room`;
CREATE TABLE `im_relation_user_room` (
  `user_id` varchar(255) NOT NULL COMMENT '用户ID',
  `room_id` varchar(255) NOT NULL COMMENT '聊天室ID',
  `create_at` datetime NOT NULL COMMENT '加入聊天室时间'
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `im_room`
-- ----------------------------
DROP TABLE IF EXISTS `im_room`;
CREATE TABLE `im_room` (
  `id` varchar(255) NOT NULL COMMENT '群的唯一标识',
  `name` varchar(255) NOT NULL COMMENT '群名称',
  `creator` varchar(255) NOT NULL COMMENT '创建者 user_id',
  `create_at` datetime NOT NULL COMMENT '创建日期',
  `user_num` int(11) NOT NULL DEFAULT '100' COMMENT '群允许的用户数量',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `im_session`
-- ----------------------------
DROP TABLE IF EXISTS `im_session`;
CREATE TABLE `im_session` (
  `id` varchar(255) NOT NULL COMMENT '会话的唯一标识',
  `creator` varchar(255) NOT NULL COMMENT '创建者 user_id',
  `receiver` varchar(255) NOT NULL COMMENT '接收人(可以是用户，群，讨论组)',
  `type` char(1) NOT NULL DEFAULT '0' COMMENT '会话类型 0用户会话，1群会话，2讨论组会话',
  `create_at` datetime NOT NULL COMMENT '创建日期',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `im_session`
-- ----------------------------
BEGIN;
INSERT INTO `im_session` VALUES ('ff668c19-6ae1-4369-b1c2-0ad1a4fb6b21', '11', '22', '0', '2015-05-04 22:29:47'), ('44be7aa6-9f8e-4226-85ef-6d0148112526', '22', '11', '0', '2015-05-04 22:29:57');
COMMIT;

-- ----------------------------
--  Table structure for `im_user`
-- ----------------------------
DROP TABLE IF EXISTS `im_user`;
CREATE TABLE `im_user` (
  `id` varchar(255) NOT NULL COMMENT '唯一标识',
  `account` varchar(255) NOT NULL COMMENT '账号',
  `password` varchar(255) NOT NULL COMMENT '密码',
  `nick` varchar(255) NOT NULL COMMENT '用户昵称',
  `sign` varchar(255) DEFAULT '' COMMENT '个人前民',
  `avatar` varchar(255) DEFAULT NULL COMMENT '头像',
  `status` char(1) NOT NULL DEFAULT '0' COMMENT '状态 0离线，1在线，2离开，3请勿打扰，4忙碌，5Q我吧，6隐身',
  `create_at` datetime NOT NULL COMMENT '注册日期',
  `update_at` datetime NOT NULL COMMENT '更新日期',
  `remark` varchar(255) DEFAULT NULL COMMENT '好友备注',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `im_user`
-- ----------------------------
BEGIN;
INSERT INTO `im_user` VALUES ('22', '22', '22', 'Itnik', '我是要成为海贼王的男人', 'http://att2.citysbs.com/hangzhou/sns01/11_2011/14/middle_20461634334ec10b53d15d53.34164698.jpg', '0', '2015-05-04 21:55:31', '2015-05-04 21:55:31', null), ('11', '11', '11', 'toy', '马上回来', 'http://img.qqbody.com/uploads/allimg/201409/02-173237_949.jpg', '0', '2015-05-04 21:57:38', '2015-05-04 21:57:38', null);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
