DROP TABLE IF EXISTS features CASCADE;
CREATE TABLE IF NOT EXISTS features (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

DROP TABLE IF EXISTS tags CASCADE;
CREATE TABLE IF NOT EXISTS tags (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

DROP TABLE IF EXISTS banners CASCADE;
CREATE TABLE IF NOT EXISTS banners (
   id SERIAL NOT NULL PRIMARY KEY,
   id_feature INT NOT NULL,
   created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
   is_active boolean DEFAULT true,
   FOREIGN KEY (id_feature) REFERENCES features(id)
);

DROP TABLE IF EXISTS versions CASCADE;
CREATE TABLE IF NOT EXISTS versions (
   id SERIAL NOT NULL PRIMARY KEY,
   id_banner INT NOT NULL,
   updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
   content jsonb DEFAULT '{}'::jsonb,
   FOREIGN KEY (id_banner) REFERENCES banners(id)
);

CREATE TABLE IF NOT EXISTS banner_tags (
    id_tag INT NOT NULL,
    id_banner INT NOT NULL,
    PRIMARY KEY (id_banner, id_tag),
    UNIQUE (id_banner, id_tag),
    FOREIGN KEY (id_banner) REFERENCES banners(id) ON DELETE CASCADE,
    FOREIGN KEY (id_tag) REFERENCES tags(id) ON DELETE CASCADE
);


DROP TABLE IF EXISTS profile_tags CASCADE;
CREATE TABLE IF NOT EXISTS profile_tags (
    id_tag INT NOT NULL,
    id_profile INT NOT NULL,
    PRIMARY KEY (id_profile, id_tag),
    UNIQUE (id_profile, id_tag),
    FOREIGN KEY (id_tag) REFERENCES tags(id)
);

INSERT INTO features(name) VALUES ('f1');
INSERT INTO features(name) VALUES ('f2');

INSERT INTO tags(name) VALUES ('t1');
INSERT INTO tags(name) VALUES ('t2');

INSERT INTO banners(id_feature, is_active) VALUES (1, true);
INSERT INTO banners(id_feature, is_active) VALUES (2, true);

INSERT INTO banner_tags(id_tag, id_banner) VALUES (1, 1);
INSERT INTO banner_tags(id_tag, id_banner) VALUES (2, 1);
INSERT INTO banner_tags(id_tag, id_banner) VALUES (1, 2);

INSERT INTO profile_tags(id_tag, id_profile) VALUES (1, 1);