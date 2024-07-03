$folderPath = "./EventHubOutput"  # Replace with your folder path
$filter = "*.json"  # Monitor all file types
$action = {
    param($newFile)
    
    # Action to take when a new file is created
    Write-Host "New file created: $($newFile.FullName)"
    
    # Example: Log file creation to a text file
    $logPath = $folderPath + "file.txt"  # Replace with your log file path
    $logMessage = "$(Get-Date) - New file created: $($newFile.FullName)"
    Add-Content -Path $logPath -Value $logMessage
    
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