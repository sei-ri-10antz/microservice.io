CREATE TABLE IF NOT EXISTS `event_log` (
  `id` varchar(36) NOT NULL,
  `version` int(11) NOT NULL,
  `type` varchar(32) NOT NULL,
  `dump` blob,
  `created_at` datetime NOT NULL,
  PRIMARY KEY (`id`, `version`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
