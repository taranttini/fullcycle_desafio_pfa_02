
CREATE TABLE `fullcycledb`.`modules` (
    `moduleId` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `active` BOOLEAN NOT NULL,
    PRIMARY KEY (`moduleId`)
);

INSERT INTO `modules` (`name`,`active`) VALUES
    ('Docker', 1),
    ('Padrões e técnicas avançadas com Git e Github', 1),
    ('Integração contínua', 1),
    ('Kubernetes', 1),
    ('Service Mesh com Istio', 0),
    ('Observabilidade', 0),
    ('Deploy nos Cloud Providers', 0),
    ('Fundamentos de Arquitetura de Software', 1),
    ('Comunicação', 1),
    ('RabbitMQ', 1),
    ('Autenticação e Keycloak', 1),
    ('Domain Driven Design e Arquitetura hexagonal', 1),
    ('Arquitetura do projeto prático - Codeflix', 1),
    ('Microsserviço: Catálogo de vídeos com Laravel ( Back-end )', 1),
    ('Microsserviço: Catálogo de vídeos com React ( Front-end )', 1),
    ('Microsserviço de Encoder de Vídeo com Go Lang', 1),
    ('Microsserviço - API do Catálogo com Node.JS (Back-end)', 1),
    ('Microsserviço - Aplicação do Assinante com React.js (Front-end)', 0),
    ('Microsserviço - Assinaturas com Django (Back-end)', 0),
    ('Apache Kafka', 0),
    ('Service Discovery', 0);

-- SELECT * FROM fullcycledb.modules;
-- SELECT * FROM fullcycledb.modules WHERE active = 1;
-- SELECT * FROM fullcycledb.modules WHERE active = 0;