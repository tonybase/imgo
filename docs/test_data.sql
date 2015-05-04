/*
-- Query: SELECT * FROM im.im_user
LIMIT 0, 1000

-- Date: 2015-05-04 22:17
*/
INSERT INTO `im_user` (`id`,`account`,`password`,`nick`,`sign`,`avatar`,`status`,`create_at`,`update_at`,`remark`) VALUES ('c70a9fb3-5eaf-432a-81df-010d1318335d','22','22','22','','','0','2015-05-03 15:17:25','2015-05-03 15:17:25',NULL);
INSERT INTO `im_user` (`id`,`account`,`password`,`nick`,`sign`,`avatar`,`status`,`create_at`,`update_at`,`remark`) VALUES ('b26f11c4-a49e-47a3-8d88-f827e708a7a8','11','11','11','','','0','2015-05-03 15:17:30','2015-05-03 15:17:30',NULL);

/*
-- Query: SELECT * FROM im.im_category
LIMIT 0, 1000

-- Date: 2015-05-04 22:18
*/
INSERT INTO `im_category` (`id`,`name`,`creator`,`create_at`) VALUES ('59d7ec9a-752f-41e5-a63f-e025efc9a4e5','我的好友','c70a9fb3-5eaf-432a-81df-010d1318335d','2015-05-03 15:17:25');
INSERT INTO `im_category` (`id`,`name`,`creator`,`create_at`) VALUES ('1330d92b-f20e-4171-a2ca-26dde50199f5','我的好友','b26f11c4-a49e-47a3-8d88-f827e708a7a8','2015-05-03 15:17:30');

/*
-- Query: SELECT * FROM im.im_relation_user_category
LIMIT 0, 1000

-- Date: 2015-05-04 22:18
*/
INSERT INTO `im_relation_user_category` (`user_id`,`category_id`,`create_at`) VALUES ('b26f11c4-a49e-47a3-8d88-f827e708a7a8','59d7ec9a-752f-41e5-a63f-e025efc9a4e5','2015-05-03 15:17:58');
INSERT INTO `im_relation_user_category` (`user_id`,`category_id`,`create_at`) VALUES ('c70a9fb3-5eaf-432a-81df-010d1318335d','1330d92b-f20e-4171-a2ca-26dde50199f5','2015-05-03 15:18:14');
