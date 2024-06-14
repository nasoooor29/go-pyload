package models

type PluginConfigModel map[string]ConfigModelBlock
type PluginConfigDictModel struct {
	Download struct {
		Chunks       ConfigDictModelIntItem    `json:"chunks"`
		Desc         string                    `json:"desc"`
		EndTime      ConfigDictModelStringItem `json:"end_time"`
		Interface    ConfigDictModelStringItem `json:"interface"`
		Ipv6         ConfigDictModelBoolItem   `json:"ipv6"`
		LimitSpeed   ConfigDictModelBoolItem   `json:"limit_speed"`
		MaxDownloads ConfigDictModelIntItem    `json:"max_downloads"`
		MaxSpeed     ConfigDictModelIntItem    `json:"max_speed"`
		SkipExisting ConfigDictModelBoolItem   `json:"skip_existing"`
		StartTime    ConfigDictModelStringItem `json:"start_time"`
	} `json:"download"`
	General struct {
		DebugLevel       ConfigItemListModel       `json:"debug_level"`
		DebugMode        ConfigDictModelBoolItem   `json:"debug_mode"`
		Desc             string                    `json:"desc"`
		FolderPerPackage ConfigDictModelBoolItem   `json:"folder_per_package"`
		Language         ConfigItemListModel       `json:"language"`
		MinFreeSpace     ConfigDictModelIntItem    `json:"min_free_space"`
		SslVerify        ConfigDictModelBoolItem   `json:"ssl_verify"`
		StorageFolder    ConfigDictModelStringItem `json:"storage_folder"`
	} `json:"general"`
	Log struct {
		Console        ConfigDictModelBoolItem   `json:"console"`
		ConsoleColor   ConfigDictModelBoolItem   `json:"console_color"`
		Desc           string                    `json:"desc"`
		Filelog        ConfigDictModelBoolItem   `json:"filelog"`
		FilelogEntries ConfigDictModelIntItem    `json:"filelog_entries"`
		FilelogFolder  ConfigDictModelStringItem `json:"filelog_folder"`
		FilelogRotate  ConfigDictModelBoolItem   `json:"filelog_rotate"`
		FilelogSize    ConfigDictModelIntItem    `json:"filelog_size"`
		Syslog         ConfigDictModelBoolItem   `json:"syslog"`
		SyslogFolder   ConfigDictModelStringItem `json:"syslog_folder"`
		SyslogHost     ConfigDictModelStringItem `json:"syslog_host"`
		SyslogLocation ConfigItemListModel       `json:"syslog_location"`
		SyslogPort     ConfigDictModelIntItem    `json:"syslog_port"`
	} `json:"log"`
	Permission struct {
		ChangeDl    ConfigDictModelBoolItem   `json:"change_dl"`
		ChangeFile  ConfigDictModelBoolItem   `json:"change_file"`
		ChangeGroup ConfigDictModelBoolItem   `json:"change_group"`
		ChangeUser  ConfigDictModelBoolItem   `json:"change_user"`
		Desc        string                    `json:"desc"`
		File        ConfigDictModelStringItem `json:"file"`
		Folder      ConfigDictModelStringItem `json:"folder"`
		Group       ConfigDictModelStringItem `json:"group"`
		User        ConfigDictModelStringItem `json:"user"`
	} `json:"permission"`
	Proxy struct {
		Desc            string                    `json:"desc"`
		Enabled         ConfigDictModelBoolItem   `json:"enabled"`
		Host            ConfigDictModelStringItem `json:"host"`
		Password        ConfigDictModelStringItem `json:"password"`
		Port            ConfigDictModelIntItem    `json:"port"`
		SocksResolveDNS ConfigDictModelBoolItem   `json:"socks_resolve_dns"`
		Type            ConfigItemListModel       `json:"type"`
		Username        ConfigDictModelStringItem `json:"username"`
	} `json:"proxy"`
	Reconnect struct {
		Desc      string                    `json:"desc"`
		Enabled   ConfigDictModelBoolItem   `json:"enabled"`
		EndTime   ConfigDictModelStringItem `json:"end_time"`
		Script    ConfigDictModelStringItem `json:"script"`
		StartTime ConfigDictModelStringItem `json:"start_time"`
	} `json:"reconnect"`
	Webui struct {
		Autologin       ConfigDictModelBoolItem   `json:"autologin"`
		Desc            string                    `json:"desc"`
		Develop         ConfigDictModelBoolItem   `json:"develop"`
		Enabled         ConfigDictModelBoolItem   `json:"enabled"`
		Host            ConfigDictModelStringItem `json:"host"`
		Port            ConfigDictModelIntItem    `json:"port"`
		Prefix          ConfigDictModelStringItem `json:"prefix"`
		SessionLifetime ConfigDictModelIntItem    `json:"session_lifetime"`
		SslCertchain    ConfigDictModelStringItem `json:"ssl_certchain"`
		SslCertfile     ConfigDictModelStringItem `json:"ssl_certfile"`
		SslKeyfile      ConfigDictModelStringItem `json:"ssl_keyfile"`
		Theme           ConfigItemListModel       `json:"theme"`
		UseSsl          ConfigDictModelBoolItem   `json:"use_ssl"`
	} `json:"webui"`
}
