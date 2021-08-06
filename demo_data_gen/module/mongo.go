package module

type MongoDataGenConfig struct {
	Host              string
	Database          string
	Collection        string
	MongoFieldConfigs []MongoFieldConfig
}

type MongoFieldConfig struct {
	FieldName string
	FiledType string
}
