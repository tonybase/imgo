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

 Date: 05/20/2015 14:48:19 PM
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
  `sender_cate_id` varchar(255) NOT NULL COMMENT '发送者好友分类ID',
  `receiver` varchar(255) NOT NULL COMMENT '接收者',
  `send_at` datetime NOT NULL COMMENT '发送请求日期',
  `status` char(1) NOT NULL DEFAULT '0' COMMENT '状态 0未读、1同意、2拒绝',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

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
INSERT INTO `im_login` VALUES ('8d859eda-dab4-403c-a65c-80d7777e5e91', 'b5672dc4-d345-43e7-bf47-610c597273a6', '48fabb35-ac44-45b3-ae8e-ffadb14506a6', '2015-05-04 22:01:25', '127.0.0.1'), ('7a057449-f620-4e02-bb27-57593fae77a6', 'c1c3fa56-e5ca-4696-8b89-77b4f6041859', 'e2912583-7262-4ab0-8ecd-898bf2b64db1', '2015-05-04 22:02:09', '127.0.0.1'), ('51800cb9-3ff6-49e2-ade6-43313223aed1', 'b5672dc4-d345-43e7-bf47-610c597273a6', '4f56e4c0-0cbb-4a7b-a7ed-961e9305fb3a', '2015-05-04 22:04:16', '127.0.0.1'), ('d2731c86-cc79-43ea-b642-c25f0213f81b', '11', '5d1b9f84-a3eb-49bd-a2a9-fbf4a46ecdd5', '2015-05-04 22:26:20', '127.0.0.1'), ('5bb43b1a-c144-49eb-b70b-e090da545c32', '22', '42343521-583d-48a6-9765-680ceea206f9', '2015-05-04 22:26:48', '127.0.0.1'), ('46f82cbf-9fcd-4edd-9bbe-f57dad9e5f61', '11', 'ef9ce40a-5428-42b8-8266-a1467077ac7e', '2015-05-04 22:27:06', '127.0.0.1'), ('3cf7e740-fcf1-4616-97e6-60caad2e58db', '11', '4e1d6c06-e195-481c-a0b3-9043db606498', '2015-05-04 22:27:28', '127.0.0.1'), ('b9970b97-ddc1-4c56-8e5c-b37b5fcc82c8', '11', '611247a1-4e72-4624-b041-a86b64c4424d', '2015-05-04 22:27:46', '127.0.0.1'), ('02ca742d-84c0-43e0-867c-6b05d0e14c8c', '11', '15642a96-b06f-4d41-a990-67463f59bfef', '2015-05-04 22:28:32', '127.0.0.1'), ('6b242cfe-dcc4-4ff0-b5b1-4377a80693c5', '22', '5825d45e-e0dc-4771-8129-388ae9c5f980', '2015-05-04 22:28:40', '127.0.0.1'), ('3eef8145-c617-4de9-a2a1-fd0d18cb32c6', '22', '0aa43f2b-dcda-4ffe-a6a4-6b00ca176461', '2015-05-04 22:29:41', '127.0.0.1'), ('31c6461c-5bb4-4513-86b9-301a84608ae7', '11', '5e3056c7-7669-4e6f-834e-7097c93f7273', '2015-05-04 22:32:12', '127.0.0.1'), ('fad31367-438a-4947-b311-e53d8d53df48', '22', '69bbd4a1-e346-4db5-95f4-b5fa18ac8a92', '2015-05-04 22:32:23', '127.0.0.1'), ('379e258a-615b-4279-a8a9-3388c1538044', '11', '559327e4-0298-406d-bfae-7079e1cfc327', '2015-05-10 22:06:54', '127.0.0.1'), ('8e5b6d05-25ab-4af9-af6c-2cd4c77b8ad0', '11', '362dd8b5-7fee-45d6-a15f-2f82c53d8449', '2015-05-18 12:33:31', '127.0.0.1'), ('f5266cd4-b71b-4a07-a29b-3e1514fee344', '22', 'f0bf5120-62b6-4430-8f8f-db3c2332b714', '2015-05-18 12:34:22', '127.0.0.1'), ('871398ee-3a3b-4aa6-b252-6eb6c7d2974f', '11', 'ee812e28-2f51-443c-8794-a2af22d17019', '2015-05-19 16:05:24', '127.0.0.1'), ('21561cf3-fe0c-48ae-b4e8-3a75d8165ee7', '22', 'f182695f-fb6e-4a67-94ad-2abaa1327de5', '2015-05-19 16:16:48', '127.0.0.1'), ('c1425f71-e4df-4349-8b07-c11ae64342c8', '22', 'b17b2fcc-3208-4f61-8e1a-0047783a69c2', '2015-05-19 16:40:21', '127.0.0.1'), ('6de75700-04c8-41be-8b72-de4f258aed87', '11', '9ebd3abc-33da-4ed1-af4b-5d7de2a58d6b', '2015-05-19 16:44:17', '127.0.0.1'), ('93ae4085-0a8d-4fc0-aa0c-fe37fdd4e8bd', '11', '472610fa-2b09-434b-ae58-592727cd0232', '2015-05-19 16:47:42', '127.0.0.1'), ('f76201b4-6f1c-40f9-bf10-083b4bd4eb55', '11', 'a75a538a-735b-4a1d-ae5f-7bd63134eed0', '2015-05-19 16:49:21', '127.0.0.1'), ('90654db6-a0d0-4e89-bd29-ec9301c7cd32', '11', '2a1d32d7-3d0f-497f-9ee7-1e9ab7461cf7', '2015-05-19 16:51:00', '127.0.0.1'), ('3cd9d6b9-a76c-42f1-8446-c826617e27bd', '11', '894f1e9b-cad2-4c02-801b-fe4191a1b0c8', '2015-05-19 16:53:54', '127.0.0.1'), ('833a2f89-b442-45ad-aff0-6cdc4c8b1d6c', '11', 'e9c2e0b9-dbf8-479f-937d-41b243796a6d', '2015-05-19 16:58:58', '127.0.0.1'), ('13196f30-000e-4abd-8364-078deb0b8c16', '11', 'd8aeec51-807f-4ac9-a22a-e69b299d77c1', '2015-05-19 17:00:22', '127.0.0.1'), ('3d7dffa1-060d-41d1-a512-c8ac6b42149b', '11', '28b5bf9e-b749-4a16-a408-531252c18022', '2015-05-19 17:00:52', '127.0.0.1'), ('fda4f2f3-2e4f-445b-9b5d-7d2ab67393c3', '11', 'c07be1d0-6b29-486c-b42e-a83ba6df9c8a', '2015-05-19 17:02:02', '127.0.0.1'), ('af45261e-c287-4898-8144-057b5da53524', '11', '6029a8b0-890a-4b0e-9a49-10867c497e5f', '2015-05-19 17:05:13', '127.0.0.1'), ('42cec161-aa35-46ae-bb96-0ea7def3a17f', '11', 'ef8ec53a-e976-4b63-96da-cb7290cec7bf', '2015-05-19 17:09:11', '127.0.0.1'), ('07db54b0-8316-40b8-a299-33e728e065f3', '11', '15c70b48-01c5-4153-8113-d0620aaf7079', '2015-05-19 17:10:01', '127.0.0.1'), ('ab357f71-5882-4b77-a497-2e28e44c1739', '11', 'ff3e4b36-c83e-4ec7-aeca-0376a35a42eb', '2015-05-19 17:10:28', '127.0.0.1'), ('7ed4c419-1efb-427e-86d8-a0d4efd96a9d', '22', 'e31e9998-748f-4851-bba5-e2545b9eac67', '2015-05-19 17:11:45', '127.0.0.1'), ('4e25facc-4612-4d10-94a6-4a3089a908fe', '11', '8e5f44be-a17f-4bab-a2b9-7116fe2aa2fe', '2015-05-19 17:13:54', '127.0.0.1'), ('b11ba3b6-05ea-4684-b175-9f331d83c470', '11', '5c357c60-b1e7-4798-93f7-02a5040ec3e5', '2015-05-19 17:15:09', '127.0.0.1'), ('c4450228-69bd-495e-99ee-68fdc3b45dce', '11', 'f61011d3-35ee-4c5e-a589-a6a54e9ea84e', '2015-05-19 17:25:50', '127.0.0.1'), ('2a881cc1-8276-4a1d-bd28-be43147cc2da', '11', 'cb9b8b98-98da-494c-8122-f7569b1b702a', '2015-05-19 17:26:34', '127.0.0.1'), ('241d2905-3e7a-4581-b1a3-539c80c99c2a', '11', 'e9cef4df-13f7-4a9e-a511-9338d1bcebeb', '2015-05-19 17:28:30', '127.0.0.1');
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
INSERT INTO `im_user` VALUES ('22', '22', '22', 'Itnik', '我是要成为海贼王的男人', '', '0', '2015-05-04 21:55:31', '2015-05-04 21:55:31', null), ('11', '11', '11', 'toy', '马上回来', '', '0', '2015-05-04 21:57:38', '2015-05-04 21:57:38', null);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
