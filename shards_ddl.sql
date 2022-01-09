START TRANSACTION;

CREATE TABLE `messages` (
  `id` varchar (36) NOT NULL,
  `shard_id` varchar (36) NOT NULL,
  `user_id_from` bigint(20) NOT NULL,
  `user_id_to` bigint(20) NOT NULL,
  `message` text NOT NULL,
  `created_at` timestamp NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

ALTER TABLE `messages`   ADD PRIMARY KEY (`id`);



COMMIT;