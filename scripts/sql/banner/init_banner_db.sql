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
   updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
   content jsonb DEFAULT '{}'::jsonb,
   is_active boolean DEFAULT true,
   FOREIGN KEY (id_feature) REFERENCES features(id)
);

CREATE TABLE IF NOT EXISTS banner_tags (
   id_tag INT NOT NULL,
   id_banner INT NOT NULL,
   PRIMARY KEY (id_banner, id_tag),
    UNIQUE (id_banner, id_tag),
    FOREIGN KEY (id_banner) REFERENCES banners(id) ON DELETE CASCADE,
    FOREIGN KEY (id_tag) REFERENCES tags(id) ON DELETE CASCADE,
    );


DROP TABLE IF EXISTS profile_tags CASCADE;
CREATE TABLE IF NOT EXISTS profile_tags (
    id_tag INT NOT NULL,
    id_profile INT NOT NULL,
    PRIMARY KEY (id_profile, id_tag),
    UNIQUE (id_profile, id_tag),
    FOREIGN KEY (id_tag) REFERENCES tags(id)
);