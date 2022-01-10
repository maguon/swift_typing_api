SET NAMES utf8mb4;

SET FOREIGN_KEY_CHECKS=0;
-- CREATE FUNCTION UPDATE_TIMESTAMP_FUNC

create or replace function update_timestamp_func() returns trigger as
$$
begin
  new.updated_on = current_timestamp;
  return new;
end
$$
language plpgsql;

--CREATE TABLE user_info
CREATE TABLE IF NOT EXISTS public.user_info
(
    id bigserial NOT NULL,
    created_on timestamp with time zone NOT NULL DEFAULT NOW(),
    updated_on timestamp with time zone NOT NULL DEFAULT NOW(),
    status smallint NOT NULL,
    user_name character varying(50),
    real_name character varying(50),
    password character varying(50) NOT NULL,
    phone character varying(50) NOT NULL,
    email character varying(100),
    gender smallint,
    type integer NOT NULL,
    PRIMARY KEY (id)
);

COMMENT ON COLUMN public.user_info.status IS '状态（0-停用，1-可用）';
COMMENT ON COLUMN public.user_info.user_name IS '用户名称';
COMMENT ON COLUMN public.user_info.real_name IS '真实姓名';
COMMENT ON COLUMN public.user_info.password IS '密码';
COMMENT ON COLUMN public.user_info.phone IS '联系方式';
COMMENT ON COLUMN public.user_info.email IS '邮箱';
COMMENT ON COLUMN public.user_info.gender IS '性别（0-女，1-男）';
COMMENT ON COLUMN public.user_info.type IS '用户类型（99-超级管理员）';

create trigger user_info_upt before update on user_info for each row execute procedure update_timestamp_func();
select setval('user_info_id_seq',1000,false);

CREATE UNIQUE INDEX uk_user_info_phone ON user_info(phone);


--CREATE TABLE app_info
CREATE TABLE IF NOT EXISTS public.app_info
(
    id smallserial NOT NULL,
    created_on timestamp with time zone NOT NULL DEFAULT NOW(),
    updated_on timestamp with time zone NOT NULL DEFAULT NOW(),
    status smallint NOT NULL,
    app_type smallint NOT NULL,
    device_type smallint NOT NULL,
    version character varying(20) NOT NULL,
    version_num smallint NOT NULL,
    min_version_num smallint NOT NULL,
    force_update smallint NOT NULL,
    url character varying(200) NOT NULL,
    remarks character varying(400),
    PRIMARY KEY (id)
);

COMMENT ON COLUMN public.app_info.status IS '状态(0:停用,1:启用)';
COMMENT ON COLUMN public.app_info.app_type IS 'app类型';
COMMENT ON COLUMN public.app_info.device_type IS '设备类型(1-安卓 2-苹果)';
COMMENT ON COLUMN public.app_info.version IS '版本号';
COMMENT ON COLUMN public.app_info.version_num IS '版本序号';
COMMENT ON COLUMN public.app_info.min_version_num IS '最小支持版本';
COMMENT ON COLUMN public.app_info.force_update IS '是否强制更新(0-不更新 1-更新)';
COMMENT ON COLUMN public.app_info.url IS '下载地址';
COMMENT ON COLUMN public.app_info.remarks IS '备注';

create trigger app_info_upt before update on app_info for each row execute procedure update_timestamp_func();

--CREATE TABLE user_device_info
CREATE TABLE IF NOT EXISTS public.user_device_info
(
    id bigserial NOT NULL,
    created_on timestamp with time zone NOT NULL DEFAULT NOW(),
    updated_on timestamp with time zone NOT NULL DEFAULT NOW(),
    status smallint NOT NULL,
    user_id bigint DEFAULT 0  ,
    device_type smallint NOT NULL,
    device_id character varying(60) NOT NULL,
    device_token character varying(60) ,
    app_type smallint NOT NULL,
    app_version character varying(20) NOT NULL,
    app_version_num smallint NOT NULL,
    PRIMARY KEY (id)
);

