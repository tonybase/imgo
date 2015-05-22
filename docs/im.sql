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

 Date: 05/22/2015 10:28:38 AM
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
INSERT INTO `im_login` VALUES ('8d859eda-dab4-403c-a65c-80d7777e5e91', 'b5672dc4-d345-43e7-bf47-610c597273a6', '48fabb35-ac44-45b3-ae8e-ffadb14506a6', '2015-05-04 22:01:25', '127.0.0.1'), ('7a057449-f620-4e02-bb27-57593fae77a6', 'c1c3fa56-e5ca-4696-8b89-77b4f6041859', 'e2912583-7262-4ab0-8ecd-898bf2b64db1', '2015-05-04 22:02:09', '127.0.0.1'), ('51800cb9-3ff6-49e2-ade6-43313223aed1', 'b5672dc4-d345-43e7-bf47-610c597273a6', '4f56e4c0-0cbb-4a7b-a7ed-961e9305fb3a', '2015-05-04 22:04:16', '127.0.0.1'), ('d2731c86-cc79-43ea-b642-c25f0213f81b', '11', '5d1b9f84-a3eb-49bd-a2a9-fbf4a46ecdd5', '2015-05-04 22:26:20', '127.0.0.1'), ('5bb43b1a-c144-49eb-b70b-e090da545c32', '22', '42343521-583d-48a6-9765-680ceea206f9', '2015-05-04 22:26:48', '127.0.0.1'), ('46f82cbf-9fcd-4edd-9bbe-f57dad9e5f61', '11', 'ef9ce40a-5428-42b8-8266-a1467077ac7e', '2015-05-04 22:27:06', '127.0.0.1'), ('3cf7e740-fcf1-4616-97e6-60caad2e58db', '11', '4e1d6c06-e195-481c-a0b3-9043db606498', '2015-05-04 22:27:28', '127.0.0.1'), ('b9970b97-ddc1-4c56-8e5c-b37b5fcc82c8', '11', '611247a1-4e72-4624-b041-a86b64c4424d', '2015-05-04 22:27:46', '127.0.0.1'), ('02ca742d-84c0-43e0-867c-6b05d0e14c8c', '11', '15642a96-b06f-4d41-a990-67463f59bfef', '2015-05-04 22:28:32', '127.0.0.1'), ('6b242cfe-dcc4-4ff0-b5b1-4377a80693c5', '22', '5825d45e-e0dc-4771-8129-388ae9c5f980', '2015-05-04 22:28:40', '127.0.0.1'), ('3eef8145-c617-4de9-a2a1-fd0d18cb32c6', '22', '0aa43f2b-dcda-4ffe-a6a4-6b00ca176461', '2015-05-04 22:29:41', '127.0.0.1'), ('31c6461c-5bb4-4513-86b9-301a84608ae7', '11', '5e3056c7-7669-4e6f-834e-7097c93f7273', '2015-05-04 22:32:12', '127.0.0.1'), ('fad31367-438a-4947-b311-e53d8d53df48', '22', '69bbd4a1-e346-4db5-95f4-b5fa18ac8a92', '2015-05-04 22:32:23', '127.0.0.1'), ('379e258a-615b-4279-a8a9-3388c1538044', '11', '559327e4-0298-406d-bfae-7079e1cfc327', '2015-05-10 22:06:54', '127.0.0.1'), ('8e5b6d05-25ab-4af9-af6c-2cd4c77b8ad0', '11', '362dd8b5-7fee-45d6-a15f-2f82c53d8449', '2015-05-18 12:33:31', '127.0.0.1'), ('f5266cd4-b71b-4a07-a29b-3e1514fee344', '22', 'f0bf5120-62b6-4430-8f8f-db3c2332b714', '2015-05-18 12:34:22', '127.0.0.1'), ('871398ee-3a3b-4aa6-b252-6eb6c7d2974f', '11', 'ee812e28-2f51-443c-8794-a2af22d17019', '2015-05-19 16:05:24', '127.0.0.1'), ('21561cf3-fe0c-48ae-b4e8-3a75d8165ee7', '22', 'f182695f-fb6e-4a67-94ad-2abaa1327de5', '2015-05-19 16:16:48', '127.0.0.1'), ('c1425f71-e4df-4349-8b07-c11ae64342c8', '22', 'b17b2fcc-3208-4f61-8e1a-0047783a69c2', '2015-05-19 16:40:21', '127.0.0.1'), ('6de75700-04c8-41be-8b72-de4f258aed87', '11', '9ebd3abc-33da-4ed1-af4b-5d7de2a58d6b', '2015-05-19 16:44:17', '127.0.0.1'), ('93ae4085-0a8d-4fc0-aa0c-fe37fdd4e8bd', '11', '472610fa-2b09-434b-ae58-592727cd0232', '2015-05-19 16:47:42', '127.0.0.1'), ('f76201b4-6f1c-40f9-bf10-083b4bd4eb55', '11', 'a75a538a-735b-4a1d-ae5f-7bd63134eed0', '2015-05-19 16:49:21', '127.0.0.1'), ('90654db6-a0d0-4e89-bd29-ec9301c7cd32', '11', '2a1d32d7-3d0f-497f-9ee7-1e9ab7461cf7', '2015-05-19 16:51:00', '127.0.0.1'), ('3cd9d6b9-a76c-42f1-8446-c826617e27bd', '11', '894f1e9b-cad2-4c02-801b-fe4191a1b0c8', '2015-05-19 16:53:54', '127.0.0.1'), ('833a2f89-b442-45ad-aff0-6cdc4c8b1d6c', '11', 'e9c2e0b9-dbf8-479f-937d-41b243796a6d', '2015-05-19 16:58:58', '127.0.0.1'), ('13196f30-000e-4abd-8364-078deb0b8c16', '11', 'd8aeec51-807f-4ac9-a22a-e69b299d77c1', '2015-05-19 17:00:22', '127.0.0.1'), ('3d7dffa1-060d-41d1-a512-c8ac6b42149b', '11', '28b5bf9e-b749-4a16-a408-531252c18022', '2015-05-19 17:00:52', '127.0.0.1'), ('fda4f2f3-2e4f-445b-9b5d-7d2ab67393c3', '11', 'c07be1d0-6b29-486c-b42e-a83ba6df9c8a', '2015-05-19 17:02:02', '127.0.0.1'), ('af45261e-c287-4898-8144-057b5da53524', '11', '6029a8b0-890a-4b0e-9a49-10867c497e5f', '2015-05-19 17:05:13', '127.0.0.1'), ('42cec161-aa35-46ae-bb96-0ea7def3a17f', '11', 'ef8ec53a-e976-4b63-96da-cb7290cec7bf', '2015-05-19 17:09:11', '127.0.0.1'), ('07db54b0-8316-40b8-a299-33e728e065f3', '11', '15c70b48-01c5-4153-8113-d0620aaf7079', '2015-05-19 17:10:01', '127.0.0.1'), ('ab357f71-5882-4b77-a497-2e28e44c1739', '11', 'ff3e4b36-c83e-4ec7-aeca-0376a35a42eb', '2015-05-19 17:10:28', '127.0.0.1'), ('7ed4c419-1efb-427e-86d8-a0d4efd96a9d', '22', 'e31e9998-748f-4851-bba5-e2545b9eac67', '2015-05-19 17:11:45', '127.0.0.1'), ('4e25facc-4612-4d10-94a6-4a3089a908fe', '11', '8e5f44be-a17f-4bab-a2b9-7116fe2aa2fe', '2015-05-19 17:13:54', '127.0.0.1'), ('b11ba3b6-05ea-4684-b175-9f331d83c470', '11', '5c357c60-b1e7-4798-93f7-02a5040ec3e5', '2015-05-19 17:15:09', '127.0.0.1'), ('c4450228-69bd-495e-99ee-68fdc3b45dce', '11', 'f61011d3-35ee-4c5e-a589-a6a54e9ea84e', '2015-05-19 17:25:50', '127.0.0.1'), ('2a881cc1-8276-4a1d-bd28-be43147cc2da', '11', 'cb9b8b98-98da-494c-8122-f7569b1b702a', '2015-05-19 17:26:34', '127.0.0.1'), ('241d2905-3e7a-4581-b1a3-539c80c99c2a', '11', 'e9cef4df-13f7-4a9e-a511-9338d1bcebeb', '2015-05-19 17:28:30', '127.0.0.1'), ('426f0318-a162-4ee4-9cac-c69f6eccad80', '11', '5404099e-dbba-4bba-a6be-de2ffeca154d', '2015-05-20 15:24:22', '127.0.0.1'), ('14814ea7-c67d-4184-939d-65969afe5cd6', '11', 'bbda33b8-f6da-475b-899c-12ce384cf79a', '2015-05-20 16:00:39', '127.0.0.1'), ('18341629-a27a-4fce-b9c8-6b0b0eeaaaee', '11', '724e09f5-19e9-4853-9212-aee054d2fc77', '2015-05-20 16:04:50', '127.0.0.1'), ('fa5ccb56-a6e7-45f5-ac60-9cdad8d9ba87', '11', 'a35a444a-321d-4f12-991b-f836bfb44548', '2015-05-20 16:07:52', '127.0.0.1'), ('a941a130-5630-4cf5-aa1a-e5ec0a6bee2c', '11', 'cde11dc4-9b9a-42c9-82bf-9a648ce87c42', '2015-05-20 16:09:37', '127.0.0.1'), ('f1223446-4a58-4c7b-991f-c689630813b1', '11', '6e8927f7-e94f-4de0-aeb3-4e02c759c56a', '2015-05-20 16:13:11', '127.0.0.1'), ('d3bdfe1d-5a37-4d46-bf52-8f9bd627e317', '11', '240b3554-ae2b-4e0d-a103-548594ec73b3', '2015-05-20 16:46:29', '127.0.0.1'), ('678728ea-4018-4619-b967-257b4625e87b', '11', '0a4acd4d-a9e1-4d27-bb4c-47184a081031', '2015-05-20 16:52:18', '127.0.0.1'), ('6b26e175-e65d-4ca7-9d29-2293d8e1149e', '11', 'af370609-6385-4486-bd26-3f7561087471', '2015-05-20 16:55:51', '127.0.0.1'), ('ee5482ae-fa92-46cc-889d-a8473f37ceb7', '11', '7743a73c-2147-4a45-b12f-d93ad86c30f1', '2015-05-20 17:00:57', '127.0.0.1'), ('10327eb5-d450-498b-b703-b7174658bfd0', '11', '519981fa-e858-4674-8911-2f3085f124f5', '2015-05-20 17:04:34', '127.0.0.1'), ('e4291aed-45e2-481e-814e-bf5cb031024e', '11', '9ae2b2bd-2f09-4c51-97f0-36fc85afa987', '2015-05-20 17:08:58', '127.0.0.1'), ('6773fe72-5bf9-40d0-9c0c-dc46bd1b06fc', '11', 'f0440849-7e32-4dd2-88b9-f2acab209f60', '2015-05-20 17:12:17', '127.0.0.1'), ('4abd20b1-5c4b-46cb-b12f-26cbbfcd73bb', '11', '7529a900-e6b0-4bd2-8748-4dcfd7ea4be8', '2015-05-20 17:16:22', '127.0.0.1'), ('18f8c6c2-3664-4ce6-8e10-adce3a2955a0', '11', '8593ce82-e6c3-437e-8bb9-91d8dff1bb60', '2015-05-20 17:17:13', '127.0.0.1'), ('6104f949-00d0-46ae-9a1c-a59025f9beb8', '11', 'd1f9ee92-1d93-4957-b41c-9235ace8fc67', '2015-05-20 17:19:16', '127.0.0.1'), ('9b83936f-78c6-438d-8403-a50b57688280', '11', 'cb00ffdc-baad-4818-984e-0d281be02a9b', '2015-05-20 17:21:05', '127.0.0.1'), ('4171837e-7ed5-43cf-a961-2611ddf75192', '11', '0d32bce0-0982-4aea-a4e6-8d749b9caf6e', '2015-05-20 17:24:21', '127.0.0.1'), ('c022cbc2-3862-4289-be74-aeaf2feabe9b', '11', '1cee6505-4e63-48b5-8625-03fbfa68a6ad', '2015-05-20 17:28:53', '127.0.0.1'), ('2a881dae-79ff-4b75-ba7c-75d0fac3fde5', '11', '5ab5fe9b-c677-4aa8-b2de-e8da46f048c9', '2015-05-20 19:11:18', '127.0.0.1'), ('5aa70bc3-e8da-4ddf-a557-fc3e6c36527b', '11', '5351a334-2a41-4e37-ae84-57ff3588afe5', '2015-05-20 19:12:56', '127.0.0.1'), ('9d64f097-622a-4479-906d-31cdd13280ef', '11', '31089b5e-54a8-4283-8ced-777504aec4e5', '2015-05-20 19:31:02', '127.0.0.1'), ('212f57e2-f476-4f70-a731-aef2e9e0906b', '11', 'a7675907-39a6-4160-a520-55f991cad662', '2015-05-20 19:32:19', '127.0.0.1'), ('701d0d25-065a-4b35-98da-a207991e91fd', '11', '1fbd764f-59e7-4fc4-8b3a-9632e946f753', '2015-05-20 19:33:54', '127.0.0.1'), ('cf63ed9a-cf5d-4284-bdda-691f1cc523a1', '11', '5064fed8-a0e6-4be2-b7dd-8d0d73dda6e7', '2015-05-20 19:42:32', '127.0.0.1'), ('c8450125-a974-4f58-b02e-8bfeb9814cdd', '11', 'c7c6f1b8-ec07-417d-b63a-28ef5d00bc57', '2015-05-20 19:49:21', '127.0.0.1'), ('ab9ac890-773b-4ce6-824b-7b6d52b9783e', '11', '88fd8f65-a8f1-44a0-8fc8-a671efbf1d5d', '2015-05-21 08:58:24', '127.0.0.1'), ('197b3d67-14de-42b8-bad2-ee0c30a7ba46', '22', '4ad08875-8e04-4aef-aff2-9eb3d46ddc55', '2015-05-21 08:59:12', '127.0.0.1'), ('af053f36-2251-4022-af7b-ca3079500146', '11', '83cd693a-1e01-46c2-aeb8-7fa339337202', '2015-05-21 09:02:46', '127.0.0.1'), ('9c7a6b18-90f5-4e77-b18a-68097617eec8', '11', '9b503490-353f-41c4-acf5-97468984a51f', '2015-05-21 09:17:24', '127.0.0.1'), ('35896b40-91a3-4c70-94b1-01d572753c8a', '22', '0e355c01-fb77-4eef-ae57-efd207b2f22d', '2015-05-21 09:17:36', '127.0.0.1'), ('d41e6c41-eb77-4935-b190-3a833c24c6ee', '11', 'b40cb976-73fc-432d-a972-1be45f87e18e', '2015-05-21 15:38:03', '127.0.0.1'), ('d4fbc61c-b56d-416c-ad66-7f5e8bc8383c', '11', '4436547b-4138-4456-b8dd-32c6b29e6c11', '2015-05-21 15:39:23', '127.0.0.1'), ('0e4af479-8dea-437c-9f92-f24c6be07be3', '11', 'fda02764-6546-48c8-b6ad-a5383e423910', '2015-05-21 15:43:57', '127.0.0.1'), ('53466dd3-c6fe-4fab-b849-286078ce57f1', '11', '48c8f71d-e656-4724-9b56-936886128856', '2015-05-21 15:54:48', '127.0.0.1'), ('ec3655d0-80e3-47dc-9c84-b8a96401330f', '11', '3bda2ac6-c691-4329-8cf1-aaa6b69266a9', '2015-05-21 15:56:48', '127.0.0.1'), ('a375fb2a-fe4f-4fab-948d-38256d4e65cc', '11', '599820a2-271e-4a45-995c-1592edca2475', '2015-05-21 15:58:00', '127.0.0.1'), ('b1b9b83a-aa18-4ad3-8a56-d05f9b814ba1', '11', '7612feef-a14d-4f4b-8989-816f0793f0dc', '2015-05-21 15:58:43', '127.0.0.1'), ('2fa93327-f384-4860-996a-5f107b82ab79', '11', '6be5e312-12d4-4acf-b1ec-409c65c917b8', '2015-05-21 16:00:48', '127.0.0.1'), ('b153bc75-90b8-4047-94f0-35a99e052b9e', '11', '0f966c80-c16a-4739-9785-45a27b105dc7', '2015-05-21 16:01:35', '127.0.0.1'), ('f76ec035-a6d7-4f35-b5a0-ad8f3cf18591', '11', '1ab6c440-9191-4a9e-adfd-e760005d289a', '2015-05-21 16:05:55', '127.0.0.1'), ('ed062914-30a3-4aa2-94d5-fcccc882aeb3', '11', '3fd83efa-fbf5-426a-bc0c-a4f0681bcb58', '2015-05-21 16:07:21', '127.0.0.1'), ('7d7a8697-797a-4e68-b621-46879874500f', '11', '7eec4518-ab04-4070-9d6d-dc7614d97f2a', '2015-05-21 16:17:44', '127.0.0.1'), ('cf801541-6f41-45c4-aded-823be829563e', '11', '49a34860-5888-4634-84cc-98b52f820aa9', '2015-05-21 16:21:47', '127.0.0.1'), ('1fe4ad4c-fef0-40f2-8639-136846d51f05', '11', 'aaa24d2f-b64d-4f68-946d-f20ca099badd', '2015-05-21 16:22:24', '127.0.0.1'), ('22aab85e-3e41-465f-8e70-fc84b04ceb38', '11', '77370e4b-0e0b-45de-8b5b-97f73cf20347', '2015-05-21 16:23:13', '127.0.0.1'), ('0ba14f2b-3cb1-47b4-8acc-5f356ee3c114', '11', '60fe8cce-6c01-481a-aeb3-cf0e45d46cb2', '2015-05-21 16:30:43', '127.0.0.1'), ('4b07dcca-7b12-463c-906d-5e142bad04dd', '11', '80f1d525-4d7b-41dd-b63d-1943388fac16', '2015-05-21 16:33:19', '127.0.0.1'), ('4ca69d47-8f67-4a05-b015-2f640a032be9', '11', '93e064f2-e442-4acd-8ee8-7101989e65a8', '2015-05-21 16:43:43', '127.0.0.1'), ('868a6808-42aa-442b-8b41-7550b479f829', '11', 'cc0225b5-41f0-4ec4-be6f-92ca0bcaa39f', '2015-05-21 16:46:53', '127.0.0.1'), ('79cfcf95-75fd-47f0-b18f-03bbfcd2d3ef', '11', 'b6dc90d4-1875-40e2-ba40-c1cfa50d368f', '2015-05-21 16:49:18', '127.0.0.1'), ('83ebfceb-b480-4a98-a215-1caab859ed25', '11', 'e1b2919a-bf57-4729-b589-6a6719d4202e', '2015-05-21 16:50:22', '127.0.0.1'), ('eb1b6822-fff2-46ab-bca2-ab53c34e6db0', '11', '56d2d4d0-67f9-4d67-acb8-8b033454878d', '2015-05-21 17:06:19', '127.0.0.1'), ('23374011-3b1b-429e-b8b1-6014f795ee67', '11', 'f798e3a7-298c-4d74-aa6d-c8d78100a7ee', '2015-05-21 17:07:27', '127.0.0.1'), ('f4c62bf0-b4dd-4b5c-9088-c42a400a46a9', '11', '389d3e44-9ae4-46b4-ab4b-6ab98f1f355c', '2015-05-21 17:08:42', '127.0.0.1'), ('adb2de44-557a-48b5-a950-48c48bc5509c', '11', 'fae315de-db71-4dd6-9b5f-071ad2472619', '2015-05-21 17:13:03', '127.0.0.1'), ('f6a64c55-f256-48d4-8c09-12edf8dfa125', '11', 'c807fd22-6675-4ac9-ac4c-295188e4ff20', '2015-05-21 17:14:33', '127.0.0.1'), ('429a7b8a-c65f-4b8e-a6d0-60e98d1df211', '11', '5d629485-f520-4976-ab0d-6be39607d9bd', '2015-05-21 17:17:21', '127.0.0.1'), ('2e77afd4-470f-4b5f-bb36-5d77d783ff89', '11', '1c7065ee-daee-495a-b079-0cfcfbbda907', '2015-05-21 17:18:26', '127.0.0.1'), ('fe7709f8-d848-42cf-9f2f-44db27df6777', '22', 'f7b7e396-837d-4588-ad42-d937b0437c94', '2015-05-21 17:23:39', '127.0.0.1'), ('f7e2b5a8-b36b-47ce-a0d1-31ab5d5cf81b', '11', 'bcf86975-7900-4cdf-8f46-ca6ff555f792', '2015-05-21 17:26:47', '127.0.0.1'), ('86c55d30-ed53-4829-8cf0-2af03c8df55b', '22', '7cb612ea-ac09-47da-a939-78286b3b930b', '2015-05-21 17:27:09', '127.0.0.1'), ('b6f62aff-fe68-4f36-a542-98f620b1e331', '11', 'c3c39c79-e4e1-4306-a89c-d1b6ca561d96', '2015-05-22 08:40:30', '127.0.0.1'), ('186b0afa-68ee-4555-98e7-fa1a8ec077c5', '22', 'b4b74256-d516-444e-8469-66b15a8c9fce', '2015-05-22 08:41:05', '127.0.0.1'), ('a53b8da7-72ee-43be-956e-c29f3895dc7c', '22', '9e2b2d31-416e-48ed-a9d8-2714726f7375', '2015-05-22 08:43:12', '127.0.0.1'), ('8eb85834-bfcf-4f7a-8de5-d6325a06ca30', '11', '9463713c-78d7-4081-8dc0-3151ba89f46d', '2015-05-22 08:47:23', '127.0.0.1'), ('aee51f29-a589-4953-ae9e-39377ff26b1a', '22', 'dbbe47a3-7165-454c-be72-d808e4b1b405', '2015-05-22 08:47:57', '127.0.0.1'), ('d3041bfb-4c90-4444-a827-73061d765634', '11', '58b5a1b6-6c8b-4b5d-8c73-15d5bb0b73e4', '2015-05-22 08:49:10', '127.0.0.1'), ('8304bac1-fd6c-46bd-9d32-aed636df2503', '22', '5be9dc18-f7b3-4d16-98be-857f0b1851b2', '2015-05-22 08:49:23', '127.0.0.1'), ('1e903525-da20-4de8-a421-f1250b7c434b', '22', 'e6929bac-31a2-452d-86ac-e7d775c6e9a4', '2015-05-22 08:49:55', '127.0.0.1');
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
