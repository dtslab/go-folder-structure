package config

// code to load data from .env file
// and from here exports these .env data to be used across the app

var JWTKey string // code to load MY_JWT_SECRET_KEY from .env
									//and the to use it jusst use config.JWTKey