COMMENT ON COLUMN public.user_device_info.status IS '状态(0:停用,1:启用)';
COMMENT ON COLUMN public.user_device_info.app_type IS 'app类型';
COMMENT ON COLUMN public.user_device_info.device_type IS '设备类型(1-安卓 2-苹果)';
COMMENT ON COLUMN public.user_device_info.app_version IS '版本号';
COMMENT ON COLUMN public.user_device_info.app_version_num IS '版本序号';
COMMENT ON COLUMN public.user_device_info.remarks IS '备注';

create trigger user_device_info_upt before update on user_device_info for each row execute procedure update_timestamp_func();

--CREATE TABLE t_word
CREATE TABLE IF NOT EXISTS public.t_word
(
    id bigserial NOT NULL,
    created_on timestamp with time zone NOT NULL DEFAULT NOW(),
    updated_on timestamp with time zone NOT NULL DEFAULT NOW(),
    status smallint NOT NULL DEFAULT 1,
    word character varying(20) NOT NULL,
    spell character varying(60) NOT NULL,
    explain character varying(300) ,
    example character varying(300) ,
    refere character varying(100) ,
    level smallint NOT NULL DEFAULT 1,
    PRIMARY KEY (id)
);

create trigger t_word_upt before update on t_word for each row execute procedure update_timestamp_func();



--CREATE TABLE t_word
CREATE TABLE IF NOT EXISTS public.t_poem
(
    id bigserial NOT NULL,
    created_on timestamp with time zone NOT NULL DEFAULT NOW(),
    updated_on timestamp with time zone NOT NULL DEFAULT NOW(),
    status smallint NOT NULL DEFAULT 1,
    author character varying(20) NOT NULL,
    title character varying(60) NOT NULL,
    content character varying(300) ,
    level smallint NOT NULL DEFAULT 1,
    PRIMARY KEY (id)
);

create trigger t_poem_upt before update on t_poem for each row execute procedure update_timestamp_func();

--CREATE TABLE t_sentence
CREATE TABLE IF NOT EXISTS public.t_sentence
(
    id bigserial NOT NULL,
    created_on timestamp with time zone NOT NULL DEFAULT NOW(),
    updated_on timestamp with time zone NOT NULL DEFAULT NOW(),
    status smallint NOT NULL DEFAULT 1,
    author character varying(20) NOT NULL,
    content character varying(300) ,
    PRIMARY KEY (id)
);

create trigger t_sentence_upt before update on t_sentence for each row execute procedure update_timestamp_func();

--CREATE TABLE t_article
CREATE TABLE IF NOT EXISTS public.t_article
(
    id bigserial NOT NULL,
    created_on timestamp with time zone NOT NULL DEFAULT NOW(),
    updated_on timestamp with time zone NOT NULL DEFAULT NOW(),
    status smallint NOT NULL DEFAULT 1,
    author character varying(20) NOT NULL,
    title character varying(60) NOT NULL,
    content character varying(3000) ,
    refere character varying(100) ,
    level smallint NOT NULL DEFAULT 1,
    PRIMARY KEY (id)
);

create trigger t_article_upt before update on t_article for each row execute procedure update_timestamp_func();

--CREATE TABLE t_score
CREATE TABLE IF NOT EXISTS public.t_score
(
    id bigserial NOT NULL,
    created_on timestamp with time zone NOT NULL DEFAULT NOW(),
    updated_on timestamp with time zone NOT NULL DEFAULT NOW(),
    status smallint NOT NULL DEFAULT 1,
    user_id bigint NOT NULL,
    corrent int NOT NULL DEFAULT 0,
    incorrect int NOT NULL DEFAULT 0,
    score decimal(12,2) NOT NULL DEFAULT 0,
    score_type smallint NOT NULL DEFAULT 1,
    level smallint NOT NULL DEFAULT 1,
    PRIMARY KEY (id)
);

create trigger t_score_upt before update on t_score for each row execute procedure update_timestamp_func();
