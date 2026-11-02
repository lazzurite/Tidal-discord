[Setup]
AppName=TIDAL Discord
AppVersion=1.1
DefaultDirName={autopf}\TIDAL Discord
DefaultGroupName=TIDAL Discord
OutputDir=.
OutputBaseFilename=TidalDiscordSetup-v1.1
SetupIconFile=tidal-discord.ico
Compression=lzma
SolidCompression=yes
CloseApplications=yes
RestartApplications=no

[Tasks]
Name: "desktopicon"; Description: "Создать ярлык на рабочем столе"; GroupDescription: "Дополнительно:"; Flags: unchecked
Name: "startupicon"; Description: "Запускать вместе с Windows"; GroupDescription: "Дополнительно:"; Flags: checkedonce

[Files]
Source: "..\tidal-discord.exe"; DestDir: "{app}"; Flags: ignoreversion

[Icons]
Name: "{autoprograms}\TIDAL Discord"; Filename: "{app}\tidal-discord.exe"
Name: "{autodesktop}\TIDAL Discord"; Filename: "{app}\tidal-discord.exe"; Tasks: desktopicon
Name: "{userstartup}\TIDAL Discord"; Filename: "{app}\tidal-discord.exe"; Tasks: startupicon

[Run]
Filename: "{app}\tidal-discord.exe"; Description: "Запустить TIDAL Discord"; Flags: nowait postinstall skipifsilent