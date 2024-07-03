# Azure Event Hub Reader

This Go application reads events from an Azure Event Hub and writes specific events to JSON files. It filters events based on predefined categories and operations, making it useful for monitoring and auditing purposes.

This command line was created specificly to monitor for users and devices creation/deletion in Entra ID. You will find a PowerShell script that also can be used to trigger any other actions based when a new JSON file (new event) is created.

## Installation

### On Linux & macOS using [Homebrew](https://brew.sh)

```bash
brew tap obay/tap
brew update
brew install azhubreader
```

### On Windows using [Scoop](https://scoop.sh)

```powershell
scoop bucket add obay https://github.com/obay/scoop-bucket.git
scoop update
scoop install obay/azhubreader
```

## Usage

### On Linux or macOS

Replace the values for `eventhub_name`, `consumer_group`, `eventhub_namespace`, `EventHubOutput`, `key_name`, and `key` with your own values.

```bash
eventhub_name="EntraID-eh"
consumer_group="EntraID-cg"
eventhub_namespace="EntraID-ns"
EventHubOutput="EventHubOutput"
key_name="RootManageSharedAccessKey"
key="EBzdTLpH8l+4H7xxxIq5rbFpbOVQNuO/f+AExPxeKGo=" # This is a fake key. Replace with your own key.

azhubreader -hub=$eventhub_name -group=$consumer_group -output=$EventHubOutput -namespace=$eventhub_namespace -keyname=$key_name -key=$key
```

### On Windows

Replace the values for `eventhub_name`, `consumer_group`, `eventhub_namespace`, `EventHubOutput`, `key_name`, and `key` with your own values.

```powershell
$eventhub_name="EntraID-eh"
$consumer_group="EntraID-cg"
$eventhub_namespace="EntraID-ns"
$EventHubOutput="EventHubOutput"
$key_name="RootManageSharedAccessKey"
$key="EBzdTLpH8l+4H7xxxIq5rbFpbOVQNuO/f+AExPxeKGo=" # This is a fake key. Replace with your own key.

azhubreader -hub $eventhub_name -group $consumer_group -output $EventHubOutput -namespace $eventhub_namespace -keyname $key_name -key $key
```

### On Docker

Replace the values for `eventhub_name`, `consumer_group`, `eventhub_namespace`, `EventHubOutput`, `key_name`, and `key` with your own values.

```bash
eventhub_name="EntraID-eh"
consumer_group="EntraID-cg"
eventhub_namespace="EntraID-ns"
EventHubOutput="EventHubOutput"
key_name="RootManageSharedAccessKey"
key="EBzdTLpH8l+4H7xxxIq5rbFpbOVQNuO/f+AExPxeKGo=" # This is a fake key. Replace with your own key.

docker run -it --rm \
  -v $(pwd)/$EventHubOutput:/app/EventHubOutput \
  xobay/azhubreader:v0.1.7 \
  ./azhubreader \
  -hub $eventhub_name \
  -group $consumer_group \
  -output /app/EventHubOutput \
  -namespace $eventhub_namespace \
  -keyname $key_name \
  -key $key
```

### Parameters

- `-hub`: Event Hub name
- `-group`: Consumer group name
- `-output`: Output directory for JSON files
- `-namespace`: Event Hub namespace
- `-keyname`: Shared Access Key name
- `-key`: Shared Access Key

All parameters are required.

## Configuration

The application is configured to monitor the following:

- Categories: "AuditLogs"
- Operations: "Add user", "Delete user"

To modify these, update the `monitoredCategories` and `monitoredOperations` slices in the `ReadEventHub` function.

## Output

The application will create JSON files in the specified output directory. Each file will be named using the format:
YYYY-MM-DD_HH-MM-SS-Operation_Name.json

## Notes

- The application will continue running until manually stopped.
