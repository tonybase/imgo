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

 Date: 04/20/2015 22:39:48 PM
*/
CREATE DATABASE IF NOT EXISTS im DEFAULT CHARSET utf8 COLLATE utf8_general_ci;
USE im;

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `im_conn`
-- ----------------------------
DROP TABLE IF EXISTS `im_conn`;
CREATE TABLE `im_conn` (
  `user_id` varchar(255) CHARACTER SET utf8 NOT NULL,
  `token` varchar(255) NOT NULL,
  `key` varchar(255) NOT NULL,
  `date` datetime NOT NULL,
  PRIMARY KEY (`key`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- ----------------------------
--  Table structure for `im_conversation`
-- ----------------------------
DROP TABLE IF EXISTS `im_conversation`;
CREATE TABLE `im_conversation` (
  `id` varchar(255) NOT NULL COMMENT '会话的唯一标识',
  `creater` varchar(255) NOT NULL COMMENT '创建人',
  `create_date` date NOT NULL COMMENT '创建日期',
  `receiver` varchar(255) CHARACTER SET utf8 NOT NULL COMMENT '接收人(可以是用户，群，讨论组)',
  `type` char(1) NOT NULL DEFAULT '0' COMMENT '会话类型 0用户会话，1群会话，2讨论组会话',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `im_crowd`
-- ----------------------------
DROP TABLE IF EXISTS `im_crowd`;
CREATE TABLE `im_crowd` (
  `id` varchar(255) NOT NULL COMMENT '群的唯一标识',
  `name` varchar(255) NOT NULL COMMENT '群名称',
  `creater` varchar(255) NOT NULL COMMENT '创建者',
  `create_date` date NOT NULL COMMENT '创建日期',
  `user_num` int(11) NOT NULL DEFAULT '100' COMMENT '群允许的用户数量',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `im_discussion`
-- ----------------------------
DROP TABLE IF EXISTS `im_discussion`;
CREATE TABLE `im_discussion` (
  `id` varchar(255) NOT NULL COMMENT '讨论组的唯一标识',
  `name` varchar(255) NOT NULL COMMENT '讨论组名',
  `creater` varchar(255) NOT NULL COMMENT '创建者',
  `create_date` date NOT NULL COMMENT '创建日期',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `im_group`
-- ----------------------------
DROP TABLE IF EXISTS `im_group`;
CREATE TABLE `im_group` (
  `id` varchar(255) NOT NULL COMMENT '唯一标识',
  `name` varchar(255) NOT NULL COMMENT '分组名',
  `creater` varchar(255) DEFAULT NULL COMMENT '创建人',
  `create_date` date DEFAULT NULL COMMENT '创建日期',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `im_group`
-- ----------------------------
BEGIN;
INSERT INTO `im_group` VALUES ('1', '我的好友', '1', '2015-04-17'), ('2', '我的好友', '2', '2015-04-17');
COMMIT;

-- ----------------------------
--  Table structure for `im_login`
-- ----------------------------
DROP TABLE IF EXISTS `im_login`;
CREATE TABLE `im_login` (
  `id` varchar(255) NOT NULL COMMENT '登录记录唯一标识',
  `user_id` varchar(255) NOT NULL COMMENT '用户ID',
  `token` varchar(255) NOT NULL COMMENT '用户token',
  `login_date` datetime NOT NULL COMMENT '登录日期',
  `login_ip` varchar(20) NOT NULL COMMENT '用户登录IP',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `im_message`
-- ----------------------------
DROP TABLE IF EXISTS `im_message`;
CREATE TABLE `im_message` (
  `id` varchar(255) NOT NULL COMMENT '消息唯一标识',
  `sender` varchar(255) NOT NULL COMMENT '发送人(用户ID)',
  `contents` varchar(255) NOT NULL COMMENT '内容(支持富文本)',
  `date` date NOT NULL COMMENT '发送日期',
  `state` char(1) NOT NULL DEFAULT '0' COMMENT '消息状态 0未发送，1送达，2已读，3取消，4删除',
  `direction` char(1) NOT NULL DEFAULT '0' COMMENT '方向 0发送，1接收',
  `type` char(1) NOT NULL DEFAULT '0' COMMENT '消息类型 0聊天信息，1系统提示信息',
  `font` varchar(255) DEFAULT NULL COMMENT '字体',
  `receiver` varchar(255) NOT NULL COMMENT '接收人(可以是用户，群，讨论组)',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `im_relation_user_crowd`
-- ----------------------------
DROP TABLE IF EXISTS `im_relation_user_crowd`;
CREATE TABLE `im_relation_user_crowd` (
  `user_id` varchar(255) NOT NULL COMMENT '用户ID',
  `crowd_id` varchar(255) NOT NULL COMMENT '群ID'
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `im_relation_user_discussion`
-- ----------------------------
DROP TABLE IF EXISTS `im_relation_user_discussion`;
CREATE TABLE `im_relation_user_discussion` (
  `user_id` varchar(255) NOT NULL COMMENT '用户ID',
  `discussion_id` varchar(255) NOT NULL COMMENT '讨论组ID'
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `im_relation_user_group`
-- ----------------------------
DROP TABLE IF EXISTS `im_relation_user_group`;
CREATE TABLE `im_relation_user_group` (
  `user_id` varchar(255) DEFAULT NULL COMMENT '用户ID',
  `group_id` varchar(255) DEFAULT NULL COMMENT '分组ID'
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `im_relation_user_group`
-- ----------------------------
BEGIN;
INSERT INTO `im_relation_user_group` VALUES ('1', '1');
INSERT INTO `im_relation_user_group` VALUES ('2', '2');
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
  `sign` varchar(255) DEFAULT NULL COMMENT '个人前民',
  `avatar` varchar(255) DEFAULT NULL COMMENT '头像',
  `status` char(1) NOT NULL DEFAULT '0' COMMENT '状态 0离线，1在线，2离开，3请勿打扰，4忙碌，5Q我吧，6隐身',
  `remark` varchar(255) DEFAULT NULL COMMENT '好友备注',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `im_user`
-- ----------------------------
BEGIN;
INSERT INTO `im_user` VALUES ('1', '11', '11', '11', '11', '11', '0', null), ('2', '22', '22', '22', '22', '22', '0', null);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
