/*
Navicat MySQL Data Transfer

Source Server         : 本地
Source Server Version : 50540
Source Host           : localhost:3306
Source Database       : im

Target Server Type    : MYSQL
Target Server Version : 50540
File Encoding         : 65001

Date: 2015-04-26 11:35:27
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for im_conn
-- ----------------------------
DROP TABLE IF EXISTS `im_conn`;
CREATE TABLE `im_conn` (
  `user_id` varchar(50) CHARACTER SET utf8 NOT NULL,
  `token` varchar(50) NOT NULL,
  `key` varchar(50) NOT NULL,
  `date` datetime NOT NULL,
  PRIMARY KEY (`key`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of im_conn
-- ----------------------------
INSERT INTO `im_conn` VALUES ('1', '8be47788-5dbf-4cdb-a68b-b3273e977f18', 'ff9b5768-ebe6-46a1-b4f1-4a57fbb88711', '2015-04-26 11:24:47');
INSERT INTO `im_conn` VALUES ('2', '8b60ff9f-a71b-4c54-9c47-96f89749afe0', 'c1feffb1-0ea7-4c2f-9c5e-0b8c73538259', '2015-04-26 11:24:57');

-- ----------------------------
-- Table structure for im_conversation
-- ----------------------------
DROP TABLE IF EXISTS `im_conversation`;
CREATE TABLE `im_conversation` (
  `id` varchar(50) NOT NULL COMMENT '会话的唯一标识',
  `creater` varchar(50) NOT NULL COMMENT '创建人',
  `create_date` datetime NOT NULL COMMENT '创建日期',
  `receiver` varchar(50) NOT NULL COMMENT '接收人(可以是用户，群，讨论组)',
  `type` char(1) NOT NULL DEFAULT '0' COMMENT '会话类型 0用户会话，1群会话，2讨论组会话',
  `token` varchar(50) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of im_conversation
-- ----------------------------
INSERT INTO `im_conversation` VALUES ('5cac3aa1-7803-4423-bf3a-5813a3fecc0f', '1', '2015-04-26 00:00:00', '2', '0', '8be47788-5dbf-4cdb-a68b-b3273e977f18');
INSERT INTO `im_conversation` VALUES ('cce2d800-6c6c-4549-be01-f09494ace0f2', '2', '2015-04-26 00:00:00', '1', '0', '8b60ff9f-a71b-4c54-9c47-96f89749afe0');

-- ----------------------------
-- Table structure for im_crowd
-- ----------------------------
DROP TABLE IF EXISTS `im_crowd`;
CREATE TABLE `im_crowd` (
  `id` varchar(50) NOT NULL COMMENT '群的唯一标识',
  `name` varchar(50) NOT NULL COMMENT '群名称',
  `creater` varchar(50) NOT NULL COMMENT '创建者',
  `create_date` datetime NOT NULL COMMENT '创建日期',
  `user_num` int(11) NOT NULL DEFAULT '100' COMMENT '群允许的用户数量',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of im_crowd
-- ----------------------------

-- ----------------------------
-- Table structure for im_discussion
-- ----------------------------
DROP TABLE IF EXISTS `im_discussion`;
CREATE TABLE `im_discussion` (
  `id` varchar(50) NOT NULL COMMENT '讨论组的唯一标识',
  `name` varchar(50) NOT NULL COMMENT '讨论组名',
  `creater` varchar(50) NOT NULL COMMENT '创建者',
  `create_date` datetime NOT NULL COMMENT '创建日期',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of im_discussion
-- ----------------------------

-- ----------------------------
-- Table structure for im_group
-- ----------------------------
DROP TABLE IF EXISTS `im_group`;
CREATE TABLE `im_group` (
  `id` varchar(50) NOT NULL COMMENT '唯一标识',
  `name` varchar(50) NOT NULL COMMENT '分组名',
  `creater` varchar(50) DEFAULT NULL COMMENT '创建人',
  `create_date` datetime DEFAULT NULL COMMENT '创建日期',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of im_group
-- ----------------------------
INSERT INTO `im_group` VALUES ('1', '我的好友', '1', '2015-04-17 00:00:00');
INSERT INTO `im_group` VALUES ('2', '我的好友', '2', '2015-04-17 00:00:00');

-- ----------------------------
-- Table structure for im_login
-- ----------------------------
DROP TABLE IF EXISTS `im_login`;
CREATE TABLE `im_login` (
  `id` varchar(50) NOT NULL COMMENT '登录记录唯一标识',
  `user_id` varchar(50) NOT NULL COMMENT '用户ID',
  `token` varchar(50) NOT NULL COMMENT '用户token',
  `login_date` datetime NOT NULL COMMENT '登录日期',
  `login_ip` varchar(20) NOT NULL COMMENT '用户登录IP',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of im_login
-- ----------------------------
INSERT INTO `im_login` VALUES ('069c5de7-3049-4440-a20c-8e65721d543d', '1', '8be47788-5dbf-4cdb-a68b-b3273e977f18', '2015-04-26 11:24:47', '127.0.0.1');
INSERT INTO `im_login` VALUES ('757560c9-4048-4b27-b629-3aec6e9bd199', '2', '8b60ff9f-a71b-4c54-9c47-96f89749afe0', '2015-04-26 11:24:57', '127.0.0.1');

-- ----------------------------
-- Table structure for im_message
-- ----------------------------
DROP TABLE IF EXISTS `im_message`;
CREATE TABLE `im_message` (
  `id` varchar(255) NOT NULL COMMENT '消息唯一标识',
  `sender` varchar(50) NOT NULL COMMENT '发送人(用户ID)',
  `contents` varchar(255) NOT NULL COMMENT '内容(支持富文本)',
  `date` datetime NOT NULL COMMENT '发送日期',
  `state` char(1) NOT NULL DEFAULT '0' COMMENT '消息状态 0未发送，1送达，2已读，3取消，4删除',
  `direction` char(1) NOT NULL DEFAULT '0' COMMENT '方向 0发送，1接收',
  `type` char(1) NOT NULL DEFAULT '0' COMMENT '消息类型 0聊天信息，1系统提示信息',
  `font` varchar(255) DEFAULT NULL COMMENT '字体',
  `receiver` varchar(50) NOT NULL COMMENT '接收人(可以是用户，群，讨论组)',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of im_message
-- ----------------------------

-- ----------------------------
-- Table structure for im_relation_user_crowd
-- ----------------------------
DROP TABLE IF EXISTS `im_relation_user_crowd`;
CREATE TABLE `im_relation_user_crowd` (
  `user_id` varchar(255) NOT NULL COMMENT '用户ID',
  `crowd_id` varchar(255) NOT NULL COMMENT '群ID'
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of im_relation_user_crowd
-- ----------------------------

-- ----------------------------
-- Table structure for im_relation_user_discussion
-- ----------------------------
DROP TABLE IF EXISTS `im_relation_user_discussion`;
CREATE TABLE `im_relation_user_discussion` (
  `user_id` varchar(255) NOT NULL COMMENT '用户ID',
  `discussion_id` varchar(255) NOT NULL COMMENT '讨论组ID'
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of im_relation_user_discussion
-- ----------------------------

-- ----------------------------
-- Table structure for im_relation_user_group
-- ----------------------------
DROP TABLE IF EXISTS `im_relation_user_group`;
CREATE TABLE `im_relation_user_group` (
  `user_id` varchar(50) DEFAULT NULL COMMENT '用户ID',
  `group_id` varchar(50) DEFAULT NULL COMMENT '分组ID'
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of im_relation_user_group
-- ----------------------------
INSERT INTO `im_relation_user_group` VALUES ('2', '1');
INSERT INTO `im_relation_user_group` VALUES ('1', '2');

-- ----------------------------
-- Table structure for im_user
-- ----------------------------
DROP TABLE IF EXISTS `im_user`;
CREATE TABLE `im_user` (
  `id` varchar(50) NOT NULL COMMENT '唯一标识',
  `account` varchar(50) NOT NULL COMMENT '账号',
  `password` varchar(50) NOT NULL COMMENT '密码',
  `nick` varchar(50) NOT NULL COMMENT '用户昵称',
  `sign` varchar(50) DEFAULT NULL COMMENT '个人前民',
  `avatar` varchar(255) DEFAULT NULL COMMENT '头像',
  `status` char(1) NOT NULL DEFAULT '0' COMMENT '状态 0离线，1在线，2离开，3请勿打扰，4忙碌，5Q我吧，6隐身',
  `remark` varchar(255) DEFAULT NULL COMMENT '好友备注',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of im_user
-- ----------------------------
INSERT INTO `im_user` VALUES ('1', '11', '11', '11', '11', '11', '1', null);
INSERT INTO `im_user` VALUES ('2', '22', '22', '22', '22', '22', '1', null);
