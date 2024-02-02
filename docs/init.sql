INSERT INTO computeshare.compute_specs (id, core, memory) VALUES (1, 2, 4);
INSERT INTO computeshare.compute_specs (id, core, memory) VALUES (2, 16, 32);
INSERT INTO computeshare.compute_specs (id, core, memory) VALUES (3, 32, 64);

INSERT INTO computeshare.compute_spec_prices (id, fk_compute_spec_id, day, price) VALUES (1, 1, 30, 50000);
INSERT INTO computeshare.compute_spec_prices (id, fk_compute_spec_id, day, price) VALUES (2, 2, 30, 100000);
INSERT INTO computeshare.compute_spec_prices (id, fk_compute_spec_id, day, price) VALUES (3, 3, 30, 400000);

INSERT INTO computeshare.compute_images (id, name, image, tag, port, command) VALUES (1, 'Ubuntu 20.04', 'ubuntu', '20.04', 22, 'tail -f /var/log/bootstrap.log');
INSERT INTO computeshare.compute_images (id, name, image, tag, port, command) VALUES (2, 'WindowsServer 2016', 'WindowsServer', '2016', 3389, ' ');
INSERT INTO computeshare.compute_images (id, name, image, tag, port, command) VALUES (3, 'Centos 7', 'Centos', '7', 22, ' ');

INSERT INTO computeshare.cycle_redeem_codes (id, fk_user_id, redeem_code, cycle, state, create_time, use_time) VALUES ('47c2a666-8100-b4cd-d89c-2bca0ae5102b', null, '4fBqexRW', 1000.00, 0, '2024-01-31 17:01:29', null);
INSERT INTO computeshare.cycle_redeem_codes (id, fk_user_id, redeem_code, cycle, state, create_time, use_time) VALUES ('59235b80-edda-43d3-517e-74dcdad1b9ed', null, 'nWNpq5vX', 1000.00, 0, '2024-01-31 17:01:29', null);
INSERT INTO computeshare.cycle_redeem_codes (id, fk_user_id, redeem_code, cycle, state, create_time, use_time) VALUES ('7ece24e2-a218-f60d-7403-0e1ef0de27b6', null, '5XxeXQdm', 1000.00, 0, '2024-01-31 17:01:29', null);
INSERT INTO computeshare.cycle_redeem_codes (id, fk_user_id, redeem_code, cycle, state, create_time, use_time) VALUES ('8218e83c-63a3-830a-1df6-cdcc846b4378', null, 'uVUgkCAM', 1000000.00, 0, '2024-01-31 17:01:29', null);
INSERT INTO computeshare.cycle_redeem_codes (id, fk_user_id, redeem_code, cycle, state, create_time, use_time) VALUES ('95e87a68-ebd9-f775-ab4b-48cf9f179479', null, 'uUo1eAqo', 1000.00, 0, '2024-01-31 17:01:29', null);
INSERT INTO computeshare.cycle_redeem_codes (id, fk_user_id, redeem_code, cycle, state, create_time, use_time) VALUES ('c1785ebf-05b2-7535-dd22-ffd338b7758f', null, 'CPPuEJ9M', 10000.00, 0, '2024-01-31 17:01:29', null);
INSERT INTO computeshare.cycle_redeem_codes (id, fk_user_id, redeem_code, cycle, state, create_time, use_time) VALUES ('c38a6c52-d92f-3474-89a0-673ea5fb8b03', null, '1fXd6OdI', 500000.00, 0, '2024-01-31 17:01:29', null);
INSERT INTO computeshare.cycle_redeem_codes (id, fk_user_id, redeem_code, cycle, state, create_time, use_time) VALUES ('cab9ba0e-a237-a002-9c07-f1109b2ea9f8', null, 'i5XfBYlx', 500000.00, 0, '2024-01-31 17:01:29', null);
INSERT INTO computeshare.cycle_redeem_codes (id, fk_user_id, redeem_code, cycle, state, create_time, use_time) VALUES ('d793b92e-3b0a-b665-88a9-f6e5ff3763a5', null, 'G37pO5kk', 500000.00, 0, '2024-01-31 17:01:29', null);
INSERT INTO computeshare.cycle_redeem_codes (id, fk_user_id, redeem_code, cycle, state, create_time, use_time) VALUES ('e76d9f0b-b6dd-ce3c-a5cb-a8b06ae09f47', null, 'EQDgZyTx', 500000.00, 0, '2024-01-31 17:01:29', null);

INSERT INTO computeshare.gateways (id, name, ip, port, internal_ip) VALUES ('db3b9c89-2eea-4ee0-a567-1f5f88d334f2', 'G1TEST', '61.172.179.6', 37000, '192.168.0.2');
