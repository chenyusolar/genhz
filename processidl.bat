@echo off
setlocal enabledelayedexpansion

set "IDL_DIR=idl"

for %%f in (%IDL_DIR%\*.thrift) do (
    set "file=%%f"
    set "base=%%~nf"
    echo Processing !file!...
    cwgo server --type HTTP --idl "!file!" --server_name "!base!_server" -module "towin"
)

endlocal
