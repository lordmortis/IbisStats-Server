$timestamp = [Math]::Floor([decimal](Get-Date(Get-Date).ToUniversalTime()-uformat "%s"))

$scriptDir = Split-Path $script:MyInvocation.MyCommand.Path

$name = Read-Host -Prompt 'Migration Name'

$upfile = "$($timestamp)_$($name).up.sql"
$downfile= "$($timestamp)_$($name).down.sql"

Write-Host "Creating $scriptDir\datasource\migrations\$upfile"
Write-Host "Creating $scriptDir\datasource\migrations\$downfile"

Read-Host -Prompt "Press Enter to continue"

New-Item -Path $scriptDir\datasource\migrations -Name $upfile -ItemType File
New-Item -Path $scriptDir\datasource\migrations -Name $downfile -ItemType File

