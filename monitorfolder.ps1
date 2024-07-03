<#
.SYNOPSIS
    Monitors a specified folder for new file creation and performs an action when a new file is created.

.DESCRIPTION
    This script uses a FileSystemWatcher to monitor a specified folder for new file creation events.
    When a new file is created, it logs the creation event to a specified log file and performs any
    custom actions defined within the script.

.PARAMETER folderPath
    The path to the folder to monitor.

.PARAMETER filter
    The file type filter for monitoring (e.g., "*.json").

.PARAMETER logFilePath
    The path to the log file where file creation events will be recorded.

.EXAMPLE
    .\Monitor-Folder.ps1 -folderPath "./EventHubOutput" -filter "*.json" -logFilePath "./EventHubOutput/file.txt"
#>

param (
    [string]$folderPath = "./EventHubOutput",
    [string]$filter = "*.json",
    [string]$logFilePath = "./EventHubOutput/log.txt"
)

$action = {
    param($newFile)
    
    # Action to take when a new file is created
    Write-Host "New file created: $($newFile.FullName)"
    
    # Log file creation to a text file
    $logMessage = "$(Get-Date) - New file created: $($newFile.FullName)"
    Add-Content -Path $using:logFilePath -Value $logMessage
    
    # Add your custom actions here
    # For example, you could process the file, move it, or trigger other scripts
}

$watcher = New-Object System.IO.FileSystemWatcher
$watcher.Path = $folderPath
$watcher.Filter = $filter
$watcher.IncludeSubdirectories = $false
$watcher.EnableRaisingEvents = $true

$onCreate = Register-ObjectEvent $watcher "Created" -Action $action

try {
    Write-Host "Monitoring folder: $folderPath"
    Write-Host "Press Ctrl+C to stop monitoring."
    while ($true) { Start-Sleep -Seconds 1 }
} finally {
    Unregister-Event -SourceIdentifier $onCreate.Name
    $watcher.Dispose()
    Write-Host "Monitoring stopped."
}
