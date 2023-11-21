INSERT INTO
    gateways(ip, port, name, id)
VALUES (
        '61.172.179.73',
        7000,
        'G1',
        'db3bec89-2eea-4ae0-a567-1b5f88d334f2'
    );

INSERT INTO
    gateway_ports(fk_gateway_id, is_use, id, port)
VALUES (
        'db3bec89-2eea-4ae0-a567-1b5f88d334f2',
        false,
        'afd38d3c-7732-4d4d-85c4-6e565dff87cc',
        40001
    );

INSERT INTO
    gateway_ports(fk_gateway_id, is_use, id, port)
VALUES (
        'db3bec89-2eea-4ae0-a567-1b5f88d334f2',
        false,
        'a5979391-adbd-432e-aa10-628dfe886ad0',
        40002
    );

INSERT INTO
    gateway_ports(fk_gateway_id, is_use, id, port)
VALUES (
        'db3bec89-2eea-4ae0-a567-1b5f88d334f2',
        false,
        '2fbb5386-bb17-433d-bb7a-861856be91b9',
        40003
    );

INSERT INTO
    gateway_ports(fk_gateway_id, is_use, id, port)
VALUES (
        'db3bec89-2eea-4ae0-a567-1b5f88d334f2',
        false,
        'bac2faa6-db3f-41fb-a2c6-3c0b424f38aa',
        40004
    );

INSERT INTO
    gateway_ports(fk_gateway_id, is_use, id, port)
VALUES (
        'db3bec89-2eea-4ae0-a567-1b5f88d334f2',
        false,
        '5740cf30-d372-4781-bfa8-abfe7692f0c1',
        40005
    );