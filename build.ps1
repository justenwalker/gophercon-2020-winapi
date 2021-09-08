#!powershell
[CmdletBinding()]
param (
    [Parameter(Position=0)]
    [ValidateSet(
        "all",
        "credenumeratew_managed",
        "credenumeratew_unmanaged",
        "logon",
        "netstat",
        "networkparams",
        "unsafe_cast",
        "volumes")]
    [string]
    $Command="all"
)

function Build-Command {
    param(
        [string]$Command
    )
    if ( -not (Test-Path -Path ".\bin") ) {
        New-Item -Path ".\bin" -ItemType Directory
    }
    Write-Host "Building: ${command}"
    & go build -o ".\bin\${command}.exe" ".\${command}" 
}

if ($Command -eq "all") {
    Build-Command -Command "credenumeratew_managed"
    Build-Command -Command "credenumeratew_unmanaged"
    Build-Command -Command "logon"
    Build-Command -Command "netstat"
    Build-Command -Command "networkparams"
    Build-Command -Command "unsafe_cast"
    Build-Command -Command "volumes"
} else {
    Build-Command -Command $Command
}