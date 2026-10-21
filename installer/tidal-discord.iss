[Setup]
AppName=TIDAL Discord
AppVersion=1.0.0
DefaultDirName={autopf}\TIDAL Discord
DefaultGroupName=TIDAL Discord
OutputDir=.
OutputBaseFilename=TidalDiscordSetup
SetupIconFile=tidal-discord.ico
Compression=lzma
SolidCompression=yes

[Files]
Source: "..\tidal-discord.exe"; DestDir: "{app}"; Flags: ignoreversion

[Icons]
Name: "{autoprograms}\TIDAL Discord"; Filename: "{app}\tidal-discord.exe"
Name: "{autodesktop}\TIDAL Discord"; Filename: "{app}\tidal-discord.exe"

[Run]
Filename: "{app}\tidal-discord.exe"; Description: "Запустить TIDAL Discord"; Flags: nowait postinstall skipifsilent