[Setup]
AppName=医院管理系统
AppVersion=1.0
DefaultDirName={pf}\医院管理系统
DefaultGroupName=医院管理系统
OutputDir=.
OutputBaseFilename=医院管理系统安装程序
Compression=lzma
SolidCompression=yes
PrivilegesRequired=admin

[Dirs]
Name: "{app}"; Permissions: users-modify

[Files]
Source: "Release\hospital-system.exe"; DestDir: "{app}"; Flags: ignoreversion
Source: "Release\hospital-client\*"; DestDir: "{app}\hospital-client"; Flags: ignoreversion recursesubdirs

[Icons]
Name: "{group}\医院管理系统"; Filename: "{app}\hospital-system.exe"
Name: "{group}\卸载医院管理系统"; Filename: "{uninstallexe}"
Name: "{commondesktop}\医院管理系统"; Filename: "{app}\hospital-system.exe"; Tasks: desktopicon

[Tasks]
Name: "desktopicon"; Description: "创建桌面快捷方式"; GroupDescription: "附加图标:"

[Run]
Filename: "{app}\hospital-system.exe"; Description: "立即运行医院管理系统"; Flags: nowait postinstall skipifsilent
