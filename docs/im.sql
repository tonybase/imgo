/*
Navicat MySQL Data Transfer

Source Server         : 本地
Source Server Version : 50540
Source Host           : localhost:3306
Source Database       : im

Target Server Type    : MYSQL
Target Server Version : 50540
File Encoding         : 65001

Date: 2015-05-10 22:55:01
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for im_category
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
-- Records of im_category
-- ----------------------------
INSERT INTO `im_category` VALUES ('33', '我的好友', '11', '2015-05-04 21:55:31');
INSERT INTO `im_category` VALUES ('44', '我的好友', '22', '2015-05-04 21:57:38');

-- ----------------------------
-- Table structure for im_conn
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
-- Records of im_conn
-- ----------------------------

-- ----------------------------
-- Table structure for im_login
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
-- Records of im_login
-- ----------------------------
INSERT INTO `im_login` VALUES ('8d859eda-dab4-403c-a65c-80d7777e5e91', 'b5672dc4-d345-43e7-bf47-610c597273a6', '48fabb35-ac44-45b3-ae8e-ffadb14506a6', '2015-05-04 22:01:25', '127.0.0.1');
INSERT INTO `im_login` VALUES ('7a057449-f620-4e02-bb27-57593fae77a6', 'c1c3fa56-e5ca-4696-8b89-77b4f6041859', 'e2912583-7262-4ab0-8ecd-898bf2b64db1', '2015-05-04 22:02:09', '127.0.0.1');
INSERT INTO `im_login` VALUES ('51800cb9-3ff6-49e2-ade6-43313223aed1', 'b5672dc4-d345-43e7-bf47-610c597273a6', '4f56e4c0-0cbb-4a7b-a7ed-961e9305fb3a', '2015-05-04 22:04:16', '127.0.0.1');
INSERT INTO `im_login` VALUES ('d2731c86-cc79-43ea-b642-c25f0213f81b', '11', '5d1b9f84-a3eb-49bd-a2a9-fbf4a46ecdd5', '2015-05-04 22:26:20', '127.0.0.1');
INSERT INTO `im_login` VALUES ('5bb43b1a-c144-49eb-b70b-e090da545c32', '22', '42343521-583d-48a6-9765-680ceea206f9', '2015-05-04 22:26:48', '127.0.0.1');
INSERT INTO `im_login` VALUES ('46f82cbf-9fcd-4edd-9bbe-f57dad9e5f61', '11', 'ef9ce40a-5428-42b8-8266-a1467077ac7e', '2015-05-04 22:27:06', '127.0.0.1');
INSERT INTO `im_login` VALUES ('3cf7e740-fcf1-4616-97e6-60caad2e58db', '11', '4e1d6c06-e195-481c-a0b3-9043db606498', '2015-05-04 22:27:28', '127.0.0.1');
INSERT INTO `im_login` VALUES ('b9970b97-ddc1-4c56-8e5c-b37b5fcc82c8', '11', '611247a1-4e72-4624-b041-a86b64c4424d', '2015-05-04 22:27:46', '127.0.0.1');
INSERT INTO `im_login` VALUES ('02ca742d-84c0-43e0-867c-6b05d0e14c8c', '11', '15642a96-b06f-4d41-a990-67463f59bfef', '2015-05-04 22:28:32', '127.0.0.1');
INSERT INTO `im_login` VALUES ('6b242cfe-dcc4-4ff0-b5b1-4377a80693c5', '22', '5825d45e-e0dc-4771-8129-388ae9c5f980', '2015-05-04 22:28:40', '127.0.0.1');
INSERT INTO `im_login` VALUES ('3eef8145-c617-4de9-a2a1-fd0d18cb32c6', '22', '0aa43f2b-dcda-4ffe-a6a4-6b00ca176461', '2015-05-04 22:29:41', '127.0.0.1');
INSERT INTO `im_login` VALUES ('31c6461c-5bb4-4513-86b9-301a84608ae7', '11', '5e3056c7-7669-4e6f-834e-7097c93f7273', '2015-05-04 22:32:12', '127.0.0.1');
INSERT INTO `im_login` VALUES ('fad31367-438a-4947-b311-e53d8d53df48', '22', '69bbd4a1-e346-4db5-95f4-b5fa18ac8a92', '2015-05-04 22:32:23', '127.0.0.1');
INSERT INTO `im_login` VALUES ('379e258a-615b-4279-a8a9-3388c1538044', '11', '559327e4-0298-406d-bfae-7079e1cfc327', '2015-05-10 22:06:54', '127.0.0.1');

-- ----------------------------
-- Table structure for im_message
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
-- Records of im_message
-- ----------------------------

-- ----------------------------
-- Table structure for im_relation_user_category
-- ----------------------------
DROP TABLE IF EXISTS `im_relation_user_category`;
CREATE TABLE `im_relation_user_category` (
  `user_id` varchar(255) NOT NULL COMMENT '用户ID',
  `category_id` varchar(255) NOT NULL COMMENT '分类ID',
  `create_at` datetime NOT NULL COMMENT '建立好友关系时间'
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of im_relation_user_category
-- ----------------------------
INSERT INTO `im_relation_user_category` VALUES ('22', '33', '2015-05-04 21:55:44');
INSERT INTO `im_relation_user_category` VALUES ('11', '44', '0000-00-00 00:00:00');

-- ----------------------------
-- Table structure for im_relation_user_room
-- ----------------------------
DROP TABLE IF EXISTS `im_relation_user_room`;
CREATE TABLE `im_relation_user_room` (
  `user_id` varchar(255) NOT NULL COMMENT '用户ID',
  `room_id` varchar(255) NOT NULL COMMENT '聊天室ID',
  `create_at` datetime NOT NULL COMMENT '加入聊天室时间'
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of im_relation_user_room
-- ----------------------------

-- ----------------------------
-- Table structure for im_room
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
-- Records of im_room
-- ----------------------------

-- ----------------------------
-- Table structure for im_session
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
-- Records of im_session
-- ----------------------------
INSERT INTO `im_session` VALUES ('ff668c19-6ae1-4369-b1c2-0ad1a4fb6b21', '11', '22', '0', '2015-05-04 22:29:47');
INSERT INTO `im_session` VALUES ('44be7aa6-9f8e-4226-85ef-6d0148112526', '22', '11', '0', '2015-05-04 22:29:57');

-- ----------------------------
-- Table structure for im_user
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
-- Records of im_user
-- ----------------------------
INSERT INTO `im_user` VALUES ('22', '22', '22', '22', '', '', '0', '2015-05-04 21:55:31', '2015-05-04 21:55:31', null);
INSERT INTO `im_user` VALUES ('11', '11', '11', '11', '', '', '0', '2015-05-04 21:57:38', '2015-05-04 21:57:38', null);
