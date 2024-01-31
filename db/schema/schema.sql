CREATE TABLE User (
    UserID STRING(MAX) NOT NULL,
    Name STRING(MAX) NOT NULL,
    Email STRING(MAX) NOT NULL,
    Password STRING(MAX) NOT NULL,
    CreatedAt TIMESTAMP NOT NULL,
    UpdatedAt TIMESTAMP NOT NULL,
) PRIMARY KEY (UserID);
CREATE INDEX User_Email ON User(Email);