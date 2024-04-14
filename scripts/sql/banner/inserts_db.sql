INSERT INTO banners(id_feature, title, text, url) VALUES (1, 'title1', 'text1', 'url1');
INSERT INTO banners(id_feature, title, text, url) VALUES (2, 'title2', 'text2', 'url2');

INSERT INTO features(name) VALUES ('f1');
INSERT INTO features(name) VALUES ('f2');

INSERT INTO tags(name) VALUES ('t1');
INSERT INTO tags(name) VALUES ('t2');

INSERT INTO banner_tags(id_tag, id_banner) VALUES (1, 1);
INSERT INTO banner_tags(id_tag, id_banner) VALUES (2, 1);
INSERT INTO banner_tags(id_tag, id_banner) VALUES (1, 2);

INSERT INTO profile_tags(id_tag, id_profile) VALUES (1, 1);

-- INSERT INTO profile(login, password, role) VALUES ('admin', '0xC7AD44CBAD762A5DA0A452F9E854FDC1E0E7A52A38015F23F3EAB1D80B931DD472634DFAC71CD34EBC35D16AB7FB8A90C81F975113D6C7538DC69DD8DE9077EC', 'admin')