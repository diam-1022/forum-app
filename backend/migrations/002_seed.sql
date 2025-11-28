-- 板块数据
INSERT INTO boards (name, description)
VALUES
    ('公告栏', '官方公告与规则（禁止吵架）'),
    ('技术讨论', '爱音的吉他是不是live两天前速成的'),
    ('灌水栏', '母鸡卡真的会有好结果吗')
ON DUPLICATE KEY UPDATE description = VALUES(description);

-- 话题数据
INSERT INTO topics (title, description)
VALUES
    ('Mujica', '祥子的狱卒和蜜雪冰城的糯香哪个更符合皱皮的口味'),
    ('MyGo', '咕咕嘎嘎旮旯game里面根本不是这样的！'),
    ('Roselia', '本格搞笑乐队')
ON DUPLICATE KEY UPDATE description = VALUES(description);