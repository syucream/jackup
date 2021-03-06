CREATE DATABASE examples;

CREATE TABLE Singers (
  SingerId   INT64 NOT NULL,
  FirstName  STRING(1024),
  LastName   STRING(1024),
  SingerInfo BYTES(MAX),
  BirthDate  DATE,
) PRIMARY KEY (SingerId);

CREATE INDEX AlbumsByAlbumTitle ON Albums(AlbumTitle);
