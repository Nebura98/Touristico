CREATE TABLE `Attraction` (
	`Id` BINARY(16) NOT NULL PRIMARY KEY,
	`Name` VARCHAR(255) NOT NULL,
	`Description` TEXT NOT NULL,
	`Schedule` JSON,
	`Type` VARCHAR(255) NOT NULL,
	`Likes` INT DEFAULT 0,
	`IsDeleted` BOOLEAN DEFAULT FALSE,
	`CreatedAt` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	`UpdatedAt` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`DeletedAt` TIMESTAMP NULL DEFAULT NULL,
	`Enabled` BOOLEAN DEFAULT TRUE,
	CHECK ( 
		JSON_SCHEMA_VALID(
			'{
  				"id": "http://json-schema.org/geo",
  				"$schema": "http://json-schema.org/draft-04/schema#",
  				"description": "Schedule for the attraction",
  				"type": "object",
  				"properties": {
    				"From": {
      					"type": "number"
    				},
    				"To": {
      					"type": "number"
    				},
    				"Weekday": {
      					"type": "string"
    				}
  				},
				"additionalProperties": false,
				"required": [ "From", "To", "Weekday" ]
			}',
			`Schedule`
		) 
	)
);


CREATE TABLE `Company` (
	`Id` BINARY(16) NOT NULL ,
	`Name` VARCHAR(255) NOT NULL,
	`Address` VARCHAR(255) NOT NULL,
	`Schedule` JSON,
	`IsDeleted` BOOLEAN DEFAULT FALSE,
	`CreatedAt` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	`UpdatedAt` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`DeletedAt` TIMESTAMP NULL DEFAULT NULL,
	`Enabled` BOOLEAN DEFAULT TRUE,
	CHECK ( 
		JSON_SCHEMA_VALID(
			'{
				"id": "http://json-schema.org/geo",
				"$schema": "http://json-schema.org/draft-04/schema#",
				"description": "Schedule for the attraction",
				"type": "object",
				"properties": {
					"From": {
						"type": "number"
					},
					"To": {
						"type": "number"
					},
					"Weekday": {
						"type": "string"
					}
				},
			"additionalProperties": false,
			"required": [ "From", "To", "Weekday" ],
			}',
			`Schedule`
		) 
	)
);

CREATE TABLE `Address` (
	`Id` BINARY(16) NOT NULL PRIMARY KEY,
	`Street` VARCHAR (255) NULL,
	`City` VARCHAR (255) NULL,
	`Province` VARCHAR (255) NULL,
	`Postal_code` VARCHAR (255) NULL,
	`Country` VARCHAR (255) NULL,
	`ISO_Country` VARCHAR (2) NULL,
	`Latitude` FLOAT NULL,
	`Longitude` FLOAT NULL
);
 
CREATE TABLE `User`(
	`Id` BINARY(16) NOT NULL PRIMARY KEY,
	`Username` VARCHAR(255) NOT NULL,
	`Fullname` VARCHAR(255) NOT NULL,
	`Email` VARCHAR(255) NOT NULL,
	`Password` VARCHAR(255) NOT NULL,
	`IsDeleted` BOOLEAN DEFAULT FALSE,
	`CreatedAt` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	`UpdatedAt` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`DeletedAt` TIMESTAMP NULL DEFAULT NULL,
	`Enabled` BOOLEAN DEFAULT TRUE,
	UNIQUE(`Email`),
	UNIQUE(`Username`)
);

CREATE TABLE `Post` (
	`Id` BINARY(16) NOT NULL PRIMARY KEY,
	`Title` VARCHAR(255) NOT NULL,
	`Post_content` JSON,
	`Likes` INT DEFAULT 0,
	`IsDeleted` BOOLEAN DEFAULT FALSE,
	`CreatedAt` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	`UpdatedAt` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`DeletedAt` TIMESTAMP NULL DEFAULT NULL,
	`Enabled` BOOLEAN DEFAULT TRUE,
	`FkUserId` BINARY(16), 
	CHECK ( 
		JSON_SCHEMA_VALID(
			'{
				"$schema": "http://json-schema.org/draft-07/schema#",
				"type": "object",
				"description": "Content type around the Post",
				"properties": {
					"post-content": {
						"type": "array",
						"items": {
							"type": "object",
							"properties": {
								"content-type": {
									"type": "string",
									"enum": [
										"image",
										"video"
									]
								},
								"url": {
									"type": "string"
								}
							},
							"required": [
								"content-type",
								"url"
							],
							"additionalProperties": false
						}
					}
				}
			}',
			`Post_content`
		) 
	),
	FOREIGN KEY (FkUserId) REFERENCES User(Id)
);

CREATE TABLE `Reviews`(
	`Id` BINARY(16) NOT NULL PRIMARY KEY,
	`Title` VARCHAR(255) NOT NULL,
	`Score` TINYINT NOT NULL,
	`Description` VARCHAR (255) NOT NULL,
	`Likes` INT DEFAULT 0 NOT NULL,
	`IsDeleted` BOOLEAN DEFAULT FALSE,
	`CreatedAt` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	`UpdatedAt` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`DeletedAt` TIMESTAMP NULL DEFAULT NULL,
	`Enabled` BOOLEAN DEFAULT TRUE,
	FOREIGN KEY (FkUserId) REFERENCES User(Id),
	FOREIGN KEY (FkAttractionId) REFERENCES Attraction(Id)
);

CREATE TABLE `Post_comment`(
	`Id` BINARY(16) NOT NULL PRIMARY KEY,
	`Comment` VARCHAR (255) NOT NULL,
	`Likes` INT DEFAULT 0 NOT NULL,
	`IsDeleted` BOOLEAN DEFAULT FALSE,
	`CreatedAt` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	`UpdatedAt` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`DeletedAt` TIMESTAMP NULL DEFAULT NULL,
	`Enabled` BOOLEAN DEFAULT TRUE,
	FOREIGN KEY (FkUserId) REFERENCES User(Id)
)