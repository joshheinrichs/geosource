DROP TYPE IF EXISTS visibility CASCADE;
DROP DOMAIN IF EXISTS username CASCADE;
DROP DOMAIN IF EXISTS avatar CASCADE;
DROP DOMAIN IF EXISTS email CASCADE;
DROP DOMAIN IF EXISTS channelname CASCADE;
DROP DOMAIN IF EXISTS title CASCADE;
DROP DOMAIN IF EXISTS comment CASCADE;
DROP DOMAIN IF EXISTS base64uuid CASCADE;
DROP DOMAIN IF EXISTS userid CASCADE;
DROP DOMAIN IF EXISTS postid CASCADE;
DROP DOMAIN IF EXISTS commentid CASCADE;
DROP DOMAIN IF EXISTS thumbnail CASCADE;

DROP TABLE IF EXISTS users CASCADE;
DROP INDEX IF EXISTS lower_username_idx CASCADE;
DROP TABLE IF EXISTS admins CASCADE;
DROP TABLE IF EXISTS channels CASCADE;
DROP TABLE IF EXISTS posts CASCADE;
DROP TABLE IF EXISTS comments CASCADE;
DROP INDEX IF EXISTS post_gix CASCADE;
DROP TABLE IF EXISTS user_favorites CASCADE;
DROP TABLE IF EXISTS user_subscriptions CASCADE;
DROP TABLE IF EXISTS channel_moderators CASCADE;
DROP TABLE IF EXISTS channel_viewers CASCADE;
DROP TABLE IF EXISTS channel_bans CASCADE;